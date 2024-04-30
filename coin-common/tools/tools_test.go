package tools

import (
	"fmt"
	"testing"
)

func TestEncryp(t *testing.T) {
	password := "123456"
	encodePass := "mXA2gZMI"
	salt := "f2b8b7423541b0d104a4bd43391febb95d9d6de911b3fc94c577ed1269dc62dcbeba1d13b649ef3105874fc42cafe4b73838a36118d1a9dcd07e320cd367821c"
	fmt.Print(Verify(password, salt, encodePass, nil))
}
