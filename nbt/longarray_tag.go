package nbt

import (
	"errors"
	"fmt"
	"reflect"
)

const longArrayTypeId longArrayType = 12

type longArrayType int8

type LongArrayTag struct {
	Value []int64
}

func (_ longArrayType) Read(reader Reader) (Tag, error) {
	length, err := reader.readInt32()

	if err != nil {
		return nil, err
	}

	data := make([]int64, length)

	for i, _ := range data {
		value, err := reader.readInt64()

		if err != nil {
			return nil, fmt.Errorf("unable to read long array at index %d. Reason: %w", i, err)
		}

		data[i] = value
	}

	return LongArrayTag{
		Value: data,
	}, nil
}

func (_ longArrayType) Write(writer Writer, tag Tag) error {
	data, ok := tag.(LongArrayTag)

	if !ok {
		return errors.New("incompatible tag. Expected LONGARRAY")
	}

	if err := writer.writeInt32(int32(len(data.Value))); err != nil {
		return err
	}

	for i, value := range data.Value {
		if err := writer.writeInt64(value); err != nil {
			return fmt.Errorf("unable to write long array at index %d. Reason: %w", i, err)
		}
	}

	return nil
}

func (_ longArrayType) GetId() int8 {
	return int8(longArrayTypeId)
}

func (_ LongArrayTag) getDataType() dataType {
	return longTypeId
}

func (dtype longArrayType) Decode(tag Tag, value reflect.Value) error {
	data, ok := tag.(LongArrayTag)
	if !ok {
		return fmt.Errorf("unable to unmarshal tag with datatype %d using datatype %d", tag.getDataType(), dtype)
	}

	if err := RequireKind(value, reflect.Slice); err != nil {
		return err
	}

	//TODO: Does this actually work?
	if err := RequireKind(value.Elem(), reflect.Int64); err != nil {
		return err
	}

	value.Set(reflect.ValueOf(data.Value))

	return nil
}