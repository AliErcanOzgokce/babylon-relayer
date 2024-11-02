package types

const (
	// ModuleName defines the module name
	ModuleName = "babylonrelayer"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_babylonrelayer"
)

var (
	ParamsKey = []byte("p_babylonrelayer")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
