package error_handling

import (
	"encoding/json"
	"log"
	"net/http"
)

func LogError(w http.ResponseWriter, vasErr *VASError) {
	if vasErr == nil {
		return
	}
	clientErr := VASError{
		ErrorId:          vasErr.ErrorId,
		ErrorName:        vasErr.ErrorName,
		ErrorDescription: vasErr.ErrorDescription,
		StatusCode:       vasErr.StatusCode,
		GoError:          nil,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(vasErr.StatusCode))
	err := json.NewEncoder(w).Encode(clientErr)
	if err != nil {
		log.Println("[ERROR] Encoding JSON:", err)
		return
	}

	log.Printf("[ERROR] [%d] %s\n", vasErr.StatusCode, vasErr.ErrorId)
	log.Println(vasErr.ErrorName)
	if vasErr.ErrorDescription != nil {
		log.Println(*vasErr.ErrorDescription)
	}
	if vasErr.GoError != nil {
		log.Println("Go error:", vasErr.GoError)
	}
}
