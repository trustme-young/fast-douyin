package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ikuraoo/fastdouyin/service"
)

type CommentListResponse struct {
	Response
	CommentList []*service.CommentProto `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	Response
	Comment *service.CommentProto `json:"comment,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	token := c.Query("token")
	actionType := c.Query("action_type")
	vid := c.Query("video_id")
	videoId, _ := strconv.ParseInt(vid, 10, 64)
	if user, exist := usersLoginInfo[token]; exist {
		if actionType == "1" {
			text := c.Query("comment_text")
			commentproto, err := service.CommentAction(user.Id, videoId, text)
			if err != nil {
				c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "评论失败，请重试"})
			} else {
				c.JSON(http.StatusOK, CommentActionResponse{Response: Response{StatusCode: 0, StatusMsg: "评论成功"},
					Comment: commentproto})
			}

			return
		} else if actionType == "2" {
			c.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: "删除成功"})
			return
		}

	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	vid := c.Query("video_id")
	VideoId, err := strconv.ParseInt(vid, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "parse int error"})
	}

	comments, err := service.CommentList(VideoId)
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "get commentList error"})
	}

	c.JSON(http.StatusOK, CommentListResponse{
		Response:    Response{StatusCode: 0, StatusMsg: "Video loads successfully"},
		CommentList: comments,
	})
}
