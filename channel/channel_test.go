package channel

import (
	"fmt"
	"testing"
	"time"
)

func TestRWChannel(t *testing.T){

}

func TestRChannel(t *testing.T){

}

func TestWChannel(t *testing.T){
	ch := make(chan int)
	cClose(ch)
}

func cClose(c <- chan int)  {
	<- c
}

func TestForRange(t *testing.T)  {
	ch := make(chan int,10)

	go func() {
		//for i := 0; i < 10; i++ {
		//	ch<-i
		//	if i== 8 {
		//		close(ch)
		//		return
		//	}
		//}
		time.Sleep(3*time.Second)
		ch <-1
	}()
	for i := range ch {
		fmt.Println(i)
	}

	fmt.Println("exit")
}

func chanRange(ch chan int)  {
	for i := range ch{
		fmt.Printf("get element for ch : %d",i)
	}
}


