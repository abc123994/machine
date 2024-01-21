@echo off
set PATH=%PATH%;%GOPATH%\bin 
protoc -I=schema/machine --go_out=. machine.proto
protoc -I=schema/common --go_out=. common.proto