package calculator_test

import (
	"testing"

	addermock "github.com/course-go/lectures/assets/lecture-05/gomock/adder/mocks"
	"github.com/course-go/lectures/assets/lecture-05/gomock/calculator"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

// START OMIT

func TestCalculator(t *testing.T) {
	calculator, mock := setUpCalculator(t)

	gomock.InOrder(
		mock.EXPECT().Add(1, 1).Return(2),
		mock.EXPECT().Add(0, -1).Return(-1),
		mock.EXPECT().Add(105, -5).Return(100),
	)

	assert.Equal(t, 2, calculator.Add(1, 1))
	assert.Equal(t, -1, calculator.Add(0, -1))
	assert.Equal(t, 100, calculator.Add(105, -5))
}

func setUpCalculator(t *testing.T) (c calculator.Calculator, m *addermock.MockAdder) {
	t.Helper()
	ctrl := gomock.NewController(t)
	mock := addermock.NewMockAdder(ctrl)
	calculator := calculator.New(mock)
	return calculator, mock
}

// END OMIT
