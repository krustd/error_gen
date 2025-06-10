package cmd

import (
	"fmt"
	"os"
)

func RunCreate() {
	const protoPath = "./error.proto"
	if _, err := os.Stat(protoPath); err == nil {
		fmt.Println("error.proto 已存在，跳过")
		return
	}
	content := `syntax = "proto3";
package errorcode;
option go_package = "./errorcode";

enum ErrorCode {
    UNKNOWN = 0; // 未知错误
}
`
	err := os.WriteFile(protoPath, []byte(content), 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("✅ 已创建:", protoPath)
}
