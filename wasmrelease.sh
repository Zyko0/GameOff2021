#cp $(go env GOROOT)/misc/wasm/wasm_exec.js .
GOOS=js GOARCH=wasm go build -o out.wasm .
#$Env:GOOS="js"; $Env:GOARCH="wasm"; go build -o out.wasm