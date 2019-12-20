# check sam for valid template
sam validate

### generator

# copy the generator source code
cp lambda/generator/main.go .

# build the go binary for linux that can run on lambda
GOOS=linux go build -ldflags="-s -w" main.go 

# compress the go binary, which reduces size by ~35%
cmd=$(which upx)

if [ ${#cmd} -ne '0' ]; then
    time upx -9 main
fi

# move the compressed binary to the lambda dir
mv main lambda/generator/

### backend

# copy the backend source code
cp lambda/backend/main.go .

# build the go binary for linux that can run on lambda
GOOS=linux go build -ldflags="-s -w" main.go 

# compress the go binary, which reduces size by ~35%

if [ ${#cmd} -ne '0' ]; then
    time upx -9 main
fi

# move the compressed binary to the lambda dir
mv main lambda/backend/

# remove the main.go file
rm main.go

# deploy sam
sam deploy -g