package formats

import (
	"encoding/json"
	"gopkg.in/yaml.v3"
	"io"
)

type FormatJson struct{}

func (f FormatJson) Marshal(v interface{}, out io.Writer) error {
	return json.NewEncoder(out).Encode(v)
}
func (f FormatJson) Unmarshal(in io.Reader) (interface{}, error) {
	v := make(map[string]interface{})
	if err := json.NewDecoder(in).Decode(&v); err != nil {
		return nil, err
	}
	return v, nil
}

type FormatYaml struct{}

func (f FormatYaml) Marshal(v interface{}, out io.Writer) error {
	return yaml.NewEncoder(out).Encode(v)
}
func (f FormatYaml) Unmarshal(in io.Reader) (interface{}, error) {
	v := make(map[string]interface{})
	if err := yaml.NewDecoder(in).Decode(&v); err != nil {
		return nil, err
	}
	return v, nil
}
