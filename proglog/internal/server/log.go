package server

import (
	"fmt"
	"sync"
)

type Record struct {
	Value  []byte `json:"value"`
	Offset uint64 `json:"offset"`
}

type Log struct {
	mu      sync.Mutex
	records []Record
}

var ErrOffsetNotFound = fmt.Errorf("offset not found")

// NewLog instantiates a Log
func NewLog() *Log {
	return &Log{}
}

// Append adds a record to the end of the existing ones
func (c *Log) Append(record Record) (uint64, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	record.Offset = uint64(len(c.records))
	c.records = append(c.records, record)
	return record.Offset, nil
}

// Read will read a log record at the specified 0-based index
func (c *Log) Read(offset uint64) (Record, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if offset >= uint64(len(c.records)) {
		return Record{}, ErrOffsetNotFound
	}
	return c.records[offset], nil
}
