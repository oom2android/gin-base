package model

type UserModel struct {
	Model
	Username string  `json:"username" xomr:"varchar(20)"`
	Mobile   string  `json:"mobile" xorm:"varchar(11) index"`
	Password string  `json:"password" xorm:"varchar(255)"`
	Avatar   string  `json:"avatar" xorm:"varchar(255)"`
	Blacnce  float64 `json:"blance" xorm:"double"`
	Status   int8    `json:"status" xorm:"tinyin"`
	City     string  `json:"city" xorm:"varchar(255)"`
}

func (model *UserModel) B() {

}
