# Development build and run script

BUILD_PATH="./build/gogrep.exe"
MAIN_PATH="./cmd/gogrep.go"
SEARCH_PATH="./"
SEARCH_PATTERN="pattern"

echo -e "-------------------------- \n"

# Build project to BUILD_PATH
echo "1/2: Building project to $BUILD_PATH"
go build -o $BUILD_PATH $MAIN_PATH

# Running BUILD_PATH
echo -e "2/2: Running $BUILD_PATH \n"
echo -e "-------------------------- \n"

$BUILD_PATH -r $SEARCH_PATTERN $SEARCH_PATH