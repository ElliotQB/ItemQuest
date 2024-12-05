@echo off

:CountLines
setlocal
set /a totalNumLines = 0
for /r %1 %%F in (*.go) do (
    for /f %%N in ('find /v /c "" ^<"%%F"') do set /a totalNumLines+=%%N 
    echo|set /p= "."
)

echo .
echo Lines of code: %totalNumLines%
echo Keep up the good work!
pause