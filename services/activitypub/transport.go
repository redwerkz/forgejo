// Copyright 2023 The Forgejo Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package activitypub

import (
	"fmt"
	"io"
	"net/http"

	user_model "code.gitea.io/gitea/models/user"
	"code.gitea.io/gitea/modules/forgefed"
	"code.gitea.io/gitea/modules/httplib"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"

	ap "github.com/go-ap/activitypub"
	"github.com/go-ap/jsonld"
)

// Fetch a URL as binary
func Fetch(iri string) (b []byte, err error) {
	req := httplib.NewRequest(iri, http.MethodGet)
	req.Header("Accept", ActivityStreamsContentType)
	req.Header("User-Agent", "Gitea/"+setting.AppVer)
	resp, err := req.Response()
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("url IRI fetch [%s] failed with status (%d): %s", iri, resp.StatusCode, resp.Status)
		return
	}
	b, err = io.ReadAll(io.LimitReader(resp.Body, setting.Federation.MaxSize))
	return b, err
}

// Fetch a remote ActivityStreams object as an object
func FetchObject(iri string) (ap.ObjectOrLink, error) {
	resp, err := Fetch(iri)
	if err != nil {
		return nil, err
	}
	ap.ItemTyperFunc = forgefed.GetItemByType
	ap.JSONItemUnmarshal = forgefed.JSONUnmarshalerFn
	ap.NotEmptyChecker = forgefed.NotEmpty
	return ap.UnmarshalJSON(resp)
}

// Send an activity
func Send(user *user_model.User, activity *ap.Activity) error {
	binary, err := jsonld.WithContext(
		jsonld.IRI(ap.ActivityBaseURI),
		jsonld.IRI(ap.SecurityContextURI),
		jsonld.IRI(forgefed.ForgeFedNamespaceURI),
	).Marshal(activity)
	if err != nil {
		return err
	}

	for _, to := range activity.To {
		client, _ := NewClient(user, user.GetIRI()+"#main-key")
		resp, _ := client.Post(binary, to.GetLink().String())
		respBody, _ := io.ReadAll(io.LimitReader(resp.Body, setting.Federation.MaxSize))
		log.Trace("Response from sending activity", string(respBody))
	}
	return err
}