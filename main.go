package main

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	genericlist "github.com/TGRZiminiar/golangFundemaltal/genericList"
	howtocontext "github.com/TGRZiminiar/golangFundemaltal/howToContext"
	safemap "github.com/TGRZiminiar/golangFundemaltal/safeMap"
)

func main() {
	
	// TestGenericList();
	// TestSafeMap(&testing.T{});
	TestContext();
}

func TestContext() {
	start := time.Now();
	ctx := context.Background();
	userId := 10;
	val, err := howtocontext.FetchUserData(ctx, userId);
	if(err != nil){
		log.Fatal(err)
	}

	fmt.Println("Result: ",val);
	fmt.Println("Took: ",time.Since(start));

}

func TestSafeMap(t *testing.T){

	m := safemap.New[int ,int]()

	for i := 0; i < 10; i++ {
		go func(i int) {
			m.Insert(i, i*2);
			value, err := m.Get(i);
			if(err != nil){
				t.Error(err)
			}

			if(value != i*2){
				t.Errorf("%d should be %d", i, i*2);
			}
		}(i)
	}

}

func TestGenericList(){
	gList := genericlist.New[string]();
	
	gList.Insert("bob");
	gList.Insert("foo");
	gList.Insert("bar");
	gList.Insert("alice");
	
	gList.Remove(1);
	gList.RemoveByValue("alice");
	fmt.Printf("%+v\n",gList);

}