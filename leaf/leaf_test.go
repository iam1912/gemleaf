package leaf

import (
	"fmt"
	"strings"
	"testing"

	"github.com/iam1912/gemseries/gemleaf/config"
	"github.com/iam1912/gemseries/gemleaf/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
	s  *GemLeaf
)

type Record struct {
	Value string
}

func setup() {
	_ = config.MustGetAppConfig()
	DB = config.MustGetDB()
	DB.AutoMigrate(&models.Leaf{})
	s = NewGemLeaf(DB)
}

func RunSQL(db *gorm.DB, sqls string) {
	for _, sql := range strings.Split(sqls, "\n") {
		if strings.TrimSpace(sql) != "" {
			db.Exec(strings.TrimSpace(sql))
		}
	}
}

func GetRecords(db *gorm.DB, tableName string, columns string, extra ...string) string {
	var (
		extraSQL = ""
		results  = []string{}
		records  = []Record{}
	)
	if len(extra) > 0 {
		extraSQL = extra[0]
	}
	db.Raw(fmt.Sprintf(`SELECT CONCAT_WS(", ", %v) AS value FROM %v %v`, columns, tableName, extraSQL)).Scan(&records)
	for _, record := range records {
		results = append(results, record.Value)
	}
	return strings.Join(results, ";")
}

type LeafCreateTestCase struct {
	Tag  string
	Desc string
}

func TestLeafCreate(t *testing.T) {
	setup()
	RunSQL(DB, "TRUNCATE leafs")
	testCases := []LeafCreateTestCase{
		{Tag: "test.api.leaf", Desc: "test.api.leaf"},
		{Tag: "test.api1.leaf", Desc: "test.api1.leaf"},
	}
	for _, testCase := range testCases {
		s.Create(testCase.Tag, 0, 0, testCase.Desc)
	}
	body := `test.api.leaf, 1, 2000, test.api.leaf;test.api1.leaf, 1, 2000, test.api1.leaf`
	leafs := GetRecords(DB, "leafs", "biz_tag, max_id, step, description")
	assert.Equal(t, body, leafs)
}

type LeafGetTestCase struct {
	Tag        string
	ExpectedID string
}

func TestLeafGet(t *testing.T) {
	setup()
	RunSQL(DB, `
		TRUNCATE leafs;
		INSERT INTO leafs (biz_tag, max_id, step) VALUES ('test.api.leaf', 2, 5); 
	`)
	testCases := []LeafGetTestCase{
		{Tag: "test.api.leaf"},
		{Tag: "test.api.leaf"},
	}
	go func() {
		for _, testCase := range testCases {
			id, err := s.GetID(testCase.Tag)
			assert.NoError(t, err)
			fmt.Println(id)
		}
	}()
	go func() {
		for _, testCase := range testCases {
			id, err := s.GetID(testCase.Tag)
			assert.NoError(t, err)
			fmt.Println(id)
		}
	}()
	go func() {
		for _, testCase := range testCases {
			id, err := s.GetID(testCase.Tag)
			assert.NoError(t, err)
			fmt.Println(id)
		}
	}()
	go func() {
		for _, testCase := range testCases {
			id, err := s.GetID(testCase.Tag)
			assert.NoError(t, err)
			fmt.Println(id)
		}
	}()
	select {}
}
