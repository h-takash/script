package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

func main() {
	// リネーム用のフォルダ「rename」のパスを取得
	dir := "rename"

	// リネーム用のフォルダ内のファイル一覧を取得
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("ファイル一覧取得エラー:", err)
		return
	}

	if len(files) == 0 {
		fmt.Println("renameフォルダにファイルがありません")
		return
	}

	// ファイルを名前順にソート
	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})

	// ファイルごとにリネーム
	for i, file := range files {
		oldPath := filepath.Join(dir, file.Name())
		ext := filepath.Ext(file.Name())
		newName := fmt.Sprintf("%s_%03d%s", file.Name()[:len(file.Name())-len(ext)], i+1, ext)
		newPath := filepath.Join(dir, newName)

		err := os.Rename(oldPath, newPath)
		if err != nil {
			fmt.Println("リネームエラー:", err)
			return
		}
	}

	fmt.Println("リネーム成功")
}
