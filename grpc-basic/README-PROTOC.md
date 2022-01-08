### .proto >> .pb.go
```
$ cd path/to/file
$ protoc {PACKAGE_FOLDER_NAME}/{FILE_NAME.proto} --go_out=plugins=grpc:.
```
```
$ cd Golang/basic/go-grpc
$ protoc greetpb/greet.proto --go_out=plugins=grpc:.
```

### possible error 1
protoc-gen-go: invalid Go import path .. The import path must contain at least one forward slash ('/') character.

### solution 1
in .proto file, set option go_package = "./{PACKAGE_FOLDER_NAME}"

```
syntax = "proto3";

package greet;
option go_package="./greetpb";
```

https://stackoverflow.com/questions/68637511/cant-generate-grpc-files-for-go