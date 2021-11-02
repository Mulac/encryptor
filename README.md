# encryptor

**warning:** this is not secure!  Server logs all messages and TLS is not enabled.

README is purposely short: directory structure, comments and the code itself should be immediatley understandable.

First time using gRPC (it's awesome!), so appreciate any feedback with how you'd set it up differently.

## Prerequisites
- Linux (not tested on anything else)
- Go (1.17)

## Testing
`go test ./...`

## Run the server
`go run server/main.go`

## Run the client
`go run client/main.go -key 3 decrypt "dsslwdo iru wkh zlq"`
