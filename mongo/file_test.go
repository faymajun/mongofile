package mongo

import (
	"testing"
	"time"
)

func TestFile(t *testing.T) {
	InitMongoFile("growth")
	WriteOne("22234")
	time.Sleep(2 * time.Second)
	CloseMongoLog()
}
