package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// runBuild 编译指定路径的 .proto 文件（例如：errorx/error.proto）
func RunBuild(protoPath string) {
	// 1. 获取 .proto 文件的目录
	protoAbsPath, err := filepath.Abs(protoPath)
	if err != nil {
		panic("❌ 无法解析 proto 路径: " + err.Error())
	}

	// 2. 确保文件存在
	if _, err := os.Stat(protoAbsPath); os.IsNotExist(err) {
		panic("❌ proto 文件不存在: " + protoAbsPath)
	}

	protoDir := filepath.Dir(protoAbsPath)

	// 3. 构建 protoc 命令
	cmd := exec.Command("protoc",
		fmt.Sprintf("--go_out=%s", protoDir),
		fmt.Sprintf("--proto_path=%s", protoDir),
		protoAbsPath,
	)
	println(cmd.String())
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Printf("🚀 编译 proto: %s\n", protoAbsPath)
	if err := cmd.Run(); err != nil {
		panic("❌ protoc 编译失败: " + err.Error())
	}

	fmt.Println("✅ 已生成 .go 文件")
}
