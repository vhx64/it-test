package gcperrors

import (
	"context"
	"os"

	"github.com/sirupsen/logrus"

	"cloud.google.com/go/errorreporting"
)

type GCPErrorReporting struct {
	errorClient *errorreporting.Client
}

func NewGCPErrorReporting() *GCPErrorReporting {
	return &GCPErrorReporting{}
}

func (r *GCPErrorReporting) ReportError(err error) {
	r.client().Report(errorreporting.Entry{
		Error: err,
	})
}

func (r *GCPErrorReporting) client() *errorreporting.Client {
	if r.errorClient != nil {
		return r.errorClient
	}
	ctx := context.Background()
	errorClient, err := errorreporting.NewClient(ctx, os.Getenv("GCP_PROJECT_ID"), errorreporting.Config{
		ServiceName:    "it-test",
		ServiceVersion: os.Getenv("APPLICATION_VERSION"),
		OnError: func(err error) {
			logrus.Error(err)
		},
	})
	if err != nil {
		logrus.Errorf("failed to instantiate error client: %+v", err)
	}
	return errorClient
}
