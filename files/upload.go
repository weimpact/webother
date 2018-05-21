package files

import (
	"io"
	"net/http"
	"strconv"

	"github.com/weimpact/webother/logger"
)

type Data struct {
	file   io.Reader
	userID int64
}

func Upload(fs Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		file, header, err := r.FormFile("file")
		if err != nil {
			logger.Errorf("[FileUpload] Error getting file %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		defer file.Close()
		logger.Infof("[FileUpload] file name: %s", err, header)

		userID, err := strconv.ParseInt(r.FormValue("user_id"), 10, 64)
		if err != nil {
			logger.Errorf("[FileUpload] Error parsing other form data %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := fs.Save(r.Context(), Data{file: file, userID: userID}); err != nil {
			logger.Errorf("[FileUpload] Error saving file %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}
