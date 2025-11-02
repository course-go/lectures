package calculator_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"path/to/local/package/calculator"
)

// START OMIT

func TestCalculatorMock(t *testing.T) {
	mock := setUpMock(t)
	assert.Equal(t, 2, mock.Add(1, 1))
	assert.Equal(t, -1, mock.Add(0, -1))
	assert.Equal(t, 100, mock.Add(105, -5))
}

func setUpMock(t *testing.T) *calculator.MockAdder {
	ctrl := gomock.NewController(t)
	mock := calculator.NewMockAdder(ctrl)

	gomock.InOrder(
		mock.EXPECT().Add(1, 1).Return(2),
		mock.EXPECT().Add(0, -1).Return(-1),
		mock.EXPECT().Add(105, -5).Return(100),
	)

	return mock
}

// END OMIT
