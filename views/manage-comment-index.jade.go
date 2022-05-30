// Code generated by "jade.go"; DO NOT EDIT.

package views

import (
	"bufio"
	"fmt"

	"github.com/ngocphuongnb/tetua/app/asset"
	"github.com/ngocphuongnb/tetua/app/cache"
	"github.com/ngocphuongnb/tetua/app/config"
	"github.com/ngocphuongnb/tetua/app/entities"
	"github.com/ngocphuongnb/tetua/app/utils"
)

const (
	managecommentindex__20 = `</div></div><div class="main"><div class="box">`
	managecommentindex__21 = `<h1>Comments</h1><form class="search-form" method="get" action="" accept-charset="UTF-8" style="width: 100%;overflow:initial;">`
	managecommentindex__22 = `<input class="search-input" type="text" name="q" placeholder="Search comments..." value="`
	managecommentindex__23 = `" style="width: auto;flex-grow: 1;"/><button class="search-btn" type="submit" aria-label="Search comments"><svg style="width:24px;height:24px" viewBox="0 0 24 24"><path fill="currentColor" d="M9.5,3A6.5,6.5 0 0,1 16,9.5C16,11.11 15.41,12.59 14.44,13.73L14.71,14H15.5L20.5,19L19,20.5L14,15.5V14.71L13.73,14.44C12.59,15.41 11.11,16 9.5,16A6.5,6.5 0 0,1 3,9.5A6.5,6.5 0 0,1 9.5,3M9.5,5C7,5 5,7 5,9.5C5,12 7,14 9.5,14C12,14 14,12 14,9.5C14,7 12,5 9.5,5Z"></path></svg></button></form><div class="comments">`
	managecommentindex__26 = `</ul></div></div></div></div><div class="mobile-menu"><div class="menu-head">`
	managecommentindex__31 = `<script>listenDeleteNodeEvents('comment', '/comments', '/manage/comments')</script></body></html>`
	managecommentindex__80 = `<input class="hidden" type="hidden" name="post" value="`
	managecommentindex__82 = `<input class="hidden" type="hidden" name="user" value="`
)

