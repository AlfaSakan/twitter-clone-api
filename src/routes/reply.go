package routes

import (
	"fmt"

	"github.com/AlfaSakan/twitter-clone-api/src/handlers"
	"github.com/AlfaSakan/twitter-clone-api/src/middlewares"
	"github.com/gin-gonic/gin"
)

const REPLY_ROUTE = "/reply"

func Reply(r *gin.RouterGroup, h *handlers.ReplyHandler) {
	// get all reply by tweet id
	r.GET(fmt.Sprintf("%s/:tweetId", REPLY_ROUTE), h.GetTweetRepliesByIdHandler)

	// create reply
	r.POST(REPLY_ROUTE, middlewares.RequireUser(), h.PostReplyHandler)

	// get one tweet by id
	// r.GET(fmt.Sprintf("%s/:id", TWEET_ROUTE), h.GetTweetByIdHandler)

	// delete tweet
	// r.DELETE(TWEET_ROUTE, h.DeleteTweetHandler)

	// likes tweet
	// r.PATCH(fmt.Sprintf("%s/like", TWEET_ROUTE), middlewares.RequireUser(), h.LikeTweetHandler)

	// get tweet like
	// r.POST(fmt.Sprintf("%s/like", TWEET_ROUTE), h.GetLikeTweetHandler)
}
