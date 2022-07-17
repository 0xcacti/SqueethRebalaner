// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package crab

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// CrabABI is the input ABI used to generate the binding from.
const CrabABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wSqueethController\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_weth\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_uniswapFactory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_ethWSqueethPool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_hedgeTimeThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_hedgePriceThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_auctionTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minPriceMultiplier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxPriceMultiplier\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wSqueethAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lpAmount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wSqueethBought\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ethSold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isHedgingOnUniswap\",\"type\":\"bool\"}],\"name\":\"ExecuteBuyAuction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wSqueethSold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ethBought\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isHedgingOnUniswap\",\"type\":\"bool\"}],\"name\":\"ExecuteSellAuction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tradedAmountOut\",\"type\":\"uint256\"}],\"name\":\"FlashDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"flashswapDebt\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"excess\",\"type\":\"uint256\"}],\"name\":\"FlashDepositCallback\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"withdrawer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"crabAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wSqueethAmount\",\"type\":\"uint256\"}],\"name\":\"FlashWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"withdrawer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"flashswapDebt\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"excess\",\"type\":\"uint256\"}],\"name\":\"FlashWithdrawCallback\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"hedger\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"auctionType\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"hedgerPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auctionPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wSqueethHedgeTargetAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ethHedgetargetAmount\",\"type\":\"uint256\"}],\"name\":\"Hedge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"hedger\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"auctionType\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auctionPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wSqueethHedgeTargetAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ethHedgetargetAmount\",\"type\":\"uint256\"}],\"name\":\"HedgeOnUniswap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"hedger\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"auctionType\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"hedgerPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auctionTriggerTimestamp\",\"type\":\"uint256\"}],\"name\":\"PriceHedge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"hedger\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"hedgeTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auctionTriggerTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minWSqueeth\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minEth\",\"type\":\"uint256\"}],\"name\":\"PriceHedgeOnUniswap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newAuctionTime\",\"type\":\"uint256\"}],\"name\":\"SetAuctionTime\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDeltaHedgeThreshold\",\"type\":\"uint256\"}],\"name\":\"SetDeltaHedgeThreshold\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newHedgePriceThreshold\",\"type\":\"uint256\"}],\"name\":\"SetHedgePriceThreshold\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newHedgeTimeThreshold\",\"type\":\"uint256\"}],\"name\":\"SetHedgeTimeThreshold\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"newHedgingTwapPeriod\",\"type\":\"uint32\"}],\"name\":\"SetHedgingTwapPeriod\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMaxPriceMultiplier\",\"type\":\"uint256\"}],\"name\":\"SetMaxPriceMultiplier\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMinPriceMultiplier\",\"type\":\"uint256\"}],\"name\":\"SetMinPriceMultiplier\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCapAmount\",\"type\":\"uint256\"}],\"name\":\"SetStrategyCap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"hedger\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"auctionType\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"hedgerPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auctionTriggerTimestamp\",\"type\":\"uint256\"}],\"name\":\"TimeHedge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"hedger\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"hedgeTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"auctionTriggerTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minWSqueeth\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minEth\",\"type\":\"uint256\"}],\"name\":\"TimeHedgeOnUniswap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"withdrawer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"crabAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wSqueethAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ethWithdrawn\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"withdrawer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"crabAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ethWithdrawn\",\"type\":\"uint256\"}],\"name\":\"WithdrawShutdown\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"POWER_PERP_PERIOD\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"auctionTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_auctionTriggerTime\",\"type\":\"uint256\"}],\"name\":\"checkPriceHedge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkTimeHedge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deltaHedgeThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ethQuoteCurrencyPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ethWSqueethPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_ethToDeposit\",\"type\":\"uint256\"}],\"name\":\"flashDeposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_crabAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxEthToPay\",\"type\":\"uint256\"}],\"name\":\"flashWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_auctionTriggerTime\",\"type\":\"uint256\"}],\"name\":\"getAuctionDetails\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStrategyVaultId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVaultDetails\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_crabAmount\",\"type\":\"uint256\"}],\"name\":\"getWsqueethFromCrabAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hedgePriceThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hedgeTimeThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hedgingTwapPeriod\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxPriceMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minPriceMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"powerTokenController\",\"outputs\":[{\"internalType\":\"contractIController\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceAtLastHedge\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_auctionTriggerTime\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isStrategySellingWSqueeth\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_limitPrice\",\"type\":\"uint256\"}],\"name\":\"priceHedge\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_auctionTriggerTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minWSqueeth\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minEth\",\"type\":\"uint256\"}],\"name\":\"priceHedgeOnUniswap\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quoteCurrency\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"redeemShortShutdown\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_auctionTime\",\"type\":\"uint256\"}],\"name\":\"setAuctionTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_deltaHedgeThreshold\",\"type\":\"uint256\"}],\"name\":\"setDeltaHedgeThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_hedgePriceThreshold\",\"type\":\"uint256\"}],\"name\":\"setHedgePriceThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_hedgeTimeThreshold\",\"type\":\"uint256\"}],\"name\":\"setHedgeTimeThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_hedgingTwapPeriod\",\"type\":\"uint32\"}],\"name\":\"setHedgingTwapPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxPriceMultiplier\",\"type\":\"uint256\"}],\"name\":\"setMaxPriceMultiplier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minPriceMultiplier\",\"type\":\"uint256\"}],\"name\":\"setMinPriceMultiplier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_capAmount\",\"type\":\"uint256\"}],\"name\":\"setStrategyCap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"strategyCap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeAtLastHedge\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_isStrategySellingWSqueeth\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_limitPrice\",\"type\":\"uint256\"}],\"name\":\"timeHedge\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minWSqueeth\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minEth\",\"type\":\"uint256\"}],\"name\":\"timeHedgeOnUniswap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0Delta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1Delta\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"uniswapV3SwapCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vaultId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wPowerPerp\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weth\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_crabAmount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_crabAmount\",\"type\":\"uint256\"}],\"name\":\"withdrawShutdown\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// Crab is an auto generated Go binding around an Ethereum contract.
type Crab struct {
	CrabCaller     // Read-only binding to the contract
	CrabTransactor // Write-only binding to the contract
	CrabFilterer   // Log filterer for contract events
}

