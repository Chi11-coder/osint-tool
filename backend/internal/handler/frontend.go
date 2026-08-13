package handler

import (
	"io/fs"
	"net/http"
	"path"
)

type SPAHandler struct {
	staticFS   http.FileSystem
	fileServer http.Handler
}

func NewSPAHandler(fsys fs.FS) *SPAHandler {
	hfs := http.FS(fsys)
	return &SPAHandler{
		staticFS:   hfs,
		fileServer: http.FileServer(hfs),
	}
}

func (h *SPAHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p := r.URL.Path

	// 指定されたパスのファイルが存在する場合はnil
	f, err := h.staticFS.Open(p)
	if err == nil {
		_ = f.Close()
		h.fileServer.ServeHTTP(w, r)
		return
	}

	// パスが存在しない場合 GET, HEAD かつ 拡張子無 であるか検査
	if (r.Method == http.MethodGet || r.Method == http.MethodHead) && path.Ext(p) == "" {
		r2 := r.Clone(r.Context())
		r2.URL.Path = "/index.html"
		h.fileServer.ServeHTTP(w, r2)
		return
	}

	http.NotFound(w, r)
}
