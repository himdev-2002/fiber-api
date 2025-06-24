#Dev Run
go run main.go

#Dev Run in Linux via air
./bin/air -c air.linux.conf

#Dev Run in Windows via air
./bin/air -c air.windows.conf

#Add Go Module
go get <package-name>

#Build App
go build -o app .

#Run App in bash
./app

#Build for Linux-64
CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o app-linux .

#Build for Windows
#OPT1
sudo apt-get install gcc-mingw-w64-i686
GOOS=windows GOARCH=386 CGO_ENABLED=1 CC=i686-w64-mingw32-gcc go build -o app.exe .
#OPT2
sudo apt-get install gcc-mingw-w64-x86-64
GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc go build -o app.exe .

#generate random string
node -e "console.log(require('crypto').randomBytes(256).toString('hex'))"

#generate cryptographically secure key
openssl rand -base64 32

#embed directory
go:embed static/**