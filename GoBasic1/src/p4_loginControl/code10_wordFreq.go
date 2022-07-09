package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"
)

func WordFreq() {
	if len(os.Args) == 1 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Printf("usage:%s <file>[<file2>[...<fileN>]]\n",
			filepath.Base(os.Args[0]))
		os.Exit(0)
	}

	frequencyForWord := map[string]int{}

	for _, filename := range commandLineFiles(os.Args[1:]) {
		updateFrequencies(filename, frequencyForWord)
	}
	reportByWords(frequencyForWord)

	wordsForFrequency := invertStringIntMap(frequencyForWord)
	reportByFrequency(wordsForFrequency)
}

/**
实现mac、linux以及windows之间的文件通配符 跨平台处理
*/
func commandLineFiles(files []string) []string {
	if runtime.GOOS == "windows" {
		args := make([]string, 0, len(files))
		for _, name := range files {
			if matches, err := filepath.Glob(name); err != nil {
				args = append(args, name)
			} else if matches != nil {
				args = append(args, matches...)
			}
		}
	}
	return files
}

// 处理文件 打开文件 - 读文件 - 处理 - 关闭文件
func updateFrequencies(filename string, frequencyForWord map[string]int) {
	var file *os.File
	var err error
	if file, err = os.Open(filename); err != nil {
		log.Println("failed to open the file:", err)
		return
	}
	defer file.Close()
	readAndUpdateFrequencies(bufio.NewReader(file), frequencyForWord)
}

// 读取文件并且统计词频
func readAndUpdateFrequencies(reader *bufio.Reader, frequencyForWord map[string]int) {
	for true {
		line, err := reader.ReadString('\n')
		for _, word := range SplitOnNonLetters(strings.TrimSpace(line)) {
			if len(word) > utf8.UTFMax || utf8.RuneCountInString(word) > 1 {
				frequencyForWord[strings.ToLower(word)] += 1
			}
		}
		if err != nil {
			if err != io.EOF {
				log.Println("failed to finish reading the file: ", err)
			}
			break
		}
	}
}

// SplitOnNonLetters 在非单词字符上对一个字符串进行切分
func SplitOnNonLetters(s string) []string {
	notALetter := func(char rune) bool {
		return !unicode.IsLetter(char)
	}
	return strings.FieldsFunc(s, notALetter)
}

// 打印词频结果 - 1
func reportByWords(frequencyForWord map[string]int) {
	words := make([]string, 0, len(frequencyForWord))
	wordWidth, frequencyWidth := 0, 0

	for word, frequency := range frequencyForWord {
		words = append(words, word)
		if width := utf8.RuneCountInString(word); width > wordWidth {
			wordWidth = width
		}
		if width := len(fmt.Sprint(frequency)); width > frequencyWidth {
			frequencyWidth = width
		}
	}

	sort.Strings(words)
	gap := wordWidth + frequencyWidth - len("Word") - len("Frequency")
	fmt.Printf("Word %*s%s\n", gap, " ", "Frequency")
	for _, word := range words {
		fmt.Printf("%-*s %*d\n", wordWidth, word, frequencyWidth, frequencyForWord[word])
	}
}

// 保存反转结果 一个频率对应不同的单词
func invertStringIntMap(intForString map[string]int) map[int][]string {
	stringsForInt := make(map[int][]string, len(intForString))
	for key, value := range intForString {
		stringsForInt[value] = append(stringsForInt[value], key)
	}
	return stringsForInt
}

// 打印词频结果 - 2
func reportByFrequency(wordsForFrequency map[int][]string) {
	frequencies := make([]int, 0, len(wordsForFrequency))
	for frequency := range wordsForFrequency {
		frequencies = append(frequencies, frequency)
	}
	sort.Ints(frequencies)
	width := len(fmt.Sprint(frequencies[len(frequencies)-1]))
	fmt.Println("Frequency → Words")
	for _, frequency := range frequencies {
		words := wordsForFrequency[frequency]
		sort.Strings(words)
		fmt.Printf("%*d %s\n", width, frequency, strings.Join(words, ", "))
	}
}
