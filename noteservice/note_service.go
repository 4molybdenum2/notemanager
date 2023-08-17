package noteservice

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/4molybdenum2/notemanager/noteservice/store"
)

type NoteController struct {
	Notes []store.Note
}

func handleNotInHeader(rw http.ResponseWriter, r *http.Request, param string) {
	rw.WriteHeader(http.StatusBadRequest)
	rw.Write([]byte(fmt.Sprintf("%s missing in header", param)))
}

func NewNoteController() *NoteController {
	notes := make([]store.Note, 0)
	return &NoteController{
		Notes: notes,
	}
}

func (n *NoteController) GetAllNote(rw http.ResponseWriter, r *http.Request) {

	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(n.Notes)
}

func (n *NoteController) GetNote(rw http.ResponseWriter, r *http.Request) {
	Note := store.Note{}

	if _, ok := r.Header["Id"]; !ok {
		handleNotInHeader(rw, r, "Id")
		return
	}

	Id, err := strconv.Atoi(r.Header["Id"][0])
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Invalid id"))
	}

	found := false
	for _, note := range n.Notes {
		if note.Id == Id {
			found = true
			Note = note
		}
	}

	if !found {
		rw.Write([]byte(fmt.Sprintf("Note not found with Id: %d", Id)))
		return
	}

	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(Note)
}

func (n *NoteController) PostNote(rw http.ResponseWriter, r *http.Request) {
	if _, ok := r.Header["Title"]; !ok {
		handleNotInHeader(rw, r, "Id")
		return
	}

	if _, ok := r.Header["Content"]; !ok {
		handleNotInHeader(rw, r, "Id")
		return
	}

	Id := len(n.Notes) + 1

	Note := store.Note{
		Id:      Id,
		Title:   r.Header["Title"][0],
		Content: r.Header["Content"][0],
	}

	n.Notes = append(n.Notes, Note)
	
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(Note)
}

func (n *NoteController) UpdateNote(rw http.ResponseWriter, r *http.Request) {
	Note := store.Note{}
	if _, ok := r.Header["Id"]; !ok {
		handleNotInHeader(rw, r, "Id")
		return
	}
	if _, ok := r.Header["Content"]; !ok {
		handleNotInHeader(rw, r, "Id")
		return
	}

	Id, err := strconv.Atoi(r.Header["Id"][0])
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Invalid id"))
	}

	found := false
	for _, note := range n.Notes {
		if note.Id == Id {
			found = true
			Note = store.Note{
				Id:      Id,
				Title:   r.Header["Title"][0],
				Content: r.Header["Content"][0],
			}
		}
	}

	if !found {
		rw.Write([]byte(fmt.Sprintf("Note not found with Id: %d", Id)))
		return
	}

	n.Notes[Id] = Note

	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(Note)
}

func (n *NoteController) DeleteNote(rw http.ResponseWriter, r *http.Request) {
	Notes := []store.Note{}
	deletedNote := store.Note{}

	if _, ok := r.Header["Id"]; !ok {
		handleNotInHeader(rw, r, "Id")
		return
	}

	Id, err := strconv.Atoi(r.Header["Id"][0])
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Invalid id"))
	}

	found := false
	for _, note := range n.Notes {
		if note.Id == Id {
			found = true
			deletedNote = note
			continue
		}
		Notes = append(Notes, note)
	}

	if !found {
		rw.Write([]byte(fmt.Sprintf("Note not found with Id: %d", Id)))
		return
	}

	n.Notes = Notes

	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(deletedNote)
}
