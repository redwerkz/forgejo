// Copyright 2022 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package activitypub

import (
	repo_model "code.gitea.io/gitea/models/repo"
	user_model "code.gitea.io/gitea/models/user"
	"code.gitea.io/gitea/modules/forgefed"

	ap "github.com/go-ap/activitypub"
)

// Create Follow activity
func Follow(actorUser, followUser *user_model.User) *ap.Follow {
	object := ap.PersonNew(ap.IRI(followUser.LoginName))
	follow := ap.FollowNew("", object)
	follow.Actor = ap.PersonNew(ap.IRI(actorUser.GetIRI()))
	follow.To = ap.ItemCollection{ap.Item(ap.IRI(followUser.LoginName + "/inbox"))}
	return follow
}

// Create Undo Follow activity
func Unfollow(actorUser, followUser *user_model.User) *ap.Undo {
	unfollow := ap.UndoNew("", Follow(actorUser, followUser))
	unfollow.Actor = ap.PersonNew(ap.IRI(actorUser.GetIRI()))
	unfollow.To = ap.ItemCollection{ap.Item(ap.IRI(followUser.LoginName + "/inbox"))}
	return unfollow
}

// Create Like activity
func Star(user *user_model.User, repo *repo_model.Repository) *ap.Like {
	like := ap.LikeNew("", forgefed.RepositoryNew(ap.IRI(repo.GetIRI())))
	like.Actor = ap.PersonNew(ap.IRI(user.GetIRI()))
	like.To = ap.ItemCollection{ap.IRI(repo.GetIRI() + "/inbox")}
	return like
}

// Create Undo Like activity
func Unstar(user *user_model.User, repo *repo_model.Repository) *ap.Undo {
	unlike := ap.UndoNew("", Star(user, repo))
	unlike.Actor = ap.PersonNew(ap.IRI(user.GetIRI()))
	unlike.To = ap.ItemCollection{ap.IRI(repo.GetIRI() + "/inbox")}
	return unlike
}
