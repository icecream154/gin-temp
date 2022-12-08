package model

import (
	"goskeleton/app/global/variable"
	"time"
)

func CreateCodeFactory(sqlType string) *CodeModel {
	return &CodeModel{BaseModel: BaseModel{DB: UseDbConn(sqlType)}}
}

type CodeModel struct {
	BaseModel `json:"-"`
	Id        int64  `gorm:"id" json:"id"`
	Phone     string `gorm:"column:phone" json:"phone"`
	Code      string `gorm:"column:code" json:"code"`
	IssueTime string `gorm:"issue_time" json:"issue_time"`
	Checked   int8   `gorm:"checked" json:"checked"`
}

// 表名
func (u *CodeModel) TableName() string {
	return "sys_code_tab"
}

func (u *CodeModel) ClearTable() {
	sql := "TRUNCATE TABLE " + u.TableName()
	u.Exec(sql)
}

func (u *CodeModel) QueryUncheckedByPhone(phone string) (temp CodeModel, err error) {
	sql := "SELECT * FROM sys_code_tab WHERE phone=? AND checked=0 ORDER BY id desc LIMIT 1"
	result := u.Raw(sql, phone).Find(&temp)
	return temp, result.Error
}

func (u *CodeModel) QueryUncheckedByPhoneAndCode(phone string, code string) (temp CodeModel, err error) {
	tx := u.Begin()
	if tx.Error != nil {
		return temp, tx.Error
	}

	sql := "SELECT * FROM sys_code_tab WHERE phone=? AND code=? AND checked=0 ORDER BY id desc LIMIT 1"
	result := tx.Raw(sql, phone, code).Find(&temp)
	if result.Error != nil {
		tx.Rollback()
		return temp, result.Error
	}

	if temp.Id != 0 {
		sql = "UPDATE sys_code_tab SET checked=1 WHERE id=?"
		result = tx.Exec(sql, temp.Id)
		if result.Error != nil {
			tx.Rollback()
			return temp, result.Error
		}
	}

	result = tx.Commit()
	if result.Error != nil {
		tx.Rollback()
		return temp, result.Error
	}
	return temp, result.Error
}

func (u *CodeModel) StoreCode(phone string, code string) (success bool, err error) {
	issueTime := time.Unix(time.Now().Unix(), 0).Format(variable.DateFormat)
	tx := u.Begin()
	if tx.Error != nil {
		return false, tx.Error
	}

	sql := "UPDATE sys_code_tab SET checked=1 WHERE phone=?"
	result := tx.Exec(sql, phone)
	if result.Error != nil {
		tx.Rollback()
		return false, result.Error
	}

	sql = "INSERT INTO sys_code_tab(phone,code,issue_time) VALUES (?,?,?)"
	result = tx.Exec(sql, phone, code, issueTime)
	if result.Error != nil {
		tx.Rollback()
		return false, result.Error
	}

	result = tx.Commit()
	if result.Error != nil {
		tx.Rollback()
		return false, result.Error
	}
	return true, nil
}
