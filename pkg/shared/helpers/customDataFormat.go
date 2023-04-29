package helpers

import "github.com/livelists/livelist-server/contracts/wsMessages"

func CustomDataFormat(data *map[string]string) *wsMessages.CustomData {
	if data == nil {
		return nil
	}
	return &wsMessages.CustomData{
		Data: *data,
	}
}
