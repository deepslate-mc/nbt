package nbt

import (
	"errors"
	"fmt"
	"reflect"
)

const stringTypeId stringType = 8

type stringType int8

type StringTag struct {
	Value string
}

func (_ stringType) Read(reader Reader) (Tag, error) {
	data, err := reader.readString()

	if err != nil {
		return nil, err
	}

	return StringTag{
		Value: data,
	}, nil
}

func (_ stringType) Write(writer Writer, tag Tag) error {
	data, ok := tag.(StringTag)

	if !ok {
		return errors.New("incompatible tag. Expected STRING")
	}

	return writer.writeString(data.Value)
}

func (_ stringType) GetId() int8 {
	return int8(stringTypeId)
}

func (_ StringTag) getDataType() dataType {
	return stringTypeId
}

func (dtype stringType) Decode(tag Tag, value reflect.Value) error {
	data, ok := tag.(StringTag)
	if !ok {
		return fmt.Errorf("unable to unmarshal tag with datatype %d using datatype %d", tag.getDataType(), dtype)
	}

	err := RequireKind(value, reflect.String)
	if err != nil {
		return err
	}

	value.SetString(data.Value)

	return nil
}