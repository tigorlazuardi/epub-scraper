package unsafeutils

import "testing"

func TestGetString(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success",
			args: args{
				b: []byte(`holy crap`),
			},
			want: "holy crap",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetString(tt.args.b); got != tt.want {
				t.Errorf("UnsafeGetString() = %v, want %v", got, tt.want)
			}
		})
	}
}
