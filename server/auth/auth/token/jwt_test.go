package token

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
)

const privateKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEApCuiVk+eC1u6xV33n8bKfkH8lyyIv4uDBRBuBuZVxIjt0n3c
YoLnSVPqIUnEXHwNRW5m9+XDv6xvmIdss4vGKrq2aQDJmkVXDeW7rd60fSEBXA6e
fQixSGN3V6OJg35r0IBdi+oEhOZ+tZjMv5Ppzr+pbBN9GMMtgHeBFKqH9vXSODcY
jb+xr3ESCpgcp7VNmOZBGazAOibM8T1R3To+KM7njigoKrBHh+3SGU2z0X2o/mZg
LKHBtY0w5NmemQxFeotQVLy0ULK2HWaT5sxSAFkZ8V60UDvC2K/ViqQGy4n4cNFT
ppqZvGCt0SrC2gasNqnv0DfydiChUGAVZJeifQIDAQABAoIBAGw3VwsN8GAMV1FC
X7gykcgoNxhvgrTAgd0Ut3JU4rRqjlnyHRPStiLf7bWDqqxMGXNrEvTVv/LY97bT
jhPMTxRbf2I34qyOiJPgm5ZA9ziESSElgdNpp3LQHhrm8f57lxYZh5By8geo6F1M
G6stFTHzpPpY3l00SwojvDKjNnzS/gjtCpEqACcrSsTBYaAEM+6iO+QPtXeXEXpF
JIf7+PVNPB0Ve7J5AdORScKvliOKSW2rjkYZU7pb/6wIN85bvYa9tdzTKVr0QANv
f+t08kBdYXwJCeLhwH6PalvmI/N3+uRLv3YB/QHEqcDDOG/37csTQBivbxqFxNTc
wOjTfQECgYEA4S3vBErUPQv4ypBa5ZsmAAAiyqvgdhaWCPX4VkxUlI/c5ErZMT+t
hsV3XHGjBO5j9K7Axtu0pcr4MpND/01RkNYgiICdrcRN0ZItzZfFguBB4VpOitBa
YuBRkNC79Du9vMpq/KPVe3ogIahZt4YDAtqZTI8JIdkdZlY6lVcq+98CgYEAuqQB
a+GEkuCnAu0VKlXpcteW1lgcipHbC/AjC8vec3Ncq8hS/VF6SfoZvxzls1MrQ1br
BwHOPGp9Iq4TtAVzVZ+D1PIHeZHoVf3bE+Wok6UeOfTeEJZgmKrnsnRngSdWAjbb
ssf56YT961JEwjUB99k24G8rvBMHN2fuwgKvLSMCgYEA1F0MkKC9nsEepnMTtnzm
jpOSPE9yhgvw/Yxz43djQSSYLxlD0kV6sqKiWiW1l5332LfCiQiFSuKrxo7KTwkt
Rooa7oFVs55xwZa56a4cELzMhOo1Locm5x8k7Y1d5Q0+eMDxu0gLusLHcNajnggd
6OAeOWygMTLgnjXuVA98M0MCgYAYjSx2sYLMA2tU0jUfjY+ZAM0hwxN2Evu6lUs2
6QmJLFq7ai2ojwhEJPcwEbARp9YgFJX0JQOpJOTzI+0JmFH0qHgiEfhzwILMlQk1
k9daHLVeMFzp1647RvgVnIRlE/SXc/nwiafRVSJFy1uqJGju3o7+dQFOyz3+xtg8
gzfh+wKBgQCMT9U4w28zJhuJh+w6RhKiaZEtE/xcSfOK3GlMZaK3asqq5/QPTD0F
iVM1eFue1vPz2AaRcGtMlGEGt6uXc7GpMlMRHsRo6EZunSuPTPFc4CfLtM7BZGe+
aKyDAfrTnW9sWOFWsp1xDc1kRuhZtJfc67KCRMcb+fOH/TDudGFhBg==
-----END RSA PRIVATE KEY-----`

func TestGenAccessToken(t *testing.T) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKey))
	if err != nil {
		t.Fatalf("cannot parse private key: %v", err)
	}
	g := NewJWTTokenGen("coolcar/auth", key, "")
	g.nowFunc = func() time.Time {
		return time.Unix(1516239022, 0)
	}
	tkn, err := g.GenAccessToken("5fe4a143289f4848ec97bdc9", 2*time.Hour)
	if err != nil {
		t.Errorf("cannot generate token: %v", err)
	}
	want := "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTYyNDYyMjIsImlhdCI6MTUxNjIzOTAyMiwiaXNzIjoiY29vbGNhci9hdXRoIiwic3ViIjoiNWZlNGExNDMyODlmNDg0OGVjOTdiZGM5In0.JygP4Df1_Rn4ojr8aVdHc7NkGcpx1ZPetkLnBfs6J3HvnHLuLcawgZ4qAiz2nFpnMSfnDnwY8IQ6bd8DeATG3oZM3EaX5YwABvt2kdPs02oRvGe4Q5tIiUkqyxrduI7pF2FNf4saLCvvElFwpJL8PnV1BOXYswfzYKSOGlEYkNQs0ua2qm2fW-28WGo_vlR6CjPjUiBqa0QI3-S3hiX8A8C9AhmY72ZykshjTierb9_A9BInIkhl-Sn6FqxzRiiUui3Y6FTN0l9FiIJ4rXjWZa2Qc0a2JEngBeTHqsfr9eR8bja8a-VKhW-qVvbI06weaNyvAkl3HUl3vVcXDerFdw"
	if tkn != want {
		t.Errorf("wrong token generated. want: %q; got: %q", want, tkn)
	}
}
