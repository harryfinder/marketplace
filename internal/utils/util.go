package utils

// Response ...
type Response struct {
	Code    int         `json:"-"`
	Message string      `json:"-"`
	Payload interface{} `json:"payload"`
}

//func IsAdmin(r *http.Request) (bool, error) {
//	const authorizationHeader = "Authorization"
//	var err error
//	auth := r.Header.Get(authorizationHeader)
//	if auth == "" {
//		return false, err
//	}
//	bearerToken := strings.Split(auth, " ")
//	if len(bearerToken) != 2 {
//		return false, err
//	}
//	_, token, err := ParseToken(bearerToken[1])
//	if err != nil {
//		return false, err
//	}
//	claims, ok := token.Claims.(*tokenClaims)
//	if !ok {
//		return false, err
//	}
//	if claims.Role == "admin" {
//		return true, nil
//	}
//	return false, err
//}
