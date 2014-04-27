package admin

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type contest struct {
	Cid      int         `json:"cid"bson:"cid"`
	Title    string      `json:"title"bson:"title"`
	Encrypt  int         `json:"encrypt"bson:"encrypt"`
	Argument interface{} `json:"argument"bson:"argument"`

	Start string `json:"start"bson:"start"`
	End   string `json:"end"bson:"end"`

	Status int    `json:"status"bson:"status"`
	Create string `'json:"create"bson:"create"`

	List []int `json:"list"bson:"list"`
}

type ContestController struct {
	Cid           int
	ContestDetail *contest
	Index         map[int]int
	class.Controller
}

func (this *ContestController) List(w http.ResponseWriter, r *http.Request) {
	log.Println("Contest List")
	this.Init(w, r)

	response, err := http.Post(config.PostHost+"/contest/list", "application", nil)
	defer response.Body.Close()
	if err != nil {
		http.Error(w, "post error", 500)
		return
	}

	one := make(map[string][]*contest)
	if response.StatusCode == 200 {
		err = this.LoadJson(response.Body, &one)
		if err != nil {
			http.Error(w, "load error", 400)
			return
		}
		this.Data["Contest"] = one["list"]
	}

	t := template.New("layout.tpl").Funcs(template.FuncMap{"ShowStatus": class.ShowStatus, "ShowExpire": class.ShowExpire, "ShowEncrypt": class.ShowEncrypt})
	t, err = t.ParseFiles("view/admin/layout.tpl", "view/admin/contest_list.tpl")
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}

	this.Data["Title"] = "Admin - Contest List"
	this.Data["IsContest"] = true
	this.Data["IsList"] = true
	err = t.Execute(w, this.Data)
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}

}

func (this *ContestController) Add(w http.ResponseWriter, r *http.Request) {
	log.Println("Admin Contest Add")
	this.Init(w, r)

	t := template.New("layout.tpl")
	t, err := t.ParseFiles("view/admin/layout.tpl", "view/admin/contest_add.tpl")
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}

	this.Data["Title"] = "Admin - Contest Add"
	this.Data["IsContest"] = true
	this.Data["IsAdd"] = true

	err = t.Execute(w, this.Data)
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}
}

func (this *ContestController) Insert(w http.ResponseWriter, r *http.Request) {
	log.Println("Admin Contest Insert")
	this.Init(w, r)

	one := make(map[string]interface{})
	one["title"] = r.FormValue("title")
	startTimeYear, err := strconv.Atoi(r.FormValue("startTimeYear"))
	if err != nil {
		http.Error(w, "conv error", 400)
		return
	}
	startTimeMonth, err := strconv.Atoi(r.FormValue("startTimeMonth"))
	if err != nil {
		http.Error(w, "conv error", 400)
		return
	}
	startTimeDay, err := strconv.Atoi(r.FormValue("startTimeDay"))
	if err != nil {
		http.Error(w, "conv error", 400)
		return
	}
	startTimeHour, err := strconv.Atoi(r.FormValue("startTimeHour"))
	if err != nil {
		http.Error(w, "conv error", 400)
		return
	}
	startTimeMinute, err := strconv.Atoi(r.FormValue("startTimeMinute"))
	if err != nil {
		http.Error(w, "conv error", 400)
		return
	}
	startTimeSecond, err := strconv.Atoi(r.FormValue("startTimeSecond"))
	if err != nil {
		http.Error(w, "conv error", 400)
		return
	}
	one["start"] = fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", startTimeYear, startTimeMonth, startTimeDay, startTimeHour, startTimeMinute, startTimeSecond)
	endTimeYear, err := strconv.Atoi(r.FormValue("endTimeYear"))
	if err != nil {
		http.Error(w, "conv error", 400)
		return
	}
	endTimeMonth, err := strconv.Atoi(r.FormValue("endTimeMonth"))
	if err != nil {
		http.Error(w, "conv error", 400)
		return
	}
	endTimeDay, err := strconv.Atoi(r.FormValue("endTimeDay"))
	if err != nil {
		http.Error(w, "conv error", 400)
		return
	}
	endTimeHour, err := strconv.Atoi(r.FormValue("endTimeHour"))
	if err != nil {
		http.Error(w, "conv error", 400)
		return
	}
	endTimeMinute, err := strconv.Atoi(r.FormValue("endTimeMinute"))
	if err != nil {
		http.Error(w, "conv error", 400)
		return
	}
	endTimeSecond, err := strconv.Atoi(r.FormValue("endTimeSecond"))
	if err != nil {
		http.Error(w, "conv error", 400)
		return
	}
	one["end"] = fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", endTimeYear, endTimeMonth, endTimeDay, endTimeHour, endTimeMinute, endTimeSecond)
	switch r.FormValue("type") {
	case "public":
		one["encrypt"] = config.EncryptPB
	case "private":
		one["encrypt"] = config.EncryptPT
	case "password":
		one["encrypt"] = config.EncryptPW
	default:
		http.Error(w, "args error", 400)
		return
	}
	/////
	one["argument"] = r.FormValue("argument")
	/////
	problemString := r.FormValue("problemList")
	problemString = strings.Trim(problemString, " ")
	problemString = strings.Trim(problemString, ";")
	problemList := strings.Split(problemString, ";")
	var list []int
	for _, v := range problemList {
		pid, err := strconv.Atoi(v)
		if err != nil {
			http.Error(w, "conv error", 400)
			return
		}
		list = append(list, pid)
	}
	one["list"] = list

	reader, err := this.PostReader(&one)
	if err != nil {
		http.Error(w, "read error", 500)
		return
	}

	response, err := http.Post(config.PostHost+"/contest/insert", "application/json", reader)
	defer response.Body.Close()
	if err != nil {
		http.Error(w, "post error", 500)
		return
	}

	ret := make(map[string]interface{})
	if response.StatusCode == 200 {
		err = this.LoadJson(response.Body, &ret)
		if err != nil {
			http.Error(w, "load error", 400)
			return
		}
		http.Redirect(w, r, "/admin/contest/list", http.StatusFound)
	}
}

