package main

import (
	"strings"
)



func main() {
	//fileName := "C:\\liuran\\tmp\\HashMap.java"

	//wordsSlice := text.ReadWords(fileName)
	//
	//for _, w := range wordsSlice {
	//	fmt.Printf("%s:%d\n", w.Word(), w.Count())
	//

}

func IsEmpty(str string) bool {
	if strings.TrimSpace(str) == ""{
		return true
	}

	return false
}
