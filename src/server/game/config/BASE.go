package config

type ItemConf struct {
	Id  uint32 `json:"id"`
	Cnt uint32 `json:"cnt"`
}

type SkillConf struct {
	Id uint32 `json:"id"`
	Lv uint32 `json:"lv"`
}

type AuraConf struct {
	Id uint32 `json:"id"`
	Lv uint32 `json:"lv"`
}

type AttrConf struct {
	Id  uint32 `json:"id"`
	Val uint32 `json:"val"`
}

func MakeUint64(l uint32, r uint32) uint64 {
	return (uint64(l) << 32) | uint64(r)
}
