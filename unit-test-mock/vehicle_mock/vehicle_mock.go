package vehicle_mock

import (
	"fmt"
)

type MaintenanceAPI interface {
	NeedsMaintenance(days int) (needsMaintenance bool, err error)
}

type MaintenanceWrapper struct {
	Maintenance MaintenanceAPI
}

func (w *MaintenanceWrapper) CheckUp(name string, days_used int) (string, error) {
	needsMaintenance, err := w.Maintenance.NeedsMaintenance(days_used)
	if err != nil {
		return "", err
	}

	if needsMaintenance {
		return fmt.Sprintf("%v is up for maintenance", name), nil
	} else {
		return fmt.Sprintf("%v is no need for maintenance", name), nil
	}
}
