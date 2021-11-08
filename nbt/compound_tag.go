package nbt

import "errors"

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