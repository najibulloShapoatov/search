package main

import (
	"fmt"
	"strings"

	"github.com/najibulloShapoatov/search/pkg/search"
)

func main() {
	/* res := search.FindAllMatchTextInFile("0000 0000 000000", "test.txt")

	for _, r := range res {

		fmt.Println("---------------")
		fmt.Println("res.Phrase) => ", r.Phrase)
		fmt.Println("res.Line) => ", r.Line)
		fmt.Println("res.LineNum) => ", r.LineNum)
		fmt.Println("res.ColNum) => ", r.ColNum)
		fmt.Println("---------------")
	}

	fmt.Println(strings.Index("В случае, когда дела не  обстоят таким образом, можно обратиться к другим методам скрейпинга.", "другим методам скрейпинга."))
 */
	s1:="В случае, когда дела не  обстоят таким образом, можно обратиться к другим методам скрейпинга."
	s2:="другим методам скрейпинга."

	 r1 := strings.Index(s1, s2)


	fmt.Println(r1) 


}
