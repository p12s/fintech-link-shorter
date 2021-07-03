package shorter

// Link - основная сущность сервиса - ссылка
type Link struct {
	Id    int    `json:"id" db:"id"`
	Url   string `json:"url" db:"-"`
	Short string `json:"short" db:"short" binding:"required"`
	Long  string `json:"long" db:"long" binding:"required"`
}

// UserLink - сущность для получения запроса клиента и отправки ответа
type UserLink struct {
	Url string `json:"url" binding:"required"`
}
