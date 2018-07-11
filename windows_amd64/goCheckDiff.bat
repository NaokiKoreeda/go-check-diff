@echo off
cd /d %~dp0

set baseDir=C:/Users/NaokiKoreeda/dir/
set targetDir=D:/Users/NaokiKoreeda/dir/

call goCheckDiff.exe %baseDir% %targetDir%

pause
