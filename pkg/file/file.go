package file

import (
	"fmt"
	"go-devops-admin/pkg/app"
	"go-devops-admin/pkg/auth"
	"go-devops-admin/pkg/helpers"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
)

// 写入文件
func Put(data []byte, to string) error {
	err := os.WriteFile(to, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

// 判断文件是否存在
func Exists(fileToCheck string) bool {
	if _, err := os.Stat(fileToCheck); os.IsNotExist(err) {
		return false
	}
	return true
}

func FileNameWithoutExtension(filename string) string {
	return strings.TrimSuffix(filename, filepath.Ext(filename))
}

func SaveUploadAvatar(c *gin.Context, file *multipart.FileHeader) (string, error) {

	var avatar string
	// 确保目录存在
	publicPath := "public"
	dirName := fmt.Sprintf("/uploads/avatars/%s/%s", app.TimeNowInTimezone().Format("2006/01/02"), auth.CurrentUID(c))
	os.MkdirAll(publicPath+dirName, 0755)

	// 保存文件
	fileName := randomNameFormUploadFile(file)
	avatarPath := publicPath + dirName + fileName
	if err := c.SaveUploadedFile(file, avatarPath); err != nil {
		return avatar, err
	}

	// 添加图片裁剪, 控制头像图片的尺寸 256 x 256
	img, err := imaging.Open(avatarPath, imaging.AutoOrientation(true))
	if err != nil {
		return avatar, err
	}

	resizeAvatar := imaging.Thumbnail(img, 256, 256, imaging.Lanczos)
	resizeAvatarFileName := randomNameFormUploadFile(file)
	resizeAvatarPath := publicPath + dirName + resizeAvatarFileName
	err = imaging.Save(resizeAvatar, resizeAvatarPath)
	if err != nil {
		return avatar, err
	}

	// 删除老文件
	err = os.Remove(avatarPath)
	if err != nil {
		return avatar, err
	}
	return avatarPath, nil
}

func randomNameFormUploadFile(file *multipart.FileHeader) string {
	return helpers.RandomString(16) + filepath.Ext(file.Filename)
}
