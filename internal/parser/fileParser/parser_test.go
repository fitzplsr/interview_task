package fileParser

import (
	"testing"
)

func TestFileParser_GetTargetCount(t *testing.T) {
	type fields struct {
		fileName string
		target   string
	}
	tests := []struct {
		name    string
		fields  fields
		want    int64
		wantErr bool
	}{
		{
			name: "SuccessGetTargetCount",
			fields: fields{
				"/etc/passwd",
				"Go",
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "FailureGetTargetCount",
			fields: fields{
				"nofile",
				"Go",
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fp := &FileParser{
				fileName: tt.fields.fileName,
				target:   tt.fields.target,
			}
			got, err := fp.GetTargetCount()
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
