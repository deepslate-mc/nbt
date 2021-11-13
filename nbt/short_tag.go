package nbt

import (
	"errors"
	"fmt"
	"reflect"
)

const shortTypeId shortType = 2

type shortType int8

type ShortTag struct {
	Value int16
}

func (_ shortType) Read(reader Reader) (Tag, error) {
	data, err := reader.readInt16()

	if err != nil {
		return nil, err
	}

	return ShortTag{
		Value: data,
	}, nil
}

func (_ shortType) Write(writer Writer, tag Tag) error {
	data, ok := tag.(ShortTag)

	if !ok {
		return errors.New("incompatible tag. Expected SHORT")
	}

	return writer.writeInt16(data.Value)
}

func (_ shortType) GetId() int8 {
	return int8(shortTypeId)
}

func (_ ShortTag) getDataType() dataType {
	return shortTypeId
}

func (dtype shortType) Decode(tag Tag, value reflect.Value) error {
	data, ok := tag.(ShortTag)
	if !ok {
		return fmt.Errorf("unable to unmarshal tag with datatype %d using datatype %d", tag.getDataType(), dtype)
	}

	err := RequireKind(value, reflect.Int16)
	if err != nil {
		return err
	}

	value.SetInt(int64(data.Value))

	return nil
}