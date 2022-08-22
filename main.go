package main

import (
	"github.com/zhangdapeng520/zdpgo_cmd"
	"github.com/zhangdapeng520/zdpgo_tar"
	"log"
	"path/filepath"
)

var (
	rootCmd   = &zdpgo_cmd.Command{}
	tar       bool           // 是否为压缩
	untar     bool           // 是否为解压缩
	directory string         // 目录
	file      string         // 文件
	Tar       *zdpgo_tar.Tar // 压缩对象
)

func init() {
	rootCmd.Flags().BoolVarP(&tar, "tar", "t", false, `compress directory to *.tar.gz`)
	rootCmd.Flags().BoolVarP(&untar, "untar", "u", false, `un compress *.tar.gz to directory`)
	rootCmd.Flags().StringVarP(&directory, "dir", "d", "./", `the directory name for compress`)
	rootCmd.Flags().StringVarP(&file, "file", "f", "", `the directory name for compress`)
}

// 主方法
func main() {
	// 创建一个根cmd对象
	rootCmd.Run = func(cmd *zdpgo_cmd.Command, args []string) {
		// create tar object
		Tar = zdpgo_tar.New()

		// handle command
		if tar {
			handleTar()
		} else if untar {
			handleUnTar()
		}
	}

	// 执行根rootCmd的命令
	rootCmd.Execute()
}

// 处理解压缩
func handleUnTar() {
	if file == "" {
		err := Tar.UnTarGzDir(directory)
		if err != nil {
			log.Println("un tar current dir error", "error", err, "directory", directory)
			return
		}
	} else {
		log.Println("un compress", "file", file)
		err := Tar.UnTarGzToSameDir(file)
		if err != nil {
			log.Println("un tar file error", "error", err, "file", file)
			return
		}
	}
}

// 处理压缩
func handleTar() {
	// compress current directory all child directory
	if directory == "./" {
		err := Tar.TarGzDirAllFiles(directory)
		if err != nil {
			log.Println("compress child directory error", "error", err)
		}
		return
	}

	// compress directory to *.tar.gz
	_, fileName := filepath.Split(directory)
	err := Tar.TarGz(directory, fileName+".tar.gz")
	if err != nil {
		log.Println("compress directory to *.tar.gz error", "error", err, "directory", directory, "fileName", fileName)
		return
	}
	log.Println("compress directory to *.tar.gz success", "directory", directory, "fileName", fileName)
}
