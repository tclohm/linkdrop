package graph


				// This file will be automatically regenerated based on the schema, any resolver implementations
				// will be copied through when generating and any unknown code will be moved to the end.

import (
"context"
"fmt"
"crypto/sha256"
"encoding/base64"
"math/rand"
"github.com/tclohm/linkdrop/graph/generated"
"github.com/tclohm/linkdrop/graph/model")


















func (r *mutationResolver) CreateLink(ctx context.Context, input model.NewLink) (*model.Link, error) {
		for _, link := range r.links {
		if link.URL == input.URL {
			return &link, nil
		}
	}

	hasher := sha256.New()
	hasher.Write([]byte(input.URL))
	shorten := base64.URLEncoding.EncodeToString(hasher.Sum(nil))[0:5]

	link := &model.Link{
		ID: fmt.Sprintf("T%d", rand.Int()),
		URL: input.URL,
		Hash: shorten,
	}

	r.links = append(r.links, *link)
	return link, nil
	}

func (r *queryResolver) Link(ctx context.Context, url string) (*model.Link, error) {
		for _, link := range r.links {
		if link.URL == url {
			return &link, nil
		}
	}
	return nil, nil
	}



// Mutation returns generated.MutationResolver implementation.
	func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }
// Query returns generated.QueryResolver implementation.
	func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }


type mutationResolver struct { *Resolver }
type queryResolver struct { *Resolver }



