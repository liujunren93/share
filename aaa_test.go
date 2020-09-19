package share

import (
	"fmt"
	"testing"
	"time"
)

type Next func() string

func Round(list *[]string) Next {
	return func() string {
		fmt.Println(list)
		return (*list)[0]
	}
}

func TestAa(t *testing.T) {
	var list =[]string{"adsa","bbb"}
	round := Round(&list)
	for  {
		time.Sleep(time.Second*1)
		list=append(list, "dsads")
		round()

	}

}
