package search

import (
	"context"
	"log"
	"testing"
)

func TestAll_user(t *testing.T) {

	res := All(context.Background(), "0000 0000 000000", []string{"test.txt",})

	if len(res) == 0 {
		t.Errorf("FindAllMatchTextInFile error   res => %v ", res)
	}

	log.Println("res => ", res)

}

func TestFindPhrase(t *testing.T) {

	res := FindAllMatchTextInFile("0000 0000 000000", "test.txt")

	if len(res) == 0 {
		t.Errorf("FindAllMatchTextInFile error   res => %v ", res)
	}

	log.Println("res => ", res)

}
