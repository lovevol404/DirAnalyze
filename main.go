package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strconv"
)

const (
	KB = 1024.0
	MB = 1024 * KB
	GB = 1024 * MB
)

type DirInfo struct {
	IsDir             bool
	Name, Path, sizeH string
	size              int64
	subDirs           DirInfoList
}

type DirInfoList []*DirInfo

func (I DirInfoList) Len() int {
	return len(I)
}
func (I DirInfoList) Less(i, j int) bool {
	return I[i].size < I[j].size
}
func (I DirInfoList) Swap(i, j int) {
	I[i], I[j] = I[j], I[i]
}

var dirMap  map[string]*DirInfo

func getDirInfo(path, name string) *DirInfo {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return nil
	}
	if fileInfo.IsDir() {
		dir := &DirInfo{
			IsDir: true,
			Name:  name,
			Path:  path,
		}
		var subDirs DirInfoList

		files, err := ioutil.ReadDir(path)
		if err != nil {
			panic(err)
		}
		for _, f := range files{
			subDir := getDirInfo(filepath.Join(path, f.Name()), f.Name())
			if subDir != nil {
				subDirs = append(subDirs, subDir)
			}
		}
		dir.subDirs = subDirs
		getAndSetSize(dir)
		sort.Sort(sort.Reverse(subDirs))
		dirMap[path] = dir
		return dir
	} else {
		return &DirInfo{
			IsDir:   false,
			Name:    name,
			Path:    path,
			size:    fileInfo.Size(),
			sizeH:   getSizeH(fileInfo.Size()),
			subDirs: nil,
		}
	}
}

func getSizeH(size int64) string {
	var f float64
	var subfix string
	if size > GB {
		f = float64(size) / GB
		subfix = "G"
	}else if size > MB {
		f = float64(size) / MB
		subfix = "M"
	}else if size > KB {
		f = float64(size) / KB
		subfix = "K"
	}else {
		f = float64(size)
		subfix = "B"
	}
	return strconv.FormatFloat(f, 'f', 3, 64) + subfix
}

func getAndSetSize(info *DirInfo) int64 {
	if info.size != 0{
		return info.size
	}
	if info.IsDir {
		if info.subDirs == nil {
			info.size = 0
			info.sizeH = "0"
			return 0
		}
		var size int64
		for _, i := range info.subDirs{
			size = getAndSetSize(i) + size
		}
		info.size = size
		info.sizeH = getSizeH(size)
		return size
	}else {
		return info.size
	}
}

func print(info *DirInfo, prefix string, level, currentLevel int) {
	fmt.Println(prefix + info.Name + " " + info.sizeH)
	prefix = "     "+ prefix
	currentLevel ++
	if currentLevel > level {
		return
	}
	if info.subDirs == nil {
		return
	}
	for  _,i := range info.subDirs{
		print(i, prefix, level, currentLevel)
	}
}

func main() {
	dirMap = make(map[string]*DirInfo)
	d := getDirInfo("/Users/tianming/foobar", "/Users/tianming/foobar")
	if d != nil {
		fmt.Println("err")
	}
	print(d, "", 3, 1)
}
