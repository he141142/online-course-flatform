package serdes

import (
	"testing"

	"drake.elearn-platform.ru/internal/registry"
	"github.com/stretchr/testify/assert"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Job  string `json:"job"`
}

type EntityAggregate struct {
	ID         string `json:"id"`
	EntityName string `json:"entity_name"`
}

type BuildingInfo struct {
	EntityAggregate
	Address    string `json:"address"`
	TowerCode  string `json:"tower_code"`
	Department string `json:"department"`
}

func NewTTTBuildingInfo(ID string) *BuildingInfo {
	return &BuildingInfo{
		EntityAggregate: EntityAggregate{
			ID:         ID,
			EntityName: "BuildingInfo",
		},
		Address:    "18 Ton That Thuyet",
		TowerCode:  "18",
		Department: "DU35",
	}
}

func (r *BuildingInfo) Key() string {
	return r.EntityName
}
func (r *User) Key() string {
	return "User"
}

func TestEncoderAndDecoder(t *testing.T) {
	reg := registry.New()
	jSerde := NewJSONSerde(reg)
	u := &User{}
	err := jSerde.RegisterKey("User", u)
	if err != nil {
		t.Error(err)
	}
	buildingInfo := NewTTTBuildingInfo("10001")

	err = jSerde.Register(buildingInfo)
	if err != nil {
		t.Error(err)
	}
	//var b BuildingInfo
	bb, err := reg.Serialize(buildingInfo.Key(), buildingInfo)
	if err != nil {
		t.Error(err)
	}
	v, err := reg.Deserialize(buildingInfo.Key(), bb)
	buildingDecode := v.(*BuildingInfo)
	assert.Equal(t, buildingDecode.Department, buildingInfo.Department)
	//fmt.Println(buildingDecode)
}
