package views

import (
	"blogGo/conf"
	"blogGo/src/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"net/http"
	"time"
)

type Authorization struct {
	username string
	password string
}

type AuthToken struct {
	accessToken  string
	refreshToken string
}

var (
	githubOauthConfig = &oauth2.Config{
		Scopes:   []string{"user:email", "repo"},
		Endpoint: github.Endpoint,
	}
)

func AuthViewRoutes(routes *gin.RouterGroup) {
	router := routes.Group("/auth")
	ginConf := conf.CFG.GinConf
	githubAuth := conf.CFG.GitHubAuth

	httpHead := "http://"
	if ginConf.SSLMode == "true" {
		httpHead = "https://"
	}

	githubOauthConfig.ClientID = githubAuth.ClientID
	githubOauthConfig.ClientSecret = githubAuth.ClientSecret
	githubOauthConfig.RedirectURL = fmt.Sprintf("%s%s/api/v1/auth/github/callback", httpHead, ginConf.GinHost)

	router.GET("/github/login", handleGithubLogin)
	router.GET("/github/callback", handleGithubCallback)
	router.GET("/user/token", getAuthToken)
}

func handleGithubLogin(c *gin.Context) {
	state := c.Query("state")
	authURL := githubOauthConfig.AuthCodeURL(state, oauth2.AccessTypeOffline)

	utils.Logging.Info(fmt.Sprintf("authUrl: %s, state: %s", authURL, state))
	c.JSON(http.StatusOK, gin.H{"redirect": authURL})
}

func handleGithubCallback(c *gin.Context) {
	state := c.Query("state")
	utils.Logging.Info(state)

	code := c.Query("code")
	if code == "" {
		utils.RC.Set(
			c.Request.Context(),
			fmt.Sprintf("github-%s", state),
			fmt.Sprintf("error-%s", "Missing authorization code"),
			time.Minute*10,
		)
		utils.Logging.Warning("Missing authorization code")
	}

	if token, err := githubOauthConfig.Exchange(c, code); err != nil {
		utils.RC.Set(
			c.Request.Context(),
			fmt.Sprintf("github-%s", state),
			fmt.Sprintf("error-%s", err.Error()),
			10*time.Minute,
		)
		utils.Logging.Warning(err.Error())
	} else {
		utils.RC.Set(
			c.Request.Context(),
			fmt.Sprintf("github-%s", state),
			token.AccessToken,
			time.Hour*24,
		)
		utils.Logging.Info(token.AccessToken)
	}
	c.Redirect(http.StatusFound, "http://localhost:3000/auth/success")
}

func getAuthToken(c *gin.Context) {
	_id := c.Query("id")
	_code := c.Query("code")

	if _id == "" {
		c.JSON(http.StatusOK, gin.H{"error": "Missing id"})
		return
	} else {
		fmt.Println(_id)
	}

	if token, err := utils.RC.Get(c.Request.Context(), fmt.Sprintf("%s-%s", _code, _id)).Result(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"access_token": token})
	}
}
