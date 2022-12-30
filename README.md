# LocalFileServer

#### Introduction

In order to realize the file transfer in the local network, it works for Linux, Windows, Mac os, through http server to download the files to mobile, other PC.



#### Install

You can download the release binary file, also you can download the source code to build. But I suggestion the second method.

Before you build the file, you should install Golang environment.

```bash
go mod init FileServer
go mod tidy
go build -o fs .
```

Or you can also add it to the PATH, there will be easier to use.

```bash
export PATH=$PATH:$(pwd)
```



#### How to use

This tool uses the -t parameter to specify the root directory of the file server. 

```bash
fs -t "/home/tumble-leap/Documents"
```

If this parameter is not specified, the current directory when starting the program is used by default

```bash
fs # The root dir is current dir
```

