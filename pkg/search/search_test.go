package search

import (
	"context"
	"log"
	"testing"
)

func TestAll_user(t *testing.T) {

	ch := All(context.Background(), "0000 0000 000001", []string{"test.txt",})

	s, ok := <-ch

	if !ok {
		t.Errorf(" method SumPaymentsWithProgress ok not closed => %v", ok)
	} 

	log.Println("=======>>>>>",s) 

}

/* func TestFindPhrase(t *testing.T) {

	res := FindAllMatchTextInFile("0000 0000 000000", "test.txt")

	if len(res) == 0 {
		t.Errorf("FindAllMatchTextInFile error   res => %v ", res)
	}

	log.Println("res => ", res)

} */
