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
			cmd.RunBuild(path)
		},
	}
	buildCmd.Flags().String("path", "", "proto文件路径")
	buildCmd.MarkFlagRequired("path")

	var genCmd = &cobra.Command{
		Use:   "gen",
		Short: "生成错误代码",
		Run: func(c *cobra.Command, args []string) {
			modelName, _ := c.Flags().GetString("modelname")
			pbFile, _ := c.Flags().GetString("pbfile")
			importPath, _ := c.Flags().GetString("importpath")
			// print(pbFile, modelName, importPath)
			cmd.RunGenerate(modelName, pbFile, importPath)
		},
	}
	genCmd.Flags().String("modelname", "", "所属包名称")
	genCmd.MarkFlagRequired("modelname")
	genCmd.Flags().String("pbfile", "", "pb.go 文件路径")
	genCmd.MarkFlagRequired("pbfile")
	genCmd.Flags().String("importpath", "", "导入的pb文件包名")
	genCmd.MarkFlagRequired("importpath")

	rootCmd.AddCommand(createCmd, buildCmd, genCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
