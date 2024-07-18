package models

type Playlist struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Name   string `json:"name"`
	UserID uint   `json:"user_id"`
	Songs  []Song `json:"songs" gorm:"many2many:playlist_songs;"`
}
