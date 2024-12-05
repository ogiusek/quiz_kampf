package match

import (
	"encoding/json"
	usersdto "lib/app/users/dto"
	"lib/app/ws/repo"
	"lib/common/id"
	"lib/common/valid"
	socketstorage "lib/services/socketStorage"
	"reflect"
)

func Parse[T any](id id.ID, ws socketstorage.Ws, callback func(T) error) func(m socketstorage.Message) {
	return func(m socketstorage.Message) {
		message, _ := json.Marshal(m.Payload)
		var args T
		json.Unmarshal(message, &args)

		reflection := reflect.ValueOf(&args).Elem()
		for i := range reflection.NumField() {
			field := reflection.Field(i)
			if field.IsValid() && field.Type() == reflect.TypeOf(usersdto.Session{}) {
				metadataRepo := repo.GetMetadataRepo()
				metadata, err := metadataRepo.GetBySocket(id)
				if err != nil {
					ws.Close()
					return
				}
				field.Set(reflect.ValueOf(metadata.Session))
			}
		}
		if err := valid.Valid(args); err != nil {
			ws.Send(socketstorage.Message{Topic: m.Topic, Payload: err})
			return
		}
		res := callback(args)
		if res != nil {
			ws.Send(socketstorage.Message{Topic: m.Topic, Payload: res})
		}
	}
}
