// Copyright 2022 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package activitypub

import (
	repo_model "code.gitea.io/gitea/models/repo"
	user_model "code.gitea.io/gitea/models/user"
	"code.gitea.io/gitea/modules/forgefed"

	ap "github.com/go-ap/activitypub"
)

// Create Follow activity
func Follow(actorUser, followUser *user_model.User) (follow *ap.Follow) {
	object := ap.PersonNew(ap.IRI(followUser.LoginName))
	follow = ap.FollowNew("", object)
	follow.Actor = ap.PersonNew(ap.IRI(actorUser.GetIRI()))
	follow.To = ap.ItemCollection{ap.Item(ap.IRI(followUser.LoginName + "/inbox"))}
	return
}

// Create Undo Follow activity
func Unfollow(actorUser, followUser *user_model.User) (unfollow *ap.Undo) {
	unfollow = ap.UndoNew("", Follow(actorUser, followUser))
	unfollow.Actor = ap.PersonNew(ap.IRI(actorUser.GetIRI()))
	unfollow.To = ap.ItemCollection{ap.Item(ap.IRI(followUser.LoginName + "/inbox"))}
	return
}

// Create Like activity
func Star(user *user_model.User, repo *repo_model.Repository) (like *ap.Like) {
	like = ap.LikeNew("", forgefed.RepositoryNew(ap.IRI(repo.GetIRI())))
	like.Actor = ap.PersonNew(ap.IRI(user.GetIRI()))
	like.To = ap.ItemCollection{ap.IRI(repo.GetIRI() + "/inbox")}
	return
}

// Create Undo Like activity
func Unstar(user *user_model.User, repo *repo_model.Repository) (unstar *ap.Undo) {
	unstar = ap.UndoNew("", Star(user, repo))
	unstar.Actor = ap.PersonNew(ap.IRI(user.GetIRI()))
	unstar.To = ap.ItemCollection{ap.IRI(repo.GetIRI() + "/inbox")}
	return
}
