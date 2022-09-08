package main

import (
	"testing"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr error
	}{
		{
			name:    "Test",
			input:   "a4bc2d5e",
			want:    "aaaabccddddde",
			wantErr: nil,
		},
		{
			name:    "Test",
			input:   "abcd",
			want:    "abcd",
			wantErr: nil,
		},
		{
			name:    "Test",
			input:   "45",
			want:    "",
			wantErr: inputErr,
		},
		{
			name:    "Test",
			input:   "",
			want:    "",
			wantErr: nil,
		},
		{
			name:    "Test",
			input:   `qwe\4\5`,
			want:    `qwe45`,
			wantErr: nil,
		},
		{
			name:    "Test",
			input:   `qwe\45`,
			want:    `qwe44444`,
			wantErr: nil,
		},
		{
			name:    "Test",
			input:   `qwe\\5`,
			want:    `qwe\\\\\`,
			wantErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := Unpack(test.input)
			if err != test.wantErr {
				t.Errorf("got %q want %q", err, test.wantErr)
			}
			if got != test.want {
				t.Errorf("got %q want %q", got, test.want)
			}
		})
	}

}
