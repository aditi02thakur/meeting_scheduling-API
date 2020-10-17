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


func getMeetings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(meetings)
}

func createMeeting(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var meeting Meeting
	_ = json.NewDecoder(r.Body).Decode(&meeting)
	meeting.ID = strconv.Itoa(rand.Intn(100000000))
	meetings = append(meetings, meeting)
	json.NewEncoder(w).Encode(meeting)
}

func updateMeeting(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range meetings {
		if item.ID == params["id"] {
			meetings = append(meetings[:index], meetings[index+1:]...)
			var meeting Meeting
			_ = json.NewDecoder(r.Body).Decode(&meeting)
			meeting.ID = params["id"]
			meetings = append(meetings, meeting)
			json.NewEncoder(w).Encode(meeting)
			return
		}
	}
}


func main() {

	r := mux.NewRouter()

	meetings = append(meetings, Meeting{ID: "1", Start_time: "16:00", End_time: "18:00", Title: "Meeting 1", Participant: &Participant{Firstname: "Harry", Lastname: "Potter"}})
	meetings = append(meetings, Meeting{ID: "2", Start_time: "12:00", End_time: "14:00", Title: "Meeting 2", Participant: &Participant{Firstname: "Ron", Lastname: "Weasley"}})
	meetings = append(meetings, Meeting{ID: "3", Start_time: "18:00", End_time: "20:00", Title: "Meeting 3", Participant: &Participant{Firstname: "Hermione", Lastname: "Granger"}})
	meetings = append(meetings, Meeting{ID: "4", Start_time: "14:00", End_time: "16:00", Title: "Meeting 4", Participant: &Participant{Firstname: "Sirius", Lastname: "Black"}})
	meetings = append(meetings, Meeting{ID: "5", Start_time: "10:00", End_time: "11:30", Title: "Meeting 5", Participant: &Participant{Firstname: "Neville", Lastname: "Longbottom"}})

	r.HandleFunc("/meetings", getMeetings).Methods("GET")
	r.HandleFunc("/meetings", createMeeting).Methods("POST")
	r.HandleFunc("/meetings/{id}", updateMeeting).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8000", r))
}