package urlParser

import (
	"reflect"
	"testing"
	"time"
)

func TestNewUrlParser(t *testing.T) {
	type args struct {
		url     string
		target  string
		timeout time.Duration
	}
	tests := []struct {
		name string
		args args
		want *UrlParser
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUrlParser(tt.args.url, tt.args.target, tt.args.timeout); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUrlParser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUrlParser_GetSourceName(t *testing.T) {
	type fields struct {
		url     string
		target  string
		timeout time.Duration
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			up := &UrlParser{
				url:     tt.fields.url,
				target:  tt.fields.target,
				timeout: tt.fields.timeout,
			}
			if got := up.GetSourceName(); got != tt.want {
				t.Errorf("GetSourceName() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
		// TODO: Add test cases.
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
