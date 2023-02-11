package main

import (
	"fmt"

	genericlist "github.com/TGRZiminiar/golangFundemaltal/genericList"
)

func main() {
	gList := genericlist.New[string]();

	gList.Insert("bob");
	gList.Insert("foo");
	gList.Insert("bar");
	gList.Insert("alice");

	gList.Remove(1);
	gList.RemoveByValue("alice");
	fmt.Printf("%+v\n",gList);


}