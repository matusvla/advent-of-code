package boardingpass

import "testing"

func TestTranslate(t *testing.T) {
	type args struct {
		code []byte
	}
	tests := []struct {
		name    string
		args    args
		wantRow int
		wantCol int
	}{
		{
			name: "success",
			args: args{
				code: []byte("FBFBBFFRLR"),
			},
			wantRow: 44,
			wantCol: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRow, gotCol := Translate(tt.args.code)
			if gotRow != tt.wantRow {
				t.Errorf("Translate() gotRow = %v, want %v", gotRow, tt.wantRow)
			}
			if gotCol != tt.wantCol {
				t.Errorf("Translate() gotCol = %v, want %v", gotCol, tt.wantCol)
			}
		})
	}
}
