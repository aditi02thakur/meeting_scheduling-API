package zoomAPI

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

Meeting  struct{
	ID int 'json:"id,omitempty"'
	Title string 'json:"title"'
	Participant string 'json:"participant"'
	StartTime *Time 'json:"start_time"'
	EndTime *Time 'json:"end_time"'
	CreationTimestamp string 'json:"creationTimestamp"'
}

Participants struct {
		Name string 'json:"name"'
		Email string 'json:"email"'
		RSVP string 'json:"rspv"'
}
Occurrence struct {
		ID  int    'json:"occurrence_id"'
		StartTime *Time  'json:"start_time"'
		EndTime *Time    'json:"end_time"'
		RSVP string 'json:"rspv"'
}

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
    http.HandleFunc("/", homePage)
    http.HandleFunc("/meetings", returnAllMeetings)
    log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
    handleRequests()
}

func returnAllMeetings(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: returnAllMeetings")
    json.NewEncoder(w).Encode(Meetings)
}