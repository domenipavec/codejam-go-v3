// Copyright © 2018 Domen Ipavec <domen@ipavec.net>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"go/build"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:     "init folder",
	Args:    cobra.ExactArgs(1),
	Aliases: []string{"i"},
	Short:   "Initialize folder with code.",
	Run: func(cmd *cobra.Command, args []string) {
		gopath := os.Getenv("GOPATH")
		if gopath == "" {
			gopath = build.Default.GOPATH
		}

		destpath := args[0]
		err := os.MkdirAll(destpath, 0755)
		if err != nil {
			log.Fatal(err)
		}

		sourcepath := path.Join(gopath, "src/github.com/matematik7/codejam-go-v3/source/")

		files, err := ioutil.ReadDir(sourcepath)
		if err != nil {
			log.Fatal(err)
		}

		for _, f := range files {
			sourcefile := path.Join(sourcepath, f.Name())
			destfile := path.Join(destpath, f.Name())
			if strings.Contains(f.Name(), "_test") {
				continue
			}
			if (f.Name() == "algorithm.go" || f.Name() == "example.in") && existsFile(destfile) {
				continue
			}
			log.Printf("Copying %s", f.Name())
			err = copyFile(sourcefile, destfile)
			if err != nil {
				log.Fatal("Error copying file: ", err)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func copyFile(source, dest string) error {
	sourceFile, err := os.Open(source)
	if err != nil {
		return err
	}
	defer sourceFile.Close()
	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return err
	}

	return destFile.Sync()
}

func existsFile(fn string) bool {
	_, err := os.Stat(fn)
	return err == nil
}
