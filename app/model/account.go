package model

func CreateAccountFactory(sqlType string) *AccountModel {
	return &AccountModel{BaseModel: BaseModel{DB: UseDbConn(sqlType)}}
}

type AccountModel struct {
	BaseModel `json:"-"`
	Id        int64  `gorm:"id" json:"id"`
	Phone     string `gorm:"phone" json:"phone"`
	Email     string `gorm:"email" json:"email"`
	Password  string `gorm:"password" json:"-"`
	Code      string `gorm:"code" json:"-"`
	CreatedAt string `gorm:"created_at" json:"created_at"`
}

// 表名
func (u *AccountModel) TableName() string {
	return "account_tab"
}

func (u *AccountModel) ClearTable() {
	sql := "TRUNCATE TABLE " + u.TableName()
	u.Exec(sql)
}

func (u *AccountModel) QueryByPhone(phone string) (temp AccountModel, err error) {
	sql := "SELECT * FROM account_tab WHERE phone=? ORDER BY id desc LIMIT 1"
	result := u.Raw(sql, phone).Find(&temp)
	return temp, result.Error
}

func (u *AccountModel) StoreAccount(phone string, email string, password, code string) (success bool, err error) {
	sql := "INSERT INTO account_tab(phone,email,password,code) VALUES (?,?,?,?)"
	result := u.Exec(sql, phone, email, password, code)
	if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}
