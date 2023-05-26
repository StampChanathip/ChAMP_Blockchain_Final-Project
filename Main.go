package main

import (
	"fmt"
	"log"

	"example/api"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
)

func main() {
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/dafdfd757bd24faeba7bb89624268497")
	if err != nil {
		log.Fatal(err)
	}
	_ = client

	poolAddress := common.HexToAddress("0xB4e16d0168e52d35CaCD2c6185b44281Ec28C9Dc")

	instance, err := api.NewApi(poolAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	dec, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	reserve, err := instance.GetReserves(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	//token0: USDC
	//token1: WETH
	usdcBalance := decimal.NewFromBigInt(reserve.Reserve0, -6)
	ethBalance := decimal.NewFromBigInt(reserve.Reserve1, -int32(dec))
	usdcPrice := ethBalance.Div(usdcBalance) //1USDC = ...ETH
	ethPrice := usdcBalance.Div(ethBalance)  //1ETH = ...USDC

	// fmt.Println(usdcBalance, ethBalance)
	// fmt.Println(usdcPrice, ethPrice)
	fmt.Println("1 USDC = " + usdcPrice.String() + " ETH")
	fmt.Println("1 ETH = " + ethPrice.String() + " USDC")

}
