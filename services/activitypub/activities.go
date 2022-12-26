// Copyright 2022 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package activitypub

import (
	user_model "code.gitea.io/gitea/models/user"

	ap "github.com/go-ap/activitypub"
)

// Create Follow activity
func Follow(actorUser, followUser *user_model.User) *ap.Follow {
	object := ap.PersonNew(ap.IRI(followUser.LoginName))
	follow := ap.FollowNew("", object)
	follow.Type = ap.FollowType
	follow.Actor = ap.PersonNew(ap.IRI(actorUser.GetIRI()))
	follow.To = ap.ItemCollection{ap.Item(ap.IRI(followUser.LoginName + "/inbox"))}
	return follow
}

// Create Undo Follow activity
func Unfollow(actorUser, followUser *user_model.User) *ap.Undo {
	object := ap.PersonNew(ap.IRI(followUser.LoginName))
	follow := ap.FollowNew("", object)
	follow.Actor = ap.PersonNew(ap.IRI(actorUser.GetIRI()))
	unfollow := ap.UndoNew("", follow)
	unfollow.Type = ap.UndoType
	unfollow.To = ap.ItemCollection{ap.Item(ap.IRI(followUser.LoginName + "/inbox"))}
	return unfollow
}
