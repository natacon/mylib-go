package sqn

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

type Type struct {
	CreateUserID int `json:"createUserId"`
	Supertype    struct {
		Name string `json:"name"`
		Oid  int    `json:"oid"`
		ID   int    `json:"id"`
		Type struct {
			TypeDiscriminator struct {
				ID int `json:"id"`
			} `json:"typeDiscriminator"`
			Oid int `json:"oid"`
			ID  int `json:"id"`
		} `json:"type"`
		PersistenceStatus struct {
			ID    int    `json:"id"`
			Value string `json:"value"`
		} `json:"persistenceStatus"`
	} `json:"supertype"`
	UpdateUserID     int `json:"updateUserId"`
	HistoryValidTerm struct {
		Start string `json:"start"`
		End   string `json:"end"`
	} `json:"historyValidTerm"`
	Description string `json:"description"`
	Oid         int64  `json:"oid"`
	Type        struct {
		TypeDiscriminator struct {
			ID int `json:"id"`
		} `json:"typeDiscriminator"`
		Oid int `json:"oid"`
		ID  int `json:"id"`
	} `json:"type"`
	PackageType struct {
		Name string `json:"name"`
		Oid  int64  `json:"oid"`
		ID   int64  `json:"id"`
		Type struct {
			TypeDiscriminator struct {
				ID int `json:"id"`
			} `json:"typeDiscriminator"`
			Oid int `json:"oid"`
			ID  int `json:"id"`
		} `json:"type"`
		PersistenceStatus struct {
			ID    int    `json:"id"`
			Value string `json:"value"`
		} `json:"persistenceStatus"`
	} `json:"packageType"`
	PersistenceStatus struct {
		Oid  int `json:"oid"`
		ID   int `json:"id"`
		Type struct {
			TypeDiscriminator struct {
				ID int `json:"id"`
			} `json:"typeDiscriminator"`
			Oid int `json:"oid"`
			ID  int `json:"id"`
		} `json:"type"`
		Value             string `json:"value"`
		PersistenceStatus struct {
			ID    int    `json:"id"`
			Value string `json:"value"`
		} `json:"persistenceStatus"`
	} `json:"persistenceStatus"`
	ObjectClassification struct {
		Oid  int `json:"oid"`
		ID   int `json:"id"`
		Type struct {
			TypeDiscriminator struct {
				ID int `json:"id"`
			} `json:"typeDiscriminator"`
			Oid int `json:"oid"`
			ID  int `json:"id"`
		} `json:"type"`
		Value             string `json:"value"`
		PersistenceStatus struct {
			ID    int    `json:"id"`
			Value string `json:"value"`
		} `json:"persistenceStatus"`
	} `json:"objectClassification"`
	PhysicalName      string        `json:"physicalName"`
	Feature           []interface{} `json:"feature"`
	BusinessValidTerm struct {
		Start string `json:"start"`
		End   string `json:"end"`
	} `json:"businessValidTerm"`
	PropertyTypeList []struct {
		Container struct {
			Name string `json:"name"`
			Oid  int64  `json:"oid"`
			ID   int64  `json:"id"`
			Type struct {
				TypeDiscriminator struct {
					ID int `json:"id"`
				} `json:"typeDiscriminator"`
				Oid int `json:"oid"`
				ID  int `json:"id"`
			} `json:"type"`
			PersistenceStatus struct {
				ID    int    `json:"id"`
				Value string `json:"value"`
			} `json:"persistenceStatus"`
		} `json:"container"`
		Multiplicity struct {
			Min int `json:"min"`
			Max int `json:"max"`
		} `json:"multiplicity"`
		Shared           bool   `json:"shared"`
		CreateUserID     int    `json:"createUserId"`
		DefaultValue     string `json:"defaultValue"`
		HistoryValidTerm struct {
			Start string `json:"start"`
			End   string `json:"end"`
		} `json:"historyValidTerm"`
		Description string `json:"description"`
		Storable    bool   `json:"storable"`
		Oid         int64  `json:"oid"`
		Type        struct {
			TypeDiscriminator struct {
				ID int `json:"id"`
			} `json:"typeDiscriminator"`
			Oid int `json:"oid"`
			ID  int `json:"id"`
		} `json:"type"`
		AssignedType struct {
			Name string `json:"name"`
			Oid  int64  `json:"oid"`
			ID   int64  `json:"id"`
			Type struct {
				TypeDiscriminator struct {
					ID int `json:"id"`
				} `json:"typeDiscriminator"`
				Oid int `json:"oid"`
				ID  int `json:"id"`
			} `json:"type"`
			PersistenceStatus struct {
				ID    int    `json:"id"`
				Value string `json:"value"`
			} `json:"persistenceStatus"`
		} `json:"assignedType"`
		PhysicalName      string        `json:"physicalName"`
		Feature           []interface{} `json:"feature"`
		Composition       bool          `json:"composition"`
		CreateTs          string        `json:"createTs"`
		ID                int64         `json:"id"`
		Derived           bool          `json:"derived"`
		UpdateUserID      int           `json:"updateUserId"`
		ReadOnly          bool          `json:"readOnly"`
		PersistenceStatus struct {
			Oid  int `json:"oid"`
			ID   int `json:"id"`
			Type struct {
				TypeDiscriminator struct {
					ID int `json:"id"`
				} `json:"typeDiscriminator"`
				Oid int `json:"oid"`
				ID  int `json:"id"`
			} `json:"type"`
			Value             string `json:"value"`
			PersistenceStatus struct {
				ID    int    `json:"id"`
				Value string `json:"value"`
			} `json:"persistenceStatus"`
		} `json:"persistenceStatus"`
		BusinessValidTerm struct {
			Start string `json:"start"`
			End   string `json:"end"`
		} `json:"businessValidTerm"`
		Unique            bool   `json:"unique"`
		TenantID          int    `json:"tenantId"`
		Name              string `json:"name"`
		TypeDiscriminator struct {
			ID int `json:"id"`
		} `json:"typeDiscriminator"`
		PropertyIndex struct {
			IDSqnModelIndex      string `json:"ID.sqnModelIndex"`
			SqnModelIndex        string `json:"sqnModelIndex"`
			TypeIDSqnModelIndex  string `json:"型ID.sqnModelIndex"`
			OIDSqnModelIndex     string `json:"OID.sqnModelIndex"`
			TypeOIDSqnModelIndex string `json:"型OID.sqnModelIndex"`
		} `json:"propertyIndex,omitempty"`
		PropertyClassification struct {
			Oid  int `json:"oid"`
			ID   int `json:"id"`
			Type struct {
				TypeDiscriminator struct {
					ID int `json:"id"`
				} `json:"typeDiscriminator"`
				Oid int `json:"oid"`
				ID  int `json:"id"`
			} `json:"type"`
			Value             string `json:"value"`
			PersistenceStatus struct {
				ID    int    `json:"id"`
				Value string `json:"value"`
			} `json:"persistenceStatus"`
		} `json:"propertyClassification"`
		UpdateTs string `json:"updateTs"`
	} `json:"propertyTypeList"`
	SqnProject          string        `json:"sqnProject"`
	ObjectIndex         int           `json:"objectIndex"`
	AssociationTypeList []interface{} `json:"associationTypeList"`
	TenantID            int           `json:"tenantId"`
	Name                string        `json:"name"`
	TypeDiscriminator   struct {
		ID int64 `json:"id"`
	} `json:"typeDiscriminator"`
	CreateTs          string        `json:"createTs"`
	OperationTypeList []interface{} `json:"operationTypeList"`
	ID                int64         `json:"id"`
	UpdateTs          string        `json:"updateTs"`
}

