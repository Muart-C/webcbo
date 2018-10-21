package middleware

import (
	"github.com/Muart-C/webcbo/controllers"
	"log"
	"net/http"
	"runtime/debug"
	"strings"
)
//
//func RequireAdminAuth(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		session := config.Session(r)
//		userID, ok := session.Values["admin_user_id"]
//
//		if !ok {
//			views.ErrorFlash(w, r, "Please login to view that resource.")
//			http.Redirect(w, r, "/admin/sessions/new", http.StatusSeeOther)
//			return
//		}
//
//		// find user
//		u, err := models.GetAdminUser(userID.(int64))
//		if err != nil {
//			views.ErrorFlash(w, r, "Please login to view that resource.")
//			http.Redirect(w, r, "/admin/sessions/new", http.StatusSeeOther)
//			return
//		}
//
//		ctx := context.WithValue(r.Context(), services.SessKey, u)
//		next.ServeHTTP(w, r.WithContext(ctx))
//	})
//}


func Recoverer(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rvr := recover(); rvr != nil {
				log.Println(rvr)
				log.Println(string(debug.Stack()))
				controllers.RespondWithError(w, http.StatusInternalServerError,"an error occurred")
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func MethodOverride(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			override := r.FormValue("_method")
			if override != "" {
				r.Method = strings.ToUpper(override)
				r.Form.Del("_method")
				r.PostForm.Del("_method")
			}
		}
		next.ServeHTTP(w, r)
	})
}