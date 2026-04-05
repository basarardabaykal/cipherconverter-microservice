@echo off 
echo Generating gRPC Go files...

if not exist internal\pb (
    mkdir internal\pb
)

protoc -I ./proto ^
  --go_out=./internal/pb --go_opt=paths=source_relative ^
  --go-grpc_out=./internal/pb --go-grpc_opt=paths=source_relative ^
  ./proto/cipher.proto

if %ERRORLEVEL% EQU 0 (
    echo Generation complete!
) else (
    echo Compilation failed. Check the errors above.
)