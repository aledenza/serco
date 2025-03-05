package middlewares

import (
	"crypto/x509"
	"encoding/pem"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/creasty/defaults"
	"github.com/danielgtaylor/huma/v2"
	"github.com/golang-jwt/jwt/v5"
)

type JWTAuthConfig struct {
	PublicKey       *string
	Issuer          string
	ServiceName     string
	AuthScheme      string `default:"Bearer"`
	TokenHeader     string `default:"Authorization"`
	Whitelist       []string
	AllowEmptyToken bool
}

func JWTAuth(config JWTAuthConfig) func(huma.Context, func(huma.Context)) {
	defaults.MustSet(&config)

	if (config.PublicKey == nil || *config.PublicKey == "") && config.AllowEmptyToken {
		return func(ctx huma.Context, next func(huma.Context)) { next(ctx) }
	}

	if config.PublicKey == nil || *config.PublicKey == "" || config.Issuer == "" {
		panic("public key and issuer required")
	}

	block, _ := pem.Decode([]byte(*config.PublicKey))
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic("Invalid key")
	}

	whiteList := make([]*regexp.Regexp, 0, len(config.Whitelist))
	for _, reg := range config.Whitelist {
		whiteList = append(whiteList, regexp.MustCompile(reg))
	}

	return func(ctx huma.Context, next func(huma.Context)) {
		forbidden := func() {
			ctx.SetStatus(http.StatusForbidden)
			ctx.BodyWriter().Write([]byte("forbidden"))
		}

		for _, route := range whiteList {
			if route.Match([]byte(ctx.URL().Path)) {
				next(ctx)
				return
			}
		}
		headerToken := ctx.Header(config.TokenHeader)
		if headerToken == "" {
			forbidden()
			return
		}
		headerToken = strings.TrimSpace(strings.Replace(headerToken, config.AuthScheme, "", -1))
		token, err := jwt.Parse(headerToken, func(token *jwt.Token) (any, error) { return pub, nil })
		if err != nil || !token.Valid {
			forbidden()
			return
		}
		iss, issErr := token.Claims.GetIssuer()
		sub, subErr := token.Claims.GetSubject()
		exp, expErr := token.Claims.GetExpirationTime()
		aud, audErr := token.Claims.GetAudience()

		if issErr != nil || subErr != nil || expErr != nil || audErr != nil ||
			iss != config.Issuer || sub == "" || exp.Before(time.Now()) || len(aud) == 0 || aud[0] != config.ServiceName {
			forbidden()
			return
		}
		externalClientMetric(sub)
		next(ctx)
	}
}
