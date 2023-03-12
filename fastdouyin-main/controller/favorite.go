package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ikuraoo/fastdouyin/service"
	"github.com/ikuraoo/fastdouyin/util"
)

func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	vid := c.Query("video_id")
	// uid, _ := c.Get("uid")
	action_type := c.Query("action_type")

	// userId := uid.(int64)
	videoId, err := strconv.ParseInt(vid, 10, 64)
	if len(token) == 0 {
		c.JSON(http.StatusOK, UserResponse{Response: Response{StatusCode: 1, StatusMsg: "Please login first"}})
		return
	} else {
		if user, exist := usersLoginInfo[token]; exist {
			userId := user.Id
			//判断action_type
			if action_type == "1" {
				//判断video_id转换
				if err != nil {
					c.JSON(http.StatusOK, UserResponse{
						Response: Response{StatusCode: 1, StatusMsg: "id convert error"},
					})
					return
				}
				//执行点赞操作
				err = service.FavouriteAction(userId, videoId)
				//判断点赞操作是否有误
				if err != nil {
					c.JSON(http.StatusOK, UserResponse{
						Response: Response{StatusCode: 1, StatusMsg: "favourite change error"},
					})
				} else {
					isFavourite, _ := util.NewFavouriteDaoInstance().QueryByVIdAndUId(videoId, userId)
					if isFavourite {
						c.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: "点赞成功"})
					} else {
						c.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: "取消点赞成功"})

					}

				}
			}
		} else {
			c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		}

	}
}

type VideoDemoListResponse struct {
	Response
	// VideoList     []Video `json:"video_list"`
	VideoDemoList []*service.VideoProto
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	uid := c.Query("user_id")
	userId, _ := strconv.ParseInt(uid, 10, 64)
	DemoVideos, err := service.FavoriteList(userId)

	if err != nil {
		c.JSON(http.StatusOK, VideoDemoListResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  "查询失败",
			},
		})
	}
	c.JSON(http.StatusOK, VideoDemoListResponse{
		Response: Response{
			StatusCode: 0,
			StatusMsg:  "喜欢列表",
		},
		VideoDemoList: DemoVideos,
	})
}
