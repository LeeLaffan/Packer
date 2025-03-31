$exe = "go"
$run = @("run", ".")

$packDir = "/home/l/.local/share/lee/vimconfig/"
rm -rf $packDir/*

./publish -p win
$env:GOOS="linux" # Reset os

& $exe $run -n mason -z -d
& $exe $run -n plugins -z -d
& $exe $run -n vim -z -d

# Copy pack.exe + config.toml 
cp ./bin/pack.exe $packDir/pack.exe
cp ./config.toml $packDir/config.toml
