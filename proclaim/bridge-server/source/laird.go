package main

// Laird Assessors API BrIDge for Pro-claim users
//
// This is a early beta and may lack proper errors
//
// Made By Robert Wiggins (robert.wiggins@laird-assessors.com)
//
// Any Errors please provIDe the full request and response to assist me in diagnosing the issue
//
//
//
//
//

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)



func tworkflows(w http.ResponseWriter, r *http.Request) {

	// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	APIKEY := r.FormValue("APIKEY")
	ID := r.FormValue("ID")


	genurl := "https://test-lairdassessors.swiftcase.co.uk/api/v2/"+APIKEY+"/workflow/"+ID+".xml"

	req, err := http.NewRequest("GET", genurl, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Host = "test-lairdassessors.swiftcase.co.uk"
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("User-Agent", "Laird API BrIDge GoLang")
	req.Header.Set("Accept", "*/*")
	//req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("Connection", "close")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	responseString := string(responseData)
	fmt.Fprint(w, responseString)

}

func lworkflows(w http.ResponseWriter, r *http.Request) {
	// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	APIKEY := r.FormValue("APIKEY")
	ID := r.FormValue("ID")


	genurl := "https://lairdassessors.swiftcase.co.uk/api/v2/"+APIKEY+"/workflow/"+ID+".xml"

	req, err := http.NewRequest("GET", genurl, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Host = "lairdassessors.swiftcase.co.uk"
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("User-Agent", "Laird API BrIDge GoLang")
	req.Header.Set("Accept", "*/*")
	//req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("Connection", "close")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	responseString := string(responseData)
	fmt.Fprint(w, responseString)
}


func tstatus(w http.ResponseWriter, r *http.Request) {
	// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	APIKEY := r.FormValue("APIKEY")
	ID := r.FormValue("ID")


	genurl := "https://test-lairdassessors.swiftcase.co.uk/api/v2/"+APIKEY+"/task/"+ID+"/status.xml"

	req, err := http.NewRequest("GET", genurl, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Host = "test-lairdassessors.swiftcase.co.uk"
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("User-Agent", "Laird API BrIDge GoLang")
	req.Header.Set("Accept", "*/*")
	//req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("Connection", "close")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	responseString := string(responseData)
	fmt.Fprint(w, responseString)
}

func lstatus(w http.ResponseWriter, r *http.Request) {
	// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	APIKEY := r.FormValue("APIKEY")
	ID := r.FormValue("ID")


	genurl := "https://lairdassessors.swiftcase.co.uk/api/v2/"+APIKEY+"/task/"+ID+"/status.xml"

	req, err := http.NewRequest("GET", genurl, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Host = "lairdassessors.swiftcase.co.uk"
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("User-Agent", "Laird API BrIDge GoLang")
	req.Header.Set("Accept", "*/*")
	//req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("Connection", "close")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	responseString := string(responseData)
	fmt.Fprint(w, responseString)
}


func tengineer(w http.ResponseWriter, r *http.Request) {
	// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	APIKEY := r.FormValue("APIKEY")
	TASK := r.FormValue("TASK")
	IDATE := r.FormValue("IDATE")
	MAXRESULT := r.FormValue("MAXRESULT")


	genurl := "https://test-lairdassessors.swiftcase.co.uk/api/v2/"+APIKEY+"/task/"+TASK+"/availability/"+IDATE+"/"+MAXRESULT+".xml"

	req, err := http.NewRequest("GET", genurl, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Host = "test-lairdassessors.swiftcase.co.uk"
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("User-Agent", "Laird API BrIDge GoLang")
	req.Header.Set("Accept", "*/*")
	//req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("Connection", "close")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	responseString := string(responseData)
	fmt.Fprint(w, responseString)
}

func lengineer(w http.ResponseWriter, r *http.Request) {
	// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	APIKEY := r.FormValue("APIKEY")
	TASK := r.FormValue("TASk")
	IDATE := r.FormValue("IDATE")
	MAXRESULT := r.FormValue("MAXRESULT")


	genurl := "https://lairdassessors.swiftcase.co.uk/api/v2/"+APIKEY+"/task/"+TASK+"/availability/"+IDATE+"/"+MAXRESULT+".xml"

	req, err := http.NewRequest("GET", genurl, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Host = "lairdassessors.swiftcase.co.uk"
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("User-Agent", "Laird API BrIDge GoLang")
	req.Header.Set("Accept", "*/*")
	//req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("Connection", "close")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	responseString := string(responseData)
	fmt.Fprint(w, responseString)
}







func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Laird Assesors API Brdige for Proclaim")
	})


	http.HandleFunc("/test/workflow", tworkflows)
	http.HandleFunc("/test/status", tstatus)
	http.HandleFunc("/test/engineer", tengineer)

	http.HandleFunc("/live/workflow", lworkflows)
	http.HandleFunc("/live/status", lstatus)
	http.HandleFunc("/live/engineer", lengineer)




	fmt.Printf("Starting Laird Assessors API Bridge Server...\n")
	err := http.ListenAndServe(":9090", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
