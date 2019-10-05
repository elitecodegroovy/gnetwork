package server

import (
	"fmt"
	"github.com/elitecodegroovy/util"
	"gopkg.in/macaron.v1"
	"io"
	"os"
	"regexp"
	"strings"

	//"github.com/elitecodegroovy/gnetwork/pkg/bus"
	m "github.com/elitecodegroovy/gnetwork/pkg/models"
	"github.com/elitecodegroovy/gnetwork/pkg/setting"
)

func AdminGetSettings(c *m.ReqContext) {
	settings := make(map[string]interface{})

	for _, section := range setting.Raw.Sections() {
		jsonSec := make(map[string]interface{})
		settings[section.Name()] = jsonSec

		for _, key := range section.Keys() {
			keyName := key.Name()
			value := key.Value()
			if strings.Contains(keyName, "secret") ||
				strings.Contains(keyName, "password") ||
				(strings.Contains(keyName, "provider_config")) {
				value = "************"
			}
			if strings.Contains(keyName, "url") {
				var rgx = regexp.MustCompile(`.*:\/\/([^:]*):([^@]*)@.*?$`)
				var subs = rgx.FindAllSubmatch([]byte(value), -1)
				if subs != nil && len(subs[0]) == 3 {
					value = strings.Replace(value, string(subs[0][1]), "******", 1)
					value = strings.Replace(value, string(subs[0][2]), "******", 1)
				}
			}

			jsonSec[keyName] = value
		}
	}

	c.JSON(200, settings)
}

func urlHandler(c *macaron.Context) {
	settings := make(map[string]interface{})
	settings["msg"] = "the request path is: " + c.Req.RequestURI
	settings["code"] = 200
	c.JSON(200, settings)
}

func uploadFile(c *m.ReqContext) {
	r := c.Req
	w := c.Resp
	// the FormFile function takes in the POST input id file
	file, header, err := r.FormFile("file")

	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	defer file.Close()

	c.Logger.Info("Uploaded File:", header.Filename)

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	out, err := os.Create("./" + util.GetCurrentTimeNumberISOStrTime() + "-" + header.Filename)
	if err != nil {
		c.JSON(300, err.Error())
		return
	}

	defer out.Close()

	// write the content from POST to the file
	_, err = io.Copy(out, file)
	if err != nil {
		c.JSON(301, err.Error())
	}

	c.JSON(200, "ok")
}

func (hs *HTTPServer) NotFoundHandler(c *m.ReqContext) {
	if c.IsApiRequest() {
		c.JsonApiErr(404, "Not found", nil)
		return
	}

	//data, err := hs.setIndexViewData(c)
	//if err != nil {
	//	c.Handle(500, "Failed to get settings", err)
	//	return
	//}

	c.HTML(404, "index", nil)
}
