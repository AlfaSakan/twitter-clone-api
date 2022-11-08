package routes

import (
	"fmt"

	"github.com/AlfaSakan/twitter-clone-api/src/handlers"
	"github.com/AlfaSakan/twitter-clone-api/src/middlewares"
	"github.com/gin-gonic/gin"
)

const TWEET_ROUTE = "/tweet"

// get all tweet

func Tweet(r *gin.RouterGroup, h *handlers.TweetHandler) {
	// get all tweet by user id
	r.GET(fmt.Sprintf("%ss", TWEET_ROUTE), h.GetAllTweetsHandler)

	// get all tweet by user id
	r.POST(fmt.Sprintf("%ss", TWEET_ROUTE), h.FindAllTweetsHandler)

	// get one tweet by id
	r.GET(fmt.Sprintf("%s/:id", TWEET_ROUTE), h.GetTweetByIdHandler)

	// create tweet
	r.POST(TWEET_ROUTE, middlewares.RequireUser(), h.PostTweetHandler)

	// delete tweet
	r.DELETE(TWEET_ROUTE, h.DeleteTweetHandler)

	// update likes tweet
	r.PATCH(fmt.Sprintf("%s/like", TWEET_ROUTE), middlewares.RequireUser(), h.LikeTweetHandler)

	// get tweet like
	r.GET(fmt.Sprintf("%s/like", TWEET_ROUTE), middlewares.RequireUser(), h.GetLikeTweetHandler)
}
