@echo off
chcp 65001

set tagV=v0.0.8

echo "------------%tagV%------------------------------"

git tag -a %tagV% -m "%tagV%"

git push origin %tagV%

