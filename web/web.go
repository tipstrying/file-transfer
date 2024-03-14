package web

import (
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"ui"

	_ "embed"

	"github.com/gin-gonic/gin"
)

type file_name struct {
	Name string `uri:"file" binding:"required"`
}

func Start() {
	r := gin.Default()

	t, e := loadTemplate()
	if e == nil {
		r.SetHTMLTemplate(t)
	}

	_, e = os.Stat("./files")
	if e != nil {
		os.Mkdir("./files", 0666)
	}

	r.GET("/", func(ctx *gin.Context) {
		st, e := os.Stat("./files")
		if e != nil {

			ctx.HTML(200, "/index.tmpl", gin.H{"files": []string{}})
			return
		}
		if st.Mode().IsRegular() {
			ctx.HTML(200, "/index.tmpl", gin.H{"files": []string{}})
			return
		}
		ctx.HTML(200, "/index.tmpl", gin.H{"files": read_files("./files")})
		ctx.AbortWithStatus(200)
	})
	r.GET("/:file", func(ctx *gin.Context) {
		fn := file_name{}
		if e := ctx.ShouldBindUri(&fn); e != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}
		st, e := os.Stat("./files/" + fn.Name)
		if e != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		if st.Mode().IsRegular() {

			// ctx.Header("Content-Disposition:", "attachment; filename="+fn.Name)
			// ctx.File()
			ctx.FileAttachment("./files/"+fn.Name, fn.Name)
		}
	})

	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	r.POST("/", func(c *gin.Context) {
		dst := "./files/"
		file, _ := c.FormFile("file")
		save_file(c, file, dst, file.Filename)

		// c.String(http.StatusOK, fmt.Sprintf("'%s' 上传成功!", file.Filename))
		c.Redirect(301, "/")
	})
	r.Run(":8083")
}

func save_file(c *gin.Context, f *multipart.FileHeader, dst, name string) {
	st1, e := os.Stat(dst)
	if e != nil {
		if os.IsNotExist(e) {
			if strings.HasSuffix(dst, "files/") {
				os.Mkdir(dst, 0666)
				save_file(c, f, dst, name)
				return
			}
		}
		save_file(c, f, "./", name)
		return
	}

	if st1.Mode().IsRegular() {
		save_file(c, f, "./", name)
		return
	}

	_, e = os.Stat(dst + name)
	if e != nil {
		c.SaveUploadedFile(f, dst+name)
		ui.NewFile(name)
		return
	}

	save_file(c, f, dst, "new_"+name)
}

func loadTemplate() (*template.Template, error) {
	t := template.New("")
	for name, file := range Assets.Files {
		if file.IsDir() || !strings.HasSuffix(name, ".tmpl") {
			continue
		}
		h, err := io.ReadAll(file)
		if err != nil {
			return nil, err
		}
		t, err = t.New(name).Parse(string(h))
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}

func read_files(path string) []string {
	fs := []string{}
	st, e := os.Stat(path)
	if e != nil {
		return fs
	}
	if st.IsDir() {
		info, e := os.ReadDir(path)
		if e != nil {
			return fs
		}
		for _, f := range info {
			if f.Type().IsRegular() {
				fs = append(fs, f.Name())
			}
		}
	}
	return fs
}
