package vehicle

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckUp(t *testing.T) {
	t.Run("up for maintenace", func(t *testing.T) {
		result, err := CheckUp("CRV", 45)
		assert.Equal(t, result, "CRV is up for maintenance")
		assert.Equal(t, err, nil)
	})

	t.Run("no need for maintenace", func(t *testing.T) {
		result, err := CheckUp("CRV", 15)
		assert.Equal(t, result, "CRV is no need for maintenance")
		assert.Equal(t, err, nil)
	})

	t.Run("checkup error", func(t *testing.T) {
		result, err := CheckUp("CRV", -15)
		assert.Equal(t, result, "")
		assert.Equal(t, err, errors.New("cannot accept less than zero days"))
	})
}
