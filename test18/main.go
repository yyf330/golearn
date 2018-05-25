package main

import (
	"log"
	"net/http"
	"github.com/emicklei/go-restful"
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/go-openapi/spec"
	//"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"test/test18/mdbsql"
	"test/test18/models"
	"fmt"

	"io/ioutil"
	//"bytes"
	"encoding/json"
	"io"
	//"database/sql"
)

// UserResource is the REST layer to the User domain
type UserResource struct {
	// normally one would use DAO (data access object)
	mdb mydbsql.DbBuilder
	//db *sql.DB
	users map[string]models.User
}

//type DBSouce struct{

//// User is just a sample type
//type User struct {
//	ID   string `json:"id" description:"identifier of the user"`
//	Name string `json:"name" description:"name of the user" default:"john"`
//	Age  int    `json:"age" description:"age of the user" default:"21"`
//}

// WebService creates a new service that can handle REST requests for User resources.
func (u UserResource) WebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.
		Path("/users").
		Consumes(restful.MIME_JSON, restful.MIME_JSON).
		//Consumes(restful.MIME_XML, restful.MIME_JSON).
		//Produces(restful.MIME_JSON, restful.MIME_XML)1
		Produces(restful.MIME_JSON, restful.MIME_JSON) // you can specify this per route as well

	tags := []string{"users"}

	ws.Route(ws.GET("/").To(u.findAllUsers).
	// docs
		Doc("get all users").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes([]models.User{}).
		Returns(200, "OK", []models.User{}))

	ws.Route(ws.GET("/{user-id}").To(u.findUser).
	// docs
		Doc("get a user").
		Param(ws.PathParameter("user-id", "identifier of the user").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(models.User{}). // on the response
		Returns(200, "OK", models.User{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.PUT("/{user-id}").To(u.updateUser).
	// docs
		Doc("update a user").
		Param(ws.PathParameter("user-id", "identifier of the user").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(models.User{})) // from the request

	ws.Route(ws.POST("/").To(u.createUser).
	// docs
		Doc("create a user").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(models.User{})) // from the request

	ws.Route(ws.DELETE("/").To(u.removeUser).
	// docs
		Doc("delete a user").
		Metadata(restfulspec.KeyOpenAPITags, tags))//.
		//Param(ws.PathParameter("user-id", "identifier of the user").DataType("string")))

	return ws
}

// GET http://localhost:8080/users
//
func (u UserResource) findAllUsers(request *restful.Request, response *restful.Response) {
	record := make(map[int]map[string]string)
	record = u.mdb.Query2()
	fmt.Println("findAllUsers")
	response.WriteEntity(record)

	//list := []User{}
	//for key, val := range record{
	//	fmt.Println(key)
	//	fmt.Println(val)
	//	list = append(list,val)
	//}
	//fmt.Println(list)
	//list := []User{}
	//for _, each := range u.users {
	//	list = append(list, each)
	//}
	//response.WriteEntity(list)
}

// GET http://localhost:8080/users/1
//
func (u UserResource) findUser(request *restful.Request, response *restful.Response) {
	//usr := new(models.User)
	id := request.PathParameter("user-id")
	fmt.Println("--------findUser-----------------")
	fmt.Println(id)

	ret := u.mdb.CheckUserInfo(id)
	if len(id) == 0 {
		response.WriteErrorString(http.StatusNotFound, "User could not be found.")
	} else {
		response.WriteEntity(ret)
	}
}

// PUT http://localhost:8080/users/1
// <User><Id>1</Id><Name>Melissa Raspberry</Name></User>
//
func (u *UserResource) updateUser(request *restful.Request, response *restful.Response) {
	fmt.Println("---------updateUser-------1---------")
	id := request.PathParameter("user-id")
	fmt.Println("---------updateUser----------------",id)
	ret := u.mdb.Update(id)
	if ret == 0 {
		response.AddHeader("Content-Type", "application/json")
		response.WriteHeaderAndEntity(http.StatusCreated, "bad id")
	}else{
		response.AddHeader("Content-Type", "application/json")
		response.WriteHeaderAndEntity(http.StatusCreated, "ok")
	}

}

// PUT http://localhost:8080/users/

func (u *UserResource) createUser(request *restful.Request, response *restful.Response) {
	//ioutil.ReadAll(request.Request.Body)
	result, _ := ioutil.ReadAll(io.LimitReader(request.Request.Body, 1048576))
	fmt.Println("---------createUser-------1---------")
	//fmt.Println(bytes.NewBuffer(result).String())
	var rep_string map[string]string

	fmt.Println("---------createUser-------2---------")
	json.Unmarshal(result,&rep_string)
	//if err != nil {
	fmt.Println(rep_string)
	fmt.Println("---------createUser-------3---------")

	usr := new(models.User)
	usr.UserName = rep_string["user-name"]
	usr.Password = rep_string["user-password"]
	usr.NickName = rep_string["user-nickName"]
	//if len(usr.UserName)==0
	fmt.Println("-------------------------",usr.UserName,usr.Password,usr.NickName)
	//fmt.Println("-------------------------",u.mdb)

	u.mdb.Insert(usr.UserName,usr.Password,usr.NickName)
	response.WriteHeaderAndEntity(http.StatusCreated, "ok")
	//} else {
	//	response.WriteError(http.StatusInternalServerError, err)
	//}

}

// DELETE http://localhost:8080/users/1
//
func (u *UserResource) removeUser(request *restful.Request, response *restful.Response) {
	result, _ := ioutil.ReadAll(io.LimitReader(request.Request.Body, 1048576))
	//fmt.Println("---------removeUser-------1---------")
	var rep_string map[string]string

	json.Unmarshal(result,&rep_string)
	fmt.Println(rep_string)
	//fmt.Println("---------removeUser-------3---------")

	//fmt.Println(rep_string)

	id := rep_string["user-id"]
	ret :=u.mdb.Remove(id)
	//fmt.Println("---------removeUser-------2---------")
	if ret == 0 {
		response.AddHeader("Content-Type", "application/json")
		response.WriteHeaderAndEntity(http.StatusCreated, "ok")
	} else {
		//response.AddHeader("Content-Type", "application/json")
		response.WriteErrorString(http.StatusInternalServerError, "remove failed!")
	}
	//delete(u.users, id)
}

func main() {
	//daolin add for mysql
	//var Mdb mydbsql.DbBuilder
	//Mdb := new(mydbsql.DbBuilder)
	u := new(UserResource)

	u.mdb = mydbsql.Mysqlinit()
	//Mdb = mydbsql.Mysqlinit()

	//fmt.Println(Mdb.DB)
	//daolin add for mysql

	restful.DefaultContainer.Add(u.WebService())

	config := restfulspec.Config{
		WebServices: restful.RegisteredWebServices(), // you control what services are visible
		APIPath:     "/apidocs.json",
		PostBuildSwaggerObjectHandler: enrichSwaggerObject}
	restful.DefaultContainer.Add(restfulspec.NewOpenAPIService(config))

	// Optionally, you can install the Swagger Service which provides a nice Web UI on your REST API
	// You need to download the Swagger HTML5 assets and change the FilePath location in the config below.
	// Open http://localhost:8080/apidocs/?url=http://localhost:8080/apidocs.json
	http.Handle("/d/", http.StripPrefix("/apidocs/", http.FileServer(http.Dir("/Users/emicklei/Projects/swagger-ui/dist"))))

	log.Printf("start listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
	u.mdb.Close()//daolin add for mysql
}

func enrichSwaggerObject(swo *spec.Swagger) {
	swo.Info = &spec.Info{
		InfoProps: spec.InfoProps{

			Title:       "UserService",
			Description: "Resource for managing Users",
			Contact: &spec.ContactInfo{
				Name:  "john",
				Email: "john@doe.rp",
				URL:   "http://johndoe.org",
			},
			License: &spec.License{
				Name: "MIT",
				URL:  "http://mit.org",
			},
			Version: "1.0.0",
		},
	}
	swo.Tags = []spec.Tag{spec.Tag{TagProps: spec.TagProps{
		Name:        "users",
		Description: "Managing users"}}}
}

