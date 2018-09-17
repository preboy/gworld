package app

type Equipment struct {
	Quality uint32 `bson:quality"` // 品质
	Level   uint32 `bson:level"`   // 等级
}
