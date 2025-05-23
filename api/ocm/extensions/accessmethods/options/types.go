package options

import (
	"fmt"

	"ocm.software/ocm/api/utils/cobrautils/flagsets"
)

type OptionType interface {
	flagsets.ConfigOptionType
	ValueType() string
	GetDescriptionText() string
}

type base = flagsets.ConfigOptionType

type option struct {
	base
	valueType string
}

func (o *option) Equal(t flagsets.ConfigOptionType) bool {
	if optionType, ok := t.(*option); ok {
		return o.valueType == optionType.valueType && o.GetName() == optionType.GetName()
	}
	return false
}

func (o *option) ValueType() string {
	return o.valueType
}

func (o *option) GetDescription() string {
	return fmt.Sprintf("[*%s*] %s", o.ValueType(), o.base.GetDescription())
}

func (o *option) GetDescriptionText() string {
	return o.base.GetDescription()
}

////////////////////////////////////////////////////////////////////////////////

func NewStringOptionType(name, desc string) OptionType {
	return &option{
		base:      flagsets.NewStringOptionType(name, desc),
		valueType: TYPE_STRING,
	}
}

func NewStringArrayOptionType(name, desc string) OptionType {
	return &option{
		base:      flagsets.NewStringArrayOptionType(name, desc),
		valueType: TYPE_STRINGARRAY,
	}
}

func NewIntOptionType(name, desc string) OptionType {
	return &option{
		base:      flagsets.NewIntOptionType(name, desc),
		valueType: TYPE_INT,
	}
}

func NewBoolOptionType(name, desc string) OptionType {
	return &option{
		base:      flagsets.NewBoolOptionType(name, desc),
		valueType: TYPE_BOOL,
	}
}

func NewYAMLOptionType(name, desc string) OptionType {
	return &option{
		base:      flagsets.NewYAMLOptionType(name, desc),
		valueType: TYPE_YAML,
	}
}

func NewValueMapYAMLOptionType(name, desc string) OptionType {
	return &option{
		base:      flagsets.NewValueMapYAMLOptionType(name, desc),
		valueType: TYPE_STRINGMAPYAML,
	}
}

func NewValueMapOptionType(name, desc string) OptionType {
	return &option{
		base:      flagsets.NewValueMapOptionType(name, desc),
		valueType: TYPE_STRING2YAML,
	}
}

func NewStringMapOptionType(name, desc string) OptionType {
	return &option{
		base:      flagsets.NewStringMapOptionType(name, desc),
		valueType: TYPE_STRING2STRING,
	}
}

func NewStringSliceMapOptionType(name, desc string) OptionType {
	return &option{
		base:      flagsets.NewStringSliceMapOptionType(name, desc),
		valueType: TYPE_STRING2STRINGSLICE,
	}
}

func NewStringSliceMapColonOptionType(name, desc string) OptionType {
	return &option{
		base:      flagsets.NewStringSliceMapColonOptionType(name, desc),
		valueType: TYPE_STRINGCOLONSTRINGSLICE,
	}
}

func NewBytesOptionType(name, desc string) OptionType {
	return &option{
		base:      flagsets.NewBytesOptionType(name, desc),
		valueType: TYPE_BYTES,
	}
}

func NewIdentityPathOptionType(name, desc string) OptionType {
	return &option{
		base:      flagsets.NewIdentityPathOptionType(name, desc),
		valueType: TYPE_IDENTITYPATH,
	}
}
