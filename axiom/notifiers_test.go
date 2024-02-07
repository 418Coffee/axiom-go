package axiom

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func TestNotifiersService_List(t *testing.T) {
	exp := []*Notifier{
		{
			DisabledUntil: "",
			ID:            "test",
			Name:          "test",
			Properties: NotifierProperties{
				Email: &EmailConfig{
					Emails: []string{"test@test.com"},
				},
			},
			Type: "email",
		},
	}

	hf := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)

		w.Header().Set("Content-Type", mediaTypeJSON)
		_, err := fmt.Fprint(w, `[{
			"id": "test",
			"name": "test",
			"properties": {
				"email": {
					"Emails": [
						"test@test.com"
					]
				}
			},
			"type": "email"
		}]`)
		assert.NoError(t, err)
	}
	client := setup(t, "/v2/notifiers", hf)

	res, err := client.Notifiers.List(context.Background())
	require.NoError(t, err)

	assert.Equal(t, exp, res)
}

func TestNotifiersService_Get(t *testing.T) {
	exp := &Notifier{
		DisabledUntil: "",
		ID:            "test",
		Name:          "test",
		Properties: NotifierProperties{
			Email: &EmailConfig{
				Emails: []string{"test@test.com"},
			},
		},
		Type: "email",
	}

	hf := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)

		w.Header().Set("Content-Type", mediaTypeJSON)
		_, err := fmt.Fprint(w, `{
			"id": "test",
			"name": "test",
			"properties": {
				"email": {
					"Emails": [
						"test@test.com"
					]
				}
			},
			"type": "email"
		}`)
		assert.NoError(t, err)
	}
	client := setup(t, "/v2/notifiers/test", hf)

	res, err := client.Notifiers.Get(context.Background(), "test")
	require.NoError(t, err)

	assert.Equal(t, exp, res)
}

func TestNotifiersService_Create(t *testing.T) {
	exp := &Notifier{
		DisabledUntil: "",
		ID:            "test",
		Name:          "test",
		Properties: NotifierProperties{
			Email: &EmailConfig{
				Emails: []string{"test@test.com"},
			},
		},
		Type: "email",
	}
	hf := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, mediaTypeJSON, r.Header.Get("Content-Type"))

		w.Header().Set("Content-Type", mediaTypeJSON)
		_, err := fmt.Fprint(w, `{
			"id": "test",
			"name": "test",
			"properties": {
				"email": {
					"Emails": [
						"test@test.com"
					]
				}
			},
			"type": "email"
		}`)
		assert.NoError(t, err)
	}
	client := setup(t, "/v2/notifiers", hf)

	res, err := client.Notifiers.Create(context.Background(), Notifier{
		Name:          "test",
		DisabledUntil: "",
		Properties: NotifierProperties{
			Email: &EmailConfig{
				Emails: []string{"test@test.com"},
			},
		},
		Type: "email",
	})
	require.NoError(t, err)

	assert.Equal(t, exp, res)
}

func TestNotifiersService_Update(t *testing.T) {
	exp := &Notifier{
		DisabledUntil: "",
		ID:            "test",
		Name:          "test",
		Properties: NotifierProperties{
			Email: &EmailConfig{
				Emails: []string{"test@test.com"},
			},
		},
		Type: "email",
	}
	hf := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Equal(t, mediaTypeJSON, r.Header.Get("Content-Type"))

		w.Header().Set("Content-Type", mediaTypeJSON)
		_, err := fmt.Fprint(w, `{
			"id": "test",
			"name": "test",
			"properties": {
				"email": {
					"Emails": [
						"test@test.com"
					]
				}
			},
			"type": "email"
		}`)
		assert.NoError(t, err)
	}
	client := setup(t, "/v2/notifiers/test", hf)

	res, err := client.Notifiers.Update(context.Background(), "test", Notifier{
		DisabledUntil: "",
		Name:          "test",
		Properties: NotifierProperties{
			Email: &EmailConfig{
				Emails: []string{"test@test.com"},
			},
		},
		Type: "email",
	})
	require.NoError(t, err)

	assert.Equal(t, exp, res)
}

func TestNotifiersService_Delete(t *testing.T) {
	hf := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)

		w.WriteHeader(http.StatusNoContent)
	}

	client := setup(t, "/v2/notifiers/testID", hf)

	err := client.Notifiers.Delete(context.Background(), "testID")
	require.NoError(t, err)
}