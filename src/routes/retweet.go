package routes

import (
	"fmt"

	"github.com/AlfaSakan/twitter-clone-api/src/handlers"
	"github.com/gin-gonic/gin"
)

const RETWEET_ROUTE = "/retweet"

func Retweet(r *gin.RouterGroup, h *handlers.RetweetHandler) {
	// get all reply by tweet id
	r.GET(fmt.Sprintf("%s/:tweetId", RETWEET_ROUTE), h.GetRetweetsByIdHandler)

	// create reply
	r.POST(RETWEET_ROUTE, h.PostRetweetHandler)

	// get one tweet by id
	// r.GET(fmt.Sprintf("%s/:id", TWEET_ROUTE), h.GetTweetByIdHandler)

	// delete tweet
	// r.DELETE(TWEET_ROUTE, h.DeleteTweetHandler)

	// likes tweet
	// r.PATCH(fmt.Sprintf("%s/like", TWEET_ROUTE), middlewares.RequireUser(), h.LikeTweetHandler)

	// get tweet like
	// r.POST(fmt.Sprintf("%s/like", TWEET_ROUTE), h.GetLikeTweetHandler)
}
