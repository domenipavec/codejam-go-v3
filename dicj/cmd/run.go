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
	"log"
	"os"
	"os/exec"
	"os/user"
	"path"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Runs locally with example.in",
	Run: func(cmd *cobra.Command, args []string) {
		inputFile := "example.in"
		if len(args) > 0 {
			inputFile = args[0]
		}
		usr, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}
		downloadFiles, err := filepath.Glob(path.Join(usr.HomeDir, "Downloads/*.in"))
		if err != nil {
			log.Fatal(err)
		}
		if len(downloadFiles) > 1 {
			log.Fatal("Multiple input files in Downloads")
		} else if len(downloadFiles) == 1 {
			inputFile = filepath.Base(downloadFiles[0])
			os.Rename(downloadFiles[0], inputFile)
		}

		if !strings.HasSuffix(inputFile, ".in") {
			if inputFile[len(inputFile)-1] != '.' {
				inputFile += ".in"
			} else {
				inputFile += "in"
			}
		}

		log.Println("Using input file:", inputFile)

		runCmd := "go run *.go < " + inputFile
		outputFile := strings.TrimSuffix(inputFile, ".in") + ".out"
		if inputFile != "example.in" {
			runCmd += " > " + outputFile
		}

		goCmd := exec.Command("bash", "-c", runCmd)
		goCmd.Stderr = os.Stderr
		goCmd.Stdout = os.Stdout
		err = goCmd.Run()
		if err != nil {
			log.Fatal(err)
		}

		if inputFile != "example.in" {
			copyFile(outputFile, path.Join(usr.HomeDir, "Downloads", outputFile))
		}
	},
}

func init() {
	RootCmd.AddCommand(runCmd)
}
