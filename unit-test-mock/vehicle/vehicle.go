package vehicle

import (
	"fmt"

	"../maintenance"
)

func CheckUp(name string, days_used int) (string, error) {
	specialist := maintenance.Specialist{}
	needsMaintenance, err := specialist.NeedsMaintenance(days_used)
	if err != nil {
		return "", err
	}

	if needsMaintenance {
		return fmt.Sprintf("%v is up for maintenance", name), nil
	} else {
		return fmt.Sprintf("%v is no need for maintenance", name), nil
	}
}
