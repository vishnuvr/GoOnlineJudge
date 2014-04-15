package controller

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"
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

	Status int    `json:"status"bson:"status"`
	Create string `json:"create"bson:'create'`
}

type UserController struct {
	class.Controller
}

func (this *UserController) Signin(w http.ResponseWriter, r *http.Request) {
	log.Println("User Login")
	this.Init(w, r)

	t := template.New("layout.tpl")
	t, err := t.ParseFiles("view/layout.tpl", "view/user_signin.tpl")
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}

	this.Data["Title"] = "User Sign In"
	this.Data["IsUserSignIn"] = true

	err = t.Execute(w, this.Data)
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}
}

func (this *UserController) Login(w http.ResponseWriter, r *http.Request) {
	log.Println("User Login")
	this.Init(w, r)

	one := make(map[string]string)
	one["uid"] = r.FormValue("user[handle]")
	one["pwd"] = r.FormValue("user[password]")

	reader, err := this.PostReader(&one)
	if err != nil {
		http.Error(w, "read error", 500)
		return
	}

	response, err := http.Post(config.PostHost+"/user/login", "application/json", reader)
	defer response.Body.Close()
	if err != nil {
		http.Error(w, "post error", 500)
		return
	}

	var ret user
	err = this.LoadJson(response.Body, &ret)
	if err != nil {
		http.Error(w, "load error", 400)
		return
	}

	if response.StatusCode == 200 {
		if ret.Uid == "" {
			w.WriteHeader(400)
		} else {
			this.SetSession(w, r, "CurrentUser", one["uid"])
			this.SetSession(w, r, "CurrentPrivilege", strconv.Itoa(ret.Privilege))
			w.WriteHeader(200)
		}
		return
	} else {
		w.WriteHeader(response.StatusCode)
		return
	}

}

func (this *UserController) Signup(w http.ResponseWriter, r *http.Request) {
	log.Println("User Sign Up")
	this.Init(w, r)

	t := template.New("layout.tpl")
	t, err := t.ParseFiles("view/layout.tpl", "view/user_signup.tpl")
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}

	this.Data["Title"] = "User Sign Up"
	this.Data["IsUserSignUp"] = true
	err = t.Execute(w, this.Data)
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}
}

func (this *UserController) Register(w http.ResponseWriter, r *http.Request) {
	log.Println("User Register")
	this.Init(w, r)

	one := make(map[string]string)
	one["uid"] = r.FormValue("user[handle]")
	one["nick"] = r.FormValue("user[nick]")
	one["pwd"] = r.FormValue("user[password]")
	one["pwdConfirm"] = r.FormValue("user[confirmPassword]")
	one["mail"] = r.FormValue("user[mail]")
	one["school"] = r.FormValue("user[school]")
	one["motto"] = r.FormValue("user[motto]")

	ok := 1
	warning := make(map[string]string)

	response, err := http.Post(config.PostHost+"/user/list/uid/"+one["uid"], "application/json", nil)
	defer response.Body.Close()
	if err != nil {
		http.Error(w, "post error", 500)
		return
	}

	if one["uid"] == "" {
		ok, warning["uid"] = 0, "Handle should not be empty."
	} else {
		ret := make(map[string][]*user)
		if response.StatusCode == 200 {
			err = this.LoadJson(response.Body, &ret)
			if err != nil {
				http.Error(w, "load error", 400)
				return
			}

			if len(ret["list"]) > 0 {
				ok, warning["uid"] = 0, "This handle is currently in use."
			}
		}
	}
	if one["nick"] == "" {
		ok, warning["nick"] = 0, "Nick should not be empty."
	}
	if len(one["pwd"]) < 6 {
		ok, warning["pwd"] = 0, "Password should contain at least six characters."
	}
	if one["pwd"] != one["pwdConfirm"] {
		ok, warning["pwdConfirm"] = 0, "Confirmation mismatched."
	}

	for k, v := range warning {
		log.Println(k + ":" + v)
	}

	if ok == 1 {
		reader, err := this.PostReader(&one)
		if err != nil {
			http.Error(w, "read error", 500)
			return
		}

		response, err = http.Post(config.PostHost+"/user/insert", "application/json", reader)
		defer response.Body.Close()
		if err != nil {
			http.Error(w, "post error", 400)
			return
		}

		this.SetSession(w, r, "CurrentUser", one["uid"])
		this.SetSession(w, r, "CurrentPrivilege", "1")
		w.WriteHeader(200)
	} else {
		b, err := json.Marshal(&warning)
		if err != nil {
			http.Error(w, "json error", 500)
			return
		}

		w.WriteHeader(400)
		w.Write(b)
	}
}

func (this *UserController) Logout(w http.ResponseWriter, r *http.Request) {
	log.Println("User Logout")
	this.Init(w, r)

	if this.GetSession(w, r, "CurrentUser") != "" {
		this.DeleteSession(w, r, "CurrentUser")
		this.DeleteSession(w, r, "CurrentPrivilege")
	}

	w.WriteHeader(200)
}

func (this *UserController) Detail(w http.ResponseWriter, r *http.Request) {
	log.Println("User Detail")
	this.Init(w, r)

	args := this.ParseURL(r.URL.Path)

	response, err := http.Post(config.PostHost+"/user/detail/uid/"+args["uid"], "application/json", nil)
	defer response.Body.Close()
	if err != nil {
		http.Error(w, "post error", 500)
		return
	}

	var one user
	if response.StatusCode == 200 {
		err = this.LoadJson(response.Body, &one)
		if err != nil {
			http.Error(w, "load error", 400)
			return
		}
		this.Data["Detail"] = one
	} else {
		http.Error(w, "resp error", response.StatusCode)
		return
	}

	t := template.New("layout.tpl")
	t, err = t.ParseFiles("view/layout.tpl", "view/user_detail.tpl")
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}

	this.Data["Title"] = "User Detail"

	err = t.Execute(w, this.Data)
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}
}
