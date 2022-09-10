package client

import (
	"fmt"

	"github.com/wtchangdm/is-mercury-in-retrograde/internal/mercury"
)

func Do() {
	m := mercury.New()
	ok, err := m.Retrograde()

	if err != nil {
		fmt.Printf("Looks like we don't know whether mercury is in retrograde: %v\n", err)
		return
	}

	if ok {
		fmt.Println("mercury is in retrograde!")
		return
	}

	fmt.Println("mercury is not in retrograde!")
}
