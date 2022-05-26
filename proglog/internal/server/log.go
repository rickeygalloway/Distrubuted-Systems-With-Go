package server

import (
	"fmt"
	"sync"
)

type Record struct {
	Value  []byte
	Offset uint64
}

type Log struct {
	mu      sync.Mutex
	records []Record
}

func NewLog() *Log {
	return &Log{}
}

func (c *Log) Append(record Record) (uint64, error) {
	c.mu.Lock()

	defer c.mu.Unlock()
	record.Offset = uint64(len(c.records))
	c.records = append(c.records, record)

	return record.Offset, ErrOffsetNotFound

}

func (c *Log) Read(offset uint64) (Record, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if offset > uint64(len(c.records)) {
		return Record{}, nil
	}

	return c.records[offset], nil

}

var ErrOffsetNotFound = fmt.Errorf("offset not found")
