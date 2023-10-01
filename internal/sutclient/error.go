package sutclient

import "errors"

var ErrFailedToInitUser = errors.New("failed to init login")
var ErrBadUser = errors.New("bad login or password")
var ErrForbidden = errors.New("forbidden sut_client")

var ErrUnexpectedStatusCode = errors.New("unexpected status code")

var ErrNoLessonToCommit = errors.New("no lesson to open")
var ErrBadLessonParams = errors.New("failed to match lesson params")
