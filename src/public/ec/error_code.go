package ec

const (

	// --------------------------------
	// 基本 [0, 100)
	// --------------------------------

	OK                 = iota + 0 // 无错误
	Failed                        // 操作失败
	Conf_Invalid                  // 配置未找到
	Item_Unknown                  // 未知道具ID
	Item_Not_Enough               // 道具数量不足
	Item_Unusable                 // 不可使用的道具
	Hero_Not_Activated            // 未获得英雄
	Level_Exceed                  // 等级超过限制

	// --------------------------------
	// 登录 [100, 200)
	// --------------------------------

	Login_Failed = iota + 100 // 用户名、密码不正确
	Login_Dup                 // 已在线，重复登录
	Login_Not                 // 未登录

	// --------------------------------
	// 任务 [200, 300)
	// --------------------------------
	QUEST_Cond_Dissatisfy   // 接任务条件不满足
	QUEST_Pass_Over         // 过去的任务不能接
	QUEST_Not_Finish        // 任务未领奖
	QUEST_Finish_Yet        // 任务已经完成了
	QUEST_Task_Unfinished   // 任务项未完成
	QUEST_Unknown           // 未知的任务
	QUEST_Tasks_Over        // 已完成所有任务项
	QUEST_Task_Invalid_r    // 提交对话任务若的索引不对
	QUEST_Task_Invalid_Kill // 击杀怪物数据不足

	// 功能错误码
	ERR_END
)
