package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"test/gintest/db"
	"test/gintest/utils"
	"fmt"
)

// AuthController handle auth request
type AuthController struct {
	UserDao db.UserDao
}

func (AuthController) ToLogin(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("user_id")
	r.HTML(c.Writer, http.StatusOK, "login.html", gin.H{})
}

// Login handle login
func (ac AuthController) Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" || password == "" {
		r.HTML(c.Writer, http.StatusUnauthorized, "login.html", gin.H{
			"tips": "用户名密码不能为空",
		})
		return
	}

	user, err := ac.UserDao.GetUser(username)
	if err != nil {
		r.HTML(c.Writer, http.StatusUnauthorized, "login.html", gin.H{
			"tips": err.Error(),
		})
		return
	}
	passwd, err := utils.Encrypt(password, user.Salt)
	if err != nil {
		r.HTML(c.Writer, http.StatusUnauthorized, "login.html", gin.H{
			"tips": err.Error(),
		})
		return
	}
	if user.Password == passwd {
		session := sessions.Default(c)
		session.Set("user_id", user.ID)
		session.Save()
		fmt.Println("--------------userid-------------",user.ID)

		c.Redirect(http.StatusMovedPermanently, "/")
	} else {
		r.HTML(c.Writer, http.StatusUnauthorized, "login.html", gin.H{
			"tips": "密码错误",
		})
	}
}

// Logout is log out system
func (AuthController) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.Redirect(http.StatusMovedPermanently, "/login")
}
