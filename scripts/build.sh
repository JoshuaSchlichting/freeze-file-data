if [[ "$OSTYPE" == "linux-gnu"* ]]; then
    OS=linux
elif [[ "$OSTYPE" == "darwin"* ]]; then
    OS=macos
elif [[ "$OSTYPE" == "cygwin" ]]; then
    OS=cygwin
elif [[ "$OSTYPE" == "msys" ]]; then
    # Lightweight shell and GNU utilities compiled for Windows (part of MinGW)
    OS=mysys
elif [[ "$OSTYPE" == "win32" ]]; then
    OS=windows
elif [[ "$OSTYPE" == "freebsd"* ]]; then
    OS=freebsd
else
    OS=unknown
fi

echo $OS
BIN_FILENAME=bin/freeze-file-data_$(echo $OS)_$(uname -m)
echo FILENAME = $BIN_FILENAME
go build -o $BIN_FILENAME .
chmod +x $BIN_FILENAME