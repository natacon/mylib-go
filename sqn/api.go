package sqn

import (
	"github.com/natacon/mylib-go/common"
	"net/url"
	"path"
	"strconv"
)

func KickService(endpoint string, serviceTypeId int) error {
	u, _ := url.Parse(endpoint)

	u.Path = path.Join(u.Path, "api", "activity")

	query := u.Query()
	query.Set("serviceTypeId", strconv.Itoa(serviceTypeId))
	query.Set("loginid", "90101")
	query.Set("password", "password")
	u.RawQuery = query.Encode()

	_, e := common.CallApi("POST", u.String())
	return e
}
