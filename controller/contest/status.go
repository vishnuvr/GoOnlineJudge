package contest

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type solution struct {
	Sid int `json:"sid"bson:"sid"`

	Pid      int    `json:"pid"bson:"pid"`
	Uid      string `json:"uid"bson:"uid"`
	Judge    int    `json:"judge"bson:"judge"`
	Time     int    `json:"time"bson:"time"`
	Memory   int    `json:"memory"bson:"memory"`
	Length   int    `json:"length"bson:"length"`
	Language int    `json:"language"bson:"language"`

	Module int `json:"module"bson:"module"`
	Mid    int `json:"mid"bson:"mid"`

	Code string `json:"code"bson:"code"`

	Status int    `json:"status"bson:"status"`
	Create string `json:"create"bson:"create"`
}

type StatusController struct {
	Contest
}

func (this *StatusController) List(w http.ResponseWriter, r *http.Request) {
	log.Println("Contest Status List")
	this.InitContest(w, r)

	response, err := http.Post(config.PostHost+"/solution/list/module/"+strconv.Itoa(config.ModuleC)+"/mid/"+strconv.Itoa(this.Cid), "application/json", nil)
	defer response.Body.Close()
	if err != nil {
		http.Error(w, "post error", 500)
		return
	}

	one := make(map[string][]solution)
	if response.StatusCode == 200 {
		err = this.LoadJson(response.Body, &one)
		if err != nil {
			http.Error(w, "load error", 400)
			return
		}
		this.Data["Solution"] = one["list"]
	}

	for _, v := range one["list"] {
		v.Pid = this.Index[v.Pid]
	}

	t := template.New("layout.tpl").Funcs(template.FuncMap{"ShowStatus": class.ShowStatus, "ShowJudge": class.ShowJudge, "ShowLanguage": class.ShowLanguage})
	t, err = t.ParseFiles("view/layout.tpl", "view/contest/status_list.tpl")
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}

	this.Data["IsContestStatus"] = true
	err = t.Execute(w, this.Data)
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}
}

func (this *StatusController) Code(w http.ResponseWriter, r *http.Request) {
	log.Println("Status Code")
	this.Init(w, r)
	this.InitContest(w, r)

	args := this.ParseURL(r.URL.Path[8:])
	sid, err := strconv.Atoi(args["sid"])
	if err != nil {
		http.Error(w, "args error", 400)
		return
	}

	response, err := http.Post(config.PostHost+"/solution/detail/sid/"+strconv.Itoa(sid), "application/json", nil)
	defer response.Body.Close()
	if err != nil {
		http.Error(w, "post error", 500)
		return
	}

	var one solution
	if response.StatusCode == 200 {
		err = this.LoadJson(response.Body, &one)
		if err != nil {
			http.Error(w, "load error", 400)
			return
		}
		this.Data["Solution"] = one
	}

	t := template.New("layout.tpl")
	t, err = t.ParseFiles("view/layout.tpl", "view/contest/status_code.tpl")
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}

	err = t.Execute(w, this.Data)
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}
}
