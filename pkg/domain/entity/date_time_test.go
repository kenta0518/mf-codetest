package entity

import (
	"database/sql/driver"
	"reflect"
	"testing"
	"time"
)

func TestDateTime_UnmarshalYAML(t *testing.T) {
	type fields struct {
		Time time.Time
	}
	type args struct {
		unmarshal func(interface{}) error
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    time.Time
		wantErr bool
	}{
		{
			name: "YAMLからパースできるか",
			args: args{
				unmarshal: func(i interface{}) error {
					v := (i).(*string)
					*v = "2023/04/05 12:34:56"
					return nil
				},
			},
			want:    time.Date(2023, 4, 5, 12, 34, 56, 0, time.Local),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &DateTime{
				Time: tt.fields.Time,
			}
			if err := tr.UnmarshalYAML(tt.args.unmarshal); (err != nil) != tt.wantErr {
				t.Errorf("DateTime.UnmarshalYAML() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tr.Time.Equal(tt.want) {
				t.Errorf("DateTime.UnmarshalYAML() error = %v, wantErr %v", tr.Time, tt.want)
			}
		})
	}
}

func TestDateTime_MarshalYAML(t *testing.T) {
	type fields struct {
		Time time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		{
			name: "YAMLに変換できるか",
			fields: fields{
				Time: time.Date(2023, 4, 5, 12, 34, 56, 0, time.Local),
			},
			want:    "2023/04/05 12:34:56",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &DateTime{
				Time: tt.fields.Time,
			}
			got, err := tr.MarshalYAML()
			if (err != nil) != tt.wantErr {
				t.Errorf("DateTime.MarshalYAML() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DateTime.MarshalYAML() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDateTime_Scan(t *testing.T) {
	type fields struct {
		Time time.Time
	}
	type args struct {
		value interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    DateTime
		wantErr bool
	}{
		{
			name: "Scanできるか",
			args: args{
				value: time.Date(2023, 4, 5, 12, 34, 56, 0, time.Local),
			},
			want: DateTime{
				Time: time.Date(2023, 4, 5, 12, 34, 56, 0, time.Local),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &DateTime{
				Time: tt.fields.Time,
			}
			if err := tr.Scan(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("DateTime.Scan() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tr.Time.Equal(tt.want.Time) {
				t.Errorf("DateTime.Scan() error = %v, wantErr %v", tr.Time, tt.want.Time)
			}
		})
	}
}

func TestDateTime_Value(t *testing.T) {
	type fields struct {
		Time time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		want    driver.Value
		wantErr bool
	}{
		{
			name: "Valueできるか",
			fields: fields{
				Time: time.Date(2023, 4, 5, 12, 34, 56, 0, time.Local),
			},
			want:    time.Date(2023, 4, 5, 12, 34, 56, 0, time.Local),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := DateTime{
				Time: tt.fields.Time,
			}
			got, err := tr.Value()
			if (err != nil) != tt.wantErr {
				t.Errorf("DateTime.Value() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DateTime.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}
