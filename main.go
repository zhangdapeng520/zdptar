package main

import (
	"github.com/zhangdapeng520/zdpgo_cmd"
	"github.com/zhangdapeng520/zdpgo_tar"
	"path/filepath"
)

var (
	rootCmd   = &zdpgo_cmd.Command{}
	tar       bool
	untar     bool
	directory string
	file      string
	Tar       *zdpgo_tar.Tar
)

func init() {
	rootCmd.Flags().BoolVarP(&tar, "tar", "t", false, `compress directory to *.tar.gz`)
	rootCmd.Flags().BoolVarP(&untar, "untar", "u", false, `un compress *.tar.gz to directory`)
	rootCmd.Flags().StringVarP(&directory, "dir", "d", "./", `the directory name for compress`)
	rootCmd.Flags().StringVarP(&file, "file", "f", "", `the directory name for compress`)
}

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

func handleUnTar() {
	if file == "" {
		err := Tar.UnTarGzDir(directory)
		if err != nil {
			c.Log.Error("un tar current dir error", "error", err, "directory", directory)
			return
		}
	} else {
		c.Log.Debug("un compress", "file", file)
		err := Tar.UnTarGzToSameDir(file)
		if err != nil {
			c.Log.Error("un tar file error", "error", err, "file", file)
			return
		}
	}
}

func handleTar() {
	// compress current directory all child directory
	if directory == "./" {
		err := Tar.TarGzDirAllFiles(directory)
		if err != nil {
			c.Log.Error("compress child directory error", "error", err)
		}
		return
	}

	// compress directory to *.tar.gz
	_, fileName := filepath.Split(directory)
	err := Tar.TarGz(directory, fileName+".tar.gz")
	if err != nil {
		c.Log.Error("compress directory to *.tar.gz error", "error", err, "directory", directory, "fileName", fileName)
		return
	}
	c.Log.Error("compress directory to *.tar.gz success", "directory", directory, "fileName", fileName)
}
