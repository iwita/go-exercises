package urlshort

import (
	"net/http"
	"reflect"
	"testing"
)

func TestMapHandler(t *testing.T) {
	type args struct {
		pathToUrls map[string]string
		fallback   http.Handler
	}
	tests := []struct {
		name string
		args args
		want http.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapHandler(tt.args.pathToUrls, tt.args.fallback); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yamlHandler(t *testing.T) {
	type args struct {
		yamlBytes []byte
		fallback  http.Handler
	}
	tests := []struct {
		name    string
		args    args
		want    http.HandlerFunc
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := yamlHandler(tt.args.yamlBytes, tt.args.fallback)
			if (err != nil) != tt.wantErr {
				t.Errorf("yamlHandler() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("yamlHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseYAML(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []pathUrl
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseYAML(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseYAML() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseYAML() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_buildMap(t *testing.T) {
	type args struct {
		pathUrls []pathUrl
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildMap(tt.args.pathUrls); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
