package search

import (
	"context"
	"io/ioutil"
	"strings"
	"sync"
)

//Result ..
type Result struct {
	Phrase  string
	Line    string
	LineNum int64
	ColNum  int64
}

//All ...
func All(ctx context.Context, phrase string, files []string) <-chan []Result {
	ch := make(chan []Result)
	//defer close(ch)
	wg := sync.WaitGroup{}

	ctx, cancel := context.WithCancel(ctx)

	for i := 0; i < len(files); i++ {
		wg.Add(1)

		go func(ctx context.Context, filename string, ch chan<- []Result) {
			defer wg.Done()
			res := FindAllMatchTextInFile(phrase, filename)

			if len(res) > 0 {
				ch <- res
			}

		}(ctx, files[i], ch)
	}
	go func() {
		defer close(ch)
		wg.Wait()
	}()

	cancel()
	return ch
}

/* //Any ...
func Any(ctx context.Context, phrase string, files []string) <-chan Result {
	ch := make(chan Result)
	defer close(ch)

	for i := 0; i < len(files); i++ {
		go func(ctx context.Context, ch chan<- Result) {

		}(ctx, ch)
	}

	return ch
} */

//FindAllMatchTextInFile ...
func FindAllMatchTextInFile(phrase, fileName string) (res []Result) {

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return res
	}

	file := string(data)

	temp := strings.Split(file, "\n")

	for i, line := range temp {
		//fmt.Println("[", i+1, "]\t", line)
		if strings.Contains(line, phrase) {

			r := Result{
				Phrase:  phrase,
				Line:    line,
				LineNum: int64(i + 1),
				ColNum:  int64(strings.Index(line, phrase)),
			}

			res = append(res, r)
		}
	}

	return res
}
/* 
//GetIndex ...
func GetIndex(s string, substr string) (n int) {

	st := []string(s)
	var str []string
	t := len(st) / len(substr)
	i := 0

	for i := 0; i <= t; i++ {

		l := i * len(substr)
		r := (i + 1) * len(substr)
		if r > len(st) {
			r = len(st)
		}

		sr := strings.Join(st[l:r], "")
		str = append(str, sr)
	}
	sr := strings.Join(st[i*len(substr):], "")
	str = append(str, sr)

	fmt.Println(str)
	return len(str)

}
 */