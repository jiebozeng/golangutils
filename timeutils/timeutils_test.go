package timeutils

import (
	"fmt"
	"testing"
)

func TestTimeUtils(t *testing.T) {
	nowTime := GetNowTime()
	fmt.Println("nowTime:", nowTime)
	threeDaysAgo := GetTimeSubDateTime(nowTime, 3)
	fmt.Println("three days ago:", threeDaysAgo)
}
