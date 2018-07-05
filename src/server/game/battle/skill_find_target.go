package battle

import (
	"fmt"
)

/*
我方：
	0
	1
	2
	3
	4
	5	// 以上为绝对位置
	6	自己
	7	自己所在排行(包括自己)
	8	自己所在列(包括自己)
	9	自己周围(包括自己)
	10	自己周围(不包括自己)
	11	已方第一排
	12	乙方第二排行
	13	乙方全体
	14	已方第一列
	15	已方第二列
	16	已方第三列
	17	乙方HP最少单位
	18	乙方HP百分比最少单位
	101	// 定向寻找Id，比如张飞、张苞(具体在另一张表中依次列表所有可能)


敌方：
	200	对面的第一个单位(简称目标)
	201	对面的第二个单位
	202	对面的所有单位
	203	对方第一排
	204	对方第二排
	205	对方第一列
	206	对方第二列
	207	对方第三列
	208	对方全体
	209	目标所在排
	210	目标所在列
	211	目标周边
	212	敌方HP最少
	213	敌方HP百分比最少
	301	//定向寻找Id，（与101类似，寻找我的仇人，比如吕蒙喜欢找关羽杀）
*/

func (self *BattleSkill) find_target_impl(ids []int32, units []*BattleUnit) {
	for _, id := range ids {
		switch id {
		// ------------------------------- 我方单位 -------------------------------

		case 1, 2, 3, 4, 5, 6:
			{
				u := self.caster.Troop.members[id-1]
				units = append(units, u)
			}

		// ------------------------------- 敌方单位 -------------------------------

		case 201, 202, 203, 204, 205, 206:
			{
				u := self.caster.Troop.GetRival().members[id-200-1]
				units = append(units, u)
			}
		default:
			{
				fmt.Println("kunown target type:", id)
			}
		}
	}
}

func (self *BattleSkill) find_target() {
	self.find_target_impl(self.proto.Target_major, self.target_major)
	self.find_target_impl(self.proto.Target_minor, self.target_minor)
}
