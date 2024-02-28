package httpServer

type httpServer struct {
	Port       int      `json:"port"`
	StorageDir string   `json:"storage_dir"`
	MediaTypes []string `json:"media_types"`
}
