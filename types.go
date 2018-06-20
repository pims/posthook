package posthook

import (
	"encoding/json"
	"time"
)

// Filters lets you filter a list query
type Filters struct {
	Limit         *int       `url:"limit,omitempty"`
	SortBy        *string    `url:"sortBy,omitempty"`
	SortOrder     *string    `url:"sortOrder,omitempty"`
	PostAtBefore  *time.Time `url:"postAtBefore,omitempty"`
	PostAtAfter   *time.Time `url:"postAtAfter,omitempty"`
	CreatedBefore *time.Time `url:"createdBefore,omitempty"`
	CreatedAfter  *time.Time `url:"createdAfter,omitempty"`
}

// Hook represents a Hook
type Hook struct {
	ID        string      `json:"id,omitempty"`
	Path      string      `json:"path"`
	PostAt    time.Time   `json:"postAt"`
	Status    string      `json:"status,omitempty"`
	CreatedAt *time.Time  `json:"createdAt,omitempty"`
	Data      interface{} `json:"data"`
}

// singleResponse is used to deserialize a single hook response
// { "data": {} }
type singleResponse struct {
	Hook Hook `json:"data"`
}

// listResponse is used to deserialize a list of hooks response
// { "data": [] }
type listResponse struct {
	Hooks []Hook `json:"data"`
}

func single(body []byte) (*Hook, error) {
	response := singleResponse{}
	err := json.Unmarshal(body, &response)
	return &response.Hook, err
}

func list(body []byte) ([]Hook, error) {
	response := listResponse{}
	err := json.Unmarshal(body, &response)
	return response.Hooks, err
}
