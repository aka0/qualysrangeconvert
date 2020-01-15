package main

import "testing"

func Test_convertRange(t *testing.T) {
	type args struct {
		qualyRange string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Qualys Network Range to NMAP Range", args: args{qualyRange: "127.0.0.1-127.0.0.254"}, want: "127.0.0.1-254"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertRange(tt.args.qualyRange); got != tt.want {
				t.Errorf("convertRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
