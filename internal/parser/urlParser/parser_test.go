package urlParser

import (
	"testing"
	"time"
)

func TestUrlParser_GetTargetCount(t *testing.T) {
	type fields struct {
		url     string
		target  string
		timeout time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		want    int64
		wantErr bool
	}{
		{
			name: "SuccessfulGetTargetCount",
			fields: fields{
				"https://golang.org",
				"Go",
				time.Second * 5,
			},
			want:    246,
			wantErr: false,
		},
		{
			name: "TimeExceedGetTargetCount",
			fields: fields{
				"https://golang.org",
				"Go",
				time.Second * 0,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "InvalidUrlTargetCount",
			fields: fields{
				"htt://golang.org",
				"Go",
				time.Second * 5,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			up := &UrlParser{
				url:     tt.fields.url,
				target:  tt.fields.target,
				timeout: tt.fields.timeout,
			}
			got, err := up.GetTargetCount()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTargetCount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetTargetCount() got = %v, want %v", got, tt.want)
			}
		})
	}
}
