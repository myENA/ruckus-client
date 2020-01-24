package vsz

// API Version: v8_1

import (
	"context"
	"net/http"
)

type WSGGDPRService struct {
	apiClient *APIClient
}

func NewWSGGDPRService(c *APIClient) *WSGGDPRService {
	s := new(WSGGDPRService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGGDPRService() *WSGGDPRService {
	return NewWSGGDPRService(ss.apiClient)
}

type WSGGDPRFtp struct {
	// FtpHost
	// IP/DN of FTP
	FtpHost *string `json:"ftpHost,omitempty"`

	// FtpPassword
	// Password for FTP login
	FtpPassword *string `json:"ftpPassword,omitempty"`

	// FtpPort
	// Port used by FTP
	// Constraints:
	//    - min:21
	//    - max:65535
	FtpPort *int `json:"ftpPort,omitempty" validate:"gte=21,lte=65535"`

	// FtpProtocol
	// Protocol used
	// Constraints:
	//    - oneof:[FTP,SFTP]
	FtpProtocol *string `json:"ftpProtocol,omitempty" validate:"oneof=FTP SFTP"`

	// FtpRemoteDirectory
	// Destination directory used for file upload
	FtpRemoteDirectory *string `json:"ftpRemoteDirectory,omitempty"`

	// FtpUserName
	// Username for FTP login
	FtpUserName *string `json:"ftpUserName,omitempty"`
}

func NewWSGGDPRFtp() *WSGGDPRFtp {
	m := new(WSGGDPRFtp)
	return m
}

type WSGGDPRReport struct {
	// Action
	// Request action
	// Constraints:
	//    - oneof:[SEARCH,DELETE,INTERRUPT,PROGRESS]
	Action *string `json:"action,omitempty" validate:"oneof=SEARCH DELETE INTERRUPT PROGRESS"`

	// ClientMac
	// Client mac
	ClientMac *string `json:"clientMac,omitempty"`

	Ftp *WSGGDPRFtp `json:"ftp,omitempty"`
}

func NewWSGGDPRReport() *WSGGDPRReport {
	m := new(WSGGDPRReport)
	return m
}

// AddGdprReport
//
// Use this API command to execute a client-related data search or delete task and upload a report to FTP. Also use this API to check task progress or to interrupt it.
//
// Request Body:
//	 - body *WSGGDPRReport
func (s *WSGGDPRService) AddGdprReport(ctx context.Context, body *WSGGDPRReport) (interface{}, error) {
	var (
		req  *APIRequest
		resp interface{}
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddGdprReport, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}
