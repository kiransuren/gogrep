# Production development script

BUILD_PATH="./build/gogrep.exe"
MAIN_PATH="./cmd/gogrep.go"

# Build project to BUILD_PATH
echo -e "-------------------------- "
echo "Building project to $BUILD_PATH"
echo -e "-------------------------- \n"
go build -o $BUILD_PATH $MAIN_PATH