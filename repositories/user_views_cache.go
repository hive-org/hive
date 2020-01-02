package repositories

import (
	"auth/enums"
	"auth/inout"
	"auth/models"
	"auth/modelsFunctools"
	"context"
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/go-redis/redis/v7"
	"github.com/golang/protobuf/proto"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"time"
)

func getUserKey(id models.UserID) string {
	return fmt.Sprintf("%s:%d", enums.UserView, id)
}

func GetUserViewFromCache(cache *redis.Client, ctx context.Context, id models.UserID) *inout.GetUserViewResponseV1 {

	span, ctx := opentracing.StartSpanFromContext(ctx, "Get user view from cache")

	key := getUserKey(id)

	value, err := cache.Get(key).Bytes()
	if err != nil {
		span.LogFields(log.Error(err))
		sentry.CaptureException(err)
		return nil
	}

	var userView inout.GetUserViewResponseV1

	err = proto.Unmarshal(value, &userView)

	if err != nil {
		span.LogFields(log.Error(err))
		sentry.CaptureException(err)
		return nil
	}

	span.Finish()

	return &userView
}

func CacheUserView(cache *redis.Client, ctx context.Context, userViews []*inout.GetUserViewResponseV1) {

	span, ctx := opentracing.StartSpanFromContext(ctx, "Cache user views")

	if userViews == nil {
		return
	}

	identifiers := make([]models.UserID, len(userViews))

	pipeline := cache.TxPipeline()
	for i, uv := range userViews {
		data, err := proto.Marshal(uv)
		if err != nil {
			span.LogFields(log.Error(err))
			sentry.CaptureException(err)
			continue
		}

		userID := models.UserID(uv.Id)

		identifiers[i] = userID
		pipeline.Set(getUserKey(userID), data, time.Hour*48)
	}

	_, err := pipeline.Exec()

	if err != nil {
		span.LogFields(log.Error(err))
		sentry.CaptureException(err)
	}

	span.LogFields(log.String("user_id", modelsFunctools.UserIDListToString(identifiers, ", ")))
	span.Finish()
}
