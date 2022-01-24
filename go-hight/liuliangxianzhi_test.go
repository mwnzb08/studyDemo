package go_hight

import (
	"fmt"
	"github.com/juju/ratelimit"
	"testing"
	"time"
)

func TestLiu (b *testing.T) {
	bucket := ratelimit.NewBucket(time.Second * 2, 10)
	fmt.Println(bucket.Available())
	go add(bucket)
	go add(bucket)
	go add(bucket)

	time.Sleep(time.Second * 15)
	fmt.Println(bucket.Available())
}

func add (bucket *ratelimit.Bucket) {
	for {
		select {
		case <-time.Tick(time.Second):
			bucket.TakeAvailable(1)
			fmt.Println(bucket.Available())

		}
	}

}
