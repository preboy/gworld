package battle

import (
	"fmt"
)

/*
formation
	4 1	   vs	 9 12
	5 2    vs    8 11
	6 3	   vs	 7 10
*/

/*
已方：
	1 ~ 6	// 绝对位置
	7	自己所在排(包括自己)
	8	自己所在列(包括自己)
	9	自己周围(包括自己)
	10	自己周围(不包括自己)
	11	已方第一排
	12	已方第二排行
	13	已方全体
	14	已方第一列
	15	已方第二列
	16	已方第三列
	17	已方HP最少单位
	18	已方HP百分比最少单位
	19  自己
	101	// 定向寻找Id，比如张飞、张苞(具体在另一张表中依次列表所有可能)


敌方：
	201 ~ 106 // 绝对位置
	207	对面的第一个单位(简称目标)
	208	对面的第二个单位
	209	对面的所有单位
	210	对方第一排
	211	对方第二排
	212	对方第一列
	213	对方第二列
	214	对方第三列
	215	对方全体
	216	目标所在排
	217	目标所在列
	218	目标周边
	219	敌方HP最少
	220	敌方HP百分比最少
	221 默认目标(直线对过去的第一个目标)
		//定向寻找Id，（与101类似，寻找我的仇人，比如吕蒙喜欢找关羽杀）
*/

func _add_target(units []*BattleUnit, us ...*BattleUnit) {
	for _, u := range us {
		if u != nil && !u.Dead() {
			units = append(units, u)
		}
	}
}

