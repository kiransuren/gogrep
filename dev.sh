# Development build and run script

echo -e "-------------------------- \n"

# Build project to ./build/gogrep.exe
echo "1/2: Building project to ./build/gogrep.exe"
go build -o ./build/gogrep.exe ./cmd/gogrep.go

# Running ./build/gogrep.exe
echo -e "2/2: Running ./build/gogrep.exe \n"
echo -e "-------------------------- \n"

./build/gogrep.exe