package knowledge

import (
	"backend/model"
	"backend/util"
	"backend/view/adapter"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sundy-li/html2article"
)

const (
	MpToken = "mp_token"
)

func Register(router *gin.RouterGroup) {
	router.GET("knowledge/daily", Daily)
	router.GET("knowledge/list", Lists)
	router.POST("knowledge/proxy", Proxy)
}

// 每日知识点
func Daily(c *gin.Context) {
	day := c.DefaultQuery("day", "")
	if day == "" {
		util.AbortError(c, 500, errors.New("day不能为空"))
		return
	}
	daily, err := model.GetDaily(day)
	if err != nil {
		util.AbortUnknownError(c, err)
		return
	}
	if daily == nil {
		util.AbortError(c, 500, errors.New("没有当日的信息"))
		return
	}
	cont, err := daily.ParseContent()
	if err != nil {
		util.AbortError(c, 500, errors.New("解析当日信息错误"))
		return
	}

	ret := adapter.ToDaily(daily, cont)

	util.ResponseSuccess(c, ret)
}

// 获取某天之前的每日信息
func Lists(c *gin.Context) {
	day := c.DefaultQuery("day", "")
	count := c.DefaultQuery("count", "5")
	token := c.DefaultQuery("token", "")
	if token != MpToken {
		util.AbortError(c, 500, errors.New("token 错误"))
		return
	}
	if day == "" {
		// 使用明天作为
		tomorow := time.Now().AddDate(0, 0, 1)
		day = tomorow.Format("2006-01-02")
	}
	cn, _ := strconv.Atoi(count)
	dailys, err := model.GetDailyList(day, cn)
	if err != nil {
		util.AbortUnknownError(c, err)
		return
	}
	util.ResponseSuccess(c, dailys)
}

type ProxyRequest struct {
	URL string `form:"url"`
}

// 获取某个地址的HTML内容
func Proxy(c *gin.Context) {
	token := c.DefaultQuery("token", "")
	url := c.PostForm("url")
	if token != MpToken {
		util.AbortError(c, 500, errors.New("token不对"))
		return
	}

	if url == "" {
		util.AbortError(c, 500, errors.New("url 为空"))
		return
	}
	fmt.Println(url)
	ext, err := html2article.NewFromUrl(url)
	if err != nil {
		util.AbortError(c, 500, errors.New("url返回错误"+err.Error()))
		return
	}
	article, err := ext.ToArticle()
	if err != nil {
		util.AbortError(c, 500, errors.New(err.Error()))
		return
	}
	article.Readable(url)

	type Article struct {
		Title       string
		PublishTime int64
		ReadContent string
	}

	art := &Article{
		Title:       article.Title,
		PublishTime: article.Publishtime,
		ReadContent: article.ReadContent,
	}

	util.ResponseSuccess(c, art)
}
