package organ

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"api/base"
	"api/db"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func GetMockDB() *gorm.DB {
	mockDB, _ := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	return mockDB
}
func setupTestRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/organ/tree", OrganTree)
	r.POST("/organ/page", OrganPage)
	r.POST("/organ/save", OrganSave)
	r.POST("/organ/del", OrganDel)
	r.POST("/organ/disable", OrganDisable)
	r.POST("/organ/enable", OrganEnable)
	r.POST("/organ/saveParent", OrganSaveParent)
	return r
}

func TestOrganTree(t *testing.T) {
	// 初始化测试数据
	mockOrgans := []BmbpRbacOrgan{{
		BaseBean: base.BaseBean{
			DataId: "1",
		},
		BaseTree: base.BaseTree{
			NodeCode:       "1",
			NodeParentCode: "",
			NodeTitle:      "根节点",
		},
	}, {
		BaseBean: base.BaseBean{
			DataId: "2",
		},
		BaseTree: base.BaseTree{
			NodeCode:       "2",
			NodeParentCode: "1",
			NodeTitle:      "子节点",
		},
	}}

	// 模拟数据库查询
	oldDB := db.DbUtil
	defer func() { db.DbUtil = oldDB }()

	db.DbUtil = GetMockDB()
	db.DbUtil.AutoMigrate(&BmbpRbacOrgan{})
	db.DbUtil.Create(&mockOrgans)

	// 创建测试请求
	router := setupTestRouter()
	body, _ := json.Marshal(BmbpRbacOrgan{})
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/organ/tree", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	// 执行请求
	router.ServeHTTP(w, req)

	// 验证结果
	assert.Equal(t, 200, w.Code)
	var resp base.Resp
	json.Unmarshal(w.Body.Bytes(), &resp)
	slog.Info("resp", "resp", resp)
	assert.Equal(t, "0", resp.Code)
	assert.NotNil(t, resp.Data)
}

func TestOrganPage(t *testing.T) {
	// 初始化测试数据
	mockOrgans := make([]BmbpRbacOrgan, 15)
	for i := 0; i < 15; i++ {
		mockOrgans[i] = BmbpRbacOrgan{
			BaseTree: base.BaseTree{
				NodeCode:  string(rune(i + 1)),
				NodeTitle: "组织" + string(rune(i+1)),
			},
		}
	}

	// 模拟数据库查询
	oldDB := db.DbUtil
	defer func() { db.DbUtil = oldDB }()

	mockDB, _ := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	db.DbUtil = mockDB
	mockDB.Create(&mockOrgans)

	// 创建测试请求
	router := setupTestRouter()
	pageReq := base.PageReq[BmbpRbacOrgan]{
		Page:   1,
		Size:   10,
		Params: &BmbpRbacOrgan{BaseTree: base.BaseTree{NodeTitle: "组织1"}},
	}
	body, _ := json.Marshal(pageReq)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/organ/page", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	// 执行请求
	router.ServeHTTP(w, req)

	// 验证结果
	assert.Equal(t, 200, w.Code)
	var resp base.Resp
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, "0", resp.Code)
	assert.NotNil(t, resp.Data)
}

// 其他测试函数可以按照类似模式编写
// TestOrganSave, TestOrganDel, TestOrganDisable, TestOrganEnable, TestOrganSaveParent
