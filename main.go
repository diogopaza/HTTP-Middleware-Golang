package main

import(

	"fmt"
	"net/http"
	"log"
)


func handleMain( w http.ResponseWriter, r *http.Request){

	

	_, err := w.Write([]byte("handleMain"))
	if err != nil{
		fmt.Println("error handlemain", err)
		
	}
}

func handleUser( w http.ResponseWriter, r *http.Request){
	_, err := w.Write([]byte("handleUser"))
	if err != nil{
		fmt.Println("error handleUser", err)
		
	}


}

func middlewareJson( h http.HandlerFunc ) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		
		w.Header().Set("Content-Type","application/json; charset=UTF-8")
		w.WriteHeader(http.StatusAccepted)
		h.ServeHTTP(w,r)

	}
}

func middlewareAuth( h http.HandlerFunc ) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){

	user, pass, ok:= r.BasicAuth()
	_ = pass
	fmt.Println("user", user)
	if !ok || user != "admin"{
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Println("Usuário inválido")
	}

	w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
	fmt.Println("Usuário autenticado")
	h.ServeHTTP(w,r)

	}
}




func main(){

	http.HandleFunc("/", middlewareAuth(handleMain))
	http.HandleFunc( "/users", handleUser )

	log.Fatal( http.ListenAndServe(":8000", nil) )

}