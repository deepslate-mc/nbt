package nbt

import (
	"errors"
	"fmt"
	"io"
	"reflect"
)

type TagDecoder interface {
	Decode(tag Tag, value reflect.Value) error
}

func Unmarshal(reader io.Reader, v interface{}) error {
	tag, err := Read(reader)
	if err != nil {
		return err
	}

	return decode(tag, tag.getDataType(), v)
}

func decode(tag Tag, decoder TagDecoder, v interface{}) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || v == nil {
		return errors.New("unable to unmarshal non pointer of nil values")
	}

	return decoder.Decode(tag, rv)
}

type IncompatibleKindError struct {
	expected reflect.Kind
	actual reflect.Kind
}

func (err IncompatibleKindError) Error() string {
	return fmt.Sprintf("unable to unmarshal value. Types incompatible: Expected type %s but got %s", err.expected.String(), err.actual.String())
}

func RequireKind(rv reflect.Value, expected reflect.Kind) error {
	if kind := rv.Kind(); kind != expected {
		return &IncompatibleKindError{
			expected: expected,
			actual: kind,
		}
	}

	return nil
}