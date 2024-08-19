package pkg

import (
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

type MockCloudinary struct {
	mock.Mock
}

func (m *MockCloudinary) UploadFile(ctx *gin.Context, file interface{}, fileName string) (*uploader.UploadResult, error) {
	args := m.Mock.Called()
	return args.Get(0).(*uploader.UploadResult), args.Error(1)
}