package types

const (
	ModuleName   = "government"
	StoreKey     = ModuleName
	RouterKey    = ModuleName
	QuerierRoute = ModuleName

	GovernmentStoreKey = StoreKey + "government"
	TumblerStoreKey    = StoreKey + "tumbler"

	MsgTypeSetTumblerAddress = "setTumblerAddress"

	QueryGovernmentAddress = "governmentAddress"
	QueryTumblerAddress    = "tumblerAddress"
)
