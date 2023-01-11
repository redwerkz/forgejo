// Copyright 2022 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package activitypub

import (
	"context"

	user_service "code.gitea.io/gitea/services/user"
	"code.gitea.io/gitea/services/activitypub"

	ap "github.com/go-ap/activitypub"
)

// Process an incoming Delete activity
func delete(ctx context.Context, delete ap.Delete) error {
	actorIRI := delete.Actor.GetLink()
	objectIRI := delete.Object.GetLink()
	// Make sure actor matches the object getting deleted
	if actorIRI != objectIRI {
		return nil
	}
	
	// Object is the user getting deleted
	objectUser, err := activitypub.PersonIRIToUser(ctx, objectIRI)
	if err != nil {
		return err
	}
	user_service.DeleteUser(ctx, objectUser, true)
	return nil
}
