package uuidpkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	m.Run()
}

func Test_NewRandom(t *testing.T) {
	uuid, err := NewRandom()
	assert.Nil(t, err)
	t.Logf("uuid = %x\n", uuid)
}


