package main

import (
	"main/pkg/bot"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	NODE_URL     = ""
	NODE_URL_RPC = ""

	CRAB_INTERFACE_ADDR     = "0xf205ad80BB86ac92247638914265887A8BAa437D"
	STRATEGY_INTERFACE_ADDR = ""
	BRIBE_NUM               = 2
	BRIBE_DENOM             = 3
)

var (
	BRIBE_NUM_BI   = big.NewInt(BRIBE_NUM)
	BRIBE_DENOM_BI = big.NewInt(BRIBE_DENOM)
)

var environment *bot.Config

func init() {
	//config to be loaded in
	environment = &bot.Config{}

	//initialize logging library
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	//load environment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load env file")
	}
	environment.NodeURL = NODE_URL
	environment.CrabInterfaceAddress = common.HexToAddress(CRAB_INTERFACE_ADDR)
	environment.StrategyInterfaceAddress = common.HexToAddress(STRATEGY_INTERFACE_ADDR)
	environment.NodeURLHTTP = NODE_URL_RPC

	privateKeyString := os.Getenv("PRIVATE_KEY")
	if len(privateKeyString) == 0 {
		log.Fatal().Msg("PRIVATE_KEY not set in .env")
	}
	signingKeyString := os.Getenv("SIGNING_KEY")
	if len(signingKeyString) == 0 {
		log.Fatal().Msg("SIGNING_KEY not set in .env")
	}

	//parse keys
	environment.PrivateKey, err = crypto.HexToECDSA(privateKeyString)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to parse PRIVATE_KEY from .env")
	}

	environment.SigningKey, err = crypto.HexToECDSA(signingKeyString)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to parse PRIVATE_KEY from .env")
	}

	//parse public key
	publicKey := crypto.PubkeyToAddress(environment.PrivateKey.PublicKey)
	signingPublicKey := crypto.PubkeyToAddress(environment.SigningKey.PublicKey)
	log.Info().Str("wallet_address", publicKey.String()).Str("flashbots_signing_address", signingPublicKey.String()).Msg("loaded wallet")

	log.Info().Msg("loaded environment")
}

func main() {

	//create bot object
	bot, err := bot.NewBot(environment)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to initialize bot")
	}

	//start time hedge
	err = bot.StartTimeHedge(BRIBE_NUM_BI, BRIBE_DENOM_BI)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start time hedge")
	}

}
