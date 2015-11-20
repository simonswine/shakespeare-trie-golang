package main

import (
	"bufio"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"strings"
	"unicode/utf8"
)

func checkFail(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func httpDownloadToFile(uri string, path string) error {
	log.Infof("download from:  %s", uri)
	res, err := http.Get(uri)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	out, err := os.Create(path)
	if err != nil {
		return err
	}

	defer out.Close()
	_, err = io.Copy(out, res.Body)
	return err
}

func readLinesFromFile(path string) (*ShakespeareTrie, error) {
	inFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	trie := NewShakespeareTrie()

	for scanner.Scan() {
		trie.AddString(strings.TrimSpace(scanner.Text()))
	}

	return trie, nil
}

func readShakespearLines(path string, url string) (*ShakespeareTrie, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {

		log.Infof("download shakespears full texts as they are not existing in '%s'", path)
		err = httpDownloadToFile(url, path)
		if err != nil {
			return nil, fmt.Errorf("Error downloading from '%s' to '%s': %s", url, path, err)
		}
	}

	return readLinesFromFile(path)
}

func main() {
	inputPath := "shakespeare.txt"
	url := "https://www.gutenberg.org/ebooks/100.txt.utf-8"

	log.Info("Reading shakespeare lines")
	trie, err := readShakespearLines(inputPath, url)
	if err != nil {
		log.Fatalf("Error building the trie: %s", err)
	}
	log.Info("Finished reading shakespeare lines")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter query: ")
		text, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		for _, result := range trie.GetMatches(strings.TrimSpace(string(text))) {
			fmt.Println(" ## " + result)
		}
	}
	fmt.Println("")
	log.Info("Goodbye")

}

type ShakespeareTrie struct {
	Nodes map[rune]*ShakespeareTrie
}

func NewShakespeareTrie() *ShakespeareTrie {
	t := ShakespeareTrie{}
	t.Nodes = make(map[rune]*ShakespeareTrie)
	return &t
}

func (t *ShakespeareTrie) GetMatches(s string) (result []string) {

	var firstRune rune
	var remainder string

	if s == "" {
		if len(t.Nodes) == 0 {
			return []string{""}
		}
		remainder = ""
	} else {
		firstRune, remainder = splitRune(s)
	}

	output := []string{}
	for key, node := range t.Nodes {
		if firstRune == key || s == "" {
			output = append(output, prependRune(key, node.GetMatches(remainder))...)
		}
	}
	return output
}

func prependRune(c rune, input []string) (output []string) {
	for _, elem := range input {
		output = append(output, string(c)+elem)
	}
	return output
}

func splitRune(s string) (rune, string) {
	firstRune, size := utf8.DecodeRuneInString(s)
	return firstRune, s[size:len(s)]
}

func (t *ShakespeareTrie) AddString(s string) {
	log.SetLevel(log.DebugLevel)
	if len(s) < 1 {
		return
	}

	firstRune, remainder := splitRune(s)

	var child *ShakespeareTrie
	if n, ok := t.Nodes[firstRune]; ok {
		child = n
	} else {
		child = NewShakespeareTrie()
		t.Nodes[firstRune] = child
	}

	// log.Debugf("char: %+c remainder: %s child: %+v", firstRune, remainder, child)

	child.AddString(remainder)
	return
}
