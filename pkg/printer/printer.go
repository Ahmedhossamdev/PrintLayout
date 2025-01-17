package printer

import (
	"fmt"
	"os"
	"path/filepath"
)

func PrintProjectStructure() {
	root := "."

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if root == path {
			return nil
		}

		relPath, _ := filepath.Rel(root, path)

		fmt.Println(relPath)

		// parts := strings.Split(relPath, string(filepath.Separator))

		return nil
	})

	if err != nil {
		fmt.Println(err)
	}

}
