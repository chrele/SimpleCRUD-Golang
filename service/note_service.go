package service

import (
	"simple-crud-golang/data/request"
	"simple-crud-golang/data/response"
)

type NoteService interface {
	Create(note request.CreateNoteRequest)
	Update(note request.UpdateNoteRequest)
	Delete(noteId int)
	FindById(noteId int) response.NoteResponse
	FindAll() []response.NoteResponse
}
