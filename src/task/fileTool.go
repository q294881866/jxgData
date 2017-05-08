package task

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync/atomic"
	"time"
	"strconv"
)

// 当前目录 ./data
var RootDir string

func SetPath(path string) {
	length := len(path)
	if length > 0 {
		// tmp
		if strings.HasSuffix(path, string(os.PathSeparator)) {
			RootDir = string([]rune(path)[0 : length-1])
		}
		fmt.Println(RootDir)
		return
	}

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	RootDir = dir + string(os.PathSeparator) + "data"
	fi, err := os.Stat(RootDir)
	if err != nil {
		os.Mkdir(RootDir, os.ModePerm)
	}
	fmt.Println(fi)
}

// 去掉最后的路径分隔符如：/
func GetPath(dir string) string {
	var root string
	if os.IsPathSeparator('\\') {
		i := strings.LastIndex(dir, "\\")
		root = string([]rune(dir)[0:i])
	} else {
		i := strings.LastIndex(dir, "/")
		root = string([]rune(dir)[0:i])
	}
	return root
}

func CreateFile(name string) (*os.File, error) {
	filename := RootDir + string(os.PathSeparator) + name

	if checkFileIsExist(name) {
		return os.Open(filename)
	} else {
		return os.Create(filename)
	}
}

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

// 生产文件名 时间戳
func filename(now time.Time) string {
	s := Timestamp(now)
	name := strings.TrimSpace(s)
	ss := strings.Split(s, " ")
	name = ss[0] + "-" + ss[1]
	return name
}

// 开始到现在经过多少分钟
var FileCount uint64 = 0

// 每分钟 异步进行 将 .csb.bak 文件改为 .bak 文件
func WritedAndRename() {
	go func() {
		dater := time.NewTicker(time.Minute * 1)
		for _ = range dater.C {
			fmt.Println("rename file task open")
			dir, _ := ioutil.ReadDir(RootDir)
			for _, fi := range dir {
				name := RootDir + string(os.PathSeparator) + fi.Name()
				reName := strings.TrimSuffix(name, ".bak")
				if fi.IsDir() { // 忽略目录
					continue
				}
				if strings.HasSuffix(name, ".csv.bak") { //匹配文件
					os.Rename(name, reName)
				}
			}
		}
	}()
}

// 每分钟生产数据
func MiTask() {
	now := GetCurrentTime()
	name := filename(now)
	name = strings.Replace(name, ":", "_", -1)
	f, err := CreateFile(name + ".csv")
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}
	w := bufio.NewWriter(f)
	// 一分钟
	for j := 0; j < 60; j++ {
		nowSecond := AddSecond(j, now)
		// 5000台机器
		for i := 1000; i < 6000; i++ {
			var id = "ZX" + fmt.Sprint(i) + "X,"
			data := DataProducter(nowSecond)
			s := Data2String(",", data)
			if len(s) != 0 {
				w.WriteString(id + s + "\n")
			}
		}
	}
	w.Flush()
}

// 当前时间加一秒
func AddSecond(i int, now time.Time) time.Time {
	newTime := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), time.Now().Nanosecond(), time.Local)
	second := int64(i) * 1000000000
	return newTime.Add(time.Duration(second))
}

// 当前时间加一分钟
// 并持久化到文件 起始时间到现在经过的分钟整数
func WriteCurrentTime() {
	abort := getAbort()
	f, err := os.OpenFile(abort, os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("create abort file")
		f, err = os.Create(abort)
	}
	defer f.Close()
	f.WriteString(fmt.Sprint(FileCount))
	f.Sync()
}

func getAbort() string {
	index := strings.LastIndex(RootDir, string(os.PathSeparator))
	parentDir := string([]rune(RootDir)[0:index])
	return parentDir + string(os.PathSeparator) + "abort"
}

//先读是否有abort文件
func readAbort() {
	abort := getAbort()
	if checkFileIsExist(abort) {
		f, _ := os.OpenFile(abort, os.O_RDWR, 0666)
		ret, _ := f.Seek(0, os.SEEK_END)
		bs := make([]byte, ret)
		f.ReadAt(bs, 0)
		s := strings.TrimSpace(string(bs))
		count, _ := strconv.ParseUint(s, 10, 0)
		FileCount = count
	}
}

func GetCurrentTime() time.Time {
	fmt.Println(FileCount)
	if FileCount > 50400 {
		fmt.Println(time.Now())
	}
	WriteCurrentTime()
	defer atomic.AddUint64(&FileCount, 1)
	return BeginTime.Add(time.Duration(60000000000 * FileCount))
}
