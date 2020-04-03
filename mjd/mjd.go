package mjd

import (
	"encoding/binary"
	"math"
	"time"
)

const secs_per_day = 24 * 60 * 60
const unix_epoch = 40587

type Mjd struct {
	day uint64
	µs  float64
}

func (t Mjd) Unix() time.Time {
	days := t.day - unix_epoch
	secs := t.µs / 1e6
	if secs > secs_per_day {
		secs -= secs_per_day
	}
	additional_secs := (int64)(math.Floor(secs))
	residue := secs - math.Floor(secs)
	nanos := (int64)(math.Floor(residue * 1e9))
	return time.Unix((int64)(days*secs_per_day)+additional_secs, nanos)
}

func (t Mjd) Day() uint64 {
	return t.day
}

func (t Mjd) Microseconds() float64 {
	return t.µs
}

func (t Mjd) RoughtimeEncoding() uint64 {
	var daybytes [4]byte
	var microbytes [8]byte
	var buf [8]byte
	binary.LittleEndian.PutUint32(daybytes[:], uint32(t.day))
	binary.LittleEndian.PutUint64(microbytes[:], uint64(math.Floor(t.µs)))
	copy(buf[0:3], daybytes[1:4])
	copy(buf[3:8], microbytes[3:8])
	return binary.BigEndian.Uint64(buf[:])
}

func RoughtimeVal(in uint64) Mjd {
	var buf [8]byte
	var daybytes [4]byte
	var microbytes [8]byte
	binary.BigEndian.PutUint64(buf[:], in)
	copy(daybytes[1:4], buf[0:3])
	copy(microbytes[3:8], buf[3:8])
	ret := Mjd{}
	ret.day = uint64(binary.LittleEndian.Uint32(daybytes[:]))
	ret.µs = float64(binary.LittleEndian.Uint64(microbytes[:]))
	return ret
}
