package models

/* Dinh nghia cau truc du lieu dia chi IP
 Dung de giao tiep voi MySQL (anh xa cac truong du lieu trong database) 
 va phuc vu ca API JSON Response (chuyen tu struct sang JSON) */

type IP struct {
	ID        int    `json:"id"` 		// STT cua dia chi IP trong danh sach
	Address   string `json:"ip_address"`	// Dia chi IP
	Allocated bool   `json:"allocated"`	// Chi so phuc vu Automatic Allcation tra ve true/false 
}
