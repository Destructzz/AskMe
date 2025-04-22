package register

import (
	"askme/cookie"
	"net/http"
	"context"
)

//сделай middleware для регистрации челоека по cookie.

//запрос -> смотришь есть ли cookie с токеном -> если есть отдаешь юзера в middleware на использование если нет то делаешь его и отдаешь. Делать токены 

//3 аргумент в функции

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		token, err := cookie.GetCookie(r)
		
		if err != nil {
			token, err := cookie.GenerateToken()
			if err != nil {
				http.Error(w, "Failed to generate token", http.StatusInternalServerError)
				return
			}
			
			cookie.SetCookie(w, token)
		}

		// 4. Добавляем токен в контекст запроса
		ctx := context.WithValue(r.Context(), "auth_token", token)
		r = r.WithContext(ctx)

		// 5. Передаем управление следующему обработчику
		next.ServeHTTP(w, r)
	}
}

