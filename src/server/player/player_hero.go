package player

import ()

type Equipment struct {
	quality uint32
	level   uint32
}

type Hero struct {
	// 这里的数据就是要存入DB的数据
	Pid   uint32    `bson:pid"`
	Uid   uint32    `bson:uid"`
	Level uint32    `bson:"level"`
	Equip Equipment `bson:"equip"`
}
