package main

import(

	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/negroni"
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
		h.ServeHTTP(w,r)

	}
}

func middlewareAuth() negroni.Handler{
	return func(w http.ResponseWriter, r *http.Request){

	if r.URL.Path == "/general"{
		fmt.Println("General")
		h.ServeHTTP(w,r)
		return
	}

	user, pass, ok:= r.BasicAuth()
	fmt.Println("user", user)
	if !ok || user != "admin" || pass!= "admin"{
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintln( w, "Usuário ou senha inválidos" )
		return
	}

	w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
	fmt.Println("Usuário autenticado")
	h.ServeHTTP(w,r)

	}
}




func main(){

	http.HandleFunc("/", middlewareJson( middlewareAuth( handleMain )))
	http.HandleFunc( "/users", handleUser )

	log.Fatal( http.ListenAndServe(":8000", nil) )

}