func (self *BattleSkill) find_target_impl(ids []int32, units []*BattleUnit) {
	for _, id := range ids {
		switch id {
		// ------------------------------- 我方单位 -------------------------------
		case 1, 2, 3, 4, 5, 6:
			{
				_add_target(units, self.caster.Troop.members[id])
			}
		case 7: // 自己所在排(包括自己)
			{
				members := self.caster.Troop.members
				pos := self.caster.Pos
				switch pos {
				case 1, 2, 3, 7, 8, 9:
					_add_target(units, members[0], members[1], members[2])
				case 4, 5, 6, 10, 11, 12:
					_add_target(units, members[3], members[4], members[5])
				}
			}
		case 8: // 自己所在列(包括自己)
			{
				members := self.caster.Troop.members
				pos := self.caster.Pos
				switch pos {
				case 1, 4, 7, 10:
					_add_target(units, members[0], members[3])
				case 2, 5, 8, 11:
					_add_target(units, members[1], members[4])
				case 3, 6, 9, 12:
					_add_target(units, members[2], members[5])
				}
			}
		case 9: // 自己周围(包括自己)
			{
			}
		case 10: // 自己周围(不包括自己)
			{
			}
		case 11: // 已方第一排
			{
				members := self.caster.Troop.members
				_add_target(units, members[0], members[1], members[2])
			}
		case 12: //	已方第二排行
			{
				members := self.caster.Troop.members
				_add_target(units, members[3], members[4], members[5])
			}
		case 13: //	已方全体
			{
				members := self.caster.Troop.members
				_add_target(units, members[0], members[1], members[2], members[3], members[4], members[5])
			}
		case 14: // 已方第一列
			{
				members := self.caster.Troop.members
				_add_target(units, members[0], members[3])
			}
		case 15: // 已方第二列
			{
				members := self.caster.Troop.members
				_add_target(units, members[1], members[4])
			}
		case 16: // 已方第三列
			{
				members := self.caster.Troop.members
				_add_target(units, members[2], members[5])
			}
		case 17: // 已方HP最少单位
			{
			}
		case 18: // 已方HP百分比最少单位
			{
			}
		case 19: // 自己
			{
				_add_target(units, self.caster)
			}

		// ------------------------------- 敌方单位 -------------------------------
		case 201, 202, 203, 204, 205, 206:
			{
				u := self.caster.Troop.members[id]
				if u != nil && !u.Dead() {
					units = append(units, u)
				}
			}
		case 207: // 对面的第一个单位(简称目标)
			{
				members := self.caster.Troop.GetRival().members
				switch self.caster.Pos {
				case 1, 4, 7, 10:
					_add_target(units, members[2])
				case 2, 5, 8, 11:
					_add_target(units, members[1])
				case 3, 6, 9, 12:
					_add_target(units, members[0])
				}
			}
		case 208: // 对面的第二个单位
			{
				members := self.caster.Troop.GetRival().members
				switch self.caster.Pos {
				case 1, 4, 7, 10:
					_add_target(units, members[5])
				case 2, 5, 8, 11:
					_add_target(units, members[4])
				case 3, 6, 9, 12:
					_add_target(units, members[3])
				}
			}
		case 209: // 对面的所有单位
			{
				members := self.caster.Troop.GetRival().members
				switch self.caster.Pos {
				case 1, 4, 7, 10:
					_add_target(units, members[2], members[5])
				case 2, 5, 8, 11:
					_add_target(units, members[1], members[4])
				case 3, 6, 9, 12:
					_add_target(units, members[0], members[3])
				}
			}
		case 210: // 对方第一排
			{
				members := self.caster.Troop.GetRival().members
				_add_target(units, members[0], members[1], members[2])
			}
		case 211: // 对方第二排
			{
				members := self.caster.Troop.GetRival().members
				_add_target(units, members[3], members[4], members[5])
			}
		case 212: // 对方第一列
			{
				members := self.caster.Troop.GetRival().members
				_add_target(units, members[0], members[3])
			}
		case 213: // 对方第二列
			{
				members := self.caster.Troop.GetRival().members
				_add_target(units, members[1], members[4])
			}
		case 214: // 对方第三列
			{
				members := self.caster.Troop.GetRival().members
				_add_target(units, members[2], members[5])
			}
		case 215: // 对方全体
			{
				members := self.caster.Troop.GetRival().members
				_add_target(units, members[0], members[1], members[2], members[3], members[4], members[5])
			}
		case 216: // 目标所在排
			{
				target := self._default_target()
				members := self.caster.Troop.GetRival().members
				if target != nil {
					switch target.Pos {
					case 1, 2, 3, 7, 8, 9:
						_add_target(units, members[0], members[1], members[2])
					case 4, 5, 6, 10, 11, 12:
						_add_target(units, members[3], members[4], members[5])
					}
				}
			}
		case 217: // 目标所在列
			{
				target := self._default_target()
				members := self.caster.Troop.GetRival().members
				if target != nil {
					switch target.Pos {
					case 1, 4, 7, 10:
						_add_target(units, members[0], members[3])
					case 2, 5, 8, 11:
						_add_target(units, members[1], members[4])
					case 3, 6, 9, 12:
						_add_target(units, members[2], members[5])
					}
				}
			}
		case 218: // 目标周边
			{
			}
		case 219: // 敌方HP最少
			{
			}
		case 220: // 敌方HP百分比最少
			{
			}
		case 221: // 默认目标(直线对过去的第一个目标)
			{
				_add_target(units, self._default_target())
			}
		default:
			{
				fmt.Println("unknown target type:", id)
			}
		}
	}
}

func (self *BattleSkill) _default_target() *BattleUnit {
	var pos int
	members := self.caster.Troop.GetRival().members
	switch self.caster.Pos {
	case 1, 4, 7, 10:
		pos = 2
	case 2, 5, 8, 11:
		pos = 1
	case 3, 6, 9, 12:
		pos = 0
	}
	for i := 0; i < 2; i++ {
		tpos := pos + i*3
		if members[tpos] != nil && !members[tpos].Dead() {
			return members[tpos]
		}
	}
	return nil
}

func (self *BattleSkill) find_target() {
	self.find_target_impl(self.proto.Target_major, self.target_major)
	self.find_target_impl(self.proto.Target_minor, self.target_minor)
	if len(self.target_major) == 0 && len(self.target_minor) == 0 {
		members := self.caster.Troop.GetRival().members
		for i := 0; i < MAX_TROOP_MEMBER; i++ {
			if members[i] != nil && !members[i].Dead() {
				self.target_major = append(self.target_major, members[i])
			}
		}
	}
	fmt.Println(self.proto.Id, "技能目标：", self.target_major, self.target_minor)
}
