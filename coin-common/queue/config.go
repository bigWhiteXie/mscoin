package queue

type KafkaConfig struct {
	Addr     string   `json:"addr,optional"`
	WriteCap int      `json:"writeCap,optional"`
	ReadCap  int      `json:"readCap,optional"`
	Groups   []string `json:"Groups,optional"`
}
