package bar

import (
	"fmt"

	"github.com/superfly/macaroon"
)

func Baz() {
	macaroon.ToAuthorizationHeader([]byte("foo"))
	fmt.Println("hi from bar.Baz")
}
