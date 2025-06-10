# proto-error-gen

> A CLI tool to generate Go error code wrappers from Protocol Buffer definitions.  
> 一个从 Protocol Buffer 定义生成 Go 错误代码包装器的命令行工具。

## Features / 特性

- Create new error proto file(`create`)  
  生成新错误定义proto文件（`create`）
- Compile `.proto` files to Go (`build`)  
  将 `.proto` 文件编译为 Go 代码（`build`）
- Generate error variables and wrapper functions (`gen`)  
  自动生成错误变量及封装逻辑（`gen`）


### 1. Create / 创建

```bash
generrorx create
```

在当前目录生成`error.proto`供你定义错误。

### 2. Build / 编译

```bash
generrorx build path/to/error.proto
```

根据 Proto 文件生成 Go 定义（`errorcode/error.pb.go`）。

### 3. Generate / 生成

```bash
generrorx gen \
  --modelname=<package_name> \
  --pbfile=errorcode/error.pb.go \
  --importpath=<import/path/to/errorcode>
```

- `--modelname`：生成的 Go 包名称  
- `--pbfile`：已生成的 `.pb.go` 文件路径  
- `--importpath`：Protobuf Go 包的导入路径  

将生成 `errors_gen.go` 和 `wrap.go`，包含错误变量和封装函数。

## Project Structure / 项目结构

```
.
├── cmd
│   ├── create.go
│   ├── build.go
│   └── generate.go
├── error.proto
├── errorcode
│   └── error.pb.go
├── main.go
├── go.mod
└── README.md
```

## Contributing / 贡献

欢迎贡献！请打开 Issue 或提交 Pull Request。  
Contributions are welcome! Please open an Issue or submit a Pull Request.

## License / 许可证

本项目依据 GPL-3.0 许可证发布。  
This project is licensed under the GNU GPLv3 License.
