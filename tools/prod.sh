# Production development script
# Build project to ./build/gogrep.exe
echo -e "--------------------------"
echo "Building project to ./build/gogrep.exe"
echo -e "-------------------------- \n"

go build -o ./build/gogrep.exe ./cmd/gogrep.go