// CrabCaller is an auto generated read-only Go binding around an Ethereum contract.
type CrabCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrabTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CrabTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrabFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CrabFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrabSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CrabSession struct {
	Contract     *Crab             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CrabCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CrabCallerSession struct {
	Contract *CrabCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CrabTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CrabTransactorSession struct {
	Contract     *CrabTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CrabRaw is an auto generated low-level Go binding around an Ethereum contract.
type CrabRaw struct {
	Contract *Crab // Generic contract binding to access the raw methods on
}

// CrabCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CrabCallerRaw struct {
	Contract *CrabCaller // Generic read-only contract binding to access the raw methods on
}

// CrabTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CrabTransactorRaw struct {
	Contract *CrabTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCrab creates a new instance of Crab, bound to a specific deployed contract.
func NewCrab(address common.Address, backend bind.ContractBackend) (*Crab, error) {
	contract, err := bindCrab(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Crab{CrabCaller: CrabCaller{contract: contract}, CrabTransactor: CrabTransactor{contract: contract}, CrabFilterer: CrabFilterer{contract: contract}}, nil
}

// NewCrabCaller creates a new read-only instance of Crab, bound to a specific deployed contract.
func NewCrabCaller(address common.Address, caller bind.ContractCaller) (*CrabCaller, error) {
	contract, err := bindCrab(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CrabCaller{contract: contract}, nil
}

// NewCrabTransactor creates a new write-only instance of Crab, bound to a specific deployed contract.
func NewCrabTransactor(address common.Address, transactor bind.ContractTransactor) (*CrabTransactor, error) {
	contract, err := bindCrab(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CrabTransactor{contract: contract}, nil
}

// NewCrabFilterer creates a new log filterer instance of Crab, bound to a specific deployed contract.
func NewCrabFilterer(address common.Address, filterer bind.ContractFilterer) (*CrabFilterer, error) {
	contract, err := bindCrab(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CrabFilterer{contract: contract}, nil
}

// bindCrab binds a generic wrapper to an already deployed contract.
func bindCrab(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CrabABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Crab *CrabRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Crab.Contract.CrabCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Crab *CrabRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Crab.Contract.CrabTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Crab *CrabRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Crab.Contract.CrabTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Crab *CrabCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Crab.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Crab *CrabTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Crab.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Crab *CrabTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Crab.Contract.contract.Transact(opts, method, params...)
}

// POWERPERPPERIOD is a free data retrieval call binding the contract method 0x395ebb69.
//
// Solidity: function POWER_PERP_PERIOD() view returns(uint32)
func (_Crab *CrabCaller) POWERPERPPERIOD(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Crab.contract.Call(opts, &out, "POWER_PERP_PERIOD")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// POWERPERPPERIOD is a free data retrieval call binding the contract method 0x395ebb69.
//
// Solidity: function POWER_PERP_PERIOD() view returns(uint32)
func (_Crab *CrabSession) POWERPERPPERIOD() (uint32, error) {
	return _Crab.Contract.POWERPERPPERIOD(&_Crab.CallOpts)
}

// POWERPERPPERIOD is a free data retrieval call binding the contract method 0x395ebb69.
//
// Solidity: function POWER_PERP_PERIOD() view returns(uint32)
func (_Crab *CrabCallerSession) POWERPERPPERIOD() (uint32, error) {
	return _Crab.Contract.POWERPERPPERIOD(&_Crab.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Crab *CrabCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Crab.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Crab *CrabSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Crab.Contract.Allowance(&_Crab.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Crab *CrabCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Crab.Contract.Allowance(&_Crab.CallOpts, owner, spender)
}

// AuctionTime is a free data retrieval call binding the contract method 0x155f586d.
//
// Solidity: function auctionTime() view returns(uint256)
func (_Crab *CrabCaller) AuctionTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Crab.contract.Call(opts, &out, "auctionTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AuctionTime is a free data retrieval call binding the contract method 0x155f586d.
//
// Solidity: function auctionTime() view returns(uint256)
func (_Crab *CrabSession) AuctionTime() (*big.Int, error) {
	return _Crab.Contract.AuctionTime(&_Crab.CallOpts)
}

// AuctionTime is a free data retrieval call binding the contract method 0x155f586d.
//
// Solidity: function auctionTime() view returns(uint256)
func (_Crab *CrabCallerSession) AuctionTime() (*big.Int, error) {
	return _Crab.Contract.AuctionTime(&_Crab.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Crab *CrabCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Crab.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Crab *CrabSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Crab.Contract.BalanceOf(&_Crab.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Crab *CrabCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Crab.Contract.BalanceOf(&_Crab.CallOpts, account)
}

// CheckPriceHedge is a free data retrieval call binding the contract method 0xb0b93461.
//
// Solidity: function checkPriceHedge(uint256 _auctionTriggerTime) view returns(bool)
func (_Crab *CrabCaller) CheckPriceHedge(opts *bind.CallOpts, _auctionTriggerTime *big.Int) (bool, error) {
	var out []interface{}
	err := _Crab.contract.Call(opts, &out, "checkPriceHedge", _auctionTriggerTime)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckPriceHedge is a free data retrieval call binding the contract method 0xb0b93461.
//
// Solidity: function checkPriceHedge(uint256 _auctionTriggerTime) view returns(bool)
func (_Crab *CrabSession) CheckPriceHedge(_auctionTriggerTime *big.Int) (bool, error) {
	return _Crab.Contract.CheckPriceHedge(&_Crab.CallOpts, _auctionTriggerTime)
}

// CheckPriceHedge is a free data retrieval call binding the contract method 0xb0b93461.
//
// Solidity: function checkPriceHedge(uint256 _auctionTriggerTime) view returns(bool)
func (_Crab *CrabCallerSession) CheckPriceHedge(_auctionTriggerTime *big.Int) (bool, error) {
	return _Crab.Contract.CheckPriceHedge(&_Crab.CallOpts, _auctionTriggerTime)
}

// CheckTimeHedge is a free data retrieval call binding the contract method 0xc2451689.
//
// Solidity: function checkTimeHedge() view returns(bool, uint256)
func (_Crab *CrabCaller) CheckTimeHedge(opts *bind.CallOpts) (bool, *big.Int, error) {
	var out []interface{}
	err := _Crab.contract.Call(opts, &out, "checkTimeHedge")

	if err != nil {
		return *new(bool), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// CheckTimeHedge is a free data retrieval call binding the contract method 0xc2451689.
//
// Solidity: function checkTimeHedge() view returns(bool, uint256)
func (_Crab *CrabSession) CheckTimeHedge() (bool, *big.Int, error) {
	return _Crab.Contract.CheckTimeHedge(&_Crab.CallOpts)
}

// CheckTimeHedge is a free data retrieval call binding the contract method 0xc2451689.
//
// Solidity: function checkTimeHedge() view returns(bool, uint256)
func (_Crab *CrabCallerSession) CheckTimeHedge() (bool, *big.Int, error) {
	return _Crab.Contract.CheckTimeHedge(&_Crab.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Crab *CrabCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Crab.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Crab *CrabSession) Decimals() (uint8, error) {
	return _Crab.Contract.Decimals(&_Crab.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Crab *CrabCallerSession) Decimals() (uint8, error) {
	return _Crab.Contract.Decimals(&_Crab.CallOpts)
}

// DeltaHedgeThreshold is a free data retrieval call binding the contract method 0x955a15e8.
//
// Solidity: function deltaHedgeThreshold() view returns(uint256)
func (_Crab *CrabCaller) DeltaHedgeThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Crab.contract.Call(opts, &out, "deltaHedgeThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DeltaHedgeThreshold is a free data retrieval call binding the contract method 0x955a15e8.
//
// Solidity: function deltaHedgeThreshold() view returns(uint256)
func (_Crab *CrabSession) DeltaHedgeThreshold() (*big.Int, error) {
	return _Crab.Contract.DeltaHedgeThreshold(&_Crab.CallOpts)
}

// DeltaHedgeThreshold is a free data retrieval call binding the contract method 0x955a15e8.
//
// Solidity: function deltaHedgeThreshold() view returns(uint256)
func (_Crab *CrabCallerSession) DeltaHedgeThreshold() (*big.Int, error) {
	return _Crab.Contract.DeltaHedgeThreshold(&_Crab.CallOpts)
}

// EthQuoteCurrencyPool is a free data retrieval call binding the contract method 0x4468c022.
//
// Solidity: function ethQuoteCurrencyPool() view returns(address)
func (_Crab *CrabCaller) EthQuoteCurrencyPool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Crab.contract.Call(opts, &out, "ethQuoteCurrencyPool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EthQuoteCurrencyPool is a free data retrieval call binding the contract method 0x4468c022.
//
// Solidity: function ethQuoteCurrencyPool() view returns(address)
func (_Crab *CrabSession) EthQuoteCurrencyPool() (common.Address, error) {
	return _Crab.Contract.EthQuoteCurrencyPool(&_Crab.CallOpts)
}

// EthQuoteCurrencyPool is a free data retrieval call binding the contract method 0x4468c022.
//
// Solidity: function ethQuoteCurrencyPool() view returns(address)
func (_Crab *CrabCallerSession) EthQuoteCurrencyPool() (common.Address, error) {
	return _Crab.Contract.EthQuoteCurrencyPool(&_Crab.CallOpts)
}

// EthWSqueethPool is a free data retrieval call binding the contract method 0x4d76e6fc.
//
// Solidity: function ethWSqueethPool() view returns(address)
func (_Crab *CrabCaller) EthWSqueethPool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Crab.contract.Call(opts, &out, "ethWSqueethPool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EthWSqueethPool is a free data retrieval call binding the contract method 0x4d76e6fc.
//
// Solidity: function ethWSqueethPool() view returns(address)
func (_Crab *CrabSession) EthWSqueethPool() (common.Address, error) {
	return _Crab.Contract.EthWSqueethPool(&_Crab.CallOpts)
}

// EthWSqueethPool is a free data retrieval call binding the contract method 0x4d76e6fc.
//
// Solidity: function ethWSqueethPool() view returns(address)
func (_Crab *CrabCallerSession) EthWSqueethPool() (common.Address, error) {
	return _Crab.Contract.EthWSqueethPool(&_Crab.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Crab *CrabCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Crab.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Crab *CrabSession) Factory() (common.Address, error) {
	return _Crab.Contract.Factory(&_Crab.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Crab *CrabCallerSession) Factory() (common.Address, error) {
	return _Crab.Contract.Factory(&_Crab.CallOpts)
}

// GetAuctionDetails is a free data retrieval call binding the contract method 0xf20e5e35.
//
// Solidity: function getAuctionDetails(uint256 _auctionTriggerTime) view returns(bool, uint256, uint256, uint256, bool)
func (_Crab *CrabCaller) GetAuctionDetails(opts *bind.CallOpts, _auctionTriggerTime *big.Int) (bool, *big.Int, *big.Int, *big.Int, bool, error) {
	var out []interface{}
	err := _Crab.contract.Call(opts, &out, "getAuctionDetails", _auctionTriggerTime)

	if err != nil {
		return *new(bool), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(bool)).(*bool)

	return out0, out1, out2, out3, out4, err

}

// GetAuctionDetails is a free data retrieval call binding the contract method 0xf20e5e35.
//
// Solidity: function getAuctionDetails(uint256 _auctionTriggerTime) view returns(bool, uint256, uint256, uint256, bool)
func (_Crab *CrabSession) GetAuctionDetails(_auctionTriggerTime *big.Int) (bool, *big.Int, *big.Int, *big.Int, bool, error) {
	return _Crab.Contract.GetAuctionDetails(&_Crab.CallOpts, _auctionTriggerTime)
}

// GetAuctionDetails is a free data retrieval call binding the contract method 0xf20e5e35.
//
// Solidity: function getAuctionDetails(uint256 _auctionTriggerTime) view returns(bool, uint256, uint256, uint256, bool)
func (_Crab *CrabCallerSession) GetAuctionDetails(_auctionTriggerTime *big.Int) (bool, *big.Int, *big.Int, *big.Int, bool, error) {
	return _Crab.Contract.GetAuctionDetails(&_Crab.CallOpts, _auctionTriggerTime)
}

// GetStrategyVaultId is a free data retrieval call binding the contract method 0x6c1040a9.
//
// Solidity: function getStrategyVaultId() view returns(uint256)
func (_Crab *CrabCaller) GetStrategyVaultId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Crab.contract.Call(opts, &out, "getStrategyVaultId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStrategyVaultId is a free data retrieval call binding the contract method 0x6c1040a9.
//
// Solidity: function getStrategyVaultId() view returns(uint256)
func (_Crab *CrabSession) GetStrategyVaultId() (*big.Int, error) {
	return _Crab.Contract.GetStrategyVaultId(&_Crab.CallOpts)
}

// GetStrategyVaultId is a free data retrieval call binding the contract method 0x6c1040a9.
//
// Solidity: function getStrategyVaultId() view returns(uint256)
func (_Crab *CrabCallerSession) GetStrategyVaultId() (*big.Int, error) {
	return _Crab.Contract.GetStrategyVaultId(&_Crab.CallOpts)
}

// GetVaultDetails is a free data retrieval call binding the contract method 0x533092ef.
//
// Solidity: function getVaultDetails() view returns(address, uint256, uint256, uint256)
func (_Crab *CrabCaller) GetVaultDetails(opts *bind.CallOpts) (common.Address, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Crab.contract.Call(opts, &out, "getVaultDetails")

	if err != nil {
		return *new(common.Address), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// GetVaultDetails is a free data retrieval call binding the contract method 0x533092ef.
//
// Solidity: function getVaultDetails() view returns(address, uint256, uint256, uint256)
func (_Crab *CrabSession) GetVaultDetails() (common.Address, *big.Int, *big.Int, *big.Int, error) {
	return _Crab.Contract.GetVaultDetails(&_Crab.CallOpts)
}

// GetVaultDetails is a free data retrieval call binding the contract method 0x533092ef.
//
// Solidity: function getVaultDetails() view returns(address, uint256, uint256, uint256)
func (_Crab *CrabCallerSession) GetVaultDetails() (common.Address, *big.Int, *big.Int, *big.Int, error) {
	return _Crab.Contract.GetVaultDetails(&_Crab.CallOpts)
}

// GetWsqueethFromCrabAmount is a free data retrieval call binding the contract method 0xf73e19c3.
//
// Solidity: function getWsqueethFromCrabAmount(uint256 _crabAmount) view returns(uint256)
func (_Crab *CrabCaller) GetWsqueethFromCrabAmount(opts *bind.CallOpts, _crabAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Crab.contract.Call(opts, &out, "getWsqueethFromCrabAmount", _crabAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWsqueethFromCrabAmount is a free data retrieval call binding the contract method 0xf73e19c3.
//
// Solidity: function getWsqueethFromCrabAmount(uint256 _crabAmount) view returns(uint256)
func (_Crab *CrabSession) GetWsqueethFromCrabAmount(_crabAmount *big.Int) (*big.Int, error) {
	return _Crab.Contract.GetWsqueethFromCrabAmount(&_Crab.CallOpts, _crabAmount)
}

// GetWsqueethFromCrabAmount is a free data retrieval call binding the contract method 0xf73e19c3.
//
// Solidity: function getWsqueethFromCrabAmount(uint256 _crabAmount) view returns(uint256)
func (_Crab *CrabCallerSession) GetWsqueethFromCrabAmount(_crabAmount *big.Int) (*big.Int, error) {
	return _Crab.Contract.GetWsqueethFromCrabAmount(&_Crab.CallOpts, _crabAmount)
}

// HedgePriceThreshold is a free data retrieval call binding the contract method 0xf101d92f.
//
// Solidity: function hedgePriceThreshold() view returns(uint256)
func (_Crab *CrabCaller) HedgePriceThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Crab.contract.Call(opts, &out, "hedgePriceThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HedgePriceThreshold is a free data retrieval call binding the contract method 0xf101d92f.
//
// Solidity: function hedgePriceThreshold() view returns(uint256)
func (_Crab *CrabSession) HedgePriceThreshold() (*big.Int, error) {
	return _Crab.Contract.HedgePriceThreshold(&_Crab.CallOpts)
}

// HedgePriceThreshold is a free data retrieval call binding the contract method 0xf101d92f.
//
// Solidity: function hedgePriceThreshold() view returns(uint256)
func (_Crab *CrabCallerSession) HedgePriceThreshold() (*big.Int, error) {
	return _Crab.Contract.HedgePriceThreshold(&_Crab.CallOpts)
}

// HedgeTimeThreshold is a free data retrieval call binding the contract method 0x66a91b76.
//
// Solidity: function hedgeTimeThreshold() view returns(uint256)
func (_Crab *CrabCaller) HedgeTimeThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Crab.contract.Call(opts, &out, "hedgeTimeThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HedgeTimeThreshold is a free data retrieval call binding the contract method 0x66a91b76.
//
// Solidity: function hedgeTimeThreshold() view returns(uint256)
func (_Crab *CrabSession) HedgeTimeThreshold() (*big.Int, error) {
	return _Crab.Contract.HedgeTimeThreshold(&_Crab.CallOpts)
}

// HedgeTimeThreshold is a free data retrieval call binding the contract method 0x66a91b76.
//
// Solidity: function hedgeTimeThreshold() view returns(uint256)
func (_Crab *CrabCallerSession) HedgeTimeThreshold() (*big.Int, error) {
	return _Crab.Contract.HedgeTimeThreshold(&_Crab.CallOpts)
}

// HedgingTwapPeriod is a free data retrieval call binding the contract method 0xf5d278e4.
//
// Solidity: function hedgingTwapPeriod() view returns(uint32)
func (_Crab *CrabCaller) HedgingTwapPeriod(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Crab.contract.Call(opts, &out, "hedgingTwapPeriod")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// HedgingTwapPeriod is a free data retrieval call binding the contract method 0xf5d278e4.
//
// Solidity: function hedgingTwapPeriod() view returns(uint32)
func (_Crab *CrabSession) HedgingTwapPeriod() (uint32, error) {
	return _Crab.Contract.HedgingTwapPeriod(&_Crab.CallOpts)
}

// HedgingTwapPeriod is a free data retrieval call binding the contract method 0xf5d278e4.
//
// Solidity: function hedgingTwapPeriod() view returns(uint32)
func (_Crab *CrabCallerSession) HedgingTwapPeriod() (uint32, error) {
	return _Crab.Contract.HedgingTwapPeriod(&_Crab.CallOpts)
}

// MaxPriceMultiplier is a free data retrieval call binding the contract method 0x8ec4b721.
//
// Solidity: function maxPriceMultiplier() view returns(uint256)
func (_Crab *CrabCaller) MaxPriceMultiplier(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Crab.contract.Call(opts, &out, "maxPriceMultiplier")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxPriceMultiplier is a free data retrieval call binding the contract method 0x8ec4b721.
//
// Solidity: function maxPriceMultiplier() view returns(uint256)
func (_Crab *CrabSession) MaxPriceMultiplier() (*big.Int, error) {
	return _Crab.Contract.MaxPriceMultiplier(&_Crab.CallOpts)
}

// MaxPriceMultiplier is a free data retrieval call binding the contract method 0x8ec4b721.
//
// Solidity: function maxPriceMultiplier() view returns(uint256)
func (_Crab *CrabCallerSession) MaxPriceMultiplier() (*big.Int, error) {
	return _Crab.Contract.MaxPriceMultiplier(&_Crab.CallOpts)
}

// MinPriceMultiplier is a free data retrieval call binding the contract method 0x2a340056.
//
// Solidity: function minPriceMultiplier() view returns(uint256)
func (_Crab *CrabCaller) MinPriceMultiplier(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Crab.contract.Call(opts, &out, "minPriceMultiplier")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinPriceMultiplier is a free data retrieval call binding the contract method 0x2a340056.
//
// Solidity: function minPriceMultiplier() view returns(uint256)
func (_Crab *CrabSession) MinPriceMultiplier() (*big.Int, error) {
	return _Crab.Contract.MinPriceMultiplier(&_Crab.CallOpts)
}

// MinPriceMultiplier is a free data retrieval call binding the contract method 0x2a340056.
//
// Solidity: function minPriceMultiplier() view returns(uint256)
func (_Crab *CrabCallerSession) MinPriceMultiplier() (*big.Int, error) {
	return _Crab.Contract.MinPriceMultiplier(&_Crab.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Crab *CrabCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Crab.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Crab *CrabSession) Name() (string, error) {
	return _Crab.Contract.Name(&_Crab.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Crab *CrabCallerSession) Name() (string, error) {
	return _Crab.Contract.Name(&_Crab.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Crab *CrabCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Crab.contract.Call(opts, &out, "oracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Crab *CrabSession) Oracle() (common.Address, error) {
	return _Crab.Contract.Oracle(&_Crab.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Crab *CrabCallerSession) Oracle() (common.Address, error) {
	return _Crab.Contract.Oracle(&_Crab.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Crab *CrabCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Crab.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Crab *CrabSession) Owner() (common.Address, error) {
	return _Crab.Contract.Owner(&_Crab.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Crab *CrabCallerSession) Owner() (common.Address, error) {
	return _Crab.Contract.Owner(&_Crab.CallOpts)
}

// PowerTokenController is a free data retrieval call binding the contract method 0x3dcb0c5d.
//
// Solidity: function powerTokenController() view returns(address)
func (_Crab *CrabCaller) PowerTokenController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Crab.contract.Call(opts, &out, "powerTokenController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PowerTokenController is a free data retrieval call binding the contract method 0x3dcb0c5d.
//
// Solidity: function powerTokenController() view returns(address)
func (_Crab *CrabSession) PowerTokenController() (common.Address, error) {
	return _Crab.Contract.PowerTokenController(&_Crab.CallOpts)
}

// PowerTokenController is a free data retrieval call binding the contract method 0x3dcb0c5d.
//
// Solidity: function powerTokenController() view returns(address)
func (_Crab *CrabCallerSession) PowerTokenController() (common.Address, error) {
	return _Crab.Contract.PowerTokenController(&_Crab.CallOpts)
}

// PriceAtLastHedge is a free data retrieval call binding the contract method 0x63bbc4b6.
//
// Solidity: function priceAtLastHedge() view returns(uint256)
func (_Crab *CrabCaller) PriceAtLastHedge(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Crab.contract.Call(opts, &out, "priceAtLastHedge")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PriceAtLastHedge is a free data retrieval call binding the contract method 0x63bbc4b6.
//
// Solidity: function priceAtLastHedge() view returns(uint256)
func (_Crab *CrabSession) PriceAtLastHedge() (*big.Int, error) {
	return _Crab.Contract.PriceAtLastHedge(&_Crab.CallOpts)
}

// PriceAtLastHedge is a free data retrieval call binding the contract method 0x63bbc4b6.
//
// Solidity: function priceAtLastHedge() view returns(uint256)
func (_Crab *CrabCallerSession) PriceAtLastHedge() (*big.Int, error) {
	return _Crab.Contract.PriceAtLastHedge(&_Crab.CallOpts)
}

// QuoteCurrency is a free data retrieval call binding the contract method 0x82564bca.
//
// Solidity: function quoteCurrency() view returns(address)
func (_Crab *CrabCaller) QuoteCurrency(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Crab.contract.Call(opts, &out, "quoteCurrency")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// QuoteCurrency is a free data retrieval call binding the contract method 0x82564bca.
//
// Solidity: function quoteCurrency() view returns(address)
func (_Crab *CrabSession) QuoteCurrency() (common.Address, error) {
	return _Crab.Contract.QuoteCurrency(&_Crab.CallOpts)
}

// QuoteCurrency is a free data retrieval call binding the contract method 0x82564bca.
//
// Solidity: function quoteCurrency() view returns(address)
func (_Crab *CrabCallerSession) QuoteCurrency() (common.Address, error) {
	return _Crab.Contract.QuoteCurrency(&_Crab.CallOpts)
}

// StrategyCap is a free data retrieval call binding the contract method 0xcae74029.
//
// Solidity: function strategyCap() view returns(uint256)
func (_Crab *CrabCaller) StrategyCap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Crab.contract.Call(opts, &out, "strategyCap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StrategyCap is a free data retrieval call binding the contract method 0xcae74029.
//
// Solidity: function strategyCap() view returns(uint256)
func (_Crab *CrabSession) StrategyCap() (*big.Int, error) {
	return _Crab.Contract.StrategyCap(&_Crab.CallOpts)
}

// StrategyCap is a free data retrieval call binding the contract method 0xcae74029.
//
// Solidity: function strategyCap() view returns(uint256)
func (_Crab *CrabCallerSession) StrategyCap() (*big.Int, error) {
	return _Crab.Contract.StrategyCap(&_Crab.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Crab *CrabCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Crab.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Crab *CrabSession) Symbol() (string, error) {
	return _Crab.Contract.Symbol(&_Crab.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Crab *CrabCallerSession) Symbol() (string, error) {
	return _Crab.Contract.Symbol(&_Crab.CallOpts)
}

// TimeAtLastHedge is a free data retrieval call binding the contract method 0x67b8c345.
//
// Solidity: function timeAtLastHedge() view returns(uint256)
func (_Crab *CrabCaller) TimeAtLastHedge(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Crab.contract.Call(opts, &out, "timeAtLastHedge")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TimeAtLastHedge is a free data retrieval call binding the contract method 0x67b8c345.
//
// Solidity: function timeAtLastHedge() view returns(uint256)
func (_Crab *CrabSession) TimeAtLastHedge() (*big.Int, error) {
	return _Crab.Contract.TimeAtLastHedge(&_Crab.CallOpts)
}

// TimeAtLastHedge is a free data retrieval call binding the contract method 0x67b8c345.
//
// Solidity: function timeAtLastHedge() view returns(uint256)
func (_Crab *CrabCallerSession) TimeAtLastHedge() (*big.Int, error) {
	return _Crab.Contract.TimeAtLastHedge(&_Crab.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Crab *CrabCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Crab.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Crab *CrabSession) TotalSupply() (*big.Int, error) {
	return _Crab.Contract.TotalSupply(&_Crab.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Crab *CrabCallerSession) TotalSupply() (*big.Int, error) {
	return _Crab.Contract.TotalSupply(&_Crab.CallOpts)
}

// VaultId is a free data retrieval call binding the contract method 0x33194c0a.
//
// Solidity: function vaultId() view returns(uint256)
func (_Crab *CrabCaller) VaultId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Crab.contract.Call(opts, &out, "vaultId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VaultId is a free data retrieval call binding the contract method 0x33194c0a.
//
// Solidity: function vaultId() view returns(uint256)
func (_Crab *CrabSession) VaultId() (*big.Int, error) {
	return _Crab.Contract.VaultId(&_Crab.CallOpts)
}

// VaultId is a free data retrieval call binding the contract method 0x33194c0a.
//
// Solidity: function vaultId() view returns(uint256)
func (_Crab *CrabCallerSession) VaultId() (*big.Int, error) {
	return _Crab.Contract.VaultId(&_Crab.CallOpts)
}

// WPowerPerp is a free data retrieval call binding the contract method 0x7f07b130.
//
// Solidity: function wPowerPerp() view returns(address)
func (_Crab *CrabCaller) WPowerPerp(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Crab.contract.Call(opts, &out, "wPowerPerp")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WPowerPerp is a free data retrieval call binding the contract method 0x7f07b130.
//
// Solidity: function wPowerPerp() view returns(address)
func (_Crab *CrabSession) WPowerPerp() (common.Address, error) {
	return _Crab.Contract.WPowerPerp(&_Crab.CallOpts)
}

// WPowerPerp is a free data retrieval call binding the contract method 0x7f07b130.
//
// Solidity: function wPowerPerp() view returns(address)
func (_Crab *CrabCallerSession) WPowerPerp() (common.Address, error) {
	return _Crab.Contract.WPowerPerp(&_Crab.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_Crab *CrabCaller) Weth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Crab.contract.Call(opts, &out, "weth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_Crab *CrabSession) Weth() (common.Address, error) {
	return _Crab.Contract.Weth(&_Crab.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_Crab *CrabCallerSession) Weth() (common.Address, error) {
	return _Crab.Contract.Weth(&_Crab.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Crab *CrabTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Crab.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Crab *CrabSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Crab.Contract.Approve(&_Crab.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Crab *CrabTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Crab.Contract.Approve(&_Crab.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Crab *CrabTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Crab.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Crab *CrabSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Crab.Contract.DecreaseAllowance(&_Crab.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Crab *CrabTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Crab.Contract.DecreaseAllowance(&_Crab.TransactOpts, spender, subtractedValue)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Crab *CrabTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Crab.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Crab *CrabSession) Deposit() (*types.Transaction, error) {
	return _Crab.Contract.Deposit(&_Crab.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Crab *CrabTransactorSession) Deposit() (*types.Transaction, error) {
	return _Crab.Contract.Deposit(&_Crab.TransactOpts)
}

// FlashDeposit is a paid mutator transaction binding the contract method 0x1a5af342.
//
// Solidity: function flashDeposit(uint256 _ethToDeposit) payable returns()
func (_Crab *CrabTransactor) FlashDeposit(opts *bind.TransactOpts, _ethToDeposit *big.Int) (*types.Transaction, error) {
	return _Crab.contract.Transact(opts, "flashDeposit", _ethToDeposit)
}

// FlashDeposit is a paid mutator transaction binding the contract method 0x1a5af342.
//
// Solidity: function flashDeposit(uint256 _ethToDeposit) payable returns()
func (_Crab *CrabSession) FlashDeposit(_ethToDeposit *big.Int) (*types.Transaction, error) {
	return _Crab.Contract.FlashDeposit(&_Crab.TransactOpts, _ethToDeposit)
}

// FlashDeposit is a paid mutator transaction binding the contract method 0x1a5af342.
//
// Solidity: function flashDeposit(uint256 _ethToDeposit) payable returns()
func (_Crab *CrabTransactorSession) FlashDeposit(_ethToDeposit *big.Int) (*types.Transaction, error) {
	return _Crab.Contract.FlashDeposit(&_Crab.TransactOpts, _ethToDeposit)
}

// FlashWithdraw is a paid mutator transaction binding the contract method 0x23ccafd9.
//
// Solidity: function flashWithdraw(uint256 _crabAmount, uint256 _maxEthToPay) returns()
func (_Crab *CrabTransactor) FlashWithdraw(opts *bind.TransactOpts, _crabAmount *big.Int, _maxEthToPay *big.Int) (*types.Transaction, error) {
	return _Crab.contract.Transact(opts, "flashWithdraw", _crabAmount, _maxEthToPay)
}

// FlashWithdraw is a paid mutator transaction binding the contract method 0x23ccafd9.
//
// Solidity: function flashWithdraw(uint256 _crabAmount, uint256 _maxEthToPay) returns()
func (_Crab *CrabSession) FlashWithdraw(_crabAmount *big.Int, _maxEthToPay *big.Int) (*types.Transaction, error) {
	return _Crab.Contract.FlashWithdraw(&_Crab.TransactOpts, _crabAmount, _maxEthToPay)
}

// FlashWithdraw is a paid mutator transaction binding the contract method 0x23ccafd9.
//
// Solidity: function flashWithdraw(uint256 _crabAmount, uint256 _maxEthToPay) returns()
func (_Crab *CrabTransactorSession) FlashWithdraw(_crabAmount *big.Int, _maxEthToPay *big.Int) (*types.Transaction, error) {
	return _Crab.Contract.FlashWithdraw(&_Crab.TransactOpts, _crabAmount, _maxEthToPay)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Crab *CrabTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Crab.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Crab *CrabSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Crab.Contract.IncreaseAllowance(&_Crab.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Crab *CrabTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Crab.Contract.IncreaseAllowance(&_Crab.TransactOpts, spender, addedValue)
}

// PriceHedge is a paid mutator transaction binding the contract method 0x70749b44.
//
// Solidity: function priceHedge(uint256 _auctionTriggerTime, bool _isStrategySellingWSqueeth, uint256 _limitPrice) payable returns()
func (_Crab *CrabTransactor) PriceHedge(opts *bind.TransactOpts, _auctionTriggerTime *big.Int, _isStrategySellingWSqueeth bool, _limitPrice *big.Int) (*types.Transaction, error) {
	return _Crab.contract.Transact(opts, "priceHedge", _auctionTriggerTime, _isStrategySellingWSqueeth, _limitPrice)
}

// PriceHedge is a paid mutator transaction binding the contract method 0x70749b44.
//
// Solidity: function priceHedge(uint256 _auctionTriggerTime, bool _isStrategySellingWSqueeth, uint256 _limitPrice) payable returns()
func (_Crab *CrabSession) PriceHedge(_auctionTriggerTime *big.Int, _isStrategySellingWSqueeth bool, _limitPrice *big.Int) (*types.Transaction, error) {
	return _Crab.Contract.PriceHedge(&_Crab.TransactOpts, _auctionTriggerTime, _isStrategySellingWSqueeth, _limitPrice)
}

// PriceHedge is a paid mutator transaction binding the contract method 0x70749b44.
//
// Solidity: function priceHedge(uint256 _auctionTriggerTime, bool _isStrategySellingWSqueeth, uint256 _limitPrice) payable returns()
func (_Crab *CrabTransactorSession) PriceHedge(_auctionTriggerTime *big.Int, _isStrategySellingWSqueeth bool, _limitPrice *big.Int) (*types.Transaction, error) {
	return _Crab.Contract.PriceHedge(&_Crab.TransactOpts, _auctionTriggerTime, _isStrategySellingWSqueeth, _limitPrice)
}

// PriceHedgeOnUniswap is a paid mutator transaction binding the contract method 0x5e5cdcd7.
//
// Solidity: function priceHedgeOnUniswap(uint256 _auctionTriggerTime, uint256 _minWSqueeth, uint256 _minEth) payable returns()
func (_Crab *CrabTransactor) PriceHedgeOnUniswap(opts *bind.TransactOpts, _auctionTriggerTime *big.Int, _minWSqueeth *big.Int, _minEth *big.Int) (*types.Transaction, error) {
	return _Crab.contract.Transact(opts, "priceHedgeOnUniswap", _auctionTriggerTime, _minWSqueeth, _minEth)
}

// PriceHedgeOnUniswap is a paid mutator transaction binding the contract method 0x5e5cdcd7.
//
// Solidity: function priceHedgeOnUniswap(uint256 _auctionTriggerTime, uint256 _minWSqueeth, uint256 _minEth) payable returns()
func (_Crab *CrabSession) PriceHedgeOnUniswap(_auctionTriggerTime *big.Int, _minWSqueeth *big.Int, _minEth *big.Int) (*types.Transaction, error) {
	return _Crab.Contract.PriceHedgeOnUniswap(&_Crab.TransactOpts, _auctionTriggerTime, _minWSqueeth, _minEth)
}

// PriceHedgeOnUniswap is a paid mutator transaction binding the contract method 0x5e5cdcd7.
//
// Solidity: function priceHedgeOnUniswap(uint256 _auctionTriggerTime, uint256 _minWSqueeth, uint256 _minEth) payable returns()
func (_Crab *CrabTransactorSession) PriceHedgeOnUniswap(_auctionTriggerTime *big.Int, _minWSqueeth *big.Int, _minEth *big.Int) (*types.Transaction, error) {
	return _Crab.Contract.PriceHedgeOnUniswap(&_Crab.TransactOpts, _auctionTriggerTime, _minWSqueeth, _minEth)
}

// RedeemShortShutdown is a paid mutator transaction binding the contract method 0xa319b29f.
//
// Solidity: function redeemShortShutdown() returns()
func (_Crab *CrabTransactor) RedeemShortShutdown(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Crab.contract.Transact(opts, "redeemShortShutdown")
}

// RedeemShortShutdown is a paid mutator transaction binding the contract method 0xa319b29f.
//
// Solidity: function redeemShortShutdown() returns()
func (_Crab *CrabSession) RedeemShortShutdown() (*types.Transaction, error) {
	return _Crab.Contract.RedeemShortShutdown(&_Crab.TransactOpts)
}

// RedeemShortShutdown is a paid mutator transaction binding the contract method 0xa319b29f.
//
// Solidity: function redeemShortShutdown() returns()
func (_Crab *CrabTransactorSession) RedeemShortShutdown() (*types.Transaction, error) {
	return _Crab.Contract.RedeemShortShutdown(&_Crab.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Crab *CrabTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Crab.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Crab *CrabSession) RenounceOwnership() (*types.Transaction, error) {
	return _Crab.Contract.RenounceOwnership(&_Crab.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Crab *CrabTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Crab.Contract.RenounceOwnership(&_Crab.TransactOpts)
}

// SetAuctionTime is a paid mutator transaction binding the contract method 0xa54c8899.
//
// Solidity: function setAuctionTime(uint256 _auctionTime) returns()
func (_Crab *CrabTransactor) SetAuctionTime(opts *bind.TransactOpts, _auctionTime *big.Int) (*types.Transaction, error) {
	return _Crab.contract.Transact(opts, "setAuctionTime", _auctionTime)
}

// SetAuctionTime is a paid mutator transaction binding the contract method 0xa54c8899.
//
// Solidity: function setAuctionTime(uint256 _auctionTime) returns()
func (_Crab *CrabSession) SetAuctionTime(_auctionTime *big.Int) (*types.Transaction, error) {
	return _Crab.Contract.SetAuctionTime(&_Crab.TransactOpts, _auctionTime)
}

// SetAuctionTime is a paid mutator transaction binding the contract method 0xa54c8899.
//
// Solidity: function setAuctionTime(uint256 _auctionTime) returns()
func (_Crab *CrabTransactorSession) SetAuctionTime(_auctionTime *big.Int) (*types.Transaction, error) {
	return _Crab.Contract.SetAuctionTime(&_Crab.TransactOpts, _auctionTime)
}

// SetDeltaHedgeThreshold is a paid mutator transaction binding the contract method 0xd4aec817.
//
// Solidity: function setDeltaHedgeThreshold(uint256 _deltaHedgeThreshold) returns()
func (_Crab *CrabTransactor) SetDeltaHedgeThreshold(opts *bind.TransactOpts, _deltaHedgeThreshold *big.Int) (*types.Transaction, error) {
	return _Crab.contract.Transact(opts, "setDeltaHedgeThreshold", _deltaHedgeThreshold)
}

// SetDeltaHedgeThreshold is a paid mutator transaction binding the contract method 0xd4aec817.
//
// Solidity: function setDeltaHedgeThreshold(uint256 _deltaHedgeThreshold) returns()
func (_Crab *CrabSession) SetDeltaHedgeThreshold(_deltaHedgeThreshold *big.Int) (*types.Transaction, error) {
	return _Crab.Contract.SetDeltaHedgeThreshold(&_Crab.TransactOpts, _deltaHedgeThreshold)
}

// SetDeltaHedgeThreshold is a paid mutator transaction binding the contract method 0xd4aec817.
//
// Solidity: function setDeltaHedgeThreshold(uint256 _deltaHedgeThreshold) returns()
func (_Crab *CrabTransactorSession) SetDeltaHedgeThreshold(_deltaHedgeThreshold *big.Int) (*types.Transaction, error) {
	return _Crab.Contract.SetDeltaHedgeThreshold(&_Crab.TransactOpts, _deltaHedgeThreshold)
}

// SetHedgePriceThreshold is a paid mutator transaction binding the contract method 0xbdd438b8.
//
// Solidity: function setHedgePriceThreshold(uint256 _hedgePriceThreshold) returns()
func (_Crab *CrabTransactor) SetHedgePriceThreshold(opts *bind.TransactOpts, _hedgePriceThreshold *big.Int) (*types.Transaction, error) {
	return _Crab.contract.Transact(opts, "setHedgePriceThreshold", _hedgePriceThreshold)
}

// SetHedgePriceThreshold is a paid mutator transaction binding the contract method 0xbdd438b8.
//
// Solidity: function setHedgePriceThreshold(uint256 _hedgePriceThreshold) returns()
func (_Crab *CrabSession) SetHedgePriceThreshold(_hedgePriceThreshold *big.Int) (*types.Transaction, error) {
	return _Crab.Contract.SetHedgePriceThreshold(&_Crab.TransactOpts, _hedgePriceThreshold)
}

// SetHedgePriceThreshold is a paid mutator transaction binding the contract method 0xbdd438b8.
//
// Solidity: function setHedgePriceThreshold(uint256 _hedgePriceThreshold) returns()
func (_Crab *CrabTransactorSession) SetHedgePriceThreshold(_hedgePriceThreshold *big.Int) (*types.Transaction, error) {
	return _Crab.Contract.SetHedgePriceThreshold(&_Crab.TransactOpts, _hedgePriceThreshold)
}

// SetHedgeTimeThreshold is a paid mutator transaction binding the contract method 0x7bcdc16e.
//
// Solidity: function setHedgeTimeThreshold(uint256 _hedgeTimeThreshold) returns()
func (_Crab *CrabTransactor) SetHedgeTimeThreshold(opts *bind.TransactOpts, _hedgeTimeThreshold *big.Int) (*types.Transaction, error) {
	return _Crab.contract.Transact(opts, "setHedgeTimeThreshold", _hedgeTimeThreshold)
}

// SetHedgeTimeThreshold is a paid mutator transaction binding the contract method 0x7bcdc16e.
//
// Solidity: function setHedgeTimeThreshold(uint256 _hedgeTimeThreshold) returns()
func (_Crab *CrabSession) SetHedgeTimeThreshold(_hedgeTimeThreshold *big.Int) (*types.Transaction, error) {
	return _Crab.Contract.SetHedgeTimeThreshold(&_Crab.TransactOpts, _hedgeTimeThreshold)
}

// SetHedgeTimeThreshold is a paid mutator transaction binding the contract method 0x7bcdc16e.
//
// Solidity: function setHedgeTimeThreshold(uint256 _hedgeTimeThreshold) returns()
func (_Crab *CrabTransactorSession) SetHedgeTimeThreshold(_hedgeTimeThreshold *big.Int) (*types.Transaction, error) {
	return _Crab.Contract.SetHedgeTimeThreshold(&_Crab.TransactOpts, _hedgeTimeThreshold)
}

// SetHedgingTwapPeriod is a paid mutator transaction binding the contract method 0x8f8b8dbc.
//
// Solidity: function setHedgingTwapPeriod(uint32 _hedgingTwapPeriod) returns()
func (_Crab *CrabTransactor) SetHedgingTwapPeriod(opts *bind.TransactOpts, _hedgingTwapPeriod uint32) (*types.Transaction, error) {
	return _Crab.contract.Transact(opts, "setHedgingTwapPeriod", _hedgingTwapPeriod)
}

// SetHedgingTwapPeriod is a paid mutator transaction binding the contract method 0x8f8b8dbc.
//
// Solidity: function setHedgingTwapPeriod(uint32 _hedgingTwapPeriod) returns()
func (_Crab *CrabSession) SetHedgingTwapPeriod(_hedgingTwapPeriod uint32) (*types.Transaction, error) {
	return _Crab.Contract.SetHedgingTwapPeriod(&_Crab.TransactOpts, _hedgingTwapPeriod)
}

// SetHedgingTwapPeriod is a paid mutator transaction binding the contract method 0x8f8b8dbc.
//
// Solidity: function setHedgingTwapPeriod(uint32 _hedgingTwapPeriod) returns()
func (_Crab *CrabTransactorSession) SetHedgingTwapPeriod(_hedgingTwapPeriod uint32) (*types.Transaction, error) {
	return _Crab.Contract.SetHedgingTwapPeriod(&_Crab.TransactOpts, _hedgingTwapPeriod)
}

// SetMaxPriceMultiplier is a paid mutator transaction binding the contract method 0xdfc19520.
//
// Solidity: function setMaxPriceMultiplier(uint256 _maxPriceMultiplier) returns()
func (_Crab *CrabTransactor) SetMaxPriceMultiplier(opts *bind.TransactOpts, _maxPriceMultiplier *big.Int) (*types.Transaction, error) {
	return _Crab.contract.Transact(opts, "setMaxPriceMultiplier", _maxPriceMultiplier)
}

// SetMaxPriceMultiplier is a paid mutator transaction binding the contract method 0xdfc19520.
//
// Solidity: function setMaxPriceMultiplier(uint256 _maxPriceMultiplier) returns()
func (_Crab *CrabSession) SetMaxPriceMultiplier(_maxPriceMultiplier *big.Int) (*types.Transaction, error) {
	return _Crab.Contract.SetMaxPriceMultiplier(&_Crab.TransactOpts, _maxPriceMultiplier)
}

// SetMaxPriceMultiplier is a paid mutator transaction binding the contract method 0xdfc19520.
//
// Solidity: function setMaxPriceMultiplier(uint256 _maxPriceMultiplier) returns()
func (_Crab *CrabTransactorSession) SetMaxPriceMultiplier(_maxPriceMultiplier *big.Int) (*types.Transaction, error) {
	return _Crab.Contract.SetMaxPriceMultiplier(&_Crab.TransactOpts, _maxPriceMultiplier)
}

// SetMinPriceMultiplier is a paid mutator transaction binding the contract method 0xfa004fbf.
//
// Solidity: function setMinPriceMultiplier(uint256 _minPriceMultiplier) returns()
func (_Crab *CrabTransactor) SetMinPriceMultiplier(opts *bind.TransactOpts, _minPriceMultiplier *big.Int) (*types.Transaction, error) {
	return _Crab.contract.Transact(opts, "setMinPriceMultiplier", _minPriceMultiplier)
}

// SetMinPriceMultiplier is a paid mutator transaction binding the contract method 0xfa004fbf.
//
// Solidity: function setMinPriceMultiplier(uint256 _minPriceMultiplier) returns()
func (_Crab *CrabSession) SetMinPriceMultiplier(_minPriceMultiplier *big.Int) (*types.Transaction, error) {
	return _Crab.Contract.SetMinPriceMultiplier(&_Crab.TransactOpts, _minPriceMultiplier)
}

// SetMinPriceMultiplier is a paid mutator transaction binding the contract method 0xfa004fbf.
//
// Solidity: function setMinPriceMultiplier(uint256 _minPriceMultiplier) returns()
func (_Crab *CrabTransactorSession) SetMinPriceMultiplier(_minPriceMultiplier *big.Int) (*types.Transaction, error) {
	return _Crab.Contract.SetMinPriceMultiplier(&_Crab.TransactOpts, _minPriceMultiplier)
}

// SetStrategyCap is a paid mutator transaction binding the contract method 0x0ca514cd.
//
// Solidity: function setStrategyCap(uint256 _capAmount) returns()
func (_Crab *CrabTransactor) SetStrategyCap(opts *bind.TransactOpts, _capAmount *big.Int) (*types.Transaction, error) {
	return _Crab.contract.Transact(opts, "setStrategyCap", _capAmount)
}

// SetStrategyCap is a paid mutator transaction binding the contract method 0x0ca514cd.
//
// Solidity: function setStrategyCap(uint256 _capAmount) returns()
func (_Crab *CrabSession) SetStrategyCap(_capAmount *big.Int) (*types.Transaction, error) {
	return _Crab.Contract.SetStrategyCap(&_Crab.TransactOpts, _capAmount)
}

// SetStrategyCap is a paid mutator transaction binding the contract method 0x0ca514cd.
//
// Solidity: function setStrategyCap(uint256 _capAmount) returns()
func (_Crab *CrabTransactorSession) SetStrategyCap(_capAmount *big.Int) (*types.Transaction, error) {
	return _Crab.Contract.SetStrategyCap(&_Crab.TransactOpts, _capAmount)
}

// TimeHedge is a paid mutator transaction binding the contract method 0x25d7707c.
//
// Solidity: function timeHedge(bool _isStrategySellingWSqueeth, uint256 _limitPrice) payable returns()
func (_Crab *CrabTransactor) TimeHedge(opts *bind.TransactOpts, _isStrategySellingWSqueeth bool, _limitPrice *big.Int) (*types.Transaction, error) {
	return _Crab.contract.Transact(opts, "timeHedge", _isStrategySellingWSqueeth, _limitPrice)
}

// TimeHedge is a paid mutator transaction binding the contract method 0x25d7707c.
//
// Solidity: function timeHedge(bool _isStrategySellingWSqueeth, uint256 _limitPrice) payable returns()
func (_Crab *CrabSession) TimeHedge(_isStrategySellingWSqueeth bool, _limitPrice *big.Int) (*types.Transaction, error) {
	return _Crab.Contract.TimeHedge(&_Crab.TransactOpts, _isStrategySellingWSqueeth, _limitPrice)
}

// TimeHedge is a paid mutator transaction binding the contract method 0x25d7707c.
//
// Solidity: function timeHedge(bool _isStrategySellingWSqueeth, uint256 _limitPrice) payable returns()
func (_Crab *CrabTransactorSession) TimeHedge(_isStrategySellingWSqueeth bool, _limitPrice *big.Int) (*types.Transaction, error) {
	return _Crab.Contract.TimeHedge(&_Crab.TransactOpts, _isStrategySellingWSqueeth, _limitPrice)
}

// TimeHedgeOnUniswap is a paid mutator transaction binding the contract method 0x9c1ab1b8.
//
// Solidity: function timeHedgeOnUniswap(uint256 _minWSqueeth, uint256 _minEth) returns()
func (_Crab *CrabTransactor) TimeHedgeOnUniswap(opts *bind.TransactOpts, _minWSqueeth *big.Int, _minEth *big.Int) (*types.Transaction, error) {
	return _Crab.contract.Transact(opts, "timeHedgeOnUniswap", _minWSqueeth, _minEth)
}

// TimeHedgeOnUniswap is a paid mutator transaction binding the contract method 0x9c1ab1b8.
//
// Solidity: function timeHedgeOnUniswap(uint256 _minWSqueeth, uint256 _minEth) returns()
func (_Crab *CrabSession) TimeHedgeOnUniswap(_minWSqueeth *big.Int, _minEth *big.Int) (*types.Transaction, error) {
	return _Crab.Contract.TimeHedgeOnUniswap(&_Crab.TransactOpts, _minWSqueeth, _minEth)
}

// TimeHedgeOnUniswap is a paid mutator transaction binding the contract method 0x9c1ab1b8.
//
// Solidity: function timeHedgeOnUniswap(uint256 _minWSqueeth, uint256 _minEth) returns()
func (_Crab *CrabTransactorSession) TimeHedgeOnUniswap(_minWSqueeth *big.Int, _minEth *big.Int) (*types.Transaction, error) {
	return _Crab.Contract.TimeHedgeOnUniswap(&_Crab.TransactOpts, _minWSqueeth, _minEth)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Crab *CrabTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Crab.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Crab *CrabSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Crab.Contract.Transfer(&_Crab.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Crab *CrabTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Crab.Contract.Transfer(&_Crab.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Crab *CrabTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Crab.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Crab *CrabSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Crab.Contract.TransferFrom(&_Crab.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Crab *CrabTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Crab.Contract.TransferFrom(&_Crab.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Crab *CrabTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Crab.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Crab *CrabSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Crab.Contract.TransferOwnership(&_Crab.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Crab *CrabTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Crab.Contract.TransferOwnership(&_Crab.TransactOpts, newOwner)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes _data) returns()
func (_Crab *CrabTransactor) UniswapV3SwapCallback(opts *bind.TransactOpts, amount0Delta *big.Int, amount1Delta *big.Int, _data []byte) (*types.Transaction, error) {
	return _Crab.contract.Transact(opts, "uniswapV3SwapCallback", amount0Delta, amount1Delta, _data)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes _data) returns()
func (_Crab *CrabSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, _data []byte) (*types.Transaction, error) {
	return _Crab.Contract.UniswapV3SwapCallback(&_Crab.TransactOpts, amount0Delta, amount1Delta, _data)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes _data) returns()
func (_Crab *CrabTransactorSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, _data []byte) (*types.Transaction, error) {
	return _Crab.Contract.UniswapV3SwapCallback(&_Crab.TransactOpts, amount0Delta, amount1Delta, _data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _crabAmount) returns()
func (_Crab *CrabTransactor) Withdraw(opts *bind.TransactOpts, _crabAmount *big.Int) (*types.Transaction, error) {
	return _Crab.contract.Transact(opts, "withdraw", _crabAmount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _crabAmount) returns()
func (_Crab *CrabSession) Withdraw(_crabAmount *big.Int) (*types.Transaction, error) {
	return _Crab.Contract.Withdraw(&_Crab.TransactOpts, _crabAmount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _crabAmount) returns()
func (_Crab *CrabTransactorSession) Withdraw(_crabAmount *big.Int) (*types.Transaction, error) {
	return _Crab.Contract.Withdraw(&_Crab.TransactOpts, _crabAmount)
}

// WithdrawShutdown is a paid mutator transaction binding the contract method 0x3d3a62ee.
//
// Solidity: function withdrawShutdown(uint256 _crabAmount) returns()
func (_Crab *CrabTransactor) WithdrawShutdown(opts *bind.TransactOpts, _crabAmount *big.Int) (*types.Transaction, error) {
	return _Crab.contract.Transact(opts, "withdrawShutdown", _crabAmount)
}

// WithdrawShutdown is a paid mutator transaction binding the contract method 0x3d3a62ee.
//
// Solidity: function withdrawShutdown(uint256 _crabAmount) returns()
func (_Crab *CrabSession) WithdrawShutdown(_crabAmount *big.Int) (*types.Transaction, error) {
	return _Crab.Contract.WithdrawShutdown(&_Crab.TransactOpts, _crabAmount)
}

// WithdrawShutdown is a paid mutator transaction binding the contract method 0x3d3a62ee.
//
// Solidity: function withdrawShutdown(uint256 _crabAmount) returns()
func (_Crab *CrabTransactorSession) WithdrawShutdown(_crabAmount *big.Int) (*types.Transaction, error) {
	return _Crab.Contract.WithdrawShutdown(&_Crab.TransactOpts, _crabAmount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Crab *CrabTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Crab.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Crab *CrabSession) Receive() (*types.Transaction, error) {
	return _Crab.Contract.Receive(&_Crab.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Crab *CrabTransactorSession) Receive() (*types.Transaction, error) {
	return _Crab.Contract.Receive(&_Crab.TransactOpts)
}

// CrabApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Crab contract.
type CrabApprovalIterator struct {
	Event *CrabApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrabApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrabApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrabApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrabApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrabApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrabApproval represents a Approval event raised by the Crab contract.
type CrabApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Crab *CrabFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*CrabApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Crab.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &CrabApprovalIterator{contract: _Crab.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Crab *CrabFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CrabApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Crab.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrabApproval)
				if err := _Crab.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Crab *CrabFilterer) ParseApproval(log types.Log) (*CrabApproval, error) {
	event := new(CrabApproval)
	if err := _Crab.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrabDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Crab contract.
type CrabDepositIterator struct {
	Event *CrabDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrabDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrabDeposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrabDeposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrabDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrabDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrabDeposit represents a Deposit event raised by the Crab contract.
type CrabDeposit struct {
	Depositor      common.Address
	WSqueethAmount *big.Int
	LpAmount       *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed depositor, uint256 wSqueethAmount, uint256 lpAmount)
func (_Crab *CrabFilterer) FilterDeposit(opts *bind.FilterOpts, depositor []common.Address) (*CrabDepositIterator, error) {

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _Crab.contract.FilterLogs(opts, "Deposit", depositorRule)
	if err != nil {
		return nil, err
	}
	return &CrabDepositIterator{contract: _Crab.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed depositor, uint256 wSqueethAmount, uint256 lpAmount)
func (_Crab *CrabFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *CrabDeposit, depositor []common.Address) (event.Subscription, error) {

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _Crab.contract.WatchLogs(opts, "Deposit", depositorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrabDeposit)
				if err := _Crab.contract.UnpackLog(event, "Deposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeposit is a log parse operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed depositor, uint256 wSqueethAmount, uint256 lpAmount)
func (_Crab *CrabFilterer) ParseDeposit(log types.Log) (*CrabDeposit, error) {
	event := new(CrabDeposit)
	if err := _Crab.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrabExecuteBuyAuctionIterator is returned from FilterExecuteBuyAuction and is used to iterate over the raw logs and unpacked data for ExecuteBuyAuction events raised by the Crab contract.
type CrabExecuteBuyAuctionIterator struct {
	Event *CrabExecuteBuyAuction // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrabExecuteBuyAuctionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrabExecuteBuyAuction)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrabExecuteBuyAuction)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrabExecuteBuyAuctionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrabExecuteBuyAuctionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrabExecuteBuyAuction represents a ExecuteBuyAuction event raised by the Crab contract.
type CrabExecuteBuyAuction struct {
	Seller             common.Address
	WSqueethBought     *big.Int
	EthSold            *big.Int
	IsHedgingOnUniswap bool
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterExecuteBuyAuction is a free log retrieval operation binding the contract event 0xc7472bd0a757f40f801e047dd4f4ec901314e95afcbfed844e62af18401e0e6b.
//
// Solidity: event ExecuteBuyAuction(address indexed seller, uint256 wSqueethBought, uint256 ethSold, bool isHedgingOnUniswap)
func (_Crab *CrabFilterer) FilterExecuteBuyAuction(opts *bind.FilterOpts, seller []common.Address) (*CrabExecuteBuyAuctionIterator, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _Crab.contract.FilterLogs(opts, "ExecuteBuyAuction", sellerRule)
	if err != nil {
		return nil, err
	}
	return &CrabExecuteBuyAuctionIterator{contract: _Crab.contract, event: "ExecuteBuyAuction", logs: logs, sub: sub}, nil
}

// WatchExecuteBuyAuction is a free log subscription operation binding the contract event 0xc7472bd0a757f40f801e047dd4f4ec901314e95afcbfed844e62af18401e0e6b.
//
// Solidity: event ExecuteBuyAuction(address indexed seller, uint256 wSqueethBought, uint256 ethSold, bool isHedgingOnUniswap)
func (_Crab *CrabFilterer) WatchExecuteBuyAuction(opts *bind.WatchOpts, sink chan<- *CrabExecuteBuyAuction, seller []common.Address) (event.Subscription, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _Crab.contract.WatchLogs(opts, "ExecuteBuyAuction", sellerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrabExecuteBuyAuction)
				if err := _Crab.contract.UnpackLog(event, "ExecuteBuyAuction", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExecuteBuyAuction is a log parse operation binding the contract event 0xc7472bd0a757f40f801e047dd4f4ec901314e95afcbfed844e62af18401e0e6b.
//
// Solidity: event ExecuteBuyAuction(address indexed seller, uint256 wSqueethBought, uint256 ethSold, bool isHedgingOnUniswap)
func (_Crab *CrabFilterer) ParseExecuteBuyAuction(log types.Log) (*CrabExecuteBuyAuction, error) {
	event := new(CrabExecuteBuyAuction)
	if err := _Crab.contract.UnpackLog(event, "ExecuteBuyAuction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrabExecuteSellAuctionIterator is returned from FilterExecuteSellAuction and is used to iterate over the raw logs and unpacked data for ExecuteSellAuction events raised by the Crab contract.
type CrabExecuteSellAuctionIterator struct {
	Event *CrabExecuteSellAuction // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrabExecuteSellAuctionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrabExecuteSellAuction)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrabExecuteSellAuction)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrabExecuteSellAuctionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrabExecuteSellAuctionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrabExecuteSellAuction represents a ExecuteSellAuction event raised by the Crab contract.
type CrabExecuteSellAuction struct {
	Buyer              common.Address
	WSqueethSold       *big.Int
	EthBought          *big.Int
	IsHedgingOnUniswap bool
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterExecuteSellAuction is a free log retrieval operation binding the contract event 0x2af3664d72ebbec5e92c3a487f09a4ecd9ef50177eced03cc2b05892b5e0c915.
//
// Solidity: event ExecuteSellAuction(address indexed buyer, uint256 wSqueethSold, uint256 ethBought, bool isHedgingOnUniswap)
func (_Crab *CrabFilterer) FilterExecuteSellAuction(opts *bind.FilterOpts, buyer []common.Address) (*CrabExecuteSellAuctionIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Crab.contract.FilterLogs(opts, "ExecuteSellAuction", buyerRule)
	if err != nil {
		return nil, err
	}
	return &CrabExecuteSellAuctionIterator{contract: _Crab.contract, event: "ExecuteSellAuction", logs: logs, sub: sub}, nil
}

// WatchExecuteSellAuction is a free log subscription operation binding the contract event 0x2af3664d72ebbec5e92c3a487f09a4ecd9ef50177eced03cc2b05892b5e0c915.
//
// Solidity: event ExecuteSellAuction(address indexed buyer, uint256 wSqueethSold, uint256 ethBought, bool isHedgingOnUniswap)
func (_Crab *CrabFilterer) WatchExecuteSellAuction(opts *bind.WatchOpts, sink chan<- *CrabExecuteSellAuction, buyer []common.Address) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Crab.contract.WatchLogs(opts, "ExecuteSellAuction", buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrabExecuteSellAuction)
				if err := _Crab.contract.UnpackLog(event, "ExecuteSellAuction", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExecuteSellAuction is a log parse operation binding the contract event 0x2af3664d72ebbec5e92c3a487f09a4ecd9ef50177eced03cc2b05892b5e0c915.
//
// Solidity: event ExecuteSellAuction(address indexed buyer, uint256 wSqueethSold, uint256 ethBought, bool isHedgingOnUniswap)
func (_Crab *CrabFilterer) ParseExecuteSellAuction(log types.Log) (*CrabExecuteSellAuction, error) {
	event := new(CrabExecuteSellAuction)
	if err := _Crab.contract.UnpackLog(event, "ExecuteSellAuction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrabFlashDepositIterator is returned from FilterFlashDeposit and is used to iterate over the raw logs and unpacked data for FlashDeposit events raised by the Crab contract.
type CrabFlashDepositIterator struct {
	Event *CrabFlashDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrabFlashDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrabFlashDeposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrabFlashDeposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrabFlashDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrabFlashDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrabFlashDeposit represents a FlashDeposit event raised by the Crab contract.
type CrabFlashDeposit struct {
	Depositor       common.Address
	DepositedAmount *big.Int
	TradedAmountOut *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterFlashDeposit is a free log retrieval operation binding the contract event 0x5d85169ff8329e90f3225f9798e0eba54d00c55d3bbfe201a0d1606febb23a8e.
//
// Solidity: event FlashDeposit(address indexed depositor, uint256 depositedAmount, uint256 tradedAmountOut)
func (_Crab *CrabFilterer) FilterFlashDeposit(opts *bind.FilterOpts, depositor []common.Address) (*CrabFlashDepositIterator, error) {

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _Crab.contract.FilterLogs(opts, "FlashDeposit", depositorRule)
	if err != nil {
		return nil, err
	}
	return &CrabFlashDepositIterator{contract: _Crab.contract, event: "FlashDeposit", logs: logs, sub: sub}, nil
}

// WatchFlashDeposit is a free log subscription operation binding the contract event 0x5d85169ff8329e90f3225f9798e0eba54d00c55d3bbfe201a0d1606febb23a8e.
//
// Solidity: event FlashDeposit(address indexed depositor, uint256 depositedAmount, uint256 tradedAmountOut)
func (_Crab *CrabFilterer) WatchFlashDeposit(opts *bind.WatchOpts, sink chan<- *CrabFlashDeposit, depositor []common.Address) (event.Subscription, error) {

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _Crab.contract.WatchLogs(opts, "FlashDeposit", depositorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrabFlashDeposit)
				if err := _Crab.contract.UnpackLog(event, "FlashDeposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFlashDeposit is a log parse operation binding the contract event 0x5d85169ff8329e90f3225f9798e0eba54d00c55d3bbfe201a0d1606febb23a8e.
//
// Solidity: event FlashDeposit(address indexed depositor, uint256 depositedAmount, uint256 tradedAmountOut)
func (_Crab *CrabFilterer) ParseFlashDeposit(log types.Log) (*CrabFlashDeposit, error) {
	event := new(CrabFlashDeposit)
	if err := _Crab.contract.UnpackLog(event, "FlashDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrabFlashDepositCallbackIterator is returned from FilterFlashDepositCallback and is used to iterate over the raw logs and unpacked data for FlashDepositCallback events raised by the Crab contract.
type CrabFlashDepositCallbackIterator struct {
	Event *CrabFlashDepositCallback // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrabFlashDepositCallbackIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrabFlashDepositCallback)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrabFlashDepositCallback)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrabFlashDepositCallbackIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrabFlashDepositCallbackIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrabFlashDepositCallback represents a FlashDepositCallback event raised by the Crab contract.
type CrabFlashDepositCallback struct {
	Depositor     common.Address
	FlashswapDebt *big.Int
	Excess        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFlashDepositCallback is a free log retrieval operation binding the contract event 0xc355ebece16d7e85e486911f0cde1074bc4bd3fec251c88cdddc7076d3e99adb.
//
// Solidity: event FlashDepositCallback(address indexed depositor, uint256 flashswapDebt, uint256 excess)
func (_Crab *CrabFilterer) FilterFlashDepositCallback(opts *bind.FilterOpts, depositor []common.Address) (*CrabFlashDepositCallbackIterator, error) {

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _Crab.contract.FilterLogs(opts, "FlashDepositCallback", depositorRule)
	if err != nil {
		return nil, err
	}
	return &CrabFlashDepositCallbackIterator{contract: _Crab.contract, event: "FlashDepositCallback", logs: logs, sub: sub}, nil
}

// WatchFlashDepositCallback is a free log subscription operation binding the contract event 0xc355ebece16d7e85e486911f0cde1074bc4bd3fec251c88cdddc7076d3e99adb.
//
// Solidity: event FlashDepositCallback(address indexed depositor, uint256 flashswapDebt, uint256 excess)
func (_Crab *CrabFilterer) WatchFlashDepositCallback(opts *bind.WatchOpts, sink chan<- *CrabFlashDepositCallback, depositor []common.Address) (event.Subscription, error) {

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _Crab.contract.WatchLogs(opts, "FlashDepositCallback", depositorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrabFlashDepositCallback)
				if err := _Crab.contract.UnpackLog(event, "FlashDepositCallback", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFlashDepositCallback is a log parse operation binding the contract event 0xc355ebece16d7e85e486911f0cde1074bc4bd3fec251c88cdddc7076d3e99adb.
//
// Solidity: event FlashDepositCallback(address indexed depositor, uint256 flashswapDebt, uint256 excess)
func (_Crab *CrabFilterer) ParseFlashDepositCallback(log types.Log) (*CrabFlashDepositCallback, error) {
	event := new(CrabFlashDepositCallback)
	if err := _Crab.contract.UnpackLog(event, "FlashDepositCallback", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrabFlashWithdrawIterator is returned from FilterFlashWithdraw and is used to iterate over the raw logs and unpacked data for FlashWithdraw events raised by the Crab contract.
type CrabFlashWithdrawIterator struct {
	Event *CrabFlashWithdraw // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrabFlashWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrabFlashWithdraw)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrabFlashWithdraw)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrabFlashWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrabFlashWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrabFlashWithdraw represents a FlashWithdraw event raised by the Crab contract.
type CrabFlashWithdraw struct {
	Withdrawer     common.Address
	CrabAmount     *big.Int
	WSqueethAmount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterFlashWithdraw is a free log retrieval operation binding the contract event 0xa13b272c1cf13ba724064d3d4809d9f557aab8da2bb582cba031a2f57e728e9d.
//
// Solidity: event FlashWithdraw(address indexed withdrawer, uint256 crabAmount, uint256 wSqueethAmount)
func (_Crab *CrabFilterer) FilterFlashWithdraw(opts *bind.FilterOpts, withdrawer []common.Address) (*CrabFlashWithdrawIterator, error) {

	var withdrawerRule []interface{}
	for _, withdrawerItem := range withdrawer {
		withdrawerRule = append(withdrawerRule, withdrawerItem)
	}

	logs, sub, err := _Crab.contract.FilterLogs(opts, "FlashWithdraw", withdrawerRule)
	if err != nil {
		return nil, err
	}
	return &CrabFlashWithdrawIterator{contract: _Crab.contract, event: "FlashWithdraw", logs: logs, sub: sub}, nil
}

// WatchFlashWithdraw is a free log subscription operation binding the contract event 0xa13b272c1cf13ba724064d3d4809d9f557aab8da2bb582cba031a2f57e728e9d.
//
// Solidity: event FlashWithdraw(address indexed withdrawer, uint256 crabAmount, uint256 wSqueethAmount)
func (_Crab *CrabFilterer) WatchFlashWithdraw(opts *bind.WatchOpts, sink chan<- *CrabFlashWithdraw, withdrawer []common.Address) (event.Subscription, error) {

	var withdrawerRule []interface{}
	for _, withdrawerItem := range withdrawer {
		withdrawerRule = append(withdrawerRule, withdrawerItem)
	}

	logs, sub, err := _Crab.contract.WatchLogs(opts, "FlashWithdraw", withdrawerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrabFlashWithdraw)
				if err := _Crab.contract.UnpackLog(event, "FlashWithdraw", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFlashWithdraw is a log parse operation binding the contract event 0xa13b272c1cf13ba724064d3d4809d9f557aab8da2bb582cba031a2f57e728e9d.
//
// Solidity: event FlashWithdraw(address indexed withdrawer, uint256 crabAmount, uint256 wSqueethAmount)
func (_Crab *CrabFilterer) ParseFlashWithdraw(log types.Log) (*CrabFlashWithdraw, error) {
	event := new(CrabFlashWithdraw)
	if err := _Crab.contract.UnpackLog(event, "FlashWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrabFlashWithdrawCallbackIterator is returned from FilterFlashWithdrawCallback and is used to iterate over the raw logs and unpacked data for FlashWithdrawCallback events raised by the Crab contract.
type CrabFlashWithdrawCallbackIterator struct {
	Event *CrabFlashWithdrawCallback // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrabFlashWithdrawCallbackIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrabFlashWithdrawCallback)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrabFlashWithdrawCallback)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrabFlashWithdrawCallbackIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrabFlashWithdrawCallbackIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrabFlashWithdrawCallback represents a FlashWithdrawCallback event raised by the Crab contract.
type CrabFlashWithdrawCallback struct {
	Withdrawer    common.Address
	FlashswapDebt *big.Int
	Excess        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFlashWithdrawCallback is a free log retrieval operation binding the contract event 0x6f3269a64126ef2a1959892f3d921e81865181e09a7f72f55d3a49550c53b48d.
//
// Solidity: event FlashWithdrawCallback(address indexed withdrawer, uint256 flashswapDebt, uint256 excess)
func (_Crab *CrabFilterer) FilterFlashWithdrawCallback(opts *bind.FilterOpts, withdrawer []common.Address) (*CrabFlashWithdrawCallbackIterator, error) {

	var withdrawerRule []interface{}
	for _, withdrawerItem := range withdrawer {
		withdrawerRule = append(withdrawerRule, withdrawerItem)
	}

	logs, sub, err := _Crab.contract.FilterLogs(opts, "FlashWithdrawCallback", withdrawerRule)
	if err != nil {
		return nil, err
	}
	return &CrabFlashWithdrawCallbackIterator{contract: _Crab.contract, event: "FlashWithdrawCallback", logs: logs, sub: sub}, nil
}

// WatchFlashWithdrawCallback is a free log subscription operation binding the contract event 0x6f3269a64126ef2a1959892f3d921e81865181e09a7f72f55d3a49550c53b48d.
//
// Solidity: event FlashWithdrawCallback(address indexed withdrawer, uint256 flashswapDebt, uint256 excess)
func (_Crab *CrabFilterer) WatchFlashWithdrawCallback(opts *bind.WatchOpts, sink chan<- *CrabFlashWithdrawCallback, withdrawer []common.Address) (event.Subscription, error) {

	var withdrawerRule []interface{}
	for _, withdrawerItem := range withdrawer {
		withdrawerRule = append(withdrawerRule, withdrawerItem)
	}

	logs, sub, err := _Crab.contract.WatchLogs(opts, "FlashWithdrawCallback", withdrawerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrabFlashWithdrawCallback)
				if err := _Crab.contract.UnpackLog(event, "FlashWithdrawCallback", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFlashWithdrawCallback is a log parse operation binding the contract event 0x6f3269a64126ef2a1959892f3d921e81865181e09a7f72f55d3a49550c53b48d.
//
// Solidity: event FlashWithdrawCallback(address indexed withdrawer, uint256 flashswapDebt, uint256 excess)
func (_Crab *CrabFilterer) ParseFlashWithdrawCallback(log types.Log) (*CrabFlashWithdrawCallback, error) {
	event := new(CrabFlashWithdrawCallback)
	if err := _Crab.contract.UnpackLog(event, "FlashWithdrawCallback", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrabHedgeIterator is returned from FilterHedge and is used to iterate over the raw logs and unpacked data for Hedge events raised by the Crab contract.
type CrabHedgeIterator struct {
	Event *CrabHedge // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrabHedgeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrabHedge)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrabHedge)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrabHedgeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrabHedgeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrabHedge represents a Hedge event raised by the Crab contract.
type CrabHedge struct {
	Hedger                    common.Address
	AuctionType               bool
	HedgerPrice               *big.Int
	AuctionPrice              *big.Int
	WSqueethHedgeTargetAmount *big.Int
	EthHedgetargetAmount      *big.Int
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterHedge is a free log retrieval operation binding the contract event 0x878fd3ca52ad322c7535f559ee7c91afc67363073783360ef1b1420589dc6174.
//
// Solidity: event Hedge(address indexed hedger, bool auctionType, uint256 hedgerPrice, uint256 auctionPrice, uint256 wSqueethHedgeTargetAmount, uint256 ethHedgetargetAmount)
func (_Crab *CrabFilterer) FilterHedge(opts *bind.FilterOpts, hedger []common.Address) (*CrabHedgeIterator, error) {

	var hedgerRule []interface{}
	for _, hedgerItem := range hedger {
		hedgerRule = append(hedgerRule, hedgerItem)
	}

	logs, sub, err := _Crab.contract.FilterLogs(opts, "Hedge", hedgerRule)
	if err != nil {
		return nil, err
	}
	return &CrabHedgeIterator{contract: _Crab.contract, event: "Hedge", logs: logs, sub: sub}, nil
}

// WatchHedge is a free log subscription operation binding the contract event 0x878fd3ca52ad322c7535f559ee7c91afc67363073783360ef1b1420589dc6174.
//
// Solidity: event Hedge(address indexed hedger, bool auctionType, uint256 hedgerPrice, uint256 auctionPrice, uint256 wSqueethHedgeTargetAmount, uint256 ethHedgetargetAmount)
func (_Crab *CrabFilterer) WatchHedge(opts *bind.WatchOpts, sink chan<- *CrabHedge, hedger []common.Address) (event.Subscription, error) {

	var hedgerRule []interface{}
	for _, hedgerItem := range hedger {
		hedgerRule = append(hedgerRule, hedgerItem)
	}

	logs, sub, err := _Crab.contract.WatchLogs(opts, "Hedge", hedgerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrabHedge)
				if err := _Crab.contract.UnpackLog(event, "Hedge", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseHedge is a log parse operation binding the contract event 0x878fd3ca52ad322c7535f559ee7c91afc67363073783360ef1b1420589dc6174.
//
// Solidity: event Hedge(address indexed hedger, bool auctionType, uint256 hedgerPrice, uint256 auctionPrice, uint256 wSqueethHedgeTargetAmount, uint256 ethHedgetargetAmount)
func (_Crab *CrabFilterer) ParseHedge(log types.Log) (*CrabHedge, error) {
	event := new(CrabHedge)
	if err := _Crab.contract.UnpackLog(event, "Hedge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrabHedgeOnUniswapIterator is returned from FilterHedgeOnUniswap and is used to iterate over the raw logs and unpacked data for HedgeOnUniswap events raised by the Crab contract.
type CrabHedgeOnUniswapIterator struct {
	Event *CrabHedgeOnUniswap // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrabHedgeOnUniswapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrabHedgeOnUniswap)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrabHedgeOnUniswap)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrabHedgeOnUniswapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrabHedgeOnUniswapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrabHedgeOnUniswap represents a HedgeOnUniswap event raised by the Crab contract.
type CrabHedgeOnUniswap struct {
	Hedger                    common.Address
	AuctionType               bool
	AuctionPrice              *big.Int
	WSqueethHedgeTargetAmount *big.Int
	EthHedgetargetAmount      *big.Int
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterHedgeOnUniswap is a free log retrieval operation binding the contract event 0x4c1a959210172325f5c6678421c3834b04ae8ce57f7a7c0c0bbfbb62bca37e34.
//
// Solidity: event HedgeOnUniswap(address indexed hedger, bool auctionType, uint256 auctionPrice, uint256 wSqueethHedgeTargetAmount, uint256 ethHedgetargetAmount)
func (_Crab *CrabFilterer) FilterHedgeOnUniswap(opts *bind.FilterOpts, hedger []common.Address) (*CrabHedgeOnUniswapIterator, error) {

	var hedgerRule []interface{}
	for _, hedgerItem := range hedger {
		hedgerRule = append(hedgerRule, hedgerItem)
	}

	logs, sub, err := _Crab.contract.FilterLogs(opts, "HedgeOnUniswap", hedgerRule)
	if err != nil {
		return nil, err
	}
	return &CrabHedgeOnUniswapIterator{contract: _Crab.contract, event: "HedgeOnUniswap", logs: logs, sub: sub}, nil
}

// WatchHedgeOnUniswap is a free log subscription operation binding the contract event 0x4c1a959210172325f5c6678421c3834b04ae8ce57f7a7c0c0bbfbb62bca37e34.
//
// Solidity: event HedgeOnUniswap(address indexed hedger, bool auctionType, uint256 auctionPrice, uint256 wSqueethHedgeTargetAmount, uint256 ethHedgetargetAmount)
func (_Crab *CrabFilterer) WatchHedgeOnUniswap(opts *bind.WatchOpts, sink chan<- *CrabHedgeOnUniswap, hedger []common.Address) (event.Subscription, error) {

	var hedgerRule []interface{}
	for _, hedgerItem := range hedger {
		hedgerRule = append(hedgerRule, hedgerItem)
	}

	logs, sub, err := _Crab.contract.WatchLogs(opts, "HedgeOnUniswap", hedgerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrabHedgeOnUniswap)
				if err := _Crab.contract.UnpackLog(event, "HedgeOnUniswap", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseHedgeOnUniswap is a log parse operation binding the contract event 0x4c1a959210172325f5c6678421c3834b04ae8ce57f7a7c0c0bbfbb62bca37e34.
//
// Solidity: event HedgeOnUniswap(address indexed hedger, bool auctionType, uint256 auctionPrice, uint256 wSqueethHedgeTargetAmount, uint256 ethHedgetargetAmount)
func (_Crab *CrabFilterer) ParseHedgeOnUniswap(log types.Log) (*CrabHedgeOnUniswap, error) {
	event := new(CrabHedgeOnUniswap)
	if err := _Crab.contract.UnpackLog(event, "HedgeOnUniswap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrabOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Crab contract.
type CrabOwnershipTransferredIterator struct {
	Event *CrabOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrabOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrabOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrabOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrabOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrabOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrabOwnershipTransferred represents a OwnershipTransferred event raised by the Crab contract.
type CrabOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Crab *CrabFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CrabOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Crab.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CrabOwnershipTransferredIterator{contract: _Crab.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Crab *CrabFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CrabOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Crab.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrabOwnershipTransferred)
				if err := _Crab.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Crab *CrabFilterer) ParseOwnershipTransferred(log types.Log) (*CrabOwnershipTransferred, error) {
	event := new(CrabOwnershipTransferred)
	if err := _Crab.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrabPriceHedgeIterator is returned from FilterPriceHedge and is used to iterate over the raw logs and unpacked data for PriceHedge events raised by the Crab contract.
type CrabPriceHedgeIterator struct {
	Event *CrabPriceHedge // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrabPriceHedgeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrabPriceHedge)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrabPriceHedge)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrabPriceHedgeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrabPriceHedgeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrabPriceHedge represents a PriceHedge event raised by the Crab contract.
type CrabPriceHedge struct {
	Hedger                  common.Address
	AuctionType             bool
	HedgerPrice             *big.Int
	AuctionTriggerTimestamp *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterPriceHedge is a free log retrieval operation binding the contract event 0xa6806e5672ec1827d23bd4a25f4a41c8d0b055720c40081e3803f6b0a957f797.
//
// Solidity: event PriceHedge(address indexed hedger, bool auctionType, uint256 hedgerPrice, uint256 auctionTriggerTimestamp)
func (_Crab *CrabFilterer) FilterPriceHedge(opts *bind.FilterOpts, hedger []common.Address) (*CrabPriceHedgeIterator, error) {

	var hedgerRule []interface{}
	for _, hedgerItem := range hedger {
		hedgerRule = append(hedgerRule, hedgerItem)
	}

	logs, sub, err := _Crab.contract.FilterLogs(opts, "PriceHedge", hedgerRule)
	if err != nil {
		return nil, err
	}
	return &CrabPriceHedgeIterator{contract: _Crab.contract, event: "PriceHedge", logs: logs, sub: sub}, nil
}

// WatchPriceHedge is a free log subscription operation binding the contract event 0xa6806e5672ec1827d23bd4a25f4a41c8d0b055720c40081e3803f6b0a957f797.
//
// Solidity: event PriceHedge(address indexed hedger, bool auctionType, uint256 hedgerPrice, uint256 auctionTriggerTimestamp)
func (_Crab *CrabFilterer) WatchPriceHedge(opts *bind.WatchOpts, sink chan<- *CrabPriceHedge, hedger []common.Address) (event.Subscription, error) {

	var hedgerRule []interface{}
	for _, hedgerItem := range hedger {
		hedgerRule = append(hedgerRule, hedgerItem)
	}

	logs, sub, err := _Crab.contract.WatchLogs(opts, "PriceHedge", hedgerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrabPriceHedge)
				if err := _Crab.contract.UnpackLog(event, "PriceHedge", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePriceHedge is a log parse operation binding the contract event 0xa6806e5672ec1827d23bd4a25f4a41c8d0b055720c40081e3803f6b0a957f797.
//
// Solidity: event PriceHedge(address indexed hedger, bool auctionType, uint256 hedgerPrice, uint256 auctionTriggerTimestamp)
func (_Crab *CrabFilterer) ParsePriceHedge(log types.Log) (*CrabPriceHedge, error) {
	event := new(CrabPriceHedge)
	if err := _Crab.contract.UnpackLog(event, "PriceHedge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrabPriceHedgeOnUniswapIterator is returned from FilterPriceHedgeOnUniswap and is used to iterate over the raw logs and unpacked data for PriceHedgeOnUniswap events raised by the Crab contract.
type CrabPriceHedgeOnUniswapIterator struct {
	Event *CrabPriceHedgeOnUniswap // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrabPriceHedgeOnUniswapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrabPriceHedgeOnUniswap)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrabPriceHedgeOnUniswap)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrabPriceHedgeOnUniswapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrabPriceHedgeOnUniswapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrabPriceHedgeOnUniswap represents a PriceHedgeOnUniswap event raised by the Crab contract.
type CrabPriceHedgeOnUniswap struct {
	Hedger                  common.Address
	HedgeTimestamp          *big.Int
	AuctionTriggerTimestamp *big.Int
	MinWSqueeth             *big.Int
	MinEth                  *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterPriceHedgeOnUniswap is a free log retrieval operation binding the contract event 0xf99cce6ee3154fddfc55959dd136665c7351fe51c0ea58392ca454be8180dcd4.
//
// Solidity: event PriceHedgeOnUniswap(address indexed hedger, uint256 hedgeTimestamp, uint256 auctionTriggerTimestamp, uint256 minWSqueeth, uint256 minEth)
func (_Crab *CrabFilterer) FilterPriceHedgeOnUniswap(opts *bind.FilterOpts, hedger []common.Address) (*CrabPriceHedgeOnUniswapIterator, error) {

	var hedgerRule []interface{}
	for _, hedgerItem := range hedger {
		hedgerRule = append(hedgerRule, hedgerItem)
	}

	logs, sub, err := _Crab.contract.FilterLogs(opts, "PriceHedgeOnUniswap", hedgerRule)
	if err != nil {
		return nil, err
	}
	return &CrabPriceHedgeOnUniswapIterator{contract: _Crab.contract, event: "PriceHedgeOnUniswap", logs: logs, sub: sub}, nil
}

// WatchPriceHedgeOnUniswap is a free log subscription operation binding the contract event 0xf99cce6ee3154fddfc55959dd136665c7351fe51c0ea58392ca454be8180dcd4.
//
// Solidity: event PriceHedgeOnUniswap(address indexed hedger, uint256 hedgeTimestamp, uint256 auctionTriggerTimestamp, uint256 minWSqueeth, uint256 minEth)
func (_Crab *CrabFilterer) WatchPriceHedgeOnUniswap(opts *bind.WatchOpts, sink chan<- *CrabPriceHedgeOnUniswap, hedger []common.Address) (event.Subscription, error) {

	var hedgerRule []interface{}
	for _, hedgerItem := range hedger {
		hedgerRule = append(hedgerRule, hedgerItem)
	}

	logs, sub, err := _Crab.contract.WatchLogs(opts, "PriceHedgeOnUniswap", hedgerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrabPriceHedgeOnUniswap)
				if err := _Crab.contract.UnpackLog(event, "PriceHedgeOnUniswap", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePriceHedgeOnUniswap is a log parse operation binding the contract event 0xf99cce6ee3154fddfc55959dd136665c7351fe51c0ea58392ca454be8180dcd4.
//
// Solidity: event PriceHedgeOnUniswap(address indexed hedger, uint256 hedgeTimestamp, uint256 auctionTriggerTimestamp, uint256 minWSqueeth, uint256 minEth)
func (_Crab *CrabFilterer) ParsePriceHedgeOnUniswap(log types.Log) (*CrabPriceHedgeOnUniswap, error) {
	event := new(CrabPriceHedgeOnUniswap)
	if err := _Crab.contract.UnpackLog(event, "PriceHedgeOnUniswap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrabSetAuctionTimeIterator is returned from FilterSetAuctionTime and is used to iterate over the raw logs and unpacked data for SetAuctionTime events raised by the Crab contract.
type CrabSetAuctionTimeIterator struct {
	Event *CrabSetAuctionTime // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrabSetAuctionTimeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrabSetAuctionTime)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrabSetAuctionTime)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrabSetAuctionTimeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrabSetAuctionTimeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrabSetAuctionTime represents a SetAuctionTime event raised by the Crab contract.
type CrabSetAuctionTime struct {
	NewAuctionTime *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSetAuctionTime is a free log retrieval operation binding the contract event 0x6405fa008c5cc5710b13a509f31e7596708bdbc2a52c85a15f9992697a791b20.
//
// Solidity: event SetAuctionTime(uint256 newAuctionTime)
func (_Crab *CrabFilterer) FilterSetAuctionTime(opts *bind.FilterOpts) (*CrabSetAuctionTimeIterator, error) {

	logs, sub, err := _Crab.contract.FilterLogs(opts, "SetAuctionTime")
	if err != nil {
		return nil, err
	}
	return &CrabSetAuctionTimeIterator{contract: _Crab.contract, event: "SetAuctionTime", logs: logs, sub: sub}, nil
}

// WatchSetAuctionTime is a free log subscription operation binding the contract event 0x6405fa008c5cc5710b13a509f31e7596708bdbc2a52c85a15f9992697a791b20.
//
// Solidity: event SetAuctionTime(uint256 newAuctionTime)
func (_Crab *CrabFilterer) WatchSetAuctionTime(opts *bind.WatchOpts, sink chan<- *CrabSetAuctionTime) (event.Subscription, error) {

	logs, sub, err := _Crab.contract.WatchLogs(opts, "SetAuctionTime")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrabSetAuctionTime)
				if err := _Crab.contract.UnpackLog(event, "SetAuctionTime", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetAuctionTime is a log parse operation binding the contract event 0x6405fa008c5cc5710b13a509f31e7596708bdbc2a52c85a15f9992697a791b20.
//
// Solidity: event SetAuctionTime(uint256 newAuctionTime)
func (_Crab *CrabFilterer) ParseSetAuctionTime(log types.Log) (*CrabSetAuctionTime, error) {
	event := new(CrabSetAuctionTime)
	if err := _Crab.contract.UnpackLog(event, "SetAuctionTime", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrabSetDeltaHedgeThresholdIterator is returned from FilterSetDeltaHedgeThreshold and is used to iterate over the raw logs and unpacked data for SetDeltaHedgeThreshold events raised by the Crab contract.
type CrabSetDeltaHedgeThresholdIterator struct {
	Event *CrabSetDeltaHedgeThreshold // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrabSetDeltaHedgeThresholdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrabSetDeltaHedgeThreshold)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrabSetDeltaHedgeThreshold)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrabSetDeltaHedgeThresholdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrabSetDeltaHedgeThresholdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrabSetDeltaHedgeThreshold represents a SetDeltaHedgeThreshold event raised by the Crab contract.
type CrabSetDeltaHedgeThreshold struct {
	NewDeltaHedgeThreshold *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterSetDeltaHedgeThreshold is a free log retrieval operation binding the contract event 0x48100eb8aecbbf59b3665ff2a7b2d7257a218196dec79a67ca870fca43cdff62.
//
// Solidity: event SetDeltaHedgeThreshold(uint256 newDeltaHedgeThreshold)
func (_Crab *CrabFilterer) FilterSetDeltaHedgeThreshold(opts *bind.FilterOpts) (*CrabSetDeltaHedgeThresholdIterator, error) {

	logs, sub, err := _Crab.contract.FilterLogs(opts, "SetDeltaHedgeThreshold")
	if err != nil {
		return nil, err
	}
	return &CrabSetDeltaHedgeThresholdIterator{contract: _Crab.contract, event: "SetDeltaHedgeThreshold", logs: logs, sub: sub}, nil
}

// WatchSetDeltaHedgeThreshold is a free log subscription operation binding the contract event 0x48100eb8aecbbf59b3665ff2a7b2d7257a218196dec79a67ca870fca43cdff62.
//
// Solidity: event SetDeltaHedgeThreshold(uint256 newDeltaHedgeThreshold)
func (_Crab *CrabFilterer) WatchSetDeltaHedgeThreshold(opts *bind.WatchOpts, sink chan<- *CrabSetDeltaHedgeThreshold) (event.Subscription, error) {

	logs, sub, err := _Crab.contract.WatchLogs(opts, "SetDeltaHedgeThreshold")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrabSetDeltaHedgeThreshold)
				if err := _Crab.contract.UnpackLog(event, "SetDeltaHedgeThreshold", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetDeltaHedgeThreshold is a log parse operation binding the contract event 0x48100eb8aecbbf59b3665ff2a7b2d7257a218196dec79a67ca870fca43cdff62.
//
// Solidity: event SetDeltaHedgeThreshold(uint256 newDeltaHedgeThreshold)
func (_Crab *CrabFilterer) ParseSetDeltaHedgeThreshold(log types.Log) (*CrabSetDeltaHedgeThreshold, error) {
	event := new(CrabSetDeltaHedgeThreshold)
	if err := _Crab.contract.UnpackLog(event, "SetDeltaHedgeThreshold", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrabSetHedgePriceThresholdIterator is returned from FilterSetHedgePriceThreshold and is used to iterate over the raw logs and unpacked data for SetHedgePriceThreshold events raised by the Crab contract.
type CrabSetHedgePriceThresholdIterator struct {
	Event *CrabSetHedgePriceThreshold // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrabSetHedgePriceThresholdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrabSetHedgePriceThreshold)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrabSetHedgePriceThreshold)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrabSetHedgePriceThresholdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrabSetHedgePriceThresholdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrabSetHedgePriceThreshold represents a SetHedgePriceThreshold event raised by the Crab contract.
type CrabSetHedgePriceThreshold struct {
	NewHedgePriceThreshold *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterSetHedgePriceThreshold is a free log retrieval operation binding the contract event 0x789e4b8ad1c375952cea7f07c9b3b6619a84b406432b948246cecb8ced96b9fa.
//
// Solidity: event SetHedgePriceThreshold(uint256 newHedgePriceThreshold)
func (_Crab *CrabFilterer) FilterSetHedgePriceThreshold(opts *bind.FilterOpts) (*CrabSetHedgePriceThresholdIterator, error) {

	logs, sub, err := _Crab.contract.FilterLogs(opts, "SetHedgePriceThreshold")
	if err != nil {
		return nil, err
	}
	return &CrabSetHedgePriceThresholdIterator{contract: _Crab.contract, event: "SetHedgePriceThreshold", logs: logs, sub: sub}, nil
}

// WatchSetHedgePriceThreshold is a free log subscription operation binding the contract event 0x789e4b8ad1c375952cea7f07c9b3b6619a84b406432b948246cecb8ced96b9fa.
//
// Solidity: event SetHedgePriceThreshold(uint256 newHedgePriceThreshold)
func (_Crab *CrabFilterer) WatchSetHedgePriceThreshold(opts *bind.WatchOpts, sink chan<- *CrabSetHedgePriceThreshold) (event.Subscription, error) {

	logs, sub, err := _Crab.contract.WatchLogs(opts, "SetHedgePriceThreshold")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrabSetHedgePriceThreshold)
				if err := _Crab.contract.UnpackLog(event, "SetHedgePriceThreshold", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetHedgePriceThreshold is a log parse operation binding the contract event 0x789e4b8ad1c375952cea7f07c9b3b6619a84b406432b948246cecb8ced96b9fa.
//
// Solidity: event SetHedgePriceThreshold(uint256 newHedgePriceThreshold)
func (_Crab *CrabFilterer) ParseSetHedgePriceThreshold(log types.Log) (*CrabSetHedgePriceThreshold, error) {
	event := new(CrabSetHedgePriceThreshold)
	if err := _Crab.contract.UnpackLog(event, "SetHedgePriceThreshold", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrabSetHedgeTimeThresholdIterator is returned from FilterSetHedgeTimeThreshold and is used to iterate over the raw logs and unpacked data for SetHedgeTimeThreshold events raised by the Crab contract.
type CrabSetHedgeTimeThresholdIterator struct {
	Event *CrabSetHedgeTimeThreshold // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrabSetHedgeTimeThresholdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrabSetHedgeTimeThreshold)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrabSetHedgeTimeThreshold)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrabSetHedgeTimeThresholdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrabSetHedgeTimeThresholdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrabSetHedgeTimeThreshold represents a SetHedgeTimeThreshold event raised by the Crab contract.
type CrabSetHedgeTimeThreshold struct {
	NewHedgeTimeThreshold *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterSetHedgeTimeThreshold is a free log retrieval operation binding the contract event 0x28e0e4ee0b14d4b056ce88e1bcd890ccd32b22e213723c8765901381b5eae705.
//
// Solidity: event SetHedgeTimeThreshold(uint256 newHedgeTimeThreshold)
func (_Crab *CrabFilterer) FilterSetHedgeTimeThreshold(opts *bind.FilterOpts) (*CrabSetHedgeTimeThresholdIterator, error) {

	logs, sub, err := _Crab.contract.FilterLogs(opts, "SetHedgeTimeThreshold")
	if err != nil {
		return nil, err
	}
	return &CrabSetHedgeTimeThresholdIterator{contract: _Crab.contract, event: "SetHedgeTimeThreshold", logs: logs, sub: sub}, nil
}

// WatchSetHedgeTimeThreshold is a free log subscription operation binding the contract event 0x28e0e4ee0b14d4b056ce88e1bcd890ccd32b22e213723c8765901381b5eae705.
//
// Solidity: event SetHedgeTimeThreshold(uint256 newHedgeTimeThreshold)
func (_Crab *CrabFilterer) WatchSetHedgeTimeThreshold(opts *bind.WatchOpts, sink chan<- *CrabSetHedgeTimeThreshold) (event.Subscription, error) {

	logs, sub, err := _Crab.contract.WatchLogs(opts, "SetHedgeTimeThreshold")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrabSetHedgeTimeThreshold)
				if err := _Crab.contract.UnpackLog(event, "SetHedgeTimeThreshold", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetHedgeTimeThreshold is a log parse operation binding the contract event 0x28e0e4ee0b14d4b056ce88e1bcd890ccd32b22e213723c8765901381b5eae705.
//
// Solidity: event SetHedgeTimeThreshold(uint256 newHedgeTimeThreshold)
func (_Crab *CrabFilterer) ParseSetHedgeTimeThreshold(log types.Log) (*CrabSetHedgeTimeThreshold, error) {
	event := new(CrabSetHedgeTimeThreshold)
	if err := _Crab.contract.UnpackLog(event, "SetHedgeTimeThreshold", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrabSetHedgingTwapPeriodIterator is returned from FilterSetHedgingTwapPeriod and is used to iterate over the raw logs and unpacked data for SetHedgingTwapPeriod events raised by the Crab contract.
type CrabSetHedgingTwapPeriodIterator struct {
	Event *CrabSetHedgingTwapPeriod // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrabSetHedgingTwapPeriodIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrabSetHedgingTwapPeriod)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrabSetHedgingTwapPeriod)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrabSetHedgingTwapPeriodIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrabSetHedgingTwapPeriodIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrabSetHedgingTwapPeriod represents a SetHedgingTwapPeriod event raised by the Crab contract.
type CrabSetHedgingTwapPeriod struct {
	NewHedgingTwapPeriod uint32
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterSetHedgingTwapPeriod is a free log retrieval operation binding the contract event 0x1cd9c7f99a5530a38c8f2b387dcc78e8a76cb5b203e0c4316a66777d993dee35.
//
// Solidity: event SetHedgingTwapPeriod(uint32 newHedgingTwapPeriod)
func (_Crab *CrabFilterer) FilterSetHedgingTwapPeriod(opts *bind.FilterOpts) (*CrabSetHedgingTwapPeriodIterator, error) {

	logs, sub, err := _Crab.contract.FilterLogs(opts, "SetHedgingTwapPeriod")
	if err != nil {
		return nil, err
	}
	return &CrabSetHedgingTwapPeriodIterator{contract: _Crab.contract, event: "SetHedgingTwapPeriod", logs: logs, sub: sub}, nil
}

// WatchSetHedgingTwapPeriod is a free log subscription operation binding the contract event 0x1cd9c7f99a5530a38c8f2b387dcc78e8a76cb5b203e0c4316a66777d993dee35.
//
// Solidity: event SetHedgingTwapPeriod(uint32 newHedgingTwapPeriod)
func (_Crab *CrabFilterer) WatchSetHedgingTwapPeriod(opts *bind.WatchOpts, sink chan<- *CrabSetHedgingTwapPeriod) (event.Subscription, error) {

	logs, sub, err := _Crab.contract.WatchLogs(opts, "SetHedgingTwapPeriod")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrabSetHedgingTwapPeriod)
				if err := _Crab.contract.UnpackLog(event, "SetHedgingTwapPeriod", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetHedgingTwapPeriod is a log parse operation binding the contract event 0x1cd9c7f99a5530a38c8f2b387dcc78e8a76cb5b203e0c4316a66777d993dee35.
//
// Solidity: event SetHedgingTwapPeriod(uint32 newHedgingTwapPeriod)
func (_Crab *CrabFilterer) ParseSetHedgingTwapPeriod(log types.Log) (*CrabSetHedgingTwapPeriod, error) {
	event := new(CrabSetHedgingTwapPeriod)
	if err := _Crab.contract.UnpackLog(event, "SetHedgingTwapPeriod", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrabSetMaxPriceMultiplierIterator is returned from FilterSetMaxPriceMultiplier and is used to iterate over the raw logs and unpacked data for SetMaxPriceMultiplier events raised by the Crab contract.
type CrabSetMaxPriceMultiplierIterator struct {
	Event *CrabSetMaxPriceMultiplier // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrabSetMaxPriceMultiplierIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrabSetMaxPriceMultiplier)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrabSetMaxPriceMultiplier)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrabSetMaxPriceMultiplierIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrabSetMaxPriceMultiplierIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrabSetMaxPriceMultiplier represents a SetMaxPriceMultiplier event raised by the Crab contract.
type CrabSetMaxPriceMultiplier struct {
	NewMaxPriceMultiplier *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterSetMaxPriceMultiplier is a free log retrieval operation binding the contract event 0xce6ae334d464afcb4f9f5f5218c4e3b0d5d227a821cb8a76bbff56da8f5f0b88.
//
// Solidity: event SetMaxPriceMultiplier(uint256 newMaxPriceMultiplier)
func (_Crab *CrabFilterer) FilterSetMaxPriceMultiplier(opts *bind.FilterOpts) (*CrabSetMaxPriceMultiplierIterator, error) {

	logs, sub, err := _Crab.contract.FilterLogs(opts, "SetMaxPriceMultiplier")
	if err != nil {
		return nil, err
	}
	return &CrabSetMaxPriceMultiplierIterator{contract: _Crab.contract, event: "SetMaxPriceMultiplier", logs: logs, sub: sub}, nil
}

// WatchSetMaxPriceMultiplier is a free log subscription operation binding the contract event 0xce6ae334d464afcb4f9f5f5218c4e3b0d5d227a821cb8a76bbff56da8f5f0b88.
//
// Solidity: event SetMaxPriceMultiplier(uint256 newMaxPriceMultiplier)
func (_Crab *CrabFilterer) WatchSetMaxPriceMultiplier(opts *bind.WatchOpts, sink chan<- *CrabSetMaxPriceMultiplier) (event.Subscription, error) {

	logs, sub, err := _Crab.contract.WatchLogs(opts, "SetMaxPriceMultiplier")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrabSetMaxPriceMultiplier)
				if err := _Crab.contract.UnpackLog(event, "SetMaxPriceMultiplier", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetMaxPriceMultiplier is a log parse operation binding the contract event 0xce6ae334d464afcb4f9f5f5218c4e3b0d5d227a821cb8a76bbff56da8f5f0b88.
//
// Solidity: event SetMaxPriceMultiplier(uint256 newMaxPriceMultiplier)
func (_Crab *CrabFilterer) ParseSetMaxPriceMultiplier(log types.Log) (*CrabSetMaxPriceMultiplier, error) {
	event := new(CrabSetMaxPriceMultiplier)
	if err := _Crab.contract.UnpackLog(event, "SetMaxPriceMultiplier", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrabSetMinPriceMultiplierIterator is returned from FilterSetMinPriceMultiplier and is used to iterate over the raw logs and unpacked data for SetMinPriceMultiplier events raised by the Crab contract.
type CrabSetMinPriceMultiplierIterator struct {
	Event *CrabSetMinPriceMultiplier // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrabSetMinPriceMultiplierIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrabSetMinPriceMultiplier)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrabSetMinPriceMultiplier)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrabSetMinPriceMultiplierIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrabSetMinPriceMultiplierIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrabSetMinPriceMultiplier represents a SetMinPriceMultiplier event raised by the Crab contract.
type CrabSetMinPriceMultiplier struct {
	NewMinPriceMultiplier *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterSetMinPriceMultiplier is a free log retrieval operation binding the contract event 0x07515fe9cd564f272668faa064c14bc4cec5f75710703ff96a21a47b680ce28f.
//
// Solidity: event SetMinPriceMultiplier(uint256 newMinPriceMultiplier)
func (_Crab *CrabFilterer) FilterSetMinPriceMultiplier(opts *bind.FilterOpts) (*CrabSetMinPriceMultiplierIterator, error) {

	logs, sub, err := _Crab.contract.FilterLogs(opts, "SetMinPriceMultiplier")
	if err != nil {
		return nil, err
	}
	return &CrabSetMinPriceMultiplierIterator{contract: _Crab.contract, event: "SetMinPriceMultiplier", logs: logs, sub: sub}, nil
}

// WatchSetMinPriceMultiplier is a free log subscription operation binding the contract event 0x07515fe9cd564f272668faa064c14bc4cec5f75710703ff96a21a47b680ce28f.
//
// Solidity: event SetMinPriceMultiplier(uint256 newMinPriceMultiplier)
func (_Crab *CrabFilterer) WatchSetMinPriceMultiplier(opts *bind.WatchOpts, sink chan<- *CrabSetMinPriceMultiplier) (event.Subscription, error) {

	logs, sub, err := _Crab.contract.WatchLogs(opts, "SetMinPriceMultiplier")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrabSetMinPriceMultiplier)
				if err := _Crab.contract.UnpackLog(event, "SetMinPriceMultiplier", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetMinPriceMultiplier is a log parse operation binding the contract event 0x07515fe9cd564f272668faa064c14bc4cec5f75710703ff96a21a47b680ce28f.
//
// Solidity: event SetMinPriceMultiplier(uint256 newMinPriceMultiplier)
func (_Crab *CrabFilterer) ParseSetMinPriceMultiplier(log types.Log) (*CrabSetMinPriceMultiplier, error) {
	event := new(CrabSetMinPriceMultiplier)
	if err := _Crab.contract.UnpackLog(event, "SetMinPriceMultiplier", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrabSetStrategyCapIterator is returned from FilterSetStrategyCap and is used to iterate over the raw logs and unpacked data for SetStrategyCap events raised by the Crab contract.
type CrabSetStrategyCapIterator struct {
	Event *CrabSetStrategyCap // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrabSetStrategyCapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrabSetStrategyCap)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrabSetStrategyCap)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrabSetStrategyCapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrabSetStrategyCapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrabSetStrategyCap represents a SetStrategyCap event raised by the Crab contract.
type CrabSetStrategyCap struct {
	NewCapAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSetStrategyCap is a free log retrieval operation binding the contract event 0x29600e2e028c8c5c2b112d021938e0d0237d8fafcbb20394c56cf9fa4661ca27.
//
// Solidity: event SetStrategyCap(uint256 newCapAmount)
func (_Crab *CrabFilterer) FilterSetStrategyCap(opts *bind.FilterOpts) (*CrabSetStrategyCapIterator, error) {

	logs, sub, err := _Crab.contract.FilterLogs(opts, "SetStrategyCap")
	if err != nil {
		return nil, err
	}
	return &CrabSetStrategyCapIterator{contract: _Crab.contract, event: "SetStrategyCap", logs: logs, sub: sub}, nil
}

// WatchSetStrategyCap is a free log subscription operation binding the contract event 0x29600e2e028c8c5c2b112d021938e0d0237d8fafcbb20394c56cf9fa4661ca27.
//
// Solidity: event SetStrategyCap(uint256 newCapAmount)
func (_Crab *CrabFilterer) WatchSetStrategyCap(opts *bind.WatchOpts, sink chan<- *CrabSetStrategyCap) (event.Subscription, error) {

	logs, sub, err := _Crab.contract.WatchLogs(opts, "SetStrategyCap")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrabSetStrategyCap)
				if err := _Crab.contract.UnpackLog(event, "SetStrategyCap", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetStrategyCap is a log parse operation binding the contract event 0x29600e2e028c8c5c2b112d021938e0d0237d8fafcbb20394c56cf9fa4661ca27.
//
// Solidity: event SetStrategyCap(uint256 newCapAmount)
func (_Crab *CrabFilterer) ParseSetStrategyCap(log types.Log) (*CrabSetStrategyCap, error) {
	event := new(CrabSetStrategyCap)
	if err := _Crab.contract.UnpackLog(event, "SetStrategyCap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrabTimeHedgeIterator is returned from FilterTimeHedge and is used to iterate over the raw logs and unpacked data for TimeHedge events raised by the Crab contract.
type CrabTimeHedgeIterator struct {
	Event *CrabTimeHedge // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrabTimeHedgeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrabTimeHedge)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrabTimeHedge)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrabTimeHedgeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrabTimeHedgeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrabTimeHedge represents a TimeHedge event raised by the Crab contract.
type CrabTimeHedge struct {
	Hedger                  common.Address
	AuctionType             bool
	HedgerPrice             *big.Int
	AuctionTriggerTimestamp *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterTimeHedge is a free log retrieval operation binding the contract event 0x00b3acebad2d25b9626850dd370eadbf46d6a94dd0fab19c061f97e1dd4a9639.
//
// Solidity: event TimeHedge(address indexed hedger, bool auctionType, uint256 hedgerPrice, uint256 auctionTriggerTimestamp)
func (_Crab *CrabFilterer) FilterTimeHedge(opts *bind.FilterOpts, hedger []common.Address) (*CrabTimeHedgeIterator, error) {

	var hedgerRule []interface{}
	for _, hedgerItem := range hedger {
		hedgerRule = append(hedgerRule, hedgerItem)
	}

	logs, sub, err := _Crab.contract.FilterLogs(opts, "TimeHedge", hedgerRule)
	if err != nil {
		return nil, err
	}
	return &CrabTimeHedgeIterator{contract: _Crab.contract, event: "TimeHedge", logs: logs, sub: sub}, nil
}

// WatchTimeHedge is a free log subscription operation binding the contract event 0x00b3acebad2d25b9626850dd370eadbf46d6a94dd0fab19c061f97e1dd4a9639.
//
// Solidity: event TimeHedge(address indexed hedger, bool auctionType, uint256 hedgerPrice, uint256 auctionTriggerTimestamp)
func (_Crab *CrabFilterer) WatchTimeHedge(opts *bind.WatchOpts, sink chan<- *CrabTimeHedge, hedger []common.Address) (event.Subscription, error) {

	var hedgerRule []interface{}
	for _, hedgerItem := range hedger {
		hedgerRule = append(hedgerRule, hedgerItem)
	}

	logs, sub, err := _Crab.contract.WatchLogs(opts, "TimeHedge", hedgerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrabTimeHedge)
				if err := _Crab.contract.UnpackLog(event, "TimeHedge", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTimeHedge is a log parse operation binding the contract event 0x00b3acebad2d25b9626850dd370eadbf46d6a94dd0fab19c061f97e1dd4a9639.
//
// Solidity: event TimeHedge(address indexed hedger, bool auctionType, uint256 hedgerPrice, uint256 auctionTriggerTimestamp)
func (_Crab *CrabFilterer) ParseTimeHedge(log types.Log) (*CrabTimeHedge, error) {
	event := new(CrabTimeHedge)
	if err := _Crab.contract.UnpackLog(event, "TimeHedge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrabTimeHedgeOnUniswapIterator is returned from FilterTimeHedgeOnUniswap and is used to iterate over the raw logs and unpacked data for TimeHedgeOnUniswap events raised by the Crab contract.
type CrabTimeHedgeOnUniswapIterator struct {
	Event *CrabTimeHedgeOnUniswap // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrabTimeHedgeOnUniswapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrabTimeHedgeOnUniswap)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrabTimeHedgeOnUniswap)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrabTimeHedgeOnUniswapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrabTimeHedgeOnUniswapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrabTimeHedgeOnUniswap represents a TimeHedgeOnUniswap event raised by the Crab contract.
type CrabTimeHedgeOnUniswap struct {
	Hedger                  common.Address
	HedgeTimestamp          *big.Int
	AuctionTriggerTimestamp *big.Int
	MinWSqueeth             *big.Int
	MinEth                  *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterTimeHedgeOnUniswap is a free log retrieval operation binding the contract event 0x6c3a0d23de8295593e3e236062f9103f4a66c6d8de92b9425a2e17ae3baca677.
//
// Solidity: event TimeHedgeOnUniswap(address indexed hedger, uint256 hedgeTimestamp, uint256 auctionTriggerTimestamp, uint256 minWSqueeth, uint256 minEth)
func (_Crab *CrabFilterer) FilterTimeHedgeOnUniswap(opts *bind.FilterOpts, hedger []common.Address) (*CrabTimeHedgeOnUniswapIterator, error) {

	var hedgerRule []interface{}
	for _, hedgerItem := range hedger {
		hedgerRule = append(hedgerRule, hedgerItem)
	}

	logs, sub, err := _Crab.contract.FilterLogs(opts, "TimeHedgeOnUniswap", hedgerRule)
	if err != nil {
		return nil, err
	}
	return &CrabTimeHedgeOnUniswapIterator{contract: _Crab.contract, event: "TimeHedgeOnUniswap", logs: logs, sub: sub}, nil
}

// WatchTimeHedgeOnUniswap is a free log subscription operation binding the contract event 0x6c3a0d23de8295593e3e236062f9103f4a66c6d8de92b9425a2e17ae3baca677.
//
// Solidity: event TimeHedgeOnUniswap(address indexed hedger, uint256 hedgeTimestamp, uint256 auctionTriggerTimestamp, uint256 minWSqueeth, uint256 minEth)
func (_Crab *CrabFilterer) WatchTimeHedgeOnUniswap(opts *bind.WatchOpts, sink chan<- *CrabTimeHedgeOnUniswap, hedger []common.Address) (event.Subscription, error) {

	var hedgerRule []interface{}
	for _, hedgerItem := range hedger {
		hedgerRule = append(hedgerRule, hedgerItem)
	}

	logs, sub, err := _Crab.contract.WatchLogs(opts, "TimeHedgeOnUniswap", hedgerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrabTimeHedgeOnUniswap)
				if err := _Crab.contract.UnpackLog(event, "TimeHedgeOnUniswap", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTimeHedgeOnUniswap is a log parse operation binding the contract event 0x6c3a0d23de8295593e3e236062f9103f4a66c6d8de92b9425a2e17ae3baca677.
//
// Solidity: event TimeHedgeOnUniswap(address indexed hedger, uint256 hedgeTimestamp, uint256 auctionTriggerTimestamp, uint256 minWSqueeth, uint256 minEth)
func (_Crab *CrabFilterer) ParseTimeHedgeOnUniswap(log types.Log) (*CrabTimeHedgeOnUniswap, error) {
	event := new(CrabTimeHedgeOnUniswap)
	if err := _Crab.contract.UnpackLog(event, "TimeHedgeOnUniswap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrabTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Crab contract.
type CrabTransferIterator struct {
	Event *CrabTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrabTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrabTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrabTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrabTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrabTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrabTransfer represents a Transfer event raised by the Crab contract.
type CrabTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Crab *CrabFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CrabTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Crab.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CrabTransferIterator{contract: _Crab.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Crab *CrabFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CrabTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Crab.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrabTransfer)
				if err := _Crab.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Crab *CrabFilterer) ParseTransfer(log types.Log) (*CrabTransfer, error) {
	event := new(CrabTransfer)
	if err := _Crab.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrabWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Crab contract.
type CrabWithdrawIterator struct {
	Event *CrabWithdraw // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrabWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrabWithdraw)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrabWithdraw)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrabWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrabWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrabWithdraw represents a Withdraw event raised by the Crab contract.
type CrabWithdraw struct {
	Withdrawer     common.Address
	CrabAmount     *big.Int
	WSqueethAmount *big.Int
	EthWithdrawn   *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x02f25270a4d87bea75db541cdfe559334a275b4a233520ed6c0a2429667cca94.
//
// Solidity: event Withdraw(address indexed withdrawer, uint256 crabAmount, uint256 wSqueethAmount, uint256 ethWithdrawn)
func (_Crab *CrabFilterer) FilterWithdraw(opts *bind.FilterOpts, withdrawer []common.Address) (*CrabWithdrawIterator, error) {

	var withdrawerRule []interface{}
	for _, withdrawerItem := range withdrawer {
		withdrawerRule = append(withdrawerRule, withdrawerItem)
	}

	logs, sub, err := _Crab.contract.FilterLogs(opts, "Withdraw", withdrawerRule)
	if err != nil {
		return nil, err
	}
	return &CrabWithdrawIterator{contract: _Crab.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x02f25270a4d87bea75db541cdfe559334a275b4a233520ed6c0a2429667cca94.
//
// Solidity: event Withdraw(address indexed withdrawer, uint256 crabAmount, uint256 wSqueethAmount, uint256 ethWithdrawn)
func (_Crab *CrabFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *CrabWithdraw, withdrawer []common.Address) (event.Subscription, error) {

	var withdrawerRule []interface{}
	for _, withdrawerItem := range withdrawer {
		withdrawerRule = append(withdrawerRule, withdrawerItem)
	}

	logs, sub, err := _Crab.contract.WatchLogs(opts, "Withdraw", withdrawerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrabWithdraw)
				if err := _Crab.contract.UnpackLog(event, "Withdraw", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdraw is a log parse operation binding the contract event 0x02f25270a4d87bea75db541cdfe559334a275b4a233520ed6c0a2429667cca94.
//
// Solidity: event Withdraw(address indexed withdrawer, uint256 crabAmount, uint256 wSqueethAmount, uint256 ethWithdrawn)
func (_Crab *CrabFilterer) ParseWithdraw(log types.Log) (*CrabWithdraw, error) {
	event := new(CrabWithdraw)
	if err := _Crab.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrabWithdrawShutdownIterator is returned from FilterWithdrawShutdown and is used to iterate over the raw logs and unpacked data for WithdrawShutdown events raised by the Crab contract.
type CrabWithdrawShutdownIterator struct {
	Event *CrabWithdrawShutdown // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrabWithdrawShutdownIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrabWithdrawShutdown)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrabWithdrawShutdown)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrabWithdrawShutdownIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrabWithdrawShutdownIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrabWithdrawShutdown represents a WithdrawShutdown event raised by the Crab contract.
type CrabWithdrawShutdown struct {
	Withdrawer   common.Address
	CrabAmount   *big.Int
	EthWithdrawn *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterWithdrawShutdown is a free log retrieval operation binding the contract event 0xe9ab9870b9093d99f8e981919f65ad09b7ae90ff80f1031639af9e0357eb9ed6.
//
// Solidity: event WithdrawShutdown(address indexed withdrawer, uint256 crabAmount, uint256 ethWithdrawn)
func (_Crab *CrabFilterer) FilterWithdrawShutdown(opts *bind.FilterOpts, withdrawer []common.Address) (*CrabWithdrawShutdownIterator, error) {

	var withdrawerRule []interface{}
	for _, withdrawerItem := range withdrawer {
		withdrawerRule = append(withdrawerRule, withdrawerItem)
	}

	logs, sub, err := _Crab.contract.FilterLogs(opts, "WithdrawShutdown", withdrawerRule)
	if err != nil {
		return nil, err
	}
	return &CrabWithdrawShutdownIterator{contract: _Crab.contract, event: "WithdrawShutdown", logs: logs, sub: sub}, nil
}

// WatchWithdrawShutdown is a free log subscription operation binding the contract event 0xe9ab9870b9093d99f8e981919f65ad09b7ae90ff80f1031639af9e0357eb9ed6.
//
// Solidity: event WithdrawShutdown(address indexed withdrawer, uint256 crabAmount, uint256 ethWithdrawn)
func (_Crab *CrabFilterer) WatchWithdrawShutdown(opts *bind.WatchOpts, sink chan<- *CrabWithdrawShutdown, withdrawer []common.Address) (event.Subscription, error) {

	var withdrawerRule []interface{}
	for _, withdrawerItem := range withdrawer {
		withdrawerRule = append(withdrawerRule, withdrawerItem)
	}

	logs, sub, err := _Crab.contract.WatchLogs(opts, "WithdrawShutdown", withdrawerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrabWithdrawShutdown)
				if err := _Crab.contract.UnpackLog(event, "WithdrawShutdown", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawShutdown is a log parse operation binding the contract event 0xe9ab9870b9093d99f8e981919f65ad09b7ae90ff80f1031639af9e0357eb9ed6.
//
// Solidity: event WithdrawShutdown(address indexed withdrawer, uint256 crabAmount, uint256 ethWithdrawn)
func (_Crab *CrabFilterer) ParseWithdrawShutdown(log types.Log) (*CrabWithdrawShutdown, error) {
	event := new(CrabWithdrawShutdown)
	if err := _Crab.contract.UnpackLog(event, "WithdrawShutdown", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
