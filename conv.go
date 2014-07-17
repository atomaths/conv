// Copyright 2013 Jongmin Kim. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package conv

import (
	"encoding/binary"
	"math"
)

func BytesToFloat64(b []byte) float64 {
	bits := binary.LittleEndian.Uint64(b)
	f := math.Float64frombits(bits)
	return f
}

func Float64ToBytes(f float64) []byte {
	bits := math.Float64bits(f)
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, bits)
	return bytes
}
