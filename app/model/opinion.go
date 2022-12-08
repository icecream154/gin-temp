package model

import (
	"strconv"
	"strings"
)

func CreateOpinionFactory(sqlType string) *OpinionModel {
	return &OpinionModel{BaseModel: BaseModel{DB: UseDbConn(sqlType)}}
}

type OpinionModel struct {
	BaseModel    `json:"-"`
	Id           int64  `gorm:"primary_key" json:"id"`
	Account      string `gorm:"account" json:"account"`
	Content      string `gorm:"content" json:"content"`
	Image        string `gorm:"image" json:"image"`
	Contact      string `gorm:"contact" json:"contact"`
	Dealt        int8   `gorm:"dealt" json:"dealt"`
	Delete       int8   `gorm:"delete" json:"delete"`
	Handler      string `gorm:"handler" json:"handler"`
	HandleResult string `gorm:"handle_result" json:"handle_result"`
	CreatedAt    string `gorm:"created_at" json:"created_at"`
	UpdatedAt    string `gorm:"updated_at" json:"updated_at"`
	// 来自 accountInfo 的信息
	Name          string `json:"name"`
	Certification int8   `json:"certification"`
	LogoUrl       string `json:"logo_url"`
	CollegeName   string `json:"college_name"`
	CollegeCode   int64  `json:"college_code"`
	YouniCode     string `json:"youni_code"`
}

func (u *OpinionModel) TableName() string {
	return "sys_opinion_tab"
}

func (u *OpinionModel) ClearTable() {
	sql := "TRUNCATE TABLE " + u.TableName()
	u.Exec(sql)
}

func (u *OpinionModel) QueryById(id int64) (temp OpinionModel, err error) {
	sql := "SELECT * FROM sys_opinion_tab WHERE id=?"
	result := u.Raw(sql, id).Find(&temp)
	if result.Error != nil {
		return temp, result.Error
	}
	return temp, nil
}

func (u *OpinionModel) StoreOpinion(account string, content string, image string,
	contact string) (success bool, err error) {
	sql := "INSERT INTO sys_opinion_tab(account,content,image,contact) VALUES (?,?,?,?)"
	result := u.Exec(sql, account, content, image, contact)
	return result.Error == nil, result.Error
}

func (u *OpinionModel) HandleOpinion(id int64, handler string, handleResult string) (success bool, err error) {
	sql := "UPDATE sys_opinion_tab SET dealt=2, handler=?, handle_result=? WHERE id=?"
	result := u.Exec(sql, handler, handleResult, id)
	return result.Error == nil, result.Error
}

func (u *OpinionModel) buildQueryOpinionSql(id int64, account string, dealt int8,
	isCount int8, pageNum int, pageSize int) string {
	sql := "select * from sys_opinion_tab where 1=1 and "
	if isCount == 1 {
		sql = "select count(*) as counts from sys_opinion_tab where 1=1 and "
	}
	if id != 0 {
		sql += "id=" + strconv.FormatInt(id, 10) + " and "
	} else {
		if account != "" {
			sql += "account=" + account + " and "
		}
		if dealt != 0 {
			sql += "dealt=" + strconv.FormatInt(int64(dealt), 10) + " and "
		}
	}
	sql = strings.TrimSuffix(sql, " and ")

	if isCount != 1 {
		offset := pageSize * (pageNum - 1)
		sql += " LIMIT " + strconv.FormatInt(int64(offset), 10) + ", " + strconv.FormatInt(int64(pageSize), 10)
	}
	return sql
}

func (u *OpinionModel) QueryOpinion(id int64, account string,
	dealt int8, pageNum int, pageSize int) (total int, temp []OpinionModel, err error) {

	countSql := u.buildQueryOpinionSql(id, account, dealt, 1, 0, 0)
	if result := u.Raw(countSql).Find(&total); result.Error != nil {
		return -1, nil, result.Error
	}

	sql := u.buildQueryOpinionSql(id, account, dealt, 0, pageNum, pageSize)
	if result := u.Raw(sql).Find(&temp); result.Error != nil {
		return -1, nil, result.Error
	}

	return total, temp, nil
}
