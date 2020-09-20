:: allways check if password is deleted before running this bat file!
@echo off

:LOOP
set /p YN=Did you delete your password in the test file? (YES/NO)?

if /i "%YN%" == "yes" goto YES
if /i "%YN%" == "no" goto NO

goto LOOP


:YES
git update-index --no-assume-unchanged "mysqlfunc_test.go"
git add .
git commit -m "."
git push origin master
git update-index --assume-unchanged "mysqlfunc_test.go"

:NO

goto QUIT

:QUIT