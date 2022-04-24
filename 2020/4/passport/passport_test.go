package passport

import "testing"

func TestProcessAndValidate(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				s: "iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719",
			},
			want: true,
		},
		{
			name: "fail",
			args: args{
				s: "hgt:59cm ecl:zzz\neyr:2038 hcl:74454a iyr:2023\npid:3556412378 byr:2007",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ProcessAndValidate(tt.args.s); got != tt.want {
				t.Errorf("ProcessAndValidate() = %v, want %v", got, tt.want)
			}
		})
	}
}
