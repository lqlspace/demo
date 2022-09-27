package service

type (
	StuRequest struct {
		Name string `json:"name"`
	}
	StuResponse struct {
		Data string `json:"data"`
	}
)

func GetStudentService(req StuRequest) (resp StuResponse, err error) {
	//TODO: coding by yourself
	return
}
