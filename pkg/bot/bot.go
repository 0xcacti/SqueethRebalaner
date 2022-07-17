package bot

import (
	"context"
	"crypto/ecdsa"
	"main/pkg/contracts/crab"
	"main/pkg/flashbots"
	"math"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

const (
	GAS_LIMIT = uint64(1_200_000)
)

var (
	PARSE_ERR      = errors.New("big int parse error")
	MIN_W_SQUEETH  = big.NewInt(0)
	MIN_ETH        = big.NewInt(0)
	ULTIMATE_BRIBE = Frac{Num: big.NewInt(90), Denom: big.NewInt(100)}
)

type Frac struct {
	Num   *big.Int
	Denom *big.Int
}

//if it simulates and you get some shit with priority fee
type Bot struct {
	//connections
	Client                   *ethclient.Client
	ClientEndpoint           string
	CrabInterface            *crab.Crab
	MyContractInterface      *mycontract.MyContract
	ChainID                  *big.Int
	Signer                   types.Signer
	PrivateKey               *ecdsa.PrivateKey
	SigningKey               *ecdsa.PrivateKey
	CrabInterfaceAddress     common.Address
	StrategyInterfaceAddress common.Address
	StrategyABI              abi.ABI
}

type Config struct {
	NodeURL                  string
	NodeURLHTTP              string
	PrivateKey               *ecdsa.PrivateKey
	SigningKey               *ecdsa.PrivateKey
	CrabInterfaceAddress     common.Address
	StrategyInterfaceAddress common.Address
}

func NewBot(config *Config) (*Bot, error) {
	//attempt connection to eth client (prerequisite)
	var err error
	client, err := ethclient.Dial(config.NodeURL)
	if err != nil {
		return nil, errors.Wrap(err, "failed to dial eth client")
	}
	log.Info().Str("address", config.NodeURL).Msg("connected to client")

	//grab chain ID
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "failed to get chain id")
	}

	//create signer
	signer := types.NewLondonSigner(chainID)

	//initialize contract interfaces
	crabInterface, err := crab.NewCrab(config.CrabInterfaceAddress, client)
	if err != nil {
		return nil, errors.Wrap(err, "failed to instantiate crab contract interface")
	}
	// create strategy
	myContract, err := MyContract.NewMyContract(config.StrategyInterfaceAddress, client)
	if err != nil {
		return nil, errors.Wrap(err, "failed to instantiate crab contract interface")
	}

	strategyABI, err := abi.JSON(strings.NewReader(mycontract.MyContractABI))
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse strategyABI")
	}

	return &Bot{
		Client:                   client,
		ClientEndpoint:           config.NodeURLHTTP,
		CrabInterface:            crabInterface,
		MyContractInterface:      myContract,
		StrategyInterfaceAddress: config.StrategyInterfaceAddress,
		StrategyABI:              strategyABI,
		ChainID:                  chainID,
		Signer:                   signer,
		PrivateKey:               config.PrivateKey,
		SigningKey:               config.SigningKey,
		CrabInterfaceAddress:     config.CrabInterfaceAddress,
	}, nil
}

func (bot *Bot) GetAuctionTriggerTime() (*big.Int, error) {
	//call check time hedge
	_, auctionTriggerTime, err := bot.CrabInterface.CheckTimeHedge(&bind.CallOpts{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to call CheckTimeHedge")
	}
	return auctionTriggerTime, nil
}

func (bot *Bot) GetTxData(minWSqueeth *big.Int, minETH *big.Int, bribeNum *big.Int, bribeDenom *big.Int, gaslimit uint64) ([]byte, error) {

	data, err := bot.StrategyABI.Pack("TimeHedge", minWSqueeth, minETH, bribeNum, bribeDenom)
	if err != nil {
		return nil, errors.Wrap(err, "failed to pack data with abi")
	}

	return data, nil
}

func (bot *Bot) SignTransactionData(to *common.Address, value *big.Int, data []byte, nonce uint64, gasLimit uint64, gasFeeCap *big.Int, gasTipCap *big.Int) (*types.Transaction, error) {
	// create tx type
	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   bot.ChainID,
		Nonce:     nonce,
		GasFeeCap: gasFeeCap,
		GasTipCap: gasTipCap,
		Gas:       gasLimit,
		To:        to,
		Value:     value,
		Data:      data,
	})

	// sign tx and return
	signedTx, err := types.SignTx(tx, bot.Signer, bot.PrivateKey)
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign transaction data during build")
	}
	return signedTx, nil
}

