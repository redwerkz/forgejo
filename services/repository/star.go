// Copyright 2023 The Forgejo Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package repository

import (
	"context"

	"code.gitea.io/gitea/models/auth"
	repo_model "code.gitea.io/gitea/models/repo"
	user_model "code.gitea.io/gitea/models/user"
	"code.gitea.io/gitea/services/activitypub"

	ap "github.com/go-ap/activitypub"
)

// StarRepo or unstar repository.
func StarRepo(ctx context.Context, userID, repoID int64, star bool) error {
	repo, err := repo_model.GetRepositoryByID(ctx, repoID)
	if err != nil {
		return err
	}
	err = repo.GetOwner(ctx)
	if err != nil {
		return err
	}
	if repo.Owner.LoginType == auth.Federated {
		// Federated repo
		user, err := user_model.GetUserByID(ctx, userID)
		if err != nil {
			return err
		}
		var activity *ap.Activity
		if star {
			activity = activitypub.Star(user, repo)
		} else {
			activity = activitypub.Unstar(user, repo)
		}
		err = activitypub.Send(user, activity)
		if err != nil {
			return err
		}
	}
	return repo_model.StarRepo(userID, repoID, star)
}
