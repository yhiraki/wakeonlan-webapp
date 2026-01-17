package config

import (
	"reflect"
	"testing"
)

func TestParseTargets(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		want    []Target
		wantErr bool
	}{
		{
			name: "valid single target",
			args: []string{"pc1=00:11:22:33:44:55"},
			want: []Target{
				{Name: "pc1", MAC: "00:11:22:33:44:55"},
			},
			wantErr: false,
		},
		{
			name: "valid multiple targets",
			args: []string{"pc1=00:11:22:33:44:55", "pc2=aa:bb:cc:dd:ee:ff"},
			want: []Target{
				{Name: "pc1", MAC: "00:11:22:33:44:55"},
				{Name: "pc2", MAC: "aa:bb:cc:dd:ee:ff"},
			},
			wantErr: false,
		},
		{
			name: "invalid format - missing separator",
			args: []string{"invalid"},
			want: nil,
			wantErr: true,
		},
		{
			name: "invalid format - missing mac",
			args: []string{"name="},
			want: nil,
			wantErr: true,
		},
		{
			name: "invalid mac address",
			args: []string{"pc1=zz:zz:zz:zz:zz:zz"},
			want: nil,
			wantErr: true,
		},
		{
			name: "invalid mac length",
			args: []string{"pc1=00:11:22"},
			want: nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseTargets(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseTargets() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseTargets() = %v, want %v", got, tt.want)
			}
		})
	}
}
