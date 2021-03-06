# check sam for valid template
sam validate

#

### Generator

# copy the generator source code
cp lambda/generator/main.go .

# build the go binary for linux that can run on lambda
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" main.go 

# compress the go binary, which reduces size by ~35%
cmd=$(which upx)

if [ ${#cmd} -ne '0' ]; then
    time upx -9 main
fi

# move the compressed binary to the lambda dir
mv main lambda/generator/

#

### HTTP

# copy the backend source code
cp lambda/http/main.go .

# build the go binary for linux that can run on lambda
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" main.go 

# compress the go binary, which reduces size by ~35%

if [ ${#cmd} -ne '0' ]; then
    time upx -9 main
fi

# move the compressed binary to the lambda dir
mv main lambda/http/

#

### sqs

# copy the backend source code
cp lambda/sqs/main.go .

# build the go binary for linux that can run on lambda
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" main.go 

# compress the go binary, which reduces size by ~35%

if [ ${#cmd} -ne '0' ]; then
    time upx -9 main
fi

# move the compressed binary to the lambda dir
mv main lambda/sqs/

###

# remove the main.go file
rm main.go

# deploy sam, check if samconfig.toml file is present
if [ ! -f samconfig.toml ]; then
    echo "no samconfig.toml found, starting guided deploy"
    sam deploy -g
else
    echo "samconfig.toml found, proceeding to deploy"
    sam deploy
fi
