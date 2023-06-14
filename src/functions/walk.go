package functions

import (
	"checklistCGREG/src/structs"
	"fmt"
	"os"
	"path/filepath"
)

func WalkFunc(path string) structs.Data {
	data := structs.Data{}
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		temp := structs.Dir{}
		if info.IsDir() {
			temp.SetData(path, "")
			data.Folder = append(data.Folder, temp)
		} else {
			path = path[:len(path)-len(info.Name())-1]
			temp.SetData(path, info.Name())
			data.Files = append(data.Files, temp)
		}
		return nil
	})
	return data
}
