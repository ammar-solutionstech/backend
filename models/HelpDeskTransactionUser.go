package models

// HelpDeskTransactionUser links users to help desk transactions.
type HelpDeskTransactionUser struct {
	UserID        int `gorm:"column:user_id" json:"user_id"`
	TransactionID int `gorm:"column:transaction_id" json:"transaction_id"`
}

// TableName overrides default.
func (HelpDeskTransactionUser) TableName() string {
	return `public."Help_desk_transaction_user"`
}
