package models

func NewServices(connectionInfo string) (*Services, error) {

}

type Services struct {
	Gallery GalleryService
	User    UserService
}
