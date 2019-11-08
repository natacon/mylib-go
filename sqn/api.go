package sqn

import (
	"github.com/natacon/mylib-go/common"
	"net/url"
	"path"
	"strconv"
)

var endpoint = "http://localhost:8080"

func KickService(activityTypeId int) error {
	u, _ := url.Parse(endpoint)

	u.Path = path.Join(u.Path, "api", "activity")

	query := u.Query()
	query.Set("serviceTypeId", strconv.Itoa(activityTypeId))
	query.Set("loginid", "90101")
	query.Set("password", "password")
	u.RawQuery = query.Encode()

	_, e := common.CallApi("POST", u.String())
	return e
}
