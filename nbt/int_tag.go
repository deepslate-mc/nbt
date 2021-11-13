package nbt

import (
	"errors"
	"fmt"
	"reflect"
)

const intTypeId intType = 3

type intType int8

type IntTag struct {
	Value int32
}

func (_ intType) Read(reader Reader) (Tag, error) {
	data, err := reader.readInt32()

	if err != nil {
		return nil, err
	}

	return IntTag{
		Value: data,
	}, nil
}

func (_ intType) Write(writer Writer, tag Tag) error {
	data, ok := tag.(IntTag)

	if !ok {
		return errors.New("incompatible tag. Expected INT")
	}

	return writer.writeInt32(data.Value)
}

func (_ intType) GetId() int8 {
	return int8(intTypeId)
}

func (_ IntTag) getDataType() dataType {
	return intTypeId
}

func (dtype intType) Decode(tag Tag, value reflect.Value) error {
	data, ok := tag.(IntTag)
	if !ok {
		return fmt.Errorf("unable to unmarshal tag with datatype %d using datatype %d", tag.getDataType(), dtype)
	}

	err := RequireKind(value, reflect.Int32)
	if err != nil {
		return err
	}

	value.SetInt(int64(data.Value))

	return nil
}