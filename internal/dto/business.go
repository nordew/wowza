package dto

type CreateBusinessRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	WebsiteURL  string `json:"websiteUrl" binding:"omitempty,url"`
	Location    string `json:"location"`
	CategoryID  string `json:"categoryId" binding:"required,uuid"`
}

type UpdateBusinessRequest struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	WebsiteURL  *string `json:"websiteUrl" binding:"omitempty,url"`
	Location    *string `json:"location"`
	CategoryID  *string `json:"categoryId" binding:"omitempty,uuid"`
}

type BusinessResponse struct {
	ID          string           `json:"id"`
	UserID      string           `json:"userId"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	WebsiteURL  string           `json:"websiteUrl"`
	Location    string           `json:"location"`
	Category    CategoryResponse `json:"category"`
} 