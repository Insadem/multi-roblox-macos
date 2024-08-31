package infoplist

import (
	"bytes"
	"os"

	"howett.net/plist" // TODO: swap this dependency with your own implementation
)

func SetMultipleInstancesProhibition(path string, prohibited bool) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	dec := plist.NewDecoder(file)

	var val map[string]interface{} = nil
	dec.Decode(&val)

	val["LSMultipleInstancesProhibited"] = prohibited

	buf := &bytes.Buffer{}
	encoder := plist.NewEncoderForFormat(buf, dec.Format)
	err = encoder.Encode(&val)
	if err != nil {
		return err
	}

	os.WriteFile(path, buf.Bytes(), 0644)

	err = file.Sync()
	if err != nil {
		return err
	}

	return nil
}
