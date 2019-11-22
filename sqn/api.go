package sqn

import (
	"fmt"
	"github.com/natacon/mylib-go/restapi"
	"net/url"
	"path"
	"strconv"
)

// TODO: endpoint, tenantId, loginid, passwordは外部設定ファイル化する
func KickService(endpoint string, tenantId int, serviceTypeId int, body string) error {
	u, _ := url.Parse(endpoint)

	u.Path = path.Join(u.Path, "api", "activity")

	query := u.Query()
	query.Set("serviceTypeId", strconv.Itoa(serviceTypeId))
	query.Set("loginid", "90101")
	query.Set("password", "password")
	query.Set("tenantId", strconv.Itoa(tenantId))
	u.RawQuery = query.Encode()

	_, e := restapi.CallApi(restapi.ApiRequest{
		Method: "POST",
		Url:    u.String(),
		Body:   body,
	})
	return e
}

func 組み立てqueryparameter() {

}

// TODO: 処理を分ける
// TODO: 実装場所も変える？
func CreateInheritedEnumValueSql(enumValueName string, dir string) (string, error) {
	if err := readAllMetaJson(dir); err != nil {
		return "", err
	}

	// supertypeを探すために検索対象のもとになるtypeを取得しておく
	targetType, err := FindTypeByName(enumValueName)
	if err != nil {
		return "", err
	}
	superTypes := targetType.FindAllSuperType()

	// 自身のtype、自身の全親クラスのtypeを条件にしたENUMVALUEテーブルを検索するクエリを整形する
	sqlStr := fmt.Sprintf("SELECT * FROM ENUMVALUE WHERE TYPE_OID IN (%d", targetType.Oid)
	for _, superType := range superTypes {
		sqlStr += fmt.Sprintf(",%d", superType.Oid)
	}
	sqlStr += ") AND TENANT_ID = #{tenantId} ORDER BY TYPE_OID, OID"
	return sqlStr, nil
}
