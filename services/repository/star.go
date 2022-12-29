// Copyright 2016 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package repository

import (
	"context"
	"strings"

	repo_model "code.gitea.io/gitea/models/repo"
	user_model "code.gitea.io/gitea/models/user"
	"code.gitea.io/gitea/services/activitypub"
)

// StarRepo or unstar repository.
func StarRepo(ctx context.Context, userID, repoID int64, star bool) error {
	repo, err := repo_model.GetRepositoryByID(ctx, repoID)
	if err != nil {
		return err
	}
	if strings.Contains(repo.Name, "@") {
		// Federated repo
		user, err := user_model.GetUserByID(ctx, userID)
		if err != nil {
			return err
		}
		err = activitypub.Send(user, activitypub.Star(user, repo))
		if err != nil {
			return err
		}
	}
	return repo_model.StarRepo(userID, repoID, star)
}
