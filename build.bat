@echo off

cd %~dp0

setlocal

go build

set PORT=3000

sample_website.exe