func (this *ContestController) Status(w http.ResponseWriter, r *http.Request) {
	log.Println("Admin Contest Status")
	this.Init(w, r)

	args := this.ParseURL(r.URL.Path[6:])
	log.Println(args)
	cid, err := strconv.Atoi(args["cid"])
	if err != nil {
		http.Error(w, "args error", 400)
		return
	}

	response, err := http.Post(config.PostHost+"/contest/status/cid/"+strconv.Itoa(cid), "application/json", nil)
	defer response.Body.Close()
	if err != nil {
		http.Error(w, "post error", 500)
		return
	}

	if response.StatusCode == 200 {
		http.Redirect(w, r, "/admin/contest/list", http.StatusFound)
	}
}

func (this *ContestController) Delete(w http.ResponseWriter, r *http.Request) {
	log.Println("Admin Contest Delete")
	this.Init(w, r)

	args := this.ParseURL(r.URL.Path[6:])
	cid, err := strconv.Atoi(args["cid"])
	if err != nil {
		http.Error(w, "args error", 400)
		return
	}

	response, err := http.Post(config.PostHost+"/contest/delete/cid/"+strconv.Itoa(cid), "application/json", nil)
	defer response.Body.Close()
	if err != nil {
		http.Error(w, "post error", 500)
		return
	}

	w.WriteHeader(response.StatusCode)
}

func (this *ContestController) Edit(w http.ResponseWriter, r *http.Request) {
	log.Println("Admin Contest Edit")
	this.Init(w, r)

	args := this.ParseURL(r.URL.Path[6:])
	cid, err := strconv.Atoi(args["cid"])
	if err != nil {
		http.Error(w, "args error", 400)
		return
	}

	response, err := http.Post(config.PostHost+"/contest/detail/cid/"+strconv.Itoa(cid), "application/json", nil)
	defer response.Body.Close()
	if err != nil {
		http.Error(w, "post error", 500)
		return
	}

	var one struct {
		contest
		StartTimeYear   string
		StartTimeMonth  string
		StartTimeDay    string
		StartTimeHour   string
		StartTimeMinute string
		StartTimeSecond string
		EndTimeYear     string
		EndTimeMonth    string
		EndTimeDay      string
		EndTimeHour     string
		EndTimeMinute   string
		EndTimeSecond   string
		ProblemList     string
	}
	if response.StatusCode == 200 {
		err = this.LoadJson(response.Body, &one)
		if err != nil {
			http.Error(w, "load error", 400)
			return
		}
		one.StartTimeYear = one.Start[0:4]
		one.StartTimeMonth = one.Start[5:7]
		one.StartTimeDay = one.Start[8:10]
		one.StartTimeHour = one.Start[11:13]
		one.StartTimeMinute = one.Start[14:16]
		one.StartTimeSecond = one.Start[17:19]
		one.EndTimeYear = one.End[0:4]
		one.EndTimeMonth = one.End[5:7]
		one.EndTimeDay = one.End[8:10]
		one.EndTimeHour = one.End[11:13]
		one.EndTimeMinute = one.End[14:16]
		one.EndTimeSecond = one.End[17:19]
		one.ProblemList = ""
		for _, v := range one.List {
			one.ProblemList += strconv.Itoa(v) + ";"
		}
		this.Data["Detail"] = one
	} else {
		http.Error(w, "resp error", 500)
		return
	}

	t := template.New("layout.tpl").Funcs(template.FuncMap{"ShowRatio": class.ShowRatio})
	t, err = t.ParseFiles("view/admin/layout.tpl", "view/admin/contest_edit.tpl")
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}

	this.Data["Title"] = "Admin - Contest Edit"
	this.Data["IsContest"] = true
	this.Data["IsEdit"] = true

	err = t.Execute(w, this.Data)
	if err != nil {
		http.Error(w, "tpl error", 500)
		return
	}
}

