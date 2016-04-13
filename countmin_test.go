package countmin

import (
	"math"
	"testing"
)

func testCM() *CountMin {
	return New(4, 200)
}

func TestNew(t *testing.T) {
	depth := 1
	width := 1
	cm := New(depth, width)
	if cm.depth != depth {
		t.Errorf("Expected depth to be %d but got %d", depth, cm.depth)
	}
	if cm.width != width {
		t.Errorf("Expected width to be %d but got %d", width, cm.width)
	}
	expectedEps := (2.0 / float64(width))
	if cm.eps != expectedEps {
		t.Errorf("Expected eps to be %f but got %f", expectedEps, cm.eps)
	}
	expectedConfidence := 1.0 - 1.0/math.Pow(2, float64(depth))
	if cm.confidence != expectedConfidence {
		t.Errorf("Expected eps to be %f but got %f", expectedConfidence, cm.confidence)
	}
}

func TestNewInitTable(t *testing.T) {
	cm := testCM()
	if len(cm.table) != cm.depth {
		t.Errorf("Expected table depth of %d but got %d", cm.depth, len(cm.table))
	}
	if len(cm.table[0]) != cm.width {
		t.Errorf("Expected table width of %d but got %d", cm.width, len(cm.table[0]))
	}
}

func TestNewWithEpsCount(t *testing.T) {
	eps := 0.001
	confidence := 0.9
	cm := NewWithEpsCount(confidence, eps)
	if cm.eps != eps {
		t.Errorf("Expected eps to be %f but got %f", eps, cm.eps)
	}
	if cm.confidence != confidence {
		t.Errorf("Expected eps to be %f but got %f", confidence, cm.confidence)
	}
	expectedWidth := int(math.Ceil(float64(2.0) / eps))
	if cm.width != expectedWidth {
		t.Errorf("Expected depth to be %d but got %d", expectedWidth, cm.width)
	}
	expectedDepth := int(math.Ceil(-math.Log(1-confidence) / math.Log(2)))
	if cm.depth != expectedDepth {
		t.Errorf("Expected width to be %d but got %d", expectedDepth, cm.depth)
	}
}

func TestNewWithEpsCountConfidenceGreaterThan1(t *testing.T) {
	eps := 0.001
	confidence := 1.2
	cm := NewWithEpsCount(confidence, eps)
	expectedConfidence := 0.99999
	if cm.confidence != expectedConfidence {
		t.Errorf("Expected eps to be %f but got %f", expectedConfidence, cm.confidence)
	}
}

func TestRelativeError(t *testing.T) {
	cm := testCM()
	if cm.eps != cm.RelativeError() {
		t.Errorf("Expected RelativeError to be %f but got %f", cm.eps, cm.RelativeError())
	}
}

func TestConfidence(t *testing.T) {
	cm := testCM()
	if cm.confidence != cm.Confidence() {
		t.Errorf("Expected Confidence to be %f but got %f", cm.confidence, cm.Confidence())
	}
}

func TestSize(t *testing.T) {
	cm := testCM()
	if cm.size != cm.Size() {
		t.Errorf("Expected Size to be %d but got %d", cm.size, cm.Size())
	}
}

func TestAddNegativeCount(t *testing.T) {
	cm := testCM()
	cm.Add([]byte("item"), -1)
	if cm.size >= 1 {
		t.Errorf("Should not be able to add negative counts")
	}
}

func TestAdd(t *testing.T) {
	cm := testCM()
	cm.Add([]byte("item"), int64(1))
	if cm.size != 1 {
		t.Errorf("Size should be %d but is %d", 1, cm.size)
	}

	cm.Add([]byte("item"), int64(10))
	if cm.size != 11 {
		t.Errorf("Size should have increased to %d but is %d", 11, cm.size)
	}

	cm.Add([]byte("another item"), int64(44))
	if cm.size != 55 {
		t.Errorf("Size should have increased to %d but is %d", 55, cm.size)
	}
}

func TestCount(t *testing.T) {
	cm := testCM()
	item1 := []byte("item")
	cm.Add(item1, int64(1))
	cm.Add(item1, int64(10))
	if cm.Count(item1) != 11 {
		t.Errorf("Expected item1 to have a count of 11 but got %d", cm.Count(item1))
	}

	item2 := []byte("another item")
	cm.Add(item2, int64(44))
	if cm.Count(item2) != 44 {
		t.Errorf("Expected item1 to have a count of 44 but got %d", cm.Count(item2))
	}
}

func TestCountOfNonExistentItem(t *testing.T) {
	cm := testCM()
	item1 := []byte("item")
	item2 := []byte("another item")
	missing := []byte("missing")
	cm.Add(item1, int64(1))
	cm.Add(item1, int64(10))
	cm.Add(item2, int64(44))

	if cm.Count(missing) != 0 {
		t.Errorf("Expected missing item to have a count of 0 but got %d", cm.Count(missing))
	}

}

func TestHasher(t *testing.T) {
	cm := testCM()
	item := []byte("item")
	hashed := cm.hasher(item)
	if len(hashed) != cm.depth {
		t.Errorf("Expected %d hash functions but got %d", cm.depth, len(hashed))
	}
}

func TestMergeMismatched(t *testing.T) {
	cm1 := testCM()
	cm2 := New(1, 1)
	_, merge1Error := Merge(cm1, cm2)
	if merge1Error == nil {
		t.Errorf("Expected merge to fail with different depths but got nil")
	}

	cm3 := New(4, 1)
	_, merge2Error := Merge(cm1, cm3)
	if merge2Error == nil {
		t.Errorf("Expected merge to fail with different widths but got nil")
	}
}
