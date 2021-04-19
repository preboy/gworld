package ec

const (

	// --------------------------------
	// 基本 [0, 100)
	// --------------------------------

	OK                 = 0  // 无错误
	Failed             = 1  // 操作失败
	Conf_Invalid       = 2  // 配置未找到
	Item_Unknown       = 3  // 未知道具ID
	Item_Not_Enough    = 4  // 道具数量不足
	Item_Unusable      = 5  // 不可使用的道具
	Hero_Not_Activated = 6  // 未获得英雄
	Level_Exceed       = 7  // 等级超过限制
	NamePunctuation    = 8  // 名字包含特殊字符
	NameSensitive      = 9  // 名字包含敏感字
	NameLengthErr      = 10 // 长度错误
	NameSame           = 11 // 与原名相同
	NameConflit        = 12 // 与其他名冲突

	// --------------------------------
	// 登录 [100, 200)
	// --------------------------------

	Login_Failed = 100 // 用户名、密码不正确
	Login_Dup    = 101 // 已在线，重复登录
	Login_Not    = 102 // 未登录

	// --------------------------------
	// 任务 [200, 300)
	// --------------------------------
	QUEST_Cond_Dissatisfy   = 200 // 接任务条件不满足
	QUEST_Pass_Over         = 201 // 过去的任务不能接
	QUEST_Not_Finish        = 202 // 任务未领奖
	QUEST_Finish_Yet        = 203 // 任务已经完成了
	QUEST_Task_Unfinished   = 204 // 任务项未完成
	QUEST_Unknown           = 205 // 未知的任务
	QUEST_Tasks_Over        = 206 // 已完成所有任务项
	QUEST_Task_Invalid_r    = 207 // 提交对话任务若的索引不对
	QUEST_Task_Invalid_Kill = 208 // 击杀怪物数据不足

	// --------------------------------
	// 关卡 [300, 400)
	// --------------------------------
	CHAPTER_LootTimeShort     = 300 // 挂机时间不足
	CHAPTER_NotAccomplish     = 301 // 未完成该章节
	CHAPTER_RewardsGot        = 302 // 已领取过了
	CHAPTER_InvalidCreateTeam = 303 // 关卡队伍生成失败

	// --------------------------------
	// 任务 [400, 500)
	// --------------------------------
	BATTLE_Hero_Cnt_Exceed = 400 // 队伍人数过多
	BATTLE_Hero_Present    = 401 // 英雄重复
	BATTLE_Hero_Zero       = 402 // 队伍无英雄
	BATTLE_Hero_NotExist   = 403 // 不存在的英雄

	// 功能错误码
	ERR_END
)
