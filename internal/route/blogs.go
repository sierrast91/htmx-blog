package route

import (
	"log"
	"os"
	"strings"
	"time"
)

type BlogItem struct {
	ID     int
	Title  string
	Author string
	Date   string
}

type BlogPage struct {
	Title     string
	Author    string
	CreatedAt time.Time
	Body      string
}

func dirToBlogList(dir string) []BlogItem {
	fs, err := os.ReadDir(dir)
	if err != nil {
		log.Println("example read:", err)
	}
	var bs []BlogItem
	for i, file := range fs {
		str := file.Name()
		str = strings.Split(str, ".")[0]
		list := strings.Split(str, "_")
		ll := len(list)
		if ll > 3 || ll == 0 {
			log.Println("file-name error")
			continue
		}
		if ll == 3 {
			var b BlogItem
			b.ID = i + 1
			b.Title, b.Author, b.Date = list[0], list[1], list[2]
			bs = append(bs, b)
		} else if ll == 2 {
			var b BlogItem
			b.ID = i + 1
			b.Title, b.Author, b.Date = list[0], list[1], ""
			bs = append(bs, b)
		} else if ll == 1 {
			var b BlogItem
			b.ID = i + 1
			b.Title, b.Author, b.Date = list[0], "", ""
			bs = append(bs, b)
		}
	}
	return bs
}
