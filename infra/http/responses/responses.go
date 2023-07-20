package responses

import (
	"encoding/json"
	"go-social-action/application/appError"
	"net/http"
)

func ResponseWithErr(w http.ResponseWriter, err error) {
	switch e := err.(type) {
	case appError.FieldsValidationError:
		UnprocessableEntity(w, e)
	case appError.NotFoundError:
		NotFound(w, e)
	default:
		InternalServerError(w, e)
	}
}

func InternalServerError(w http.ResponseWriter, i interface{}) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(i)
}

func UnprocessableEntity(w http.ResponseWriter, i interface{}) {
	w.WriteHeader(http.StatusUnprocessableEntity)
	json.NewEncoder(w).Encode(i)
}

func Created(w http.ResponseWriter, i interface{}) {
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(i)
}

func BadRequest(w http.ResponseWriter, i interface{}) {
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(i)
}

func NotFound(w http.ResponseWriter, i interface{}) {
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(i)
}

func NoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

func Ok(w http.ResponseWriter, i interface{}) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(i)
}