func ManageCommentIndex(paginate *entities.Paginate[entities.Comment], search string, userID, postID int) func(meta *entities.Meta, wr *bufio.Writer) {
	return func(meta *entities.Meta, wr *bufio.Writer) {
		buffer := &WriterAsBuffer{wr}

		buffer.WriteString(commentlist__0)

		var title = meta.GetTitle()
		var appName = config.Setting("app_name")
		var appLogo = config.Setting("app_logo")
		buffer.WriteString(commentlist__1)
		WriteAll(title, true, buffer)
		buffer.WriteString(commentlist__2)
		WriteAll(meta.Canonical, true, buffer)
		buffer.WriteString(commentlist__3)
		WriteAll(meta.Type, true, buffer)
		buffer.WriteString(commentlist__4)
		WriteAll(meta.Canonical, true, buffer)
		buffer.WriteString(commentlist__5)
		WriteAll(title, true, buffer)
		buffer.WriteString(commentlist__6)
		WriteAll(appName, true, buffer)
		buffer.WriteString(commentlist__7)
		WriteAll(config.Setting("twitter_site"), true, buffer)
		buffer.WriteString(commentlist__8)
		WriteAll(title, true, buffer)
		buffer.WriteString(commentlist__9)
		WriteAll(appName, true, buffer)
		buffer.WriteString(commentlist__10)
		WriteAll(appName, true, buffer)
		buffer.WriteString(commentlist__11)
		WriteAll(appName+" Feed", true, buffer)
		buffer.WriteString(commentlist__12)
		WriteAll(utils.Url("/feed"), true, buffer)
		buffer.WriteString(commentlist__13)
		if appLogo != "" {
			buffer.WriteString(commentlist__30)
			WriteAll(appLogo, true, buffer)
			buffer.WriteString(commentlist__31)
			WriteAll(appLogo, true, buffer)
			buffer.WriteString(commentlist__13)
		}
		if meta.Description != "" {
			buffer.WriteString(commentlist__33)
			WriteAll(meta.Description, true, buffer)
			buffer.WriteString(commentlist__34)
			WriteAll(meta.Description, true, buffer)
			buffer.WriteString(commentlist__35)
			WriteAll(meta.Description, true, buffer)
			buffer.WriteString(commentlist__13)
		}
		if meta.Image != "" {
			buffer.WriteString(commentlist__37)
			WriteAll(meta.Image, true, buffer)
			buffer.WriteString(commentlist__38)
			WriteAll(meta.Image, true, buffer)
			buffer.WriteString(commentlist__13)
		}
		WriteAll(asset.CssFile("css/light.min.css"), false, buffer)
		WriteAll(asset.CssFile("css/style.css"), false, buffer)
		WriteAll(config.Setting("inject_header"), false, buffer)
		buffer.WriteString(commentlist__14)
		WriteAll(utils.Url(""), true, buffer)
		buffer.WriteString(commentlist__15)
		var logoUrl = config.Setting("app_logo")
		if logoUrl != "" {
			buffer.WriteString(commentlist__40)
			WriteAll(logoUrl, true, buffer)
			buffer.WriteString(commentlist__41)
			WriteAll(config.Setting("app_name"), true, buffer)
			buffer.WriteString(commentlist__13)
		} else {
			buffer.WriteString(commentlist__43)

		}
		buffer.WriteString(commentlist__16)
		WriteAll(meta.Query, true, buffer)
		buffer.WriteString(commentlist__17)
		WriteAll(utils.Url("/search"), true, buffer)
		buffer.WriteString(commentlist__18)

		if meta.User == nil || meta.User.ID == 0 {
			buffer.WriteString(commentlist__44)
			WriteAll(utils.Url("/login"), true, buffer)
			buffer.WriteString(commentlist__45)
			WriteAll(utils.Url("/register"), true, buffer)
			buffer.WriteString(commentlist__46)

		} else {
			buffer.WriteString(commentlist__44)
			WriteAll(utils.Url("/posts/new"), true, buffer)
			buffer.WriteString(commentlist__48)
			WriteAll(meta.User.Url(), true, buffer)
			buffer.WriteString(commentlist__49)
			WriteAll(meta.User.Username, true, buffer)
			buffer.WriteString(commentlist__50)
			if meta.User.AvatarImageUrl != "" {
				buffer.WriteString(commentlist__57)
				WriteAll(meta.User.AvatarImageUrl, true, buffer)
				buffer.WriteString(commentlist__41)
				WriteAll(meta.User.Username, true, buffer)
				buffer.WriteString(commentlist__13)
			} else {
				buffer.WriteString(commentlist__60)

			}
			buffer.WriteString(commentlist__51)

			if meta.User != nil && meta.User.IsRoot() {
				buffer.WriteString(commentlist__44)
				WriteAll(utils.Url("/manage"), true, buffer)
				buffer.WriteString(commentlist__62)

			}
			buffer.WriteString(commentlist__44)
			WriteAll(meta.User.Url(), true, buffer)
			buffer.WriteString(commentlist__53)
			WriteAll(utils.Url("/posts"), true, buffer)
			buffer.WriteString(commentlist__54)
			WriteAll(utils.Url("/settings"), true, buffer)
			buffer.WriteString(commentlist__55)
			WriteAll(utils.Url("/logout"), true, buffer)
			buffer.WriteString(commentlist__56)

		}
		buffer.WriteString(commentlist__19)

		{
			buffer.WriteString(commentlist__119)
			WriteAll(utils.Url("/manage"), true, buffer)
			buffer.WriteString(commentlist__120)
			WriteAll(utils.Url("/manage/topics"), true, buffer)
			buffer.WriteString(commentlist__121)
			WriteAll(utils.Url("/manage/posts"), true, buffer)
			buffer.WriteString(commentlist__122)
			WriteAll(utils.Url("/manage/pages"), true, buffer)
			buffer.WriteString(commentlist__123)
			WriteAll(utils.Url("/manage/roles"), true, buffer)
			buffer.WriteString(commentlist__124)
			WriteAll(utils.Url("/manage/users"), true, buffer)
			buffer.WriteString(commentlist__125)
			WriteAll(utils.Url("/manage/comments"), true, buffer)
			buffer.WriteString(commentlist__126)
			WriteAll(utils.Url("/manage/files"), true, buffer)
			buffer.WriteString(commentlist__127)
			WriteAll(utils.Url("/manage/settings"), true, buffer)
			buffer.WriteString(commentlist__72)

		}

		buffer.WriteString(managecommentindex__20)

		{
			var (
				msgs = meta.Messages
			)

			if msgs.Length() > 0 {
				buffer.WriteString(commentlist__73)
				var messages = msgs.Get()
				for _, msg := range messages {
					buffer.WriteString(commentlist__75)
					WriteAll(msg.Type, true, buffer)
					buffer.WriteString(commentlist__50)
					WriteAll(msg.Message, true, buffer)
					buffer.WriteString(commentlist__77)
				}
				buffer.WriteString(commentlist__74)
			}
		}

		buffer.WriteString(managecommentindex__21)

		if postID > 0 {
			buffer.WriteString(managecommentindex__80)
			WriteInt(int64(postID), buffer)
			buffer.WriteString(commentlist__13)
		}
		if userID > 0 {
			buffer.WriteString(managecommentindex__82)
			WriteInt(int64(userID), buffer)
			buffer.WriteString(commentlist__13)
		}
		buffer.WriteString(managecommentindex__22)
		WriteEscString(search, buffer)
		buffer.WriteString(managecommentindex__23)

		for _, comment := range paginate.Data {
			{
				var (
					comment       = comment
					postID        = comment.Post.ID
					editCondition = true
					extraInfo     = true
				)

				if extraInfo {
					buffer.WriteString(commentlist__86)
					WriteAll(comment.Post.Url(), true, buffer)
					buffer.WriteString(commentlist__87)
					WriteAll(comment.Post.Name, true, buffer)
					buffer.WriteString(commentlist__88)

				}
				buffer.WriteString(commentlist__78)
				WriteEscString(fmt.Sprintf("comment-%d", comment.ID), buffer)
				buffer.WriteString(commentlist__50)
				WriteAll(comment.User.AvatarElm("30", "30", false), false, buffer)
				buffer.WriteString(commentlist__64)
				WriteAll(comment.User.Url(), true, buffer)
				buffer.WriteString(commentlist__50)
				WriteAll(comment.User.Name(), true, buffer)
				buffer.WriteString(commentlist__82)
				WriteAll(comment.CreatedAt.Format("January 2, 2006 15:04 MST"), true, buffer)
				buffer.WriteString(commentlist__83)
				WriteAll(comment.ContentHTML, false, buffer)
				if editCondition {
					buffer.WriteString(commentlist__89)
					WriteAll(comment.ID, true, buffer)
					buffer.WriteString(commentlist__90)
					WriteAll(comment.ID, true, buffer)
					buffer.WriteString(commentlist__91)

					if extraInfo {
						var commentUrl = fmt.Sprintf("%s#comment-%d", comment.Post.Url(), comment.ID)
						var postCommentsUrl = fmt.Sprintf("/manage/comments?post=%d", postID)
						var userCommentsUrl = fmt.Sprintf("/manage/comments?user=%d", postID)
						buffer.WriteString(commentlist__93)
						WriteEscString(commentUrl, buffer)
						buffer.WriteString(commentlist__94)

						if meta.User.IsRoot() {
							buffer.WriteString(commentlist__95)
							WriteEscString(postCommentsUrl, buffer)
							buffer.WriteString(commentlist__96)
							WriteEscString(userCommentsUrl, buffer)
							buffer.WriteString(commentlist__97)

						}
					}
					buffer.WriteString(commentlist__22)
				}
				buffer.WriteString(commentlist__22)
				if editCondition {
					buffer.WriteString(commentlist__98)
					WriteEscString(fmt.Sprintf("/comments/%d", comment.ID), buffer)
					buffer.WriteString(commentlist__99)
					WriteAll(postID, true, buffer)
					buffer.WriteString(commentlist__100)
					WriteAll(comment.Content, true, buffer)
					buffer.WriteString(commentlist__101)

				}
				buffer.WriteString(commentlist__85)
			}

		}
		buffer.WriteString(commentlist__22)
		var links = paginate.Links()
		buffer.WriteString(commentlist__23)
		for _, link := range links {
			buffer.WriteString(commentlist__44)
			WriteAll(link.Link, true, buffer)
			buffer.WriteString(commentlist__103)
			WriteAll(link.Class, true, buffer)
			buffer.WriteString(commentlist__50)
			WriteAll(link.Label, true, buffer)
			buffer.WriteString(commentlist__105)

		}
		buffer.WriteString(managecommentindex__26)
		WriteAll(config.Setting("app_name"), true, buffer)
		buffer.WriteString(commentlist__25)

		if meta.User == nil || meta.User.ID == 0 {
			buffer.WriteString(commentlist__106)
			WriteAll(utils.Url("/login"), true, buffer)
			buffer.WriteString(commentlist__107)
			WriteAll(utils.Url("/register"), true, buffer)
			buffer.WriteString(commentlist__108)

		} else {
			{
				buffer.WriteString(commentlist__63)
				WriteAll(meta.User.AvatarElm("32", "32", false), false, buffer)
				buffer.WriteString(commentlist__64)
				WriteAll(meta.User.Url(), true, buffer)
				buffer.WriteString(commentlist__50)
				WriteAll(meta.User.Name(), true, buffer)
				buffer.WriteString(commentlist__66)
				WriteAll("@"+meta.User.Username, true, buffer)
				buffer.WriteString(commentlist__67)
				WriteAll(utils.Url("/posts/new"), true, buffer)
				buffer.WriteString(commentlist__68)
				WriteAll(utils.Url("/posts"), true, buffer)
				buffer.WriteString(commentlist__69)
				WriteAll(utils.Url("/comments"), true, buffer)
				buffer.WriteString(commentlist__70)
				WriteAll(utils.Url("/files"), true, buffer)
				buffer.WriteString(commentlist__71)
				WriteAll(utils.Url("/settings"), true, buffer)
				buffer.WriteString(commentlist__72)

			}

			if meta.User.IsRoot() {
				{
					buffer.WriteString(commentlist__119)
					WriteAll(utils.Url("/manage"), true, buffer)
					buffer.WriteString(commentlist__120)
					WriteAll(utils.Url("/manage/topics"), true, buffer)
					buffer.WriteString(commentlist__121)
					WriteAll(utils.Url("/manage/posts"), true, buffer)
					buffer.WriteString(commentlist__122)
					WriteAll(utils.Url("/manage/pages"), true, buffer)
					buffer.WriteString(commentlist__123)
					WriteAll(utils.Url("/manage/roles"), true, buffer)
					buffer.WriteString(commentlist__124)
					WriteAll(utils.Url("/manage/users"), true, buffer)
					buffer.WriteString(commentlist__125)
					WriteAll(utils.Url("/manage/comments"), true, buffer)
					buffer.WriteString(commentlist__126)
					WriteAll(utils.Url("/manage/files"), true, buffer)
					buffer.WriteString(commentlist__127)
					WriteAll(utils.Url("/manage/settings"), true, buffer)
					buffer.WriteString(commentlist__72)

				}

			}
		}
		buffer.WriteString(commentlist__26)

		for _, topic := range cache.Topics {
			buffer.WriteString(commentlist__106)
			WriteAll(topic.Url(), true, buffer)
			buffer.WriteString(commentlist__49)
			WriteAll(topic.Name, true, buffer)
			buffer.WriteString(commentlist__50)
			WriteAll("#"+topic.Name, true, buffer)
			buffer.WriteString(commentlist__132)
		}
		buffer.WriteString(commentlist__27)
		WriteAll(config.Setting("footer_content"), false, buffer)
		buffer.WriteString(commentlist__28)
		WriteAll(config.Setting("inject_footer"), false, buffer)
		WriteAll(asset.JsFile("js/layout.js"), false, buffer)
		WriteAll(asset.JsFile("js/main.js"), false, buffer)
		buffer.WriteString(managecommentindex__31)

	}
}