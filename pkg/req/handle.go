package req

import "net/http"

func HandleBody[T any](w *http.ResponseWriter, r *http.Request) (*T, error) {
	body, err := decode[T](r.Body)
	if err != nil {
		return nil, err
	}

	err = validate(body)
	if err != nil {
		return nil, err
	}

	return &body, err
}
