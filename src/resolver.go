//go:generate go run scripts/gqlgen.go -v

package keywords_server

import (
	"context"

	"github.com/tgfjt/keywords-server/src/models"
)

type Resolver struct {
	tags     []models.Tag
	keywords []models.Keyword
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct {
	*Resolver
}

func (r *mutationResolver) CreateKeyword(ctx context.Context, input NewKeyword) (models.Keyword, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateTag(ctx context.Context, input NewTag) (models.Tag, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Keywords(ctx context.Context) ([]models.Keyword, error) {
	return r.keywords, nil
}
func (r *queryResolver) Tags(ctx context.Context) ([]models.Tag, error) {
	return r.tags, nil
}
