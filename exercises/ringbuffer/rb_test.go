package ringbuffer_test

import (
	"grokking/exercises/ringbuffer"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSensorBuf(t *testing.T) {
	buf := ringbuffer.NewSensorBuf(10)

	for i := 0; i < 7; i++ {
		buf.Add(makeRandomSensorData())
	}

	out := buf.Emit()
	logSensorData(t, out)

	newOut := buf.Emit()
	logSensorData(t, newOut)
	assert.Len(t, newOut, 0)
}

func TestTmp(t *testing.T) {
	r := 'a'
	maxIncr := 1000
	curIncr := 0
	for {
		curIncr++
		r++
		if r == 'a' || maxIncr == curIncr {
			break
		}

		t.Log(string(r))
	}
}

func logSensorData(t *testing.T, data []*ringbuffer.SensorData) {
	t.Log("Logging sensor data.")

	for _, sd := range data {
		t.Log(sd)
	}
}

func makeRandomSensorData() ringbuffer.SensorData {
	return ringbuffer.SensorData{
		TempratureC: rand.Intn(100),
		Ts:          time.Now(),
	}
}

func TestModulo(t *testing.T) {
	exp := 16 % 6
	act := modulo(16, 6)

	assert.Equal(t, exp, act)
}

func modulo(dividend, devisor int) int {
	noRemainder := dividend / devisor
	nextNoRemainder := noRemainder * devisor
	remainder := dividend - nextNoRemainder

	return remainder
}
