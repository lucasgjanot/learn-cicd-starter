package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		input   http.Header
		want    string
		wantErr bool
	}{
		{
			name: "valid key",
			input: http.Header{
				"Authorization": []string{"ApiKey testapikey"},
			},
			want:    "testapikey",
			wantErr: false,
		},
		{
			name: "malformed header",
			input: http.Header{
				"Authorization": []string{"ApiKeytestapikey"},
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "inexistent header",
			input: http.Header{},
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.input)

			if (err != nil) != tt.wantErr {
				t.Fatalf("error = %v, wantErr %v", err, tt.wantErr)
			}

			if got != tt.want {
				t.Fatalf("got %q, want %q", got, tt.want)
			}
		})
	}
}
