package tvdb

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/go-openapi/runtime"
	runtimeclient "github.com/go-openapi/runtime/client"
	"github.com/rbtr/go-tvdb/generated/client"
	"github.com/rbtr/go-tvdb/generated/client/authentication"
	"github.com/rbtr/go-tvdb/generated/client/episodes"
	"github.com/rbtr/go-tvdb/generated/client/languages"
	"github.com/rbtr/go-tvdb/generated/client/movies"
	"github.com/rbtr/go-tvdb/generated/client/search"
	"github.com/rbtr/go-tvdb/generated/client/series"
	"github.com/rbtr/go-tvdb/generated/client/updates"
	"github.com/rbtr/go-tvdb/generated/models"
)

const (
	// tvdbTokenLifespanHours is the full valid lifespan of the JWT
	tvdbTokenLifespanHours = 24 * time.Hour
	// tvdbTokenRefreshAge is the age of the JWT at which we will attempt to refresh
	tvdbTokenRefreshThreshold = 1 * time.Hour
)

// authorization holds the bearer token and expiry data
type authorization struct {
	bearerToken string
	expiration  time.Time
}

// expired returns if the expiration of the token has elapsed
func (authorization *authorization) expired() bool {
	return time.Now().After(authorization.expiration)
}

// shouldRefresh returns if the token is in its refreshable threshold
func (authorization *authorization) shouldRefresh() bool {
	return !authorization.expired() && time.Now().After(authorization.expiration.Add(-tvdbTokenRefreshThreshold))
}

type Client struct {
	Authn *models.Auth
	authz authorization
	lang  *string
	c     *client.TheTVDBAPIV3
}

func (c *Client) SetLang(lang string) *Client {
	c.lang = &lang
	return c
}

// authenticator(ctx) provides an auth decorator func for api calls
func (c *Client) authenticator(ctx context.Context) (runtime.ClientAuthInfoWriter, error) {
	if c.authz.shouldRefresh() {
		resp, err := c.c.Authentication.GetRefreshToken(&authentication.GetRefreshTokenParams{Context: ctx}, runtimeclient.BearerToken(c.authz.bearerToken))
		if err != nil {
			return nil, err
		}

		if resp != nil && resp.Payload != nil {
			c.authz = authorization{
				bearerToken: resp.Payload.Token,
				expiration:  time.Now().Add(tvdbTokenLifespanHours),
			}
		}
	}

	if c.authz.expired() {
		if c.Authn == nil {
			return nil, errors.New("no authentication provided")
		}

		resp, err := c.c.Authentication.PostLogin(&authentication.PostLoginParams{
			AuthenticationString: c.Authn,
			Context:              ctx,
		})
		if err != nil {
			return nil, err
		}

		if resp != nil && resp.Payload != nil {
			c.authz = authorization{
				bearerToken: resp.Payload.Token,
				expiration:  time.Now().Add(tvdbTokenLifespanHours),
			}
		}
	}

	return runtimeclient.BearerToken(c.authz.bearerToken), nil
}

func DefaultClient(auth *models.Auth) *Client {
	return &Client{
		Authn: auth,
		c:     client.Default,
	}
}

type Episodes interface {
	GetEpisodeByID(context.Context, int64) (*models.Episode, *models.JSONErrors, error)
}

var _ Episodes = &Client{}

func (c *Client) GetEpisodeByID(ctx context.Context, id int64) (*models.Episode, *models.JSONErrors, error) {
	auth, err := c.authenticator(ctx)
	if err != nil {
		return nil, nil, err
	}

	params := episodes.NewGetEpisodesIDParams().WithAcceptLanguage(c.lang).WithContext(ctx)
	params.SetID(id)
	res, err := c.c.Episodes.GetEpisodesID(params, auth)

	if res == nil {
		return nil, nil, err
	}

	if res.Payload == nil {
		return nil, nil, err
	}

	return res.Payload.Data, res.Payload.Errors, err
}

