package basic

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddOne(t *testing.T) {
// 	var (
// 		input = 1
// 		output = 2
// 	) 
// 	actual := AddOne(input)
// 	if actual != output {
// 		t.Errorf("AddOne(%d), output = %d, actual = %d", input, output, actual)
// 	}
	assert.Equal(t, AddOne(2), 3, "AddOne(2) should be 3")
}

func TestWithRequire(t *testing.T) {
	require.Equal(t, 2, 3)
	fmt.Println("Not executing")
}

func TestWithAssert(t *testing.T) {
	assert.Equal(t, 2, 3)
	fmt.Println("executing")
}