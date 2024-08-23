@echo off
setlocal enabledelayedexpansion

:: Number of concurrent clients
set NUM_CLIENTS=1000000

:: Port to connect to
set PORT=7379
set HOST=127.0.0.1

for /L %%i in (1,1,%NUM_CLIENTS%) do (
    start "" ncat %HOST% %PORT% -e "cmd /c echo Hello from client %%i"
    timeout /t 1 >nul
)