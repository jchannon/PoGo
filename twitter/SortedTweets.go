package twitter

import (
	"fmt"

	"github.com/chimeracoder/anaconda"
)

//Tweets sorted
type Tweets []anaconda.Tweet

func (slice Tweets) Len() int {
	return len(slice)
}

func (slice Tweets) Less(i, j int) bool {
	firstTime, err := slice[i].CreatedAtTime()
	if err != nil {
		fmt.Println("oops")
	}
	secondTime, err2 := slice[j].CreatedAtTime()
	if err2 != nil {
		fmt.Println("oops")
	}

	return firstTime.Before(secondTime)
}

func (slice Tweets) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}
