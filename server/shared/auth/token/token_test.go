package token

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
)

const publicKey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEApCuiVk+eC1u6xV33n8bK
fkH8lyyIv4uDBRBuBuZVxIjt0n3cYoLnSVPqIUnEXHwNRW5m9+XDv6xvmIdss4vG
Krq2aQDJmkVXDeW7rd60fSEBXA6efQixSGN3V6OJg35r0IBdi+oEhOZ+tZjMv5Pp
zr+pbBN9GMMtgHeBFKqH9vXSODcYjb+xr3ESCpgcp7VNmOZBGazAOibM8T1R3To+
KM7njigoKrBHh+3SGU2z0X2o/mZgLKHBtY0w5NmemQxFeotQVLy0ULK2HWaT5sxS
AFkZ8V60UDvC2K/ViqQGy4n4cNFTppqZvGCt0SrC2gasNqnv0DfydiChUGAVZJei
fQIDAQAB
-----END PUBLIC KEY-----`

func TestVerify(t *testing.T) {
	pubKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKey))
	if err != nil {
		t.Fatalf("cannot parse public key: %v", err)
	}

	v := &JWTTokenVerifier{
		PublicKey: pubKey,
	}

	cases := []struct {
		name    string
		tkn     string
		now     time.Time
		want    string
		wantErr bool
	}{
		{
			name:    "valid_token",
			tkn:     "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTYyNDYyMjIsImlhdCI6MTUxNjIzOTAyMiwiaXNzIjoiY29vbGNhci9hdXRoIiwic3ViIjoiNWZlNGExNDMyODlmNDg0OGVjOTdiZGM5In0.JygP4Df1_Rn4ojr8aVdHc7NkGcpx1ZPetkLnBfs6J3HvnHLuLcawgZ4qAiz2nFpnMSfnDnwY8IQ6bd8DeATG3oZM3EaX5YwABvt2kdPs02oRvGe4Q5tIiUkqyxrduI7pF2FNf4saLCvvElFwpJL8PnV1BOXYswfzYKSOGlEYkNQs0ua2qm2fW-28WGo_vlR6CjPjUiBqa0QI3-S3hiX8A8C9AhmY72ZykshjTierb9_A9BInIkhl-Sn6FqxzRiiUui3Y6FTN0l9FiIJ4rXjWZa2Qc0a2JEngBeTHqsfr9eR8bja8a-VKhW-qVvbI06weaNyvAkl3HUl3vVcXDerFdw",
			now:     time.Unix(1516239122, 0),
			want:    "5fe4a143289f4848ec97bdc9",
			wantErr: false,
		},
		{
			name:    "token_expried",
			tkn:     "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTYyNDYyMjIsImlhdCI6MTUxNjIzOTAyMiwiaXNzIjoiY29vbGNhci9hdXRoIiwic3ViIjoiNWZlNGExNDMyODlmNDg0OGVjOTdiZGM5In0.JygP4Df1_Rn4ojr8aVdHc7NkGcpx1ZPetkLnBfs6J3HvnHLuLcawgZ4qAiz2nFpnMSfnDnwY8IQ6bd8DeATG3oZM3EaX5YwABvt2kdPs02oRvGe4Q5tIiUkqyxrduI7pF2FNf4saLCvvElFwpJL8PnV1BOXYswfzYKSOGlEYkNQs0ua2qm2fW-28WGo_vlR6CjPjUiBqa0QI3-S3hiX8A8C9AhmY72ZykshjTierb9_A9BInIkhl-Sn6FqxzRiiUui3Y6FTN0l9FiIJ4rXjWZa2Qc0a2JEngBeTHqsfr9eR8bja8a-VKhW-qVvbI06weaNyvAkl3HUl3vVcXDerFdw",
			now:     time.Unix(1517239122, 0),
			wantErr: true,
		},
		{
			name:    "bad_token",
			tkn:     "bad_token",
			now:     time.Unix(1516239122, 0),
			wantErr: true,
		},
		{
			name:    "wrong_signature",
			tkn:     "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTYyNDYyMjIsImlhdCI6MTUxNjIzOTAyMiwiaXNzIjoiY29vbGNhci9hdXRoIiwic3ViIjoiNWZlNGExNDMyODlmNDg0OGVjOTdiZGM5In0.a3gB7f-a85dTGqvCN3vx8wJ3aBxoGTHQlgKz79LkIw-gjhSbGsw16URCWNn943l4hiiCvsAaDcC38Y1HjrDrz8GqZHLaBxb_Jf0XQaI0kmOFszwvUjEaC4wnf6XWnretp1mdfKTQTmMoBeOv87BR40gISJEIjxkBdUwfd6JEHaFDpcqCTHzNG7M8RsYjofXrxbj8r3KzfTZe1Q5oLJoaKAOkEDKMkGGVrCB76A2YdA-kCeqNZfk5jBmPq7lLxwuDP6g0AITDO6fYQ51sw8riSHYXy67-gkz4DmK2nUvck2aH-FGb4Phom0hhywMhcGKRRD-WClBhfmAJHdcD6VehRg",
			now:     time.Unix(1516239122, 0),
			wantErr: true,
		},
	}

	for _, cc := range cases {
		t.Run(cc.name, func(t *testing.T) {
			jwt.TimeFunc = func() time.Time {
				return cc.now
			}

			accountID, err := v.VerifyAccessToken(cc.tkn)
			if !cc.wantErr && err != nil {
				t.Errorf("verification failed: %v", err)
			}

			if cc.wantErr && err == nil {
				t.Errorf("verification failed: %v", err)
			}

			if accountID != cc.want {
				t.Errorf("wrong account id. want: %q; got %q", cc.want, accountID)
			}
		})
	}
}
