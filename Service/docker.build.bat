set /p version= < ver
set /A version+=1
@echo off
echo %version% > ver

docker build --tag fullstack:2.%version% .
pause