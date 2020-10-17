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
	Participants string 'json:"participants"'
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