func (bot *Bot) SimulateTx(txn *types.Transaction) (*flashbots.CallBundleResponse, error) {
	//create flashbots provider
	start := time.Now()
	fb := flashbots.NewProvider(bot.SigningKey, bot.PrivateKey, bot.ClientEndpoint)

	//marshal binary data
	data, err := txn.MarshalBinary()
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal tx binary data")
	}
	tx := hexutil.Encode(data)

	//get block header
	block, err := bot.Client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get block number")
	}
	//simulate tx
	resp, err := fb.Simulate([]string{tx}, block.Number, "latest", 0)
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute simulation code on bundle")
	}

	err = resp.HasError()
	if err != nil {
		return nil, errors.Wrap(err, "error from transaction simulation")
	}

	cb, _ := new(big.Float).SetString(resp.Result.CoinbaseDiff)
	eth := new(big.Float).Quo(cb, big.NewFloat(math.Pow10(18)))
	wei, _ := resp.EffectiveGasPrice()
	gwei := new(big.Int).Div(wei, big.NewInt(1_000_000_000))

	log.Info().Dur("time_ms", time.Since(start)).Str("cost", eth.String()).Str("effective_price", gwei.String()).Msg("mint simulation completed")

	return resp, nil
}

func (bot *Bot) SubmitTxRelay(txn *types.Transaction, attempts int64) error { // I don't exactly remember how to do this
	fb := flashbots.NewProvider(bot.SigningKey, bot.PrivateKey, flashbots.DefaultRelayURL)
	data, err := txn.MarshalBinary()
	if err != nil {
		return errors.Wrap(err, "failed to marshal tx binary data")
	}
	tx := hexutil.Encode(data)
	block, err := bot.Client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return errors.Wrap(err, "failed to get block number")
	}
	for i := int64(0); i < attempts; i++ {
		targetBlockNumber := new(big.Int).Add(block.Number, big.NewInt(int64(i)))
		_, err := fb.SendBundle([]string{tx}, targetBlockNumber, &flashbots.Options{})
		if err != nil {
			return errors.Wrap(err, "issue in submit bundle")
		}
	}
	return nil
}

