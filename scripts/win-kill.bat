for /f "tokens=5" %%i in ('netstat -ano ^| findstr /C:"LISTENING" ^| findstr /C:"443" ' ) do set "PID=%%i"
echo The PID of the process listening on port 443 is: %PID%
taskkill /PID %PID% /F 