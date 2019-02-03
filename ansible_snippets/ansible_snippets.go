package ansible_snippets

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func Get_snippet(snip string) string {
	if snip == "" {
		return "You need to provide modulename to get info"
	}
	file, err := os.Open("/opt/telegram-fun-bot/ansible.snippets")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Append ":" to search snippet for search part to work properly
	snip += ":"

	scanner := bufio.NewScanner(file)
	scanner.Split(crunchSplitFunc)
	a := ""
	for scanner.Scan() {
		a = scanner.Text()
		b := strings.Split(a, "\n")
		for _, i := range b {
			if i == snip {
				return a
			}
		}
	}
	return "Sorry I can't find anything"
}

func crunchSplitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {

	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	var i int
	i = strings.Index(string(data), "\nsnippet")
	if i >= 0 {
		return i + 1, data[0:i], nil
	}

	if atEOF {
		return len(data), data, nil
	}
	return
}

// func main() {
// 	line := "git"
// 	snp := Get_snippet(line)
// 	fmt.Println(snp)
// }
