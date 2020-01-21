package main

import "testing"

const mega = 1024 * 1024

func Test_parseMemoryString(t *testing.T) {
	tests := []struct {
		name            string
		buildString     string
		wantMemUsed     uint64
		wantMemTotal    uint64
		wantMemMaxUsage uint64
	}{
		{"R6.0", "C5 Heap Health: OK  - Mem used: 18%  - Mem used: 383MB  - Mem total: 2048MB  - Max: 18% - UpdCtr: 60793", 383 * mega, 2048 * mega, 18},
		{"R6.2", "C5 Heap Health: OK  - Mem used: 3%  76MB  (min: 76 max: 76)  - Mem total: 2048MB  - MAX: 3% - UpdCtr: 92205", 76 * mega, 2048 * mega, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMemUsed, gotMemTotal, gotMemMaxUsage := parseMemoryString(tt.buildString)
			if gotMemUsed != tt.wantMemUsed {
				t.Errorf("parseMemoryString() gotMemUsed = %v, want %v", gotMemUsed, tt.wantMemUsed)
			}
			if gotMemTotal != tt.wantMemTotal {
				t.Errorf("parseMemoryString() gotMemTotal = %v, want %v", gotMemTotal, tt.wantMemTotal)
			}
			if gotMemMaxUsage != tt.wantMemMaxUsage {
				t.Errorf("parseMemoryString() gotMemMaxUsage = %v, want %v", gotMemMaxUsage, tt.wantMemMaxUsage)
			}
		})
	}
}

func Test_parseMemoryStringRegex(t *testing.T) {
	tests := []struct {
		name            string
		buildString     string
		wantMemUsed     uint64
		wantMemTotal    uint64
		wantMemMaxUsage uint64
	}{
		{"R6.0", "C5 Heap Health: OK  - Mem used: 18%  - Mem used: 383MB  - Mem total: 2048MB  - Max: 18% - UpdCtr: 60793", 383 * mega, 2048 * mega, 18},
		{"R6.2", "C5 Heap Health: OK  - Mem used: 3%  76MB  (min: 76 max: 76)  - Mem total: 2048MB  - MAX: 3% - UpdCtr: 92205", 76 * mega, 2048 * mega, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMemUsed, gotMemTotal, gotMemMaxUsage := parseMemoryStringRegex(tt.buildString)
			if gotMemUsed != tt.wantMemUsed {
				t.Errorf("parseMemoryString() gotMemUsed = %v, want %v", gotMemUsed, tt.wantMemUsed)
			}
			if gotMemTotal != tt.wantMemTotal {
				t.Errorf("parseMemoryString() gotMemTotal = %v, want %v", gotMemTotal, tt.wantMemTotal)
			}
			if gotMemMaxUsage != tt.wantMemMaxUsage {
				t.Errorf("parseMemoryString() gotMemMaxUsage = %v, want %v", gotMemMaxUsage, tt.wantMemMaxUsage)
			}
		})
	}
}

func Benchmark_parseMemoryString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		parseMemoryString("C5 Heap Health: OK  - Mem used: 18%  - Mem used: 383MB  - Mem total: 2048MB  - Max: 18% - UpdCtr: 60793")
		parseMemoryString("C5 Heap Health: OK  - Mem used: 3%  76MB  (min: 76 max: 76)  - Mem total: 2048MB  - MAX: 3% - UpdCtr: 92205")
	}
}

func Benchmark_parseMemoryStringRegex(b *testing.B) {
	for n := 0; n < b.N; n++ {
		parseMemoryStringRegex("C5 Heap Health: OK  - Mem used: 18%  - Mem used: 383MB  - Mem total: 2048MB  - Max: 18% - UpdCtr: 60793")
		parseMemoryStringRegex("C5 Heap Health: OK  - Mem used: 3%  76MB  (min: 76 max: 76)  - Mem total: 2048MB  - MAX: 3% - UpdCtr: 92205")
	}
}
