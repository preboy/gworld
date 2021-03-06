@echo off


del /S /Q json\*.*
del /S /Q lua\*.*


rem: record export

exporter jl excel/levelup_等级表.xlsx               Levelup

exporter jl excel/hero_英雄表.xlsx                  Hero
exporter jl excel/hero_英雄表.xlsx                  HeroProp
exporter jl excel/hero_英雄表.xlsx                  HeroTalent

exporter jl excel/creature_怪物表.xlsx              Creature
exporter jl excel/creature_team_怪物队伍表.xlsx     CreatureTeam

exporter jl excel/skill_技能表.xlsx                 Skill
exporter jl excel/aura_光环表.xlsx                  Aura

exporter jl excel/item_道具表.xlsx                  Item

exporter jl excel/market_集市表.xlsx                Market

exporter jl excel/activity_活动表.xlsx              Activity
exporter jl excel/task_个人活动表.xlsx              Task

exporter jl excel/refine_精炼表.xlsx                RefineNormal
exporter jl excel/refine_精炼表.xlsx                RefineSuper

exporter jl excel/achv_成就表.xlsx                  Growth
exporter jl excel/achv_成就表.xlsx                  Achv

exporter jl excel/world_场景对象表.xlsx             Scene
exporter jl excel/world_场景对象表.xlsx             Object

exporter jl excel/drop_掉落表.xlsx                  Drop
exporter jl excel/cond_条件表.xlsx                  Cond

exporter jl excel/break_关卡表.xlsx                 Break
exporter jl excel/break_关卡表.xlsx                 Chapter


rem: map export
exporter_map jl excel/global_全局表.xlsx            Global


echo "Checking lua files ..."
forfiles /p .\lua  /c "cmd /c lua @file"

pause
