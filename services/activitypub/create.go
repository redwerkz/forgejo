// Copyright 2022 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package activitypub

import (
	ap "github.com/go-ap/activitypub"
)

func Create(to string, object ap.ObjectOrLink) *ap.Create {
	return &ap.Create{
		Type:   ap.CreateType,
		Object: object,
		To:     ap.ItemCollection{ap.Item(ap.IRI(to))},
	}
}
