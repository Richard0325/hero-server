package data

import "errors"

var ErrHahowServer1000 error = errors.New("backend server error")
var ErrUnknown error = errors.New("unknown error")
var ErrIdNotFound error = errors.New("id not found")
var ErrNotAuthed error = errors.New("not authorizedd")
