package main

// Laird Assessors API BrIDge for Pro-claim users
//
// This is a early beta and may lack proper errors
//
// Made By Robert Wiggins (robert.wiggins@laird-assessors.com)
//
// Any Errors please provide the full request to assist me in diagnosing the issue
//


import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")

}


func tworkflows(w http.ResponseWriter, r *http.Request) {

	// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	APIKEY := r.FormValue("APIKEY")
	ID := r.FormValue("ID")
	enableCors(&w)


	genurl := "https://test-lairdassessors.swiftcase.co.uk/api/v2/"+APIKEY+"/workflow/"+ID+".xml"

	req, err := http.NewRequest("GET", genurl, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Host = "test-lairdassessors.swiftcase.co.uk"
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("User-Agent", "Laird API Bridge GoLang")
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
	if resp.StatusCode == 200 {
		w.Header().Set("Content-Type", "application/xml")
	}

	w.WriteHeader(resp.StatusCode)
	fmt.Printf("Laird API URL:"+genurl+"\n")
	fmt.Printf("Laird API Status Code Response: %d\n", resp.StatusCode)
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
	enableCors(&w)

	genurl := "https://lairdassessors.swiftcase.co.uk/api/v2/"+APIKEY+"/workflow/"+ID+".xml"

	req, err := http.NewRequest("GET", genurl, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Host = "lairdassessors.swiftcase.co.uk"
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("User-Agent", "Laird API Bridge GoLang")
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
	if resp.StatusCode == 200 {
		w.Header().Set("Content-Type", "application/xml")
	}
	w.WriteHeader(resp.StatusCode)
	fmt.Printf("Laird API URL: "+genurl+"\n")
	fmt.Printf("Laird API Status Code Response: %d\n", resp.StatusCode)
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
	TASK := r.FormValue("TASK")
	enableCors(&w)

	genurl := "https://test-lairdassessors.swiftcase.co.uk/api/v2/"+APIKEY+"/task/"+TASK+"/status.xml"

	req, err := http.NewRequest("GET", genurl, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Host = "test-lairdassessors.swiftcase.co.uk"
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("User-Agent", "Laird API Bridge GoLang")
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
	if resp.StatusCode == 200 {
		w.Header().Set("Content-Type", "application/xml")
	}
	w.WriteHeader(resp.StatusCode)
	fmt.Printf("Laird API URL: "+genurl+"\n")
	fmt.Printf("Laird API Status Code Response: %d\n", resp.StatusCode)
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
	TASK := r.FormValue("TASK")
	enableCors(&w)

	genurl := "https://lairdassessors.swiftcase.co.uk/api/v2/"+APIKEY+"/task/"+TASK+"/status.xml"

	req, err := http.NewRequest("GET", genurl, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Host = "lairdassessors.swiftcase.co.uk"
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("User-Agent", "Laird API Bridge GoLang")
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
	if resp.StatusCode == 200 {
		w.Header().Set("Content-Type", "application/xml")
	}
	w.WriteHeader(resp.StatusCode)
	fmt.Printf("Laird API URL: "+genurl+"\n")
	fmt.Printf("Laird API Status Code Response: %d\n", resp.StatusCode)
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
	enableCors(&w)

	genurl := "https://test-lairdassessors.swiftcase.co.uk/api/v2/"+APIKEY+"/task/"+TASK+"/availability/"+IDATE+"/"+MAXRESULT+".xml"

	req, err := http.NewRequest("GET", genurl, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Host = "test-lairdassessors.swiftcase.co.uk"
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("User-Agent", "Laird API Bridge GoLang")
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
	if resp.StatusCode == 200 {
		w.Header().Set("Content-Type", "application/xml")
	}
	w.WriteHeader(resp.StatusCode)
	fmt.Printf("Laird API URL: "+genurl+"\n")
	fmt.Printf("Laird API Status Code Response: %d\n", resp.StatusCode)
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
	enableCors(&w)

	genurl := "https://lairdassessors.swiftcase.co.uk/api/v2/"+APIKEY+"/task/"+TASK+"/availability/"+IDATE+"/"+MAXRESULT+".xml"

	req, err := http.NewRequest("GET", genurl, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Host = "lairdassessors.swiftcase.co.uk"
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("User-Agent", "Laird API Bridge GoLang")
	req.Header.Set("Accept", "*/*")
	//req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("Connection", "close")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Laird API Status Code Response: %d\n", resp.StatusCode)
	defer resp.Body.Close()
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	if resp.StatusCode == 200 {
		w.Header().Set("Content-Type", "application/xml")
	}
	w.WriteHeader(resp.StatusCode)
	fmt.Printf("Laird API URL: "+genurl+"\n")
	fmt.Printf("Laird API Status Code Response: %d\n", resp.StatusCode)
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
