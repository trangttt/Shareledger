package app

import (
	cmn "github.com/tendermint/tmlibs/common"
	dbm "github.com/tendermint/tmlibs/db"
	"github.com/tendermint/tmlibs/log"

	bapp "github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/wire"

	"github.com/sharering/shareledger/types"
	"github.com/sharering/shareledger/x/bank"
	"github.com/sharering/shareledger/x/asset"
	"github.com/sharering/shareledger/x/asset/handlers"

	"github.com/sharering/shareledger/x/booking"
	bhandlers "github.com/sharering/shareledger/x/booking/handlers"
)

const (
	ShareLedgerApp = "ShareLedger_v0.0.1"
)

func NewShareLedgerApp(logger log.Logger, db dbm.DB) *bapp.BaseApp {

	cdc := MakeCodec()

	// Create the base application object.
	app := bapp.NewBaseApp(ShareLedgerApp, cdc, logger, db)

	assetKey := sdk.NewKVStoreKey("asset")
	bookingKey := sdk.NewKVStoreKey("booking")
	accountKey := sdk.NewKVStoreKey("account")

	SetupAsset(app, cdc, assetKey)
	SetupBank(app, cdc, accountKey)
	SetupBooking(app, cdc, bookingKey, assetKey, accountKey)

	// Determine how transactions are decoded.
	app.SetTxDecoder(types.GetTxDecoder(cdc))

	return app
}

func MakeCodec() *wire.Codec {
	cdc := wire.NewCodec()
	cdc.RegisterInterface((*sdk.Msg)(nil), nil)
	return cdc
}

func SetupBank(app *bapp.BaseApp, cdc *wire.Codec, accountKey *sdk.KVStoreKey) {
	// Bank module
	// Create a key for accessing the account store.
	cdc = bank.RegisterCodec(cdc)

	// Register message routes.
	// Note the handler gets access to the account store.
	app.Router().
		AddRoute("bank", bank.NewHandler(accountKey))

	// Mount stores and load the latest state.
	app.MountStoresIAVL(accountKey)
	err := app.LoadLatestVersion(accountKey)
	if err != nil {
		cmn.Exit(err.Error())
	}
}

func SetupAsset(app *bapp.BaseApp, cdc *wire.Codec, assetKey *sdk.KVStoreKey) {

	keeper := asset.NewKeeper(assetKey, cdc)

	cdc = asset.RegisterCodec(cdc)


	app.Router().
		AddRoute("asset", handlers.NewHandler(keeper))

	app.MountStoresIAVL(assetKey)
	err := app.LoadLatestVersion(assetKey)
	if err != nil {
		cmn.Exit(err.Error())
	}
}


func SetupBooking(app *bapp.BaseApp, cdc *wire.Codec, bookingKey *sdk.KVStoreKey,
	              assetKey *sdk.KVStoreKey, accountKey *sdk.KVStoreKey){

	cdc = booking.RegisterCodec(cdc)

	k := booking.NewKeeper(bookingKey,
						   assetKey,
		 				   accountKey,
						   cdc)

	app.Router().
		AddRoute("booking", bhandlers.NewHandler(k))

	app.MountStoresIAVL(bookingKey)
	err := app.LoadLatestVersion(bookingKey)
	if err != nil {
		cmn.Exit(err.Error())
	}

}
