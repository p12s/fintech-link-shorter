package shorter

type Link struct {
	Id    int    `json:"id" db:"id"`
	Url   string `json:"url" db:"-"`
	Short string `json:"short" db:"short" binding:"required"`
	Long  string `json:"long" db:"long" binding:"required"`
}

type UserLink struct {
	Url   string `json:"url" binding:"required"`
}