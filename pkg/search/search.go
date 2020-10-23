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
	wg := sync.WaitGroup{}

	mu := sync.Mutex{}
	var results []Result

	ctx, cancel := context.WithCancel(ctx)

	for i := 0; i < len(files); i++ {
		wg.Add(1)

		go func(ctx context.Context, filename string, i int, ch chan<- []Result) {
			defer wg.Done()
			res := FindAllMatchTextInFile(phrase, filename)

			mu.Lock()
			results = append(results, res...)
			mu.Unlock()

			//if len(results) > 0 && i == len(files)-1{
			ch <- results
			//}

		}(ctx, files[i], i, ch)
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
