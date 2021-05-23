package typectors

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func copyDirectory(sourceRoot string, destinationRoot string) error {
	fmt.Println("Generated source files:")
	return filepath.WalkDir(sourceRoot, func(source string, entry fs.DirEntry, err error) error {
		if err == nil {
			relativePath := strings.Replace(source, sourceRoot, "", 1)
			dest := filepath.Join(destinationRoot, relativePath)
			// if the destination already exists and either it or the source is a file, remove it
			var stat fs.FileInfo
			if stat, err = os.Stat(dest); err == nil && (!stat.IsDir() || !entry.IsDir()) {
				removeErr := os.Remove(dest)
				if removeErr != nil {
					log.Printf(`unable to remove the file "%s" to replace it with the new generated one: %+v`, dest, removeErr)
				}
			}
			if !entry.IsDir() {
				err = copyFile(source, dest)
				if err == nil {
					log.Printf("%s\n", dest)
				}
			} else if err != nil { // here `err != nil` if the destination doesn't exist, or can't be read
				err = os.Mkdir(dest, os.ModeDir|0664)
			}
		}
		return err
	})
}

func copyFile(source string, destination string) error {
	content, err := ioutil.ReadFile(source)
	if err == nil {
		err = ioutil.WriteFile(destination, content, 0664)
	}
	return err
}
