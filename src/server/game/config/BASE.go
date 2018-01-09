package config

type ENC struct {
	ItemId uint32 `json:"Id"`
	ItemCt uint32 `json:"Cnt"`
}

func MakeUint64(l uint32, r uint32) uint64 {
	return (uint64(l) << 32) | uint64(r)
}
