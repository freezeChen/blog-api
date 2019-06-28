/*
   @Time : 2019-06-28 13:49:25
   @Author :
   @File : server
   @Software: blog
*/
package http

import (
	"blog/model"
	"blog/service"
	"github.com/gin-gonic/gin"

	"time"
)

var (
	svc *service.Service
)

func New(s *service.Service) (engine *gin.Engine) {
	engine = gin.Default()
	svc = s
	initRouter(engine)
	return
}

func initRouter(e *gin.Engine) {
	g := e.Group("/")
	{
		g.GET("/getHomeList", getHomeList)
	}
}

func getHomeList(ctx *gin.Context) {
	var list = make([]*model.Article, 0)

	list = append(list, &model.Article{
		Id:       1,
		Title:    "文章1",
		Detail:   detail,
		Language: "golang",
		Author:   "freeze",
		Time:     time.Now().Format("2006年01月02日"),
	})

	list = append(list, &model.Article{
		Id:       2,
		Title:    detail,
		Detail:   "java 内容",
		Language: "Java",
		Author:   "freeze",
		Time:     time.Now().Format("2006年01月02日"),
	})

	ctx.JSON(200, list)

}

var detail = `  <div class="article-entry" itemprop="articleBody">


                <a href="https://colobu.com/images/logos/golang7.png" title="" class="fancybox" rel="article6"><img class="article-post-image" src="/images/logos/golang7.png"></a>


                <p>原文: <a href="https://www.redisgreen.net/blog/reading-and-writing-redis-protocol/" target="_blank" rel="external">Reading and Writing Redis Protocol in Go</a><br>翻译整理: <a href="https://colobu.com/" target="_blank" rel="external">smallnest</a>, 译文连接: <a href="https://colobu.com/2019/04/16/Reading-and-Writing-Redis-Protocol-in-Go/" target="_blank" rel="external">使用 Go 语言读写Redis协议</a>。 转载请保留原文出处和译文译者和出处。</p>
                <p>这篇文章使用两个简单的Reader和Writer实现了redis客户端的读写协议，通过这两个实现可以容易地学习Redis协议是如何工作的。</p>
                <p>如果你想寻找一个全功能的、产品级的Redis client, 推荐你看看 Gary Burd的 <a href="https://github.com/gomodule/redigo" target="_blank" rel="external">redigo</a></p>
                <p>开始之前，建议你先阅读一下 Redis协议的介绍。</p>
                <p>官方的协议可以在其网站上找到: <a href="https://redis.io/topics/protocol" target="_blank" rel="external">protocol</a>。 Redis的协议叫做 <strong>RESP</strong> (<strong>RE</strong>dis <strong>S</strong>erialization <strong>P</strong>rotocol)，客户端和服务器端通过基于文本的协议进行通讯。</p>
                <p>所有的服务器和客户端之间的通讯都使用以下5中基本类型：</p>
                <ul>
                    <li><strong>简单字符串</strong>: 服务器用来返回简单的结果，比如"OK"或者"PONG"</li>
                    <li><strong>bulk string</strong>: 大部分单值命令的返回结果，比如 GET, LPOP, and HGET</li>
                    <li><strong>整数</strong>: 查询长度的命令的返回结果</li>
                    <li><strong>数组</strong>: 可以包含其它RESP对象，设置数组，用来发送命令给服务器，也用来返回多个值的命令</li>
                    <li><strong>Error</strong>: 服务器返回错误信息</li>
                </ul>
                <p>RESP的第一个字节表示数据的类型：</p>
                <ul>
                    <li><strong>简单字符串</strong>: 第一个字节是 "+", 比如 "+OK\r\n"</li>
                    <li><strong>bulk string</strong>: 第一个字节是 "\$", 比如 "$6\r\nfoobar\r\n"</li>
                    <li><strong>整数</strong>: 第一个字节是 ":"， 比如 ":1000\r\n"</li>
                    <li><strong>数组</strong>: 第一个字节是 "<em>", 比如 "</em>2\r\n\$3\r\nfoo\r\n\$3\r\nbar\r\n"</li>
                    <li><strong>Error</strong>: 第一个字节是 "-"， 比如 "-Error message\r\n"</li>
                </ul>
                <p>基本了解Redis的协议之后，我们就可以实现它的读写器了。</p>


                <p class="article-more-link">
                    <a href="/2019/04/16/Reading-and-Writing-Redis-Protocol-in-Go/#more">阅读全文</a>
                </p>


            </div>`
