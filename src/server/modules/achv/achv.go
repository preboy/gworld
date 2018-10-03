package achv

type achv_t struct {
	Id        int32 // achv id
	GrowthId  int32
	Completed bool
}

func IsCompleted(plr iPlayer, id int32) bool {
	if v, ok := plr.GetGrowth().achv[id]; ok && v {
		return true
	}
	return false
}

func AchvVal(plr iPlayer, id int32) int32 {
	// growth := plr.GetGrowth()
	// achv := growth.achv[id]
	// if achv != nil {
	// 	item := growth.GrowthV[achv.GrowthId]
	// 	if item != nil {
	// 		return item.Val
	// 	}
	// }

	return 0
}
