package config

import (
	"testing"

	"github.com/magiconair/properties/assert"
	"github.com/spf13/viper"
)

func TestEnvConfigMerge(t *testing.T) {
	SetEnv(Test)
	err := Load()
	if err != nil {
		t.Fatal(err)
	}

	spn := "postgresql://eca:crater10lake@localhost:5432/ear_monitor2_test?sslmode=disable"
	assert.Equal(t, viper.GetString("db.host"), spn)
}
