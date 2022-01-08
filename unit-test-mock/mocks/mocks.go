package mocks

import "github.com/stretchr/testify/mock"

type MockMaintenanceWrapper struct {
	mock.Mock
}

func (w *MockMaintenanceWrapper) NeedsMaintenance(days int) (needsMaintenance bool, err error) {
	args := w.Called(days)
	if len(args) > 1 {
		return args.Get(0).(bool), args.Get(1).(error)
	}
	return args.Get(0).(bool), nil
}
