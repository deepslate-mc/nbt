package nbt

import (
	"errors"
	"fmt"
	"reflect"
)

const compoundTypeId compoundType = 10

type compoundType int8

type CompoundTag struct {
	Tags map[string]Tag
}

func (_ compoundType) Read(reader Reader) (Tag, error) {
	compound := CompoundTag{
		Tags: map[string]Tag{},
	}

	for {
		name, tag, err := reader.Read()

		if err != nil {
			return nil, err
		}

		if tag.getDataType() == endTypeId {
			break
		}

		compound.Tags[name] = tag
	}

	return compound, nil
}

func (_ compoundType) Write(writer Writer, tag Tag) error {
	data, ok := tag.(CompoundTag)

	if !ok {
		return errors.New("incompatible tag. Expected COMPOUND")
	}

	for name, value := range data.Tags {
		err := writer.Write(name, value)

		if err != nil {
			return err
		}
	}

	err := writer.writeInt8(endTypeId.GetId())
	if err != nil {
		return err
	}

	err = endTypeId.Write(writer, endTag{})
	return err
}

func (_ compoundType) GetId() int8 {
	return int8(compoundTypeId)
}

func (_ CompoundTag) getDataType() dataType {
	return compoundTypeId
}

func (dtype compoundType) Decode(tag Tag, value reflect.Value) error {
	data, ok := tag.(CompoundTag)
	if !ok {
		return fmt.Errorf("unable to unmarshal tag with datatype %d using datatype %d", tag.getDataType(), dtype)
	}

	if err := RequireKind(reflect.Indirect(value), reflect.Struct); err != nil {
		return err
	}

	fields := readFields(reflect.Indirect(value))

	for name, t := range data.Tags {
		if field, ok := fields[name]; ok {
			if err := t.getDataType().Decode(t, field); err != nil {
				return err
			}
		}
	}

	return nil
}

func readFields(v reflect.Value) map[string]reflect.Value {
	fields := make(map[string]reflect.Value)

	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		if !field.IsExported() {
			continue
		}

		tag, ok := field.Tag.Lookup("nbt")

		if !ok {
			fields[field.Name] = v.Field(i)
		} else {
			fields[tag] = v.Field(i)
		}
	}

	return fields
}