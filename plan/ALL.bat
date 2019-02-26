@echo off


del /S /Q json\*.*
del /S /Q lua\*.*


rem: record export

exporter jl excel/levelup_等级表.xlsx               Levelup
exporter jl excel/hero_英雄表.xlsx                  Hero
exporter jl excel/creature_怪物表.xlsx              Creature
exporter jl excel/creature_team_怪物队伍表.xlsx     CreatureTeam
exporter jl excel/skill_技能表.xlsx                 Skill
exporter jl excel/aura_光环表.xlsx                  Aura
exporter jl excel/item_道具表.xlsx                  Item
exporter jl excel/market_集市表.xlsx                Market
exporter jl excel/activity_活动表.xlsx              Activity

exporter jl excel/refine_精炼表.xlsx                RefineNormal
exporter jl excel/refine_精炼表.xlsx                RefineSuper

exporter jl excel/achv_成就表.xlsx                  Growth
exporter jl excel/achv_成就表.xlsx                  Achv

exporter jl excel/world_场景对象表.xlsx             Scene
exporter jl excel/world_场景对象表.xlsx             Object


rem: map export
exporter_map jl excel/global_全局表.xlsx            Global


pause
