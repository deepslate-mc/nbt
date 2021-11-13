package nbt

import (
	"errors"
	"fmt"
	"reflect"
)

const byteArrayTypeId byteArrayType = 7

type byteArrayType int8

type ByteArrayTag struct {
	Value []int8
}

func (_ byteArrayType) Read(reader Reader) (Tag, error) {
	data, err := reader.readByteArray()

	if err != nil {
		return nil, err
	}

	return ByteArrayTag{
		Value: data,
	}, nil
}

func (_ byteArrayType) Write(writer Writer, tag Tag) error {
	data, ok := tag.(ByteArrayTag)

	if !ok {
		return errors.New("incompatible tag. Expected BYTE_ARRAY")
	}

	return writer.writeByteArray(data.Value)
}

func (_ byteArrayType) GetId() int8 {
	return int8(byteArrayTypeId)
}

func (_ ByteArrayTag) getDataType() dataType {
	return byteArrayTypeId
}

func (dtype byteArrayType) Decode(tag Tag, value reflect.Value) error {
	data, ok := tag.(ByteArrayTag)
	if !ok {
		return fmt.Errorf("unable to unmarshal tag with datatype %d using datatype %d", tag.getDataType(), dtype)
	}

	if err := RequireKind(value, reflect.Slice); err != nil {
		return err
	}

	//TODO: Does this actually work?
	if err := RequireKind(value.Elem(), reflect.Int8); err != nil {
		return err
	}

	value.Set(reflect.ValueOf(data.Value))

	return nil
}