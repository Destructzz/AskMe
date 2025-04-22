package cookie;
import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"time"
	"errors"
)
const (
	// TokenLength defines the length of the generated token in bytes
	TokenLength = 45
	// CookieName is the name of the cookie where the token will be stored
	CookieName = "token"
	// CookieExpiration defines how long the cookie should last (2 years)
	CookieExpiration = 730 * 24 * time.Hour
	
)

func GenerateToken() (string, error){// 45< symbols
	b := make([]byte, TokenLength)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func SetCookie(w http.ResponseWriter, token string) { // by http.SetCookie - кладет в request токен в куки
	http.SetCookie(w, &http.Cookie{
		Name:     CookieName,
		Value:    token,
		Path:     "/",
		Expires:  time.Now().Add(CookieExpiration),
		HttpOnly: true,
		Secure:   false, // Set to false if not using HTTPS
		SameSite: http.SameSiteLaxMode,
	})
}

func GetCookie(r *http.Request) (string, error) { // берет токен из cookie либо ошибку
	cookie, err := r.Cookie(CookieName)
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			return "", errors.New("no cookie found")
		}
		return "", err
	}
	return cookie.Value, nil
}