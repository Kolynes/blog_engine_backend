package server

type JsonResponseError struct {
	Summary string              `json:"summary"`
	Fields  map[string][]string `json:"fields,omitempty"`
}

type JsonResponse[T any] struct {
	Status        int               `json:"status"`
	Data          T                 `json:"data,omitempty"`
	NumberOfPages uint              `json:"numberOfPage,omitempty"`
	NextPage      uint              `json:"nextPage,omitempty"`
	PreviousPage  uint              `json:"previousPage,omitempty"`
	Error         JsonResponseError `json:"error,omitempty"`
}
