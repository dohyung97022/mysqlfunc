:: allways check if password is deleted before running this bat file!

:LOOP
set /p YN=(Y/N)?

if /i "%YN%" == "y" goto YES
if /i "%YN%" == "n" goto NO

goto LOOP


:YES
echo YES
git update-index --no--assume-unchanged "mysqlfunc_test.go"
git add .
git commit -m "."
git push origin master
git update-index --assume-unchanged "mysqlfunc_test.go"

:NO
echo NO
goto QUIT

:QUIT