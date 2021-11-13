package nbt

import (
	"errors"
	"fmt"
	"reflect"
)

const floatTypeId floatType = 5

type floatType int8

type FloatTag struct {
	Value float32
}

func (_ floatType) Read(reader Reader) (Tag, error) {
	data, err := reader.readFloat32()

	if err != nil {
		return nil, err
	}

	return FloatTag{
		Value: data,
	}, nil
}

func (_ floatType) Write(writer Writer, tag Tag) error {
	data, ok := tag.(FloatTag)

	if !ok {
		return errors.New("incompatible tag. Expected FLOAT")
	}

	return writer.writeFloat32(data.Value)
}

func (_ floatType) GetId() int8 {
	return int8(floatTypeId)
}

func (_ FloatTag) getDataType() dataType {
	return floatTypeId
}

func (dtype floatType) Decode(tag Tag, value reflect.Value) error {
	data, ok := tag.(FloatTag)
	if !ok {
		return fmt.Errorf("unable to unmarshal tag with datatype %d using datatype %d", tag.getDataType(), dtype)
	}

	err := RequireKind(value, reflect.Float32)
	if err != nil {
		return err
	}

	value.SetFloat(float64(data.Value))

	return nil
}