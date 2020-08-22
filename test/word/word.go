package text

import (
	"os"
	"bufio"
	"io"
	"strings"
)

var symbol = []string{"*", "(",")","/", ".", "@", "<", ">",",",":",";","{","}"}

type Text []string

type Point struct {
	X int
	Y int
}

type Word struct {
	Point
	word string
	count int
}

func (w *Word) SetWord(word string) *Word {
	w.word = word
	return w
}

func (w *Word) SetCount(count int) *Word {
	w.count = count

	return w
}

func (w *Word) Word() string{
	return w.word
}

func (w *Word) Count() int{
	return w.count
}

type WordSlice []Word


func (w WordSlice) Len() int {
	return len(w)
}

func (w WordSlice) Less(i, j int) bool {
	return w[i].count > w[j].count
}

func (w WordSlice) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}

func readAllWord(fileName string, lineCall func(line string) []string) WordSlice {
	file, _:= os.Open(fileName)
	reader := bufio.NewReader(file)

	var ws WordSlice
	lineNum := 0
	for {
		line, err := reader.ReadString('\n')
		if err != nil && err == io.EOF {
			 break
		}
		if lineCall == nil {
			continue
		}

		for idx, val := range lineCall(line) {
			ws = append(ws, Word{
				word:val,
				count:1,
				Point :Point{
					lineNum,idx,
				},
			})
		}
		lineNum ++
	}

	return ws
}

func getLineWords(line string) (words []string){
	line = strings.TrimSpace(line)

	if !strings.HasPrefix(line, "*") {
		return nil
	}
	for _, val := range symbol {
		line = strings.ReplaceAll(line, val, " ")
	}

	for _, val := range strings.Split(line, " ")  {
		if isEmpty(val) {
			continue
		}
		words = append(words, val)
	}

	return
}

func ReadWords(fileName string) WordSlice {
	return readAllWord(fileName, getLineWords)
}

func isEmpty(str string) bool{
	if strings.TrimSpace(str) == "" {
		return true
	}

	return false
}


