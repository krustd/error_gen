# generrorx

> generrorx is a CLI tool for generating Go error code wrappers from Protocol Buffer definitions.  
> generrorx 是一个从 Protocol Buffer 定义生成 Go 错误代码包装器的命令行工具。

**即使用proto文件中定义的错误码和错误名生成对应的错误变量和错误包装函数，减少定义错误码和错误名后还需要定义go的error变量的重复工作**。

## Features / 特性

- Initialize a new error.proto template (`create`)  
  在当前目录生成 `error.proto` 模板 (`create`)
- Compile `.proto` to Go using `protoc` (`build`)  
  使用 `protoc` 将 `.proto` 文件编译为 Go 代码 (`build`)
- Generate error variables and wrapper functions (`gen`)  
  自动生成错误变量及封装函数 (`gen`)

## Prerequisites / 前置条件

- Go 1.x  
- Protocol Buffer Compiler (`protoc`)  
- Go `protoc-gen-go` plugin

## Installation / 安装

```bash
# Install as a Go CLI tool (like goctl)
go install github.com/yourusername/generrorx@latest

# For Go <1.17, use go get
go get -u github.com/yourusername/generrorx

# Run without installation
go run github.com/yourusername/generrorx create
```
Ensure `$GOPATH/bin` or `$HOME/go/bin` is in your PATH. / 确保 `$GOPATH/bin` 或 `$HOME/go/bin` 在你的 PATH 中。
## From Source / 从源码构建

> 直接从源码构建并安装，可像 `goctl` 一样将 `generrorx` 作为全局 CLI 工具使用。

```bash
# Clone the repository
git clone https://github.com/yourusername/generrorx.git
cd generrorx

# Build the CLI binary
go build -o generrorx main.go

# Install to your Go bin (e.g., $GOPATH/bin or $HOME/go/bin)
mv generrorx $GOPATH/bin

# Verify installation
generrorx --help
```

从源码克隆后编译并移动可执行文件到 `$GOPATH/bin` 或 `$HOME/go/bin`，即可像 `goctl` 一样全局使用。

## Usage / 使用方法

### 1. Create / 创建 error.proto

```bash
generrorx create
```

在当前目录生成 `error.proto` 模板文件，内容示例：

```proto
syntax = "proto3";

package errorcode;
option go_package = "./errorcode";

enum ErrorCode {
    UNKNOWN = 0; // 未知错误
}
```

### 2. Build / 编译 .proto

```bash
generrorx build [-p path/to/error.proto]
```

- `-p`, `--path`：指定 `.proto` 文件路径，默认为 `./error.proto`

生成 Go 代码至同一目录，例如 `errorcode/error.pb.go`。

### 3. Generate / 生成 Go 错误包装

```bash
generrorx gen -m <modelName> [-p <pbfile>] [-i <importPath>]
```

- `-m`, `--modelname`：生成文件的 Go 包名（必填）  
- `-p`, `--pbfile`：`.pb.go` 文件路径，默认为 `./errorcode/error.pb.go`  
- `-i`, `--importpath`：生成的 protobuf 包导入路径，默认为 `<modelName>/errorcode`

生成文件：  
- `errors_gen.go`：错误变量定义  
- `wrap.go`：包含 gRPC/HTTP 错误封装逻辑

## Project Structure / 项目结构

```
.
├── cmd
│   ├── create.go
│   ├── build.go
│   └── generate.go
├── main.go
├── go.mod
└── README.md
```

## Contributing / 贡献

欢迎提出 Issue 或提交 PR。

## License / 许可证

本项目基于 GPL-3.0 发布。 / Licensed under the GNU GPLv3 License.
