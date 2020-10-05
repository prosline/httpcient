package examples

import (
	"fmt"
	"testing"
)

func TestGetEndPoints(t *testing.T) {
	//Initialization already available by singleton

	//Execution
	ep, err := GetEndPoints()

	//Validation
	fmt.Println(err)
	fmt.Println(ep)
}
