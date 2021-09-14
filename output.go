package student

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var words string
var font string
var flag string
var data []string
var newSplitWord []string
var text string

func Output(str1, str2 string) string {
	words, font, text, data, newSplitWord = "", "", "", nil, nil
	words = str1
	font = str2
	ReadFile()
	return AsciiArt(newSplitWord, data)
}

func ReadFile() {
	file, err := os.Open(font + ".txt")
	defer file.Close()

	if err != nil {

		log.Println(err.Error())
		return
	}
	newSplitWord = strings.Split(words, "\\n")

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
}

func AsciiArt(words []string, font []string) string {
	var fil *os.File
	var er error
	if flag != "" {
		fil, er = os.Create(flag + ".txt")
		defer fil.Close()
		if er != nil {
			log.Println(er.Error())
			return ""
		}
	}
	for i := 0; i < len(words); i++ {
		word := ""
		word = words[i]
		for index1 := 0; index1 < 8; index1++ {
			line := ""
			if word != "" {
				for index2 := 0; index2 < len(word); index2++ {
					line = line + font[int((word[index2]-32))*9+index1+1]
				}
			} else {
				if flag != "" {
					fil.Write([]byte("\n"))
				} else {
					fmt.Println(line)
				}
				break
			}
			if flag != "" {
				fil.Write([]byte(line))
				fil.Write([]byte("\n"))
			} else {
				text = text + line
				text = text + "\n"
			}
		}
	}
	return text
}
