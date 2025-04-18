package types_test

import (
	"testing"

	"github.com/FlowerWrong/gotwi/resources"
	"github.com/FlowerWrong/gotwi/space/searchspace/types"
	"github.com/stretchr/testify/assert"
)

func Test_SearchSpaces_HasPartialError(t *testing.T) {
	var errorTitle string = "test partical error"
	cases := []struct {
		name   string
		res    *types.ListOutput
		expect bool
	}{
		{
			name: "has partical error",
			res: &types.ListOutput{
				Errors: []resources.PartialError{
					{Title: &errorTitle},
				}},
			expect: true,
		},
		{
			name: "has no partical error",
			res: &types.ListOutput{
				Errors: []resources.PartialError{}},
			expect: false,
		},
		{
			name: "partical error is nil",
			res: &types.ListOutput{
				Errors: []resources.PartialError{}},
			expect: false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			hpe := c.res.HasPartialError()
			assert.Equal(tt, c.expect, hpe)
		})
	}
}
