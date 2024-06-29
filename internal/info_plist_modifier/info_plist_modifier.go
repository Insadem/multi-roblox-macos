package info_plist_modifier

import (
	"bytes"
	"os"

	"howett.net/plist"
)

func SetMultipleInstancesProhibition(prohibited bool) error {
	filePath := "/Applications/Roblox.app/Contents/Info.plist"
	openedFile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer openedFile.Close()

	decoder := plist.NewDecoder(openedFile)

	var val map[string]interface{} = nil
	decoder.Decode(&val)

	val["LSMultipleInstancesProhibited"] = prohibited

	buf := &bytes.Buffer{}
	encoder := plist.NewEncoderForFormat(buf, decoder.Format)
	err = encoder.Encode(&val)
	if err != nil {
		return err
	}

	os.WriteFile(filePath, buf.Bytes(), 0644)

	err = openedFile.Sync()
	if err != nil {
		return err
	}

	return nil
}
