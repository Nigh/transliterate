package transliterate

import "sync"

var memoryPool = sync.Pool{
	New: func() interface{} {
		m := make([]byte, 1024)
		return &m
	},
}
