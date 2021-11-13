package nbt

import (
	"errors"
	"fmt"
	"reflect"
)

const longTypeId longType = 4

type longType int8

type LongTag struct {
	Value int64
}

func (_ longType) Read(reader Reader) (Tag, error) {
	data, err := reader.readInt64()

	if err != nil {
		return nil, err
	}

	return LongTag{
		Value: data,
	}, nil
}

func (_ longType) Write(writer Writer, tag Tag) error {
	data, ok := tag.(LongTag)

	if !ok {
		return errors.New("incompatible tag. Expected LONG")
	}

	return writer.writeInt64(data.Value)
}

func (_ longType) GetId() int8 {
	return int8(longTypeId)
}

func (_ LongTag) getDataType() dataType {
	return longTypeId
}

func (dtype longType) Decode(tag Tag, value reflect.Value) error {
	data, ok := tag.(LongTag)
	if !ok {
		return fmt.Errorf("unable to unmarshal tag with datatype %d using datatype %d", tag.getDataType(), dtype)
	}

	err := RequireKind(value, reflect.Int64)
	if err != nil {
		return err
	}

	value.SetInt(data.Value)

	return nil
}