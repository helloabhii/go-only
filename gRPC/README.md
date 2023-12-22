
###install gRPC in Arch Linux

##Install Protocol Buffers:

sudo pacman -S protobuf

##Install gRPC:

yay -S grpc

##Install gRPC Tools (Optional but recommended):

yay -S grpcurl

##Verify Installation:

grpcurl --version
protoc --version

##file extension

.proto

##vs code extension 

vscode-proto (sanketh.dev)
