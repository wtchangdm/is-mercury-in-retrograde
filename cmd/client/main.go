package client

import (
	"fmt"

	"github.com/wtchangdm/is-mercury-in-retrograde/internal/mercury"
)

func Do() {
	m := mercury.New()
	ok, err := m.Retrograde()

	if err != nil {
		fmt.Printf("Looks like we don't know whether Mercury is in retrograde: %v\n", err)
		return
	}

	if ok {
		fmt.Println("Mercury is in retrograde!")
		return
	}

	fmt.Println("Mercury is not in retrograde!")
}
