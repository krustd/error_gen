package main

import (
	"fmt"
	"os"

	"generrorx/cmd"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "generrorx",
		Short: "gen_errorx CLI tool",
	}

	var createCmd = &cobra.Command{
		Use:   "create",
		Short: "创建文件",
		Run: func(c *cobra.Command, args []string) {
			cmd.RunCreate()
		},
	}

	var buildCmd = &cobra.Command{
		Use:   "build [proto file]",
		Short: "编译 proto 文件",
		Run: func(c *cobra.Command, args []string) {
			path, _ := c.Flags().GetString("path")
			if path == "" {
				path = "./error.proto"
			}
			cmd.RunBuild(path)
		},
	}
	buildCmd.Flags().StringP("path", "p", "", "proto 文件路径（可选，默认 ./error.proto）")

	var genCmd = &cobra.Command{
		Use:   "gen",
		Short: "生成错误代码",
		Run: func(c *cobra.Command, args []string) {
			modelName, _ := c.Flags().GetString("modelname")
			pbFile, _ := c.Flags().GetString("pbfile")
			if pbFile == "" {
				pbFile = "./errorcode/error.pb.go"
			}
			importPath, _ := c.Flags().GetString("importpath")
			if importPath == "" {
				importPath = modelName + "/errorcode"
			}
			cmd.RunGenerate(modelName, pbFile, importPath)
		},
	}
	genCmd.Flags().StringP("modelname", "m", "", "当前目录所属包名")
	genCmd.MarkFlagRequired("modelname")
	genCmd.Flags().StringP("pbfile", "p", "", ".pb.go 文件路径（可选，默认 ./errorcode/error.pb.go）")
	genCmd.Flags().StringP("importpath", "i", "", ".pb.go文件所属的包名（可选，默认 当前目录所属报名/errorcode）")

	var auto = &cobra.Command{
		Use:   "auto",
		Short: "全流程共同进行",
		Run: func(c *cobra.Command, args []string) {
			path, _ := c.Flags().GetString("path")
			if path == "" {
				path = "./error.proto"
			}
			cmd.RunBuild(path)
			modelName, _ := c.Flags().GetString("modelname")
			pbFile := "./errorcode/error.pb.go"
			importPath := modelName + "/errorcode"
			cmd.RunGenerate(modelName, pbFile, importPath)
		},
	}
	auto.Flags().StringP("path", "p", "", "proto 文件路径（可选，默认 ./error.proto）")
	auto.Flags().StringP("modelname", "m", "", "当前目录所属包名")
	auto.MarkFlagRequired("modelname")
	rootCmd.AddCommand(createCmd, buildCmd, genCmd, auto)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
