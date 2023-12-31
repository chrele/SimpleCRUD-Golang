package repository

import "simple-crud-golang/model"

type NoteRepository interface {
	Save(note model.Note)
	Update(note model.Note)
	Delete(noteId int)
	FindById(noteId int) (model.Note, error)
	FindAll() []model.Note
}
