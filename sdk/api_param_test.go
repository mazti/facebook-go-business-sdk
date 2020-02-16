package sdk

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateTimeRangeOK(t *testing.T) {
	since, _ := time.Parse(TimeFormat, "2020-01-01")
	until, _ := time.Parse(TimeFormat, "2020-02-01")
	expected := `{"since":"2020-01-01","until":"2020-02-01"}`

	timeRange := CreateTimeRange(since, until)

	bytes, err := json.Marshal(timeRange)
	assert.Nil(t, err)
	assert.Equal(t, expected, string(bytes))
}
