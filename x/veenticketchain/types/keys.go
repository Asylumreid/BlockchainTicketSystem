package types

const (
	// ModuleName defines the module name
	ModuleName = "veenticketchain"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_veenticketchain"
)

var (
	ParamsKey = []byte("p_veenticketchain")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
