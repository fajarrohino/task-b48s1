package middleware

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
)

func UploadFile(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		file, errImg := c.FormFile("input-img") //dibuat multiplatfom data di form add project -> enctype
		if errImg != nil {
			return c.JSON(http.StatusInternalServerError, errImg.Error())
		}
		// fmt.Println("file :", file)
		// return c.String(http.StatusOK, "berhasil")

		src, err := file.Open() // -> untuk mengambil path
		if err != nil {
			return c.JSON(http.StatusInternalServerError, errImg.Error()) 
		}
		fmt.Println("src :", src)

		defer src.Close()

		tmpFile, err := ioutil.TempFile("uploads", "img-*.png") // utk upload di folder mana dan templ-namenya
		if err != nil {
			return c.JSON(http.StatusInternalServerError, errImg.Error()) 
		}
		fmt.Println("TmpFile:", tmpFile)

		textCopy, err := io.Copy(tmpFile, src)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		fmt.Println("written copy:", textCopy)

		data := tmpFile.Name() // uploads/img-123124212.png -> dari tmpFile dan ambil namenya
		fmt.Println("Data Name utuh:", data)

		filename := data[8:] //dari data mengambil ke 8 = img-123124212.png dari (uploads/img-123124212.png)
		fmt.Println("Nama File terpotong", filename)
		// return c.String(http.StatusOK, "berhasil")

		c.Set("dataFile", filename)
		return next(c)
	}
}