func (this *ContestController) Update(w http.ResponseWriter, r *http.Request) {
	log.Println("Admin Contest Update")
	this.Init(w, r)

	args := this.ParseURL(r.URL.Path[6:])
	cid, err := strconv.Atoi(args["cid"])
	if err != nil {
		http.Error(w, "args error", 400)
		return
	}

	one := make(map[string]interface{})
	one["title"] = r.FormValue("title")
	startTimeYear, err := strconv.Atoi(r.FormValue("startTimeYear"))
	if err != nil {
		http.Error(w, "conv error", 400)
		return
	}
	startTimeMonth, err := strconv.Atoi(r.FormValue("startTimeMonth"))
	if err != nil {
		http.Error(w, "conv error", 400)
		return
	}
	startTimeDay, err := strconv.Atoi(r.FormValue("startTimeDay"))
	if err != nil {
		http.Error(w, "conv error", 400)
		return
	}
	startTimeHour, err := strconv.Atoi(r.FormValue("startTimeHour"))
	if err != nil {
		http.Error(w, "conv error", 400)
		return
	}
	startTimeMinute, err := strconv.Atoi(r.FormValue("startTimeMinute"))
	if err != nil {
		http.Error(w, "conv error", 400)
		return
	}
	startTimeSecond, err := strconv.Atoi(r.FormValue("startTimeSecond"))
	if err != nil {
		http.Error(w, "conv error", 400)
		return
	}
	one["start"] = fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", startTimeYear, startTimeMonth, startTimeDay, startTimeHour, startTimeMinute, startTimeSecond)
	endTimeYear, err := strconv.Atoi(r.FormValue("endTimeYear"))
	if err != nil {
		http.Error(w, "conv error", 400)
		return
	}
	endTimeMonth, err := strconv.Atoi(r.FormValue("endTimeMonth"))
	if err != nil {
		http.Error(w, "conv error", 400)
		return
	}
	endTimeDay, err := strconv.Atoi(r.FormValue("endTimeDay"))
	if err != nil {
		http.Error(w, "conv error", 400)
		return
	}
	endTimeHour, err := strconv.Atoi(r.FormValue("endTimeHour"))
	if err != nil {
		http.Error(w, "conv error", 400)
		return
	}
	endTimeMinute, err := strconv.Atoi(r.FormValue("endTimeMinute"))
	if err != nil {
		http.Error(w, "conv error", 400)
		return
	}
	endTimeSecond, err := strconv.Atoi(r.FormValue("endTimeSecond"))
	if err != nil {
		http.Error(w, "conv error", 400)
		return
	}
	one["end"] = fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", endTimeYear, endTimeMonth, endTimeDay, endTimeHour, endTimeMinute, endTimeSecond)
	switch r.FormValue("type") {
	case "public":
		one["encrypt"] = config.EncryptPB
	case "private":
		one["encrypt"] = config.EncryptPT
	case "password":
		one["encrypt"] = config.EncryptPW
	default:
		http.Error(w, "args error", 400)
		return
	}
	/////
	one["argument"] = r.FormValue("argument")
	/////
	problemString := r.FormValue("problemList")
	problemString = strings.Trim(problemString, " ")
	problemString = strings.Trim(problemString, ";")
	problemList := strings.Split(problemString, ";")
	var list []int
	for _, v := range problemList {
		pid, err := strconv.Atoi(v)
		if err != nil {
			http.Error(w, "conv error", 400)
			return
		}
		list = append(list, pid)
	}
	one["list"] = list

	reader, err := this.PostReader(&one)
	if err != nil {
		http.Error(w, "read error", 500)
		return
	}

	response, err := http.Post(config.PostHost+"/contest/update/cid/"+strconv.Itoa(cid), "application/json", reader)
	defer response.Body.Close()
	if err != nil {
		http.Error(w, "post error", 500)
		return
	}

	if response.StatusCode == 200 {
		http.Redirect(w, r, "/admin/contest/list", http.StatusFound)
	}
}
