package functools

import (
	"encoding/json"
	"github.com/getsentry/sentry-go"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"net/http"
)

type Request struct {
	*http.Request
}

const (
	JSONContentType   = "application/json"
	BinaryContentType = "application/octet-stream"
)

func (request *Request) GetContentType() string {
	contentType := request.Header.Get("Content-Type")

	if contentType == "" {
		contentType = BinaryContentType
	}

	return contentType
}

func (request *Request) IsContentTypeAllowed(allowedTypes *[]string, headers http.Header) bool {
	if allowedTypes == nil {
		allowedTypes = &[]string{BinaryContentType, JSONContentType}
	}

	contentType := request.GetContentType()
	return Contains(contentType, *allowedTypes)
}

func (request *Request) IsProto() bool {
	return request.GetContentType() == BinaryContentType
}

func (request *Request) ParseBody(message proto.Message) error {

	b, err := ioutil.ReadAll(request.Body)
	defer request.Body.Close()
	if err != nil {
		sentry.CaptureException(err)
		return err
	}

	if request.IsProto() {
		err = proto.Unmarshal(b, message)
	} else {
		err = json.Unmarshal(b, message)
	}

	if err != nil {
		sentry.CaptureException(err)
		return err
	}

	return err
}