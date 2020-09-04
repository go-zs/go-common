package date

import "testing"

var (
	testCases = []struct{
		mysqlDate string
		utcDate string
		nanoTs int64
		milliTs int64
	}{
		{
			mysqlDate: "",
			utcDate:   "",
			nanoTs:    0,
			milliTs:   0,
		},
	}
)


func TestParseMysqlDateMilliTs(t *testing.T) {
	
}

func TestParseMysqlDateNanoTs(t *testing.T) {
	type args struct {
		ds string
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseMysqlDateNanoTs(tt.args.ds)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseMysqlDateNanoTs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseMysqlDateNanoTs() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseUtcDateMilliTs(t *testing.T) {
	type args struct {
		ds string
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseUtcDateMilliTs(tt.args.ds)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseUtcDateMilliTs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseUtcDateMilliTs() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseUtcDateNanoTs(t *testing.T) {
	type args struct {
		ds string
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseUtcDateNanoTs(tt.args.ds)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseUtcDateNanoTs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseUtcDateNanoTs() got = %v, want %v", got, tt.want)
			}
		})
	}
}