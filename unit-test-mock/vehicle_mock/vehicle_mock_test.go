package vehicle_mock

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"../mocks"
)

func TestCheckUp(t *testing.T) {
	testWrapper := new(mocks.MockMaintenanceWrapper)
	maintenanceRepository := MaintenanceWrapper{Maintenance: testWrapper}

	t.Run("up for maintenace", func(t *testing.T) {
		testDays := 45
		testWrapper.On("NeedsMaintenance", testDays).Return(true).Once()
		result, err := maintenanceRepository.CheckUp("CRV", testDays)
		assert.Equal(t, result, "CRV is up for maintenance")
		assert.Equal(t, err, nil)
	})

	t.Run("no need for maintenace", func(t *testing.T) {
		testDays := 15
		testWrapper.On("NeedsMaintenance", testDays).Return(false).Once()
		result, err := maintenanceRepository.CheckUp("CRV", testDays)
		assert.Equal(t, result, "CRV is no need for maintenance")
		assert.Equal(t, err, nil)
	})

	t.Run("checkup error", func(t *testing.T) {
		testDays := -15
		expectedError := errors.New("cannot accept less than zero days")
		testWrapper.On("NeedsMaintenance", testDays).Return(false, expectedError).Once()
		result, err := maintenanceRepository.CheckUp("CRV", testDays)
		testWrapper.AssertExpectations(t)
		assert.Equal(t, result, "")
		assert.Equal(t, err, expectedError)
	})
}
