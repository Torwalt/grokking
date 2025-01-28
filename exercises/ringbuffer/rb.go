package ringbuffer

import (
	"fmt"
	"time"
)

const (
	// insertInitIdx is the initial value for lastInsertIdx. It will never be
	// -1 after first insert again.
	insertInitIdx = -1
)

type SensorData struct {
	TempratureC int
	Ts          time.Time
}

func (sd SensorData) String() string {
	return fmt.Sprintf("Temp: %v, ts: %v", sd.TempratureC, sd.Ts.Format(time.RFC3339Nano))
}

type SensorBuf struct {
	buf           []*SensorData
	size          int
	lastInsertIdx int
	nextReadIdx   int
}

func NewSensorBuf(size uint) *SensorBuf {
	return &SensorBuf{
		buf:           make([]*SensorData, size),
		size:          int(size),
		lastInsertIdx: insertInitIdx,
		nextReadIdx:   0,
	}
}

func (sb *SensorBuf) Add(sd SensorData) {
	sb.lastInsertIdx = sb.increment(sb.lastInsertIdx)
	sb.buf[sb.lastInsertIdx] = &sd
}

func (sb *SensorBuf) Emit() []*SensorData {
	out := []*SensorData{}

	for {
		if sb.nextReadIdx == sb.lastInsertIdx || sb.lastInsertIdx == insertInitIdx {
			break
		}

		if sb.buf[sb.nextReadIdx] != nil {
			out = append(out, sb.buf[sb.nextReadIdx])
			// "remove" from buf
			sb.buf[sb.nextReadIdx] = nil
		}

		sb.nextReadIdx = sb.increment(sb.nextReadIdx)
	}

	return out
}

func (sb *SensorBuf) increment(idx int) int {
	// idx: -1, size 5 -> (-1 + 1) % 5 = 0
	// idx: 0, size 5 -> (0 + 1) % 5 = 1
	// ...
	// idx: 4, size 5 -> (4 + 1) % 5 = 0
	// idx: 5, size 5 -> (5 + 1) % 5 = 1
	// as soon as size is reached, restart from index 0.
	return (idx + 1) % sb.size
}
