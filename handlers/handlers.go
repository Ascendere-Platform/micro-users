package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/ascendere/micro-users/middlew"
	asignaturarouters "github.com/ascendere/micro-users/routers/asignatura_routers"
	postrouters "github.com/ascendere/micro-users/routers/post_routers"
	relacionrouters "github.com/ascendere/micro-users/routers/relacion_routers"
	rolrouters "github.com/ascendere/micro-users/routers/rol_routers"
	usuariosAsignaturasrouters "github.com/ascendere/micro-users/routers/usuariosAsignaturas_routers"
	usuariosrouters "github.com/ascendere/micro-users/routers/usuarios_routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Manejadores() {
	router := mux.NewRouter()

	//Llamadas al crud del Usuario
	router.HandleFunc("/registro", middlew.ChequeoBD(usuariosrouters.Registro)).Methods("POST")
	router.HandleFunc("/login", middlew.ChequeoBD(usuariosrouters.Login)).Methods("POST")
	router.HandleFunc("/verperfil", middlew.ChequeoBD(middlew.ValidoJWT(usuariosrouters.VerPerfil))).Methods("GET")
	router.HandleFunc("/modificarPerfil", middlew.ChequeoBD(middlew.ValidoJWT(usuariosrouters.ModificarPerfil))).Methods("PUT")
	router.HandleFunc("/listaUsuarios", middlew.ChequeoBD(middlew.ValidoJWT(usuariosrouters.ListaUsuarios))).Methods("GET")
	router.HandleFunc("/modificarRol", middlew.ChequeoBD(middlew.ValidoJWT(middlew.ValidoAdmin(usuariosrouters.ModificarRolUsuario)))).Methods("PUT")

	//Llamadas el crud de Post
	router.HandleFunc("/post", middlew.ChequeoBD(middlew.ValidoJWT(postrouters.GraboPost))).Methods("POST")
	router.HandleFunc("/leoPost", middlew.ChequeoBD(middlew.ValidoJWT(postrouters.LeoPost))).Methods("GET")
	router.HandleFunc("/eliminarPost", middlew.ChequeoBD(middlew.ValidoJWT(postrouters.EliminarPost))).Methods("DELETE")
	router.HandleFunc("/leoPostSeguidores", middlew.ChequeoBD(middlew.ValidoJWT(postrouters.LeoPostSeguidores))).Methods("GET")

	//Llamadas al crud de Relaciones
	router.HandleFunc("/seguirUsuario", middlew.ChequeoBD(middlew.ValidoJWT(relacionrouters.SeguirUsuario))).Methods("POST")
	router.HandleFunc("/eliminarRelacion", middlew.ChequeoBD(middlew.ValidoJWT(relacionrouters.EliminarRelacion))).Methods("DELETE")
	router.HandleFunc("/consultoRelacion", middlew.ChequeoBD(middlew.ValidoJWT(relacionrouters.ConsultoRelacion))).Methods("GET")
	
	
	//Llamadas al crud de Roles
	router.HandleFunc("/registroRol", middlew.ChequeoBD(middlew.ValidoJWT(middlew.ValidoAdmin(rolrouters.RegistroRol)))).Methods("POST")
	router.HandleFunc("/eliminarRol", middlew.ChequeoBD(middlew.ValidoJWT(middlew.ValidoAdmin(rolrouters.EliminarRol)))).Methods("DELETE")
	router.HandleFunc("/verRol", middlew.ChequeoBD(middlew.ValidoJWT(rolrouters.VerRol))).Methods("GET")
	router.HandleFunc("/listaRoles", middlew.ChequeoBD(middlew.ValidoJWT(rolrouters.ListaRoles))).Methods("GET")

	//Llamadas al crud de Facultades
	router.HandleFunc("/registroFacultad", middlew.ChequeoBD(middlew.ValidoJWT(middlew.ValidoAdmin(asignaturarouters.IngresarFacultad)))).Methods("POST")
	router.HandleFunc("/eliminarFacultad", middlew.ChequeoBD(middlew.ValidoJWT(middlew.ValidoAdmin(asignaturarouters.EliminarFacultad)))).Methods("DELETE")
	router.HandleFunc("/listarFacultades", middlew.ChequeoBD(middlew.ValidoJWT(asignaturarouters.ListarFacutlades))).Methods("GET")
	router.HandleFunc("/verFacultad", middlew.ChequeoBD(middlew.ValidoJWT(asignaturarouters.BuscarFacultad))).Methods("GET")

	//Llamadas al crud de Asignaturas
	router.HandleFunc("/registroAsignatura", middlew.ChequeoBD(middlew.ValidoJWT(asignaturarouters.IngresarAsignatura))).Methods("POST")
	router.HandleFunc("/buscarAsignatura", middlew.ChequeoBD(middlew.ValidoJWT(asignaturarouters.BuscarAsignatura))).Methods("GET")
	router.HandleFunc("/eliminarAsignatura", middlew.ChequeoBD(middlew.ValidoJWT(asignaturarouters.EliminarAsignatura))).Methods("DELETE")
	router.HandleFunc("/listarAsignaturas", middlew.ChequeoBD(middlew.ValidoJWT(asignaturarouters.ListarAsignaturas))).Methods("GET")
	router.HandleFunc("/modificarAsignatura", middlew.ChequeoBD(middlew.ValidoJWT(asignaturarouters.ModificarAsignatura))).Methods("PUT")


	//Llamadas al crud de Asignar Asignaturas a Usuarios
	router.HandleFunc("/asignarAsignatura", middlew.ChequeoBD(middlew.ValidoJWT(usuariosAsignaturasrouters.AsignarAsignaturaUsuario))).Methods("POST")
	router.HandleFunc("/eliminarAsginaturaUsuario", middlew.ChequeoBD(middlew.ValidoJWT(usuariosAsignaturasrouters.EliminarAsignaturaUsuario))).Methods("DELETE")
	router.HandleFunc("/consultoAsignaturaUsuario", middlew.ChequeoBD(middlew.ValidoJWT(usuariosAsignaturasrouters.ConsultarAsignaturaUsuario))).Methods("GET")
	router.HandleFunc("/leoAsignaturasUsuario", middlew.ChequeoBD(middlew.ValidoJWT(usuariosAsignaturasrouters.ListarAsignaturasUsuario))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
