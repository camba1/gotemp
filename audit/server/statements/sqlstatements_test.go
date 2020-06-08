package statements

import "testing"

func Test_usrSql_String(t *testing.T) {
	tests := []struct {
		name string
		p    auditSql
		want string
	}{
		{name: "Find Statement", p: TestStatement, want: `select 1`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
