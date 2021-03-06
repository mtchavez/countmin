package countmin

import (
	"errors"
	"hash/fnv"
	"math"
)

// MAXINT64 is max 64 bit integer size
const MAXINT64 = 1<<63 - 1

// CountMin holds count table for calculating
// frequencies from a given input (usually a stream of data)
type CountMin struct {
	depth      int
	width      int
	size       int64
	eps        float64
	confidence float64
	table      [][]int64
}

// New CountMin given a depth and width of the frequency
// table to use. Returns the newly created CountMin
func New(depth, width int) *CountMin {
	cm := &CountMin{
		depth:      depth,
		width:      width,
		eps:        2.0 / float64(width),
		confidence: 1.0 - 1.0/math.Pow(2, float64(depth)),
	}
	cm.initTable()
	return cm
}

// NewWithEpsCount initializes a CountMin given a confidence
// to ensure between 0 < 1.0 which effects number of hashes used
// and an epsilon between 0 and 1 with a smaller value for
// higher precision as it affects the width of the table
func NewWithEpsCount(confidence, eps float64) *CountMin {
	if confidence >= 1.0 {
		confidence = 0.99999
	}
	cm := &CountMin{
		eps:        eps,
		confidence: confidence,
		width:      int(math.Ceil(float64(2.0) / eps)),
		depth:      int(math.Ceil(-math.Log(1-confidence) / math.Log(2))),
	}
	cm.initTable()
	return cm
}

func (cm *CountMin) initTable() {
	cm.table = make([][]int64, cm.depth)
	for i := range cm.table {
		cm.table[i] = make([]int64, cm.width)
	}
}

// RelativeError returns epsilon of CountMin
func (cm *CountMin) RelativeError() float64 {
	return cm.eps
}

// Confidence returns confidence of CountMin
func (cm *CountMin) Confidence() float64 {
	return cm.confidence
}

// Size returns total size of CountMin
func (cm *CountMin) Size() int64 {
	return cm.size
}

// Add inserts an item in bytes and a count associated
// which will also increment size of CountMin
func (cm *CountMin) Add(item []byte, count int64) {
	if count < 0 {
		return
	}
	hashed := cm.hasher(item)
	for i := 0; i < cm.depth; i++ {
		cm.table[i][hashed[i]] += count
	}
	cm.size += count
}

// Count gets an associated total count for an item
// in the CountMin
func (cm *CountMin) Count(item []byte) int64 {
	var answer int64 = MAXINT64
	var val int64
	hashed := cm.hasher(item)
	for i := 0; i < cm.depth; i++ {
		val = cm.table[i][hashed[i]]
		if val < answer {
			answer = val
		}
	}
	return answer
}

// Merge two CountMin tables and return a new table of their union
func Merge(cm1, cm2 *CountMin) (*CountMin, error) {
	if cm1.depth != cm2.depth {
		return nil, errors.New("Depth does not match for merging")
	}
	if cm1.width != cm2.width {
		return nil, errors.New("Width does not match for merging")
	}
	newCm := New(cm1.depth, cm2.width)
	var val int64
	for i := 0; i < newCm.depth; i++ {
		for j := 0; j < newCm.width; j++ {
			val = cm1.table[i][j] + cm2.table[i][j]
			newCm.table[i][j] = val
		}
	}
	newCm.size = cm1.size + cm2.size
	return newCm, nil
}

func (cm *CountMin) hasher(item []byte) []int64 {
	total := cm.depth
	max := int64(cm.width)
	result := make([]int64, total)
	hash := fnv.New32a()
	hash.Write(item)
	sum := int64(hash.Sum32())
	for i := 0; i < total; i++ {
		result[i] = int64(math.Abs(float64((sum * int64(i)) % max)))
	}
	return result
}
