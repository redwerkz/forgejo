// Copyright 2022 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package activitypub

import (
	user_model "code.gitea.io/gitea/models/user"
	
	ap "github.com/go-ap/activitypub"
)

func Create(user *user_model.User, object ap.ObjectOrLink, to string) *ap.Create {
	return &ap.Create{
		Type:   ap.CreateType,
		Actor:  ap.IRI(user.GetIRI()),
		Object: object,
		To:     ap.ItemCollection{ap.Item(ap.IRI(to))},
	}
}
