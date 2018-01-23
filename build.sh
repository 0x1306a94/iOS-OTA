#!/bin/sh

GOOS=linux GOARCH=amd64 go build -o iOS-OTA-Server && upx iOS-OTA-Server

rm -rf ~/Desktop/iOS-OTA-Docker/iOS-OTA-Server

sleep 1

cp -rf iOS-OTA-Server ~/Desktop/iOS-OTA-Docker/iOS-OTA-Server

#cp -rf conf ~/Desktop/iOS-OTA-Docker/OTA
#cp -rf download ~/Desktop/iOS-OTA-Docker/OTA
#cp -rf static ~/Desktop/iOS-OTA-Docker/OTA
#cp -rf views ~/Desktop/iOS-OTA-Docker/OTA

#GOOS=linux GOARCH=386 go build -o package-notice-linux-i386 && upx package-notice-linux-i386
#GOOS=windows GOARCH=amd64 go build -o package-notice-windows-64.exe && upx package-notice-windows-64.exe
#GOOS=windows GOARCH=386 go build -o package-notice-windows-i386.exe && upx package-notice-windows-i386.exe
#GOOS=darwin GOARCH=amd64 go build -o package-notice-darwin && upx package-notice-darwin