package common

import (
	"errors"
	"testing"
)

func TestGraphErrorIs(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		Name     string
		Err      GraphError
		Err2     error
		ExpectIs bool
	}{
		{
			Name:     "nil",
			Err:      GraphError{},
			Err2:     nil,
			ExpectIs: false,
		},
		{
			Name:     "other",
			Err:      GraphError{},
			Err2:     errors.New("an error"),
			ExpectIs: false,
		},
		{
			Name: "same error code",
			Err: GraphError{
				Code:    "code1",
				Message: "err",
			},
			Err2: GraphError{
				Code:    "code1",
				Message: "another err",
			},
			ExpectIs: true,
		},
		{
			Name: "different error code",
			Err: GraphError{
				Code:    "code1",
				Message: "err",
			},
			Err2: GraphError{
				Code:    "code2",
				Message: "err",
			},
			ExpectIs: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			if errors.Is(testCase.Err, testCase.Err2) != testCase.ExpectIs {
				t.Errorf(
					"Expected %t but got %t for testing if errors.Is(%#v, %#v)",
					testCase.ExpectIs,
					!testCase.ExpectIs,
					testCase.Err,
					testCase.Err2,
				)
			}
		})
	}
}
