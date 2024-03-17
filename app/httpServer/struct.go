package httpServer

type httpServer struct {
	Port            int      `json:"port"`
	RefreshInterval int      `json:"refresh_interval"`
	StorageDir      string   `json:"storage_dir"`
	MediaTypes      []string `json:"media_types"`
}
