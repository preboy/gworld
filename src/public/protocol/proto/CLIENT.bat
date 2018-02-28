@echo off
protoc --descriptor_set_out ../protocol.pb game.proto login.proto session.proto gm.proto