type Languages interface {
	GetLanguages(context.Context) ([]*models.Language, error)
	GetLanguageByID(context.Context, string) (*models.Language, error)
}

var _ Languages = &Client{}

func (c *Client) GetLanguages(ctx context.Context) ([]*models.Language, error) {
	auth, err := c.authenticator(ctx)
	if err != nil {
		return nil, err
	}

	params := languages.NewGetLanguagesParams().WithContext(ctx)
	res, err := c.c.Languages.GetLanguages(params, auth)

	if res == nil {
		return nil, err
	}

	if res.Payload == nil {
		return nil, err
	}

	return res.Payload.Data, err
}

func (c *Client) GetLanguageByID(ctx context.Context, id string) (*models.Language, error) {
	auth, err := c.authenticator(ctx)
	if err != nil {
		return nil, err
	}

	params := languages.NewGetLanguagesIDParams().WithContext(ctx)
	params.SetID(id)
	res, err := c.c.Languages.GetLanguagesID(params, auth)

	if res == nil {
		return nil, err
	}

	if res.Payload == nil {
		return nil, err
	}

	return res.Payload, err
}

type Movies interface {
	GetMovieByID(context.Context, int64) (*models.Movie, error)
	GetMovieUpdates(context.Context, time.Time) ([]int64, error)
}

var _ Movies = &Client{}

func (c *Client) GetMovieByID(ctx context.Context, id int64) (*models.Movie, error) {
	auth, err := c.authenticator(ctx)
	if err != nil {
		return nil, err
	}

	params := movies.NewGetMoviesIDParams().WithAcceptLanguage(c.lang).WithContext(ctx)
	params.SetID(id)
	res, err := c.c.Movies.GetMoviesID(params, auth)

	if res == nil {
		return nil, err
	}

	if res.Payload == nil {
		return nil, err
	}

	return res.Payload, err
}

func (c *Client) GetMovieUpdates(ctx context.Context, since time.Time) ([]int64, error) {
	auth, err := c.authenticator(ctx)
	if err != nil {
		return nil, err
	}

	params := movies.NewGetMovieupdatesParams().WithContext(ctx)
	params.SetSince(strconv.FormatInt(since.Unix(), 10))
	res, err := c.c.Movies.GetMovieupdates(params, auth)

	if res == nil {
		return nil, err
	}

	if res.Payload == nil {
		return nil, err
	}

	return res.Payload.Movies, err
}

type Search interface {
	SearchSeries(context.Context, map[string]string) ([]*models.SeriesSearchResult, error)
	GetSearchParams(context.Context) ([]string, error)
}

var _ Search = &Client{}

func (c *Client) SearchSeries(ctx context.Context, searchParams map[string]string) ([]*models.SeriesSearchResult, error) {
	auth, err := c.authenticator(ctx)
	if err != nil {
		return nil, err
	}

	params := search.NewGetSearchSeriesParams().WithAcceptLanguage(c.lang).WithContext(ctx)
	params.SetImdbID(Str(searchParams["imdbId"]))
	params.SetName(Str(searchParams["name"]))
	params.SetSlug(Str(searchParams["slug"]))
	params.SetZap2itID(Str(searchParams["zap2itId"]))
	res, err := c.c.Search.GetSearchSeries(params, auth)

	if res == nil {
		return nil, err
	}

	if res.Payload == nil {
		return nil, err
	}

	return res.Payload.Data, nil
}

func (c *Client) GetSearchParams(ctx context.Context) ([]string, error) {
	auth, err := c.authenticator(ctx)
	if err != nil {
		return nil, err
	}

	params := search.NewGetSearchSeriesParamsParams().WithContext(ctx)
	res, err := c.c.Search.GetSearchSeriesParams(params, auth)

	if res == nil {
		return nil, err
	}

	if res.Payload == nil {
		return nil, err
	}

	return res.Payload.Data, nil
}

