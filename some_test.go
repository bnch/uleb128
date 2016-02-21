package uleb128

import (
	"fmt"
	"testing"
)

func TestLength(t *testing.T) {
	testBytes := []byte("\x93\x03ssdadadihjqrjqheqj")
	fmt.Println(Unmarshal(testBytes))
}
