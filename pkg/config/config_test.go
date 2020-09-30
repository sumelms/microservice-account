package config

import (
	"reflect"
	"testing"

	_ "github.com/sumelms/microservice-account/tests"
)

func TestNewConfig(t *testing.T) {
	type args struct {
		configPath string
	}
	tests := []struct {
		name    string
		args    args
		want    *Config
		wantErr bool
	}{
		{name: "Invalid path", args: args{configPath: "config.yml"}, wantErr: true},
		{name: "Correct path", args: args{configPath: "config/config.yml"}, wantErr: false},
		// @TODO Check with sampleConfig
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewConfig(tt.args.configPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConfig() got = %v, want %v", got, tt.want)
			}
		})
	}
}
