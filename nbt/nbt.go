package nbt

import (
	"io"
)


func Read(reader io.Reader) (Tag, error) {
	_, tag, err := NewReader(reader).Read()
	return tag, err
}

func Write(writer io.Writer, tag Tag) error {
	return NewWriter(writer).Write("", tag)
}

func Marshal(v interface{}) ([]byte, error) {
	return nil, nil
}