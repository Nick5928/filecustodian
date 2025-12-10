package commands

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

type walkDirWrapper struct {
	totalDiskUsage int64
	sizeMap        map[string]int64
	parentDir      string
}

func (w *walkDirWrapper) visit(path string, d fs.DirEntry, err error) error {
	if err != nil {
		if os.IsPermission(err) {
			return fs.SkipDir
		}
		return err
	}
	relPath := strings.Replace(path, fmt.Sprintf("%s/", w.parentDir), "", 1)

	oneChildDownPath := strings.Split(relPath, string(os.PathSeparator))[0]
	info, err := d.Info()
	size := info.Size()
	blockSize := int64(4096)
	if err != nil {
		return err
	}
	//Block sizes cant be half so take the ceiling
	directorySize := (((size + blockSize - 1) / blockSize) * blockSize)
	_, ok := w.sizeMap[oneChildDownPath]
	if !ok {
		w.sizeMap[oneChildDownPath] = directorySize
	} else {
		w.sizeMap[oneChildDownPath] += directorySize
	}
	w.totalDiskUsage += directorySize
	return nil
}
func calcSize(args ...string) error {
	directory := "."
	if len(args) >= 1 {
		directory = args[0]
	}

	wrapper := walkDirWrapper{
		totalDiskUsage: 0,
		sizeMap:        make(map[string]int64),
		parentDir:      directory,
	}
	err := filepath.WalkDir(directory, wrapper.visit)
	if err != nil {
		return err
	}

	paths := sortedPaths(wrapper.sizeMap)
	n := min(len(paths)-1, 10)
	topNPaths := paths[:n]
	fmt.Printf("Total Space Used for %s: %.2fGB\n", directory, bytesToGb(wrapper.totalDiskUsage))
	fmt.Println("Top 10 Directories with most usage")

	//base padding for length of Column headers
	longestPathSize := 8
	sizePadding := 8
	for _, path := range topNPaths {
		longestPathSize = max(len(path), longestPathSize)
		// size := wrapper.sizeMap[path]
		// sizeFormatString := fmt.Sprintf("%.2fGB", bytesToGb(size))
		// sizePadding = max(sizePadding, len(sizeFormatString))
	}

	fmt.Printf("%-*s  %*s  %6s\n", longestPathSize, "PATHNAME", sizePadding, "TOTAL SIZE", "USAGE")
	for _, path := range topNPaths {
		size := wrapper.sizeMap[path]
		usage := (float64(size) / float64(wrapper.totalDiskUsage)) * 100
		fmt.Printf("%-*s  %*.2fGB  %5.2f%%\n", longestPathSize, path, sizePadding, bytesToGb(size), usage)
	}

	return nil
}
func getSizeMap(directory string) (map[string]int64, error) {
	du_args := []string{"-b", "-d", "1", directory}
	cmd := exec.Command("du", du_args...)
	result, _ := cmd.Output()
	directories := strings.Split(string(result), "\n")
	size_map := make(map[string]int64)
	for _, directory := range directories[:len(directories)-1] {
		res := strings.Split(directory, "\t")
		bytesString, path := res[0], res[1]
		size, err := strconv.ParseInt(bytesString, 10, 64)
		if err != nil {
			return nil, err
		}
		size_map[path] = size

	}

	return size_map, nil
}

func sortedPaths(sizeMap map[string]int64) []string {
	paths := make([]string, 0, len(sizeMap))

	for path := range sizeMap {
		paths = append(paths, path)
	}

	sort.Slice(paths, func(i, j int) bool { return sizeMap[paths[i]] > sizeMap[paths[j]] })

	return paths
}

func bytesToGb(bytes int64) float64 {
	return float64(bytes) / 1000000000
}
