package concurrency

import (
	"fmt"
	"sync"
	"time"
)

//ทำให้สามารถทำงานหลายๆอย่างพร้อมกันได้โดยใช้เวลามากสุดแค่ตัวที่จะทำงานมากสุด
func ExampleConCurrency() {

	now := time.Now(); 

	userId := 10;

	respch := make(chan string, 128);

	wg := &sync.WaitGroup{};

	go fetchUserData(userId, respch, wg);
	go fetchUserRecommendations(userId, respch, wg);
	go fetchUserLikes(userId, respch, wg);
	// 3 อย่างที่ต้องถูกทำ
	wg.Add(3)
	wg.Wait();

	close(respch);

	for resp := range respch {
		fmt.Println(resp);
	}

	fmt.Println(time.Since(now));

}

func fetchUserData(userId int, respch chan string, wg *sync.WaitGroup) {
	time.Sleep(80 * time.Millisecond);
	
	respch <- "user data"

	wg.Done();
}

func fetchUserRecommendations(userId int, respch chan string, wg *sync.WaitGroup) {
	time.Sleep(120 * time.Millisecond);
	
	respch <- "user recommentdation"
	
	wg.Done();
}

func fetchUserLikes(userId int, respch chan string, wg *sync.WaitGroup) {
	time.Sleep(50 * time.Millisecond);
	
	respch <- "user likes"
	
	wg.Done();
}