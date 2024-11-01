package bf_json

import (
    "encoding/json"
    "io"
)

type BFDecoder struct {
    reader io.Reader
}

func NewBFDecoder(r io.Reader) *BFDecoder {
    return &BFDecoder{reader: r}
}

func (d *BFDecoder) Model(target interface{}) (map[string]bool, error) {
    var rawData map[string]json.RawMessage
    decoder := json.NewDecoder(d.reader)
    if err := decoder.Decode(&rawData); err != nil {
        return nil, err
    }

    fields := make(map[string]bool)
    for k := range rawData {
        fields[k] = true
    }

    data, err := json.Marshal(rawData)
    if err != nil {
        return nil, err
    }

    err = json.Unmarshal(data, target)
    return fields, err
}
