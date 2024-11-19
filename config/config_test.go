package config

import "testing"

func TestSettings_IsDevelopment(t *testing.T) {
	type fields struct {
		Environment string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "Developmentならtrue",
			fields: fields{Environment: "Development"},
			want:   true,
		},
		{
			name:   "Development以外ならfalse",
			fields: fields{Environment: "Production"},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Settings{
				Environment: tt.fields.Environment,
			}
			if got := s.IsDevelopment(); got != tt.want {
				t.Errorf("Settings.IsDevelopment() = %v, want %v", got, tt.want)
			}
		})
	}
}
