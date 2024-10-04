package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Appointments struct {
	ID          int    `json:"id"`
	PName       string `json:"pname"`
	Date        string `json:"date"`
	Description string `json:"description"`
}

var appoint []Appointments

// var nextID int = 1

func postAppointments(w http.ResponseWriter, r *http.Request) {
	var newAppoint Appointments
	json.NewDecoder(r.Body).Decode(&newAppoint)
	newAppoint.ID = len(appoint) + 1
	appoint = append(appoint, newAppoint)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newAppoint)
	fmt.Fprintf(w, "\n\nAppointment successsfully booked")
}

func getAppointments(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/appointments/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	for _, appoint := range appoint {
		if appoint.ID == id {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, "Appointment Requested is as below\n\n")
			json.NewEncoder(w).Encode(appoint)
			return
		}
	}
	http.Error(w, "Task not found", http.StatusNotFound)
}

func getALLAppointments(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "All Appointment Showed\n\n")
	json.NewEncoder(w).Encode(appoint)

}

func updateAppointments(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/appointments/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid appointment ID", http.StatusBadRequest)
		return
	}
	var updatedAppointment Appointments
	json.NewDecoder(r.Body).Decode(&updatedAppointment)
	for i, appointment := range appoint {
		if appointment.ID == id {
			json.NewDecoder(r.Body).Decode(&updatedAppointment)
			updatedAppointment.ID = id
			appoint[i] = updatedAppointment
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedAppointment)
			fmt.Fprintf(w, "\n\nRequested Appointment Updated Successfully")
			return
		}
	}
	http.Error(w, "Appointment not found", http.StatusNotFound)
}

func deleteAppointment(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/appointments/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Appointment not found", http.StatusNotFound)
		return
	}
	for i, appointment := range appoint {
		//getting the correct id to delete
		if appointment.ID == id {
			appoint = append(appoint[:i], appoint[i+1:]...)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, "\nRequested Appointment Deleted Successfully\n")
			return
		}
	}
	http.Error(w, "Appointment not found", http.StatusNotFound)
}

func handleAppointmnets(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "ALL":
		getALLAppointments(w)
	case "GET":
		getAppointments(w, r)
	case "POST":
		postAppointments(w, r)
	case "PUT":
		updateAppointments(w, r)
	case "DELETE":
		deleteAppointment(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}

func main() {
	// http.HandlerFunc("/appointments/All", getALLAppointments)
	http.HandleFunc("/appointments/", handleAppointmnets)
	fmt.Println("Server successfully started...")
	http.ListenAndServe(":5455", nil)
}
