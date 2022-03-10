package unsafeutils

import (
	"reflect"
	"testing"
)

func TestGetBytes(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "success",
			args: args{
				s: "boo",
			},
			want: []byte(`boo`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetBytes(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnsafeGetBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}
