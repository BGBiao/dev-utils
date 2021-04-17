package time

import (
	"fmt"
	"testing"
)

func TestTimeFormat(t *testing.T) {

	timeUnix := StringTsToUnix("2021-04-17T13:53:00")
	fmt.Println(timeUnix)

	timeStr := UnixTsToString(timeUnix + 10)
	fmt.Println(timeStr)

	fmt.Printf("now:%v\n", NowTime())

	// after 50s
	fmt.Printf("nowafter:%v\n", NowAfterTime(50))

}
