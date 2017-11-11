package math

import "testing"

func TestAverage(t *testing.T) {
	_, _, avg, _ := GetStats([]float32{1, 2, 3})
	if avg != 2 {
		t.Error("Expected 2, got ", avg)
	}
}

func TestMin(t *testing.T) {
	_, min, _, _ := GetStats([]float32{1, 2, 3})
	if min != 1 {
		t.Error("Expected 1, got ", min)
	}
}

func TestMax(t *testing.T) {
	max, _, _, _ := GetStats([]float32{1, 2, 3})
	if max != 3 {
		t.Error("Expected 3, got ", max)
	}
}

func TestErr(t *testing.T) {
	_, _, _, err := GetStats(nil)
	if err != 1 {
		t.Error("Expected 1, got ", err)
	}
}

func TestErr2(t *testing.T) {
	_, _, _, err := GetStats([]float32{})
	if err != 1 {
		t.Error("Expected 1, got ", err)
	}
}

func BenchmarkStats(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetStats([]float32{1, 2, 3, 34, 13, 35, 42534.2, 3, 3423})
	}
}
