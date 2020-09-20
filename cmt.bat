@echo off
set /p comment="Comment : "

IF "%comment%"=="" (SET comment=".")

git config --global user.email "dohyung97022@gmail.com"
git config --global user.name "Doe"
git update-index --assume-unchanged "mysqlfunc_test.go"
git add .
git commit -m "%comment%"
git push origin master