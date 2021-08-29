package search

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"time"
)

func isDir(file string) (bool, error) {
	maybeDir, err := os.Stat(file)
	if err != nil {
		return false, err
	}
	return maybeDir.IsDir(), nil
}

func Search(pattern string, fileName string) error {
	start := time.Now()
	matches := 0
	reg, err := regexp.Compile(pattern)
	if err != nil {
		return err
	}
	isDir, err := isDir(fileName)
	if err != nil {
		return err
	}
	if isDir {
		matches, err = searchDir(reg, fileName)
		if err != nil {
			return err
		}
	} else {
		matches, err = searchFile(reg, fileName)
		if err != nil {
			return err
		}
	}
	fmt.Println("\nMatches:", matches)
	fmt.Println("Elapsed time:", time.Since(start))
	return nil
}

func searchDir(pattern *regexp.Regexp, dirName string) (int, error) {
	var matches int
	err := searchDirRec(pattern, dirName, &matches)
	if err != nil {
		return matches, err
	}
	return matches, nil
}

func searchDirRec(pattern *regexp.Regexp, dirName string, matches *int) error {
	dir, err := os.ReadDir(dirName)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	if err != nil {
		fmt.Printf("ERROR: can't read: %s: %s\n", dirName, err)
	}
	for i := range dir {
		if dir[i].IsDir() {
			err = searchDirRec(pattern, dirName+string(os.PathSeparator)+dir[i].Name(), matches)
			if err != nil {
				return err
			}
		} else {
			if pattern.MatchString(dir[i].Name()) {
				*matches++
				//fmt.Println(dirName + string(os.PathSeparator) + dir[i].Name())
				fmt.Fprintln(f, dirName+string(os.PathSeparator)+dir[i].Name())
			}
		}
	}
	return nil
}

func searchFile(pattern *regexp.Regexp, fileName string) (int, error) {
	file, err := os.Open(fileName)
	matches := 0
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		if pattern.MatchString(scanner.Text()) {
			matches++
			fmt.Println(scanner.Text())
		}
	}
	return matches, nil
}
