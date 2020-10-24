package main

import (
	"context"
	"fmt"


	"github.com/najibulloShapoatov/search/pkg/search"
)

func main() {

	 res := search.Any(context.Background(),"HTTP", []string{"./test.txt","./test copy.txt"})


	 r, ok := <- res
	 if !ok {
		 fmt.Println("error ok =>", ok)
	 }

	//for _, r := range res {

		fmt.Println("---------------")
		fmt.Println("res.Phrase) => ", r.Phrase)
		fmt.Println("res.Line) => ", r.Line)
		fmt.Println("res.LineNum) => ", r.LineNum)
		fmt.Println("res.ColNum) => ", r.ColNum)
		fmt.Println("---------------")
	//}


}
