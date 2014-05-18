package admin

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"
	"html/template"
	"log"
	"net/http"
	"time"
)

type user struct {
	Uid string `json:"uid"bson:"uid"`
	Pwd string `json:"pwd"bson:"pwd"`

	Nick   string `json:"nick"bson:"nick"`
	Mail   string `json:"mail"bson:"mail"`
	School string `json:"school"bson:"school"`
	Motto  string `json:"motto"bson:"motto"`

	Privilege int `json:"privilege"bson:"privilege"`

	Solve  int `json:"solve"bson:"solve"`
	Submit int `json:"submit"bson:"submit"`

	Last   time.Time `json:"last"bson:"last"`
	Logout bool      `json:"logout"bson:"logout"`
	Status int       `json:"status"bson:"status"`
	Create string    `json:"create"bson:'create'`
}

type UserController struct {
	class.Controller
}

// URL /admin/user/list
// Show User List Page
func (this *UserController) List(w http.ResponseWriter, r *http.Request) {
	log.Print("Admin - User List")
	this.Init(w, r)

	response, err := http.Post(config.PostHost+"/user/list", "application/json", nil)
	defer response.Body.Close()

	if err != nil {
		http.Error(w, "post error", 500)
		return
	}

	one := make(map[string][]user)
	if response.StatusCode == 200 {
		err = this.LoadJson(response.Body, &one)
		if err != nil {
			http.Error(w, "load error", 400)
			return
		}
		this.Data["User"] = one["list"]
	}

	t := template.New("layout.tpl").Funcs(template.FuncMap{"ShowStatus": class.ShowStatus})
	t, err = t.ParseFiles("view/admin/layout.tpl", "view/admin/user_list.tpl")
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}

	this.Data["Title"] = "Admin - User List"
	this.Data["IsUser"] = true
	this.Data["IsList"] = true
	err = t.Execute(w, this.Data)
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}
}

func (this *UserController) Onlinelist(w http.ResponseWriter, r *http.Request) {
	log.Print("Admin - User Online List")
	this.Init(w, r)

	response, err := http.Post(config.PostHost+"/user/onlinelist", "application/json", nil)
	defer response.Body.Close()

	if err != nil {
		http.Error(w, "post error", 500)
		return
	}

	one := make(map[string][]user)
	if response.StatusCode == 200 {
		err = this.LoadJson(response.Body, &one)
		if err != nil {
			http.Error(w, "load error", 400)
			return
		}
		this.Data["User"] = one["list"]
	}

	t := template.New("layout.tpl").Funcs(template.FuncMap{"ShowStatus": class.ShowStatus})
	t, err = t.ParseFiles("view/admin/layout.tpl", "view/admin/user_onlinelist.tpl")
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}

	this.Data["Title"] = "Admin - User List"
	this.Data["IsUser"] = true
	this.Data["IsOnlineList"] = true
	err = t.Execute(w, this.Data)
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}
}
