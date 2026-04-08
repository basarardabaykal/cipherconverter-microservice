@echo off 
echo Generating gRPC Go files...

if not exist internal\pb (
    mkdir internal\pb
)

protoc -I ./proto ^
  --go_out=. --go_opt=module=github.com/basarardabaykal/cipherconverter-microservice ^
  --go-grpc_out=. --go-grpc_opt=module=github.com/basarardabaykal/cipherconverter-microservice ^
  ./proto/*.proto

if %ERRORLEVEL% EQU 0 (
    echo Generation complete!
) else (
    echo Compilation failed. Check the errors above.
)