package formatter

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func Test_pad(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "PKCS#7パディング",
			args: args{data: make([]byte, 5)},
			want: []byte{
				0x00, 0x00, 0x00, 0x00, 0x00, 0x0B, 0x0B, 0x0B,
				0x0B, 0x0B, 0x0B, 0x0B, 0x0B, 0x0B, 0x0B, 0x0B,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pad(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pad() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_unpad(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "PKCS#7アンパディング",
			args: args{
				data: []byte{
					0x00, 0x00, 0x00, 0x00, 0x00, 0x0B, 0x0B, 0x0B,
					0x0B, 0x0B, 0x0B, 0x0B, 0x0B, 0x0B, 0x0B, 0x0B,
				},
			},
			want: []byte{
				0x00, 0x00, 0x00, 0x00, 0x00,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := unpad(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("unpad() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_encrypt(t *testing.T) {
	type args struct {
		data []byte
		key  []byte
		iv   []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "暗号化できるか",
			args: args{
				data: []byte("Hello World!"),
				key:  make([]byte, 32),
				iv:   make([]byte, 16),
			},
			want: []byte{
				0xc1, 0x0e, 0x6d, 0x56, 0x1c, 0x7a, 0xfa, 0xa5,
				0x7f, 0x3f, 0x44, 0xb2, 0xf7, 0x9d, 0xbf, 0x14,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := Encrypt(tt.args.data, tt.args.key, tt.args.iv)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Encrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_decrypt(t *testing.T) {
	type args struct {
		data []byte
		key  []byte
		iv   []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "復号化できるか",
			args: args{
				data: []byte{
					0xc1, 0x0e, 0x6d, 0x56, 0x1c, 0x7a, 0xfa, 0xa5,
					0x7f, 0x3f, 0x44, 0xb2, 0xf7, 0x9d, 0xbf, 0x14,
				},
				key: make([]byte, 32),
				iv:  make([]byte, 16),
			},
			want:    []byte("Hello World!"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := Decrypt(tt.args.data, tt.args.key, tt.args.iv)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getKeyAndIV(t *testing.T) {
	tests := []struct {
		name    string
		prepare func(ctx *gin.Context)
		want    []byte
		want1   []byte
	}{
		{
			name: "keyとivとれるか",
			prepare: func(ctx *gin.Context) {
				ctx.Set("cryptoKey", []byte("key"))
				ctx.Set("cryptoIV", []byte("iv"))
			},
			want:  []byte("key"),
			want1: []byte("iv"),
		},
		{
			name: "設定されてない時はデフォルト",
			prepare: func(ctx *gin.Context) {
			},
			want: []byte{
				0xC3, 0xFE, 0x43, 0x85, 0xD3, 0x21, 0x19, 0x10, 0x1A, 0x38, 0xE1, 0x38, 0xE0, 0x09, 0x03, 0x9D,
				0x8E, 0x09, 0x6B, 0xD0, 0x40, 0x43, 0xDF, 0xAB, 0x31, 0xE6, 0x97, 0x40, 0x5E, 0x4B, 0x86, 0xA8,
			},
			want1: []byte{
				0xD9, 0x71, 0x7B, 0xB4, 0x3D, 0x02, 0x51, 0xFA, 0xCE, 0x3C, 0x7D, 0xB7, 0xDC, 0xDF, 0x33, 0x00,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
			tt.prepare(ctx)

			got, got1 := getKeyAndIV(ctx)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getKeyAndIV() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("getKeyAndIV() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

type (
	child struct {
		V1 string `json:"v1"`
	}

	parent struct {
		V1 string `json:"v1"`
		V2 uint   `json:"v2"`
		V3 child  `json:"v3"`
	}
)

func TestRespond(t *testing.T) {
	type args struct {
		status int
		v      parent
	}
	tests := []struct {
		name    string
		prepare func(ctx *gin.Context)
		args    args
	}{
		{
			name: "暗号化してレスポンスできるか",
			prepare: func(ctx *gin.Context) {
				ctx.Request, _ = http.NewRequest(http.MethodGet, "/", nil)
				ctx.Request.Header.Set("Accept", "application/octet-stream")
			},
			args: args{
				status: 200,
				v:      parent{V1: "hello", V2: 1000, V3: child{V1: "child"}},
			},
		},
		{
			name: "SwaggerではJsonでレスポンスできるか",
			prepare: func(ctx *gin.Context) {
				ctx.Request, _ = http.NewRequest(http.MethodGet, "/", nil)
				ctx.Request.Header.Set("Accept", gin.MIMEJSON)
			},
			args: args{
				status: 200,
				v:      parent{V1: "hello", V2: 1000, V3: child{V1: "child"}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
			tt.prepare(ctx)
			Respond(ctx, tt.args.status, tt.args.v)
		})
	}
}

func TestShouldBind(t *testing.T) {
	type args struct {
		v parent
	}
	tests := []struct {
		name    string
		args    args
		prepare func(ctx *gin.Context)
		want    parent
		wantErr bool
	}{
		{
			name: "暗号化Jsonバインドできるか",
			prepare: func(ctx *gin.Context) {
				data := []byte{
					0x90, 0x19, 0x72, 0x3b, 0x54, 0x20, 0x74, 0x23, 0x7e, 0x06, 0xf2, 0xfc, 0x4c, 0x44, 0xae, 0x2e,
					0xf1, 0x49, 0x6d, 0xfe, 0xe7, 0x7a, 0x0c, 0x0f, 0xf0, 0x21, 0xc4, 0x4e, 0xfc, 0xb3, 0x6a, 0x30,
					0x93, 0xa9, 0xc2, 0x59, 0xeb, 0xae, 0x62, 0x24, 0x9e, 0x42, 0x67, 0xd6, 0x73, 0x4e, 0x12, 0x24,
				}
				ctx.Request, _ = http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(data))
				ctx.Request.Header.Add("Content-Type", "application/octet-stream")
			},
			want:    parent{V1: "hello", V2: 1000, V3: child{V1: "child"}},
			wantErr: false,
		},
		{
			name: "SwaggerではJsonでバインドできるか",
			prepare: func(ctx *gin.Context) {
				data := []byte("{\"v1\":\"hello\",\"v2\":1000,\"v3\":{\"v1\":\"child\"}}")
				ctx.Request, _ = http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(data))
				ctx.Request.Header.Add("Content-Type", gin.MIMEJSON)
			},
			want:    parent{V1: "hello", V2: 1000, V3: child{V1: "child"}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
			tt.prepare(ctx)

			if err := ShouldBind(ctx, &tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("ShouldBind() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.args.v, tt.want) {
				t.Errorf("ShouldBind() ttargs.v = %v, want %v", tt.args.v, tt.want)
			}
		})
	}
}
