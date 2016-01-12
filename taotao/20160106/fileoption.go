package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"time"
)

func main() {

}

func copyfile(filename1 string, filename2 string, limitsize int64) (int, error) {
	//记录开始时间
	start := time.Now().UnixNano()

	fileInfo, err := os.Stat(filename1)
	if err != nil {
		log.Printf(err)
		return _, err
	}

	//文件大小除以限制大小,然后取整
	num := int(math.Ceil(float64(fileInfo.Size()) / float64(limitsize)))

	fin1, err := os.OpenFile(filename1, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Printf(err)
		return _, err
	}
	defer fin1.Close()

	b := make([]byte, limitsize)
	var i int64 = 1
	for ; i <= int64(num); i++ {

		fin1.Seek((i-1)*(limitsize), 0)

		if len(b) > int((fileInfo.Size() - (i-1)*limitsize)) {
			b = make([]byte, fileInfo.Size()-(i-1)*limitsize)
		}

		fin1.Read(b)

		f, err := os.OpenFile("./"+strconv.Itoa(int(i))+".mem", os.O_CREATE|os.O_WRONLY, os.ModePerm)
		if err != nil {
			log.Println(err)
			return _, err
		}
		f.Write(b)
		f.Close() //放在write前使用defer,提示在循环内~~~
		time.Sleep(time.Second)
	}

	fii, err := os.OpenFile(filename2, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		log.Println(err)
		return _, err
	}
	for i := 1; i <= num; i++ {
		f, err := os.OpenFile("./"+strconv.Itoa(int(i))+".db", os.O_RDONLY, os.ModePerm)
		if err != nil {
			log.Println(err)
			return _, err
		}
		b, err := io.Reader() //ioutil.ReadAll(f)
		if err != nil {
			log.Println(err)
			return _, err
		}
		fii.Write(b)
		f.Close()
	}

	end := time.Now().UnixNano()

	fmt.Printf(end - start)
}