// targetTypeに指定したtypeのsupertypeを検索して返却する。
// 最上位クラスをtargetTypeに指定した場合はerrorに errSuperTypeNotFound を返却する。
func (targetType *Type) FindSuperType() (Type, error) {
	for _, t := range cachedAllType {
		if targetType.Supertype.Name == t.Name {
			return t, nil
		}
	}
	return Type{}, errors.New("SuperType is not found.")
}

// 全親クラスのsupertypeを検索して返す。
func (targetType *Type) FindAllSuperType() []Type {
	return targetType.findSuperTypeRecursively(nil)
}

// 全親クラスのsupertypeを検索して返す。
func (targetType *Type) findSuperTypeRecursively(superTypes []Type) []Type {
	superType, _ := targetType.FindSuperType()
	// superTypeがないなら終わり
	if superType.Name == "" {
		return superTypes
	}
	superTypes = append(superTypes, superType)
	return superType.findSuperTypeRecursively(superTypes)
}

// targetTypeに指定したtypeの最上位typeを検索して返却する。
// targetTypeが最上位クラスの場合、自身を返す。
func (targetType *Type) FindTopSuperType() *Type {
	superType, err := targetType.FindSuperType()
	if err != nil {
		return targetType
	}

	// まだ上位クラスがある場合、親クラスを再帰的に検索する。
	return superType.FindTopSuperType()
}

type Types []Type

var cachedAllType []Type

// メタメタ定義のjsonフォルダの全ファイルを読み込み、それらを構造体のリスト（cachedAllType）に格納する。
// cachedAllTypeはTypeのメソッド内で対象のTypeを検索するために都度利用される。
func readAllMetaJson(dir string) error {
	// メタメタ定義のjsonフォルダの全ファイルを読み込む
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	// jsonファイルを読み込み、構造体のリストへと変換する
	for _, file := range files {
		filePath := filepath.Join(dir, file.Name())
		bytes, err := ioutil.ReadFile(filePath)
		if err != nil {
			return err
		}
		var meta Type
		err = json.Unmarshal(bytes, &meta)
		if err != nil {
			return err
		}
		cachedAllType = append(cachedAllType, meta)
	}
	return nil
}

func FindTypeByName(typeName string) (Type, error) {
	for _, t := range cachedAllType {
		if t.Name == typeName {
			return t, nil
		}
	}
	return Type{}, fmt.Errorf("%s is not found.", typeName)
}
