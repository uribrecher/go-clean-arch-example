package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockReaderWriter struct {
	mock.Mock
}

func (r *MockReaderWriter) Read(bytes []byte) (int, error) {
	return 0, nil
}

func (r *MockReaderWriter) Write(bytes []byte) (int, error) {
	return 0, nil
}

func TestEncrypt2(t *testing.T) {
	a := assert.New(t)
	rw := new(MockReaderWriter)

	err := encrypt2(rw, rw)
	a.NoError(err)
}
