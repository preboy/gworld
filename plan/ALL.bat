@echo off


del /S /Q json\*.*
del /S /Q lua\*.*


rem: record export
exporter jl excel/hero_英雄表.xlsx  HeroProto
exporter jl excel/creature_怪物表.xlsx  CreatureProto


rem: map export
exporter_map jl excel/global_全局表.xlsx  global


pause