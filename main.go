package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/theykk/gitlab-adapter/common"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{}))
	//e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
	//	fmt.Println(string(reqBody))
	//}))

	e.Logger.SetLevel(log.DEBUG)
	e.Any("*", func(c echo.Context) error {
		return nil
	})
	e.POST("/api/v4/runners", func(c echo.Context) error {

		return c.JSON(http.StatusCreated, common.RegisterRunnerResponse{
			Token: "sabeybi",
		})
	})
	al := true
	e.POST("/api/v4/jobs/request", func(c echo.Context) error {
		if !al {
			return c.JSON(http.StatusNoContent, "")
		}
		al = false
		return c.JSON(http.StatusCreated, common.JobResponse{
			ID:            2,
			Token:         "jobtoken",
			AllowGitFetch: false,
			JobInfo: common.JobInfo{
				Name:        "test",
				Stage:       "build",
				ProjectID:   6,
				ProjectName: "gitlab-ci-test",
			},
			GitInfo: common.GitInfo{
				RepoURL:   "https://github.com/TheYkk/docker-nginx-static",
				Ref:       "refs/heads/v3",
				Sha:       "160213f6706711a9eb70a76bd22740489a0a4021",
				BeforeSha: "0000000000000000000000000000000000000000",
				RefType:   "",
				Refspecs:  nil,
				Depth:     0,
			},
			RunnerInfo: common.RunnerInfo{},
			Variables: common.JobVariables{
				common.JobVariable{
					Key:      "SA",
					Value:    "AS",
					Public:   false,
					Internal: false,
					File:     false,
					Masked:   false,
					Raw:      false,
				},
			},
			Steps: common.Steps{
				common.Step{
					Name: "list",
					Script: common.StepScript{
						`T='gYw'   # The test text

echo -e "\n                 40m     41m     42m     43m\
     44m     45m     46m     47m";

for FGs in '    m' '   1m' '  30m' '1;30m' '  31m' '1;31m' '  32m' \
           '1;32m' '  33m' '1;33m' '  34m' '1;34m' '  35m' '1;35m' \
           '  36m' '1;36m' '  37m' '1;37m';
  do FG=${FGs// /}
  echo -en " $FGs \033[$FG  $T  "
  for BG in 40m 41m 42m 43m 44m 45m 46m 47m;
    do echo -en "$EINS \033[$FG\033[$BG  $T  \033[0m";
  done
  echo;
done
echo`,
					},
					Timeout:      0,
					When:         "",
					AllowFailure: false,
				},
			},
			Image: common.Image{
				Name: "alpine",
			},
		})
	})
	e.PATCH("/api/v4/jobs/2/trace", func(c echo.Context) error {
		// At every log patch we should record date

		c.Logger().Info(c.Request().Header.Get("Content-Range"))
		body, _ := ioutil.ReadAll(c.Request().Body)
		// write the whole body at once
		f, _ := os.OpenFile("text.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		defer f.Close()
		_, _ = f.WriteString(string(body))

		c.Logger().Info(string(body))
		c.Response().Header().Set("Range", c.Request().Header.Get("Content-Range"))
		return c.HTML(http.StatusAccepted, "OK")
	})
	e.Logger.Fatal(e.Start(":9090"))
}
