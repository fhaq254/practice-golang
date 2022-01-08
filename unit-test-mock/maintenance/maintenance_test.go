package maintenance

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNeedsMaintenance(t *testing.T) {
	specialist := Specialist{}

	t.Run("when days is just 30", func(t *testing.T) {
		needsMaintenance, err := specialist.NeedsMaintenance(30)
		assert.Equal(t, needsMaintenance, true)
		assert.Equal(t, err, nil)
	})

	t.Run("when days is more than 30", func(t *testing.T) {
		needsMaintenance, err := specialist.NeedsMaintenance(100)
		assert.Equal(t, needsMaintenance, true)
		assert.Equal(t, err, nil)
	})

	t.Run("when days is less than 30", func(t *testing.T) {
		needsMaintenance, err := specialist.NeedsMaintenance(1)
		assert.Equal(t, needsMaintenance, false)
		assert.Equal(t, err, nil)
	})

	t.Run("when days is just 0", func(t *testing.T) {
		needsMaintenance, err := specialist.NeedsMaintenance(0)
		assert.Equal(t, needsMaintenance, false)
		assert.Equal(t, err, nil)
	})

	t.Run("when days is less than 0", func(t *testing.T) {
		needsMaintenance, err := specialist.NeedsMaintenance(-1)
		assert.Equal(t, needsMaintenance, false)
		assert.Equal(t, err, errors.New("cannot accept less than zero days"))
	})
}
