package howtocontext

import (
	"context"
	"fmt"
	"time"
)

type Response struct {
	value int;
	err error;
}

func FetchUserData(ctx context.Context, userId int) (int, error){
	ctx, calcel := context.WithTimeout(ctx, time.Millisecond * 200);
	defer calcel();
	
	respch := make(chan Response);

	go func() {
		val, err := FetchThirdPartyStuffWhiceCanBeSlow();
		respch <- Response{
			value	:	val,
			err		:	err,
		}
	}()
	
	for {
		select {
			case <- ctx.Done():
				return 0, fmt.Errorf("fetching data from third party took too long");
			case resp := <- respch:
				return resp.value, resp.err 
		}
	}

	// return val, nil;

}

func FetchThirdPartyStuffWhiceCanBeSlow() (int, error) {

	time.Sleep(time.Millisecond * 100);


	return 666, nil;

}