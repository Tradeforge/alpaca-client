package model

type Document struct {
	DocumentType    string
	DocumentSubType *string
	Content         string
	ContentData     any
	MimeType        *string
}

type UploadDocumentRequest struct {
	DocumentType    string
	DocumentSubType *string
	Content         string
	ContentData     any
	MimeType        *string
}

type UploadDocumentParams struct {
	AccountID string `path:"id"`
}

type UploadDocumentResponse struct{}
