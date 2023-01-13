@echo off
chcp 65001

set tagV=v0.0.4

echo "------------%tagV%------------------------------"

git tag -a %tagV% -m "%tagV%"

git push origin %tagV%

