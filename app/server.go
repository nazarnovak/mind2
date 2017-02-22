package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (a *App) Server() *mux.Router {
	http.Handle("/css/",http.StripPrefix("/css", http.FileServer(http.Dir("public/css"))))
	http.Handle("/js/", http.StripPrefix("/js", http.FileServer(http.Dir("public/js"))))
	http.Handle("/glue/", a.Glue)

	r := mux.NewRouter()

	r.HandleFunc("/", a.IndexHandler).Methods("GET")
	//r.NewRoute().Methods("POST").Path("/login").HandlerFunc(controllers.Login)
	//r.NewRoute().Methods("GET").Path("/logout").HandlerFunc(controllers.Logout)
	//r.NewRoute().Methods("GET").Path("/cases/{caseId:[0-9]+}").HandlerFunc(controllers.GetCase)
	//r.NewRoute().Methods("POST").Path("/cases/new").HandlerFunc(controllers.NewCase)
	//r.NewRoute().Methods("GET").Path("/api/cases/{caseId:[0-9]+}/events").HandlerFunc(controllers.GetCaseMessages)
	//r.NewRoute().Methods("POST").Path("/api/cases/{caseId:[0-9]+}/events").HandlerFunc(controllers.CreateCaseMessage)
	return r
}

func (a *App) IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Howdy from Index"))
	//var userName string
	//var userRole int
	//var casesIds []int
	//var otherCasesIds []int
	//
	//user, err := getSessionUser(r)
	//if err != nil {
	//	log.Println("Error while getting user from session: " + err.Error())
	//	serveInternalServerError(w, r)
	//	return
	//}
	//
	//if user != nil {
	//	userName = user.Name
	//	userRole = user.Role
	//
	//	switch user.Role {
	//	case data.ROLEPATIENT:
	//		casesIds, err = data.GetCasesByCreatorId(user.ID)
	//		if err != nil {
	//			log.Println("Error getting patient cases:" + err.Error())
	//			serveInternalServerError(w, r)
	//			return
	//		}
	//		otherCasesIds, err = data.GetPatientAlienCasesId(user.ID)
	//	case data.ROLEDOCTOR:
	//		casesIds, err = data.GetCasesByDoctorId(user.ID)
	//		if err != nil {
	//			log.Println("Error getting patient cases:" + err.Error())
	//			serveInternalServerError(w, r)
	//			return
	//		}
	//		otherCasesIds, err = data.GetDoctorAlienCasesId(user.ID)
	//	default:
	//		err = errors.New("Unknown role")
	//	}
	//
	//	if err != nil {
	//		log.Println(err)
	//		serveInternalServerError(w, r)
	//		return
	//	}
	//
	//}
	//
	//data := struct{
	//	UserName string
	//	UserRole int
	//	Cases []int
	//	OtherCases []int
	//}{
	//	userName,
	//	userRole,
	//	casesIds,
	//	otherCasesIds,
	//}
	//
	//tpl := template.Must(template.ParseFiles(
	//"public/templates/layout.gohtml", "public/templates/index.gohtml"))
	//tpl.Execute(w, data)
}