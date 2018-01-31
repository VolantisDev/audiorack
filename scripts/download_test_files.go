// +build scripts

package main

import (
	"os"
	"net/http"
	"fmt"
	"io"
	"path/filepath"
	"runtime"
	"errors"
	"github.com/mholt/archiver"
)

func main() {
	dirs := []string {
		"vst2-32bit",
		"vst2-64bit",
		"vst3-32bit",
		"vst3-64bit",
	}

	files := make(map[string]string)
	files["vst2-32bit/mda-vst-bin-win-32-2010-02-14.zip"] = "https://downloads.sourceforge.net/project/mda-vst/mda-vst/mda-vst-src%20100214/mda-vst-bin-win-32-2010-02-14.zip?r=https%3A%2F%2Fsourceforge.net%2Fprojects%2Fmda-vst%2Ffiles%2Fmda-vst%2Fmda-vst-src%2520100214%2Fmda-vst-bin-win-32-2010-02-14.zip%2Fdownload&ts=1517288875"
	files["vst2-64bit/mda-vst-bin-win-64-2010-02-14.zip"] = "https://downloads.sourceforge.net/project/mda-vst/mda-vst/mda-vst-src%20100214/mda-vst-bin-win-64-2010-02-14.zip?r=https%3A%2F%2Fsourceforge.net%2Fprojects%2Fmda-vst%2Ffiles%2Fmda-vst%2Fmda-vst-src%2520100214%2Fmda-vst-bin-win-64-2010-02-14.zip%2Fdownload&ts=1517288694"
	files["vst3-32bit/TimeLagFilter.dll"] = "https://github.com/sauraen/timelagfilter/raw/master/bin/TimeLagFilter32.dll"
	files["vst3-64bit/TimeLagFilter.dll"] = "https://github.com/sauraen/timelagfilter/raw/master/bin/TimeLagFilter.dll"

	baseDir := basePath()

	zipFiles := []string {
		filepath.Join(baseDir, "vst2-32bit/mda-vst-bin-win-32-2010-02-14.zip"),
		filepath.Join(baseDir, "vst2-64bit/mda-vst-bin-win-64-2010-02-14.zip"),
	}

	err := deleteFiles(dirs, baseDir)
	if err != nil {
		panic(err)
	}

	err = createDirectories(dirs, baseDir)
	if err != nil {
		panic(err)
	}

	for filePath, url := range files {
		err = downloadFile(filepath.Join(baseDir, filePath), url)
		if err != nil {
			panic(err)
		}
	}

	for _, zipFile := range zipFiles {
		err = extractZipFile(zipFile)
		if err != nil {
			panic(err)
		}
	}

	err = deleteFiles(zipFiles, "")
}

func basePath() string {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		panic(errors.New("Could not detect path"))
	}

	return filepath.Join(filepath.Dir(filename), "../test")
}

func deleteFiles(files []string, baseDir string) error {
	for _, dir := range files {
		err := os.RemoveAll(filepath.Join(baseDir, dir))
		if err != nil {
			return err
		}
	}

	return nil
}

func createDirectories(dirs []string, baseDir string) error {
	for _, dir := range dirs {
		os.MkdirAll(filepath.Join(baseDir, dir), os.ModePerm)
	}

	return nil
}

func extractZipFile(file string) error {
	dir := filepath.Dir(file)

	err := archiver.Zip.Open(file, dir)
	if err != nil {
		return err
	}

	return nil
}

func downloadFile(filePath string, url string) error {
	fmt.Println(filePath)
	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
