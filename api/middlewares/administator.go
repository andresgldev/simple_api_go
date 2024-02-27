package middlewares

import "net/http"

func OnlyAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		/*
			ctx := r.Context()
			perm, ok := ctx.Value("acl.permission").(YourPermissionType)
			if !ok || !perm.IsAdmin() {
				http.Error(w, http.StatusText(403), 403)
				return
			}
		*/
		next.ServeHTTP(w, r)
	})
}