type Series interface {
	GetSeriesByID(context.Context, int64) (*models.Series, *models.JSONErrors, error)
	GetSeriesEpisodes(context.Context, int64, int64) ([]*models.Episode, *models.Links, *models.JSONErrors, error)
}

var _ Series = &Client{}

func (c *Client) GetSeriesByID(ctx context.Context, id int64) (*models.Series, *models.JSONErrors, error) {
	auth, err := c.authenticator(ctx)
	if err != nil {
		return nil, nil, err
	}

	params := series.NewGetSeriesIDParams().WithAcceptLanguage(c.lang).WithContext(ctx)
	params.WithID(id)
	res, err := c.c.Series.GetSeriesID(params, auth)

	if res == nil {
		return nil, nil, err
	}

	if res.Payload == nil {
		return nil, nil, err
	}

	return res.Payload.Data, res.Payload.Errors, err
}

func (c *Client) GetSeriesEpisodes(ctx context.Context, id, page int64) ([]*models.Episode, *models.Links, *models.JSONErrors, error) {
	auth, err := c.authenticator(ctx)
	if err != nil {
		return nil, nil, nil, err
	}

	params := series.NewGetSeriesIDEpisodesParams().WithContext(ctx)
	params.SetID(id)
	params.SetPage(Str(strconv.FormatInt(page, 10)))
	res, err := c.c.Series.GetSeriesIDEpisodes(params, auth)

	if res == nil {
		return nil, nil, nil, err
	}

	if res.Payload == nil {
		return nil, nil, nil, err
	}

	return res.Payload.Data, res.Payload.Links, res.Payload.Errors, err
}

func (c *Client) GetSeriesEpisode(ctx context.Context, id, page int64, queryParams map[string]string) ([]*models.Episode, *models.Links, *models.JSONErrors, error) {
	auth, err := c.authenticator(ctx)
	if err != nil {
		return nil, nil, nil, err
	}

	params := series.NewGetSeriesIDEpisodesQueryParams().WithAcceptLanguage(c.lang).WithContext(ctx)
	params.SetID(id)
	params.SetPage(Str(strconv.FormatInt(page, 10)))
	params.SetAbsoluteNumber(Str(queryParams["absoluteNumber"]))
	params.SetAiredEpisode(Str(queryParams["airedEpisode"]))
	params.SetAiredSeason(Str(queryParams["airedSeason"]))
	params.SetDvdEpisode(Str(queryParams["dvdEpisode"]))
	params.SetDvdSeason(Str(queryParams["dvdSeason"]))
	params.SetImdbID(Str(queryParams["imdbId"]))
	res, err := c.c.Series.GetSeriesIDEpisodesQuery(params, auth)

	if res == nil {
		return nil, nil, nil, err
	}

	if res.Payload == nil {
		return nil, nil, nil, err
	}

	return res.Payload.Data, res.Payload.Links, res.Payload.Errors, err
}

type Update interface {
	GetUpdates(context.Context, time.Time, time.Time) ([]*models.Update, *models.JSONErrors, error)
}

var _ Update = &Client{}

func (c *Client) GetUpdates(ctx context.Context, after, before time.Time) ([]*models.Update, *models.JSONErrors, error) {
	auth, err := c.authenticator(ctx)
	if err != nil {
		return nil, nil, err
	}

	params := updates.NewGetUpdatedQueryParams().WithAcceptLanguage(c.lang).WithContext(ctx)
	params.SetFromTime(strconv.FormatInt(after.Unix(), 10))
	params.SetToTime(Str(strconv.FormatInt(before.Unix(), 10)))
	res, err := c.c.Updates.GetUpdatedQuery(params, auth)

	if res == nil {
		return nil, nil, err
	}

	if res.Payload == nil {
		return nil, nil, err
	}

	return res.Payload.Data, res.Payload.Errors, err
}

func Str(s string) *string {
	return &s
}