func (bot *Bot) StartTimeHedge(bribeNum *big.Int, bribeDenom *big.Int) error {
	//get trigger time
	auctionTriggerTime, err := bot.GetAuctionTriggerTime()
	if err != nil {
		return errors.Wrap(err, "failed to get auction trigger time")
	}

	// TESTING ONLY
	// auctionTriggerTime := big.NewInt(time.Now().Add(time.Second * 60).Unix())

	log.Info().Int64("unix_time", auctionTriggerTime.Int64()).Msg("received auction trigger time")

	//parse times, start a little bit before trigger time
	triggerTimeParsed := time.Unix(auctionTriggerTime.Int64(), 0)
	triggerTimeParsed = triggerTimeParsed.Add(time.Minute * -2)

	log.Info().Int64("duration_min", int64(time.Until(triggerTimeParsed).Minutes())).Msg("waiting until almost auction trigger time")

	//wait until squeeth time
	time.Sleep(time.Until(triggerTimeParsed))
	log.Info().Dur("dur", 5*time.Second).Msg("it's squeeth time, starting gas updater")

	//create stop timer
	stopTimer := time.NewTimer(time.Minute * 30)

	//get wallet noncey boy
	nonce, err := bot.Client.PendingNonceAt(context.Background(), crypto.PubkeyToAddress(bot.PrivateKey.PublicKey))
	if err != nil {
		return errors.Wrap(err, "failed to get wallet nonce")
	}

	//generate tx with data
	txData, err := bot.GetTxData(MIN_W_SQUEETH, MIN_ETH, bribeNum, bribeDenom, GAS_LIMIT)
	if err != nil {
		return errors.Wrap(err, "failed to get tx data")
	}

	//get initial gas fee cap
	gasFeeCap, err := bot.Client.SuggestGasPrice(context.Background()) // new(big.Float).Mul(big.NewFloat(gasPrice.BlockPrices[0].BaseFeePerGas), new(big.Float).SetFloat64(math.Pow(10.0, 9.0))).Int(nil)
	if err != nil {
		return errors.Wrap(err, "failed to get gas fee cap")
	}

	//get initial gas tip cap
	gasTipCap, err := bot.Client.SuggestGasTipCap(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed to get gas tip cap")
	}

simulateLoop:
	for {
		select {
		case <-stopTimer.C:
			break simulateLoop

		default:
			tx, err := bot.SignTransactionData(&bot.StrategyInterfaceAddress, common.Big0, txData, nonce, GAS_LIMIT, gasFeeCap, gasTipCap)
			if err != nil {
				log.Err(err).Msg("failed to get sign tx data")
				time.Sleep(time.Second)
				break
			}

			//smimulate transaction
			resp, err := bot.SimulateTx(tx)
			var percentBumpGas Frac
			if err != nil {
				log.Err(err).Msg("failed to simulate tx")
				if strings.Contains(err.Error(), "less than") { // pay more gas
					percentBumpGas = Frac{big.NewInt(1), big.NewInt(10)}
				} else if strings.Contains(err.Error(), "insufficient") { // pay more gas
					percentBumpGas = Frac{big.NewInt(-1), big.NewInt(10)}
				}

				if (percentBumpGas != Frac{}) {
					bumpAmt := new(big.Int).Mul(percentBumpGas.Num, new(big.Int).Div(gasFeeCap, percentBumpGas.Denom))
					gasFeeCap = new(big.Int).Add(gasFeeCap, bumpAmt)
				}
				//exit select loop
				break
			}
			log.Info().Msg("successfully simulated tx")

			tx, err = bot.ApplyGasStrat(resp.Result, bribeNum, bribeDenom, &bot.StrategyInterfaceAddress, common.Big0, txData, nonce, GAS_LIMIT, big.NewInt(97), big.NewInt(100)) // add arg for gas strat
			if err != nil {
				log.Err(err).Msg("failed to update txn")
				time.Sleep(time.Second)
				break
			}

			err = bot.SubmitTxRelay(tx, 5)
			if err != nil {
				log.Err(err).Msg("failed to submit to flashbots")
				break
			}

			// stop the bot?
			log.Info().Msg("success. time to make love")
			time.Sleep(time.Second * 120)
		}

	}

	return nil
}

func (bot *Bot) ApplyGasStrat(simData *flashbots.CallResult, bribeNum *big.Int, bribeDenom *big.Int, to *common.Address, value *big.Int, data []byte, nonce uint64, gasLimit uint64, newBribeNum *big.Int, newBribeDenom *big.Int) (*types.Transaction, error) {
	// get eth sent to coinbase in gwei
	ethToCB, ok := new(big.Int).SetString(simData.EthSentToCoinbase, 10)
	if !ok {
		log.Err(nil).Msg("failed to read EthSentToCoinbase into *big.Int")
		return nil, PARSE_ERR
	}

	estGasUsed := big.NewInt(int64(simData.TotalGasUsed))
	revenue := new(big.Int).Div(new(big.Int).Mul(ethToCB, bribeDenom), bribeNum)
	revShare := new(big.Int).Mul(newBribeNum, new(big.Int).Div(revenue, newBribeDenom))
	gasFeeCapGwei := new(big.Int).Mul(new(big.Int).Sub(revShare, ethToCB), estGasUsed)
	gasFeeCap := new(big.Int).Mul(gasFeeCapGwei, big.NewInt(10e9))
	gasTipCap := common.Big0

	// create tx type
	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   bot.ChainID,
		Nonce:     nonce,
		GasFeeCap: gasFeeCap,
		GasTipCap: gasTipCap,
		Gas:       gasLimit,
		To:        to,
		Value:     value,
		Data:      data,
	})

	// sign tx and return
	signedTx, err := types.SignTx(tx, bot.Signer, bot.PrivateKey)
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign transaction data during build")
	}
	return signedTx, nil

}
