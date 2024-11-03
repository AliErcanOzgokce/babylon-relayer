package types

const (
	// ModuleName defines the module name
	ModuleName = "btcstaking"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_btcstaking"

	// Version defines the current version the IBC module supports
	Version = "btcstaking-1"

	// PortID is the default port id that module binds to
	PortID = "btcstaking"
)

var (
	ParamsKey = []byte("p_btcstaking")
)

var (
	// PortKey defines the key to store the port ID in store
	PortKey = KeyPrefix("btcstaking-port-")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
