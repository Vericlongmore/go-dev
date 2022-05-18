package utils

// Kafka会话相关配置
type KafkaConfig struct {
	// for producer and consumer
	Version string   `json:"Version"`
	Brokers []string `json:"Brokers"`
	Topic   string   `json:"Topic"`
	// for producer
	ProducerReturnSuccesses bool `json:"ProducerReturnSuccesses"`
	ProducerReturnErrors    bool `json:"ProducerReturnErrors"`
	// for consumer
	ConsumerGroup        string `json:"ConsumerGroup"`
	OldestOffset         bool   `json:"OldestOffset"`
	ConsumeThreads       int    `json:"ConsumeThreads"`
	ConsumerReturnErrors bool   `json:"ConsumerReturnErrors"`
}

// IPS/OSG/TMS服务连接相关配置
type ServiceConnConfig struct {
	GrpcEndpoint   string `json:"GrpcEndpoint"`
	RequestTimeout int    `json:"RequestTimeout"`
	MaxRecvMsgSize int    `json:"MaxRecvMsgSize"` // 单位为Mb
}

// 人像库连接配置
type NamingRepositoryConfig struct {
	GrpcEndpoint   string   `json:"GrpcEndpoint"`
	DBNames        []string `json:"DBNames"`
	FeatureVersion int32    `json:"FeatureVersion"`
	MinScore       float32  `json:"MinScore"`
}

// 人档交互模块相关配置
type DocumentUpdateConfig struct {
	FeatureVersion int32        `json:"FeatureVersion"`
	KafkaConfig    *KafkaConfig `json:"KafkaConfig"`
}

// SearchEngine相关配置
type ClusteringConfig struct {
	DeviceID           int32   `json:"DeviceID"`
	FeatureDim         int32   `json:"FeatureDim"`
	UseInt8            int32   `json:"UseInt8"`
	ThreadNum          int     `json:"ThreadNum"`
	Strategy           int     `json:"Strategy"`
	TopK               int     `json:"TopK"`
	RepeatTimes        int     `json:"RepeatTimes"`
	AggrThreshold      float32 `json:"AggrThreshold"`
	Beta               float32 `json:"Beta"`
	AggrTopK           int     `json:"AggrTopK"`  // TODO: aggregate topk未和研究侧对齐
	Threshold          float32 `json:"Threshold"` // TODO: clusterer threshold需要sdk对齐研究侧
	MinPoints          int     `json:"MinPoints"`
	SubThreshold       float32 `json:"SubThreshold"`
	NorThreshold       float32 `json:"NorThreshold"`
	NTSThreshold       float32 `json:"NTSThreshold"`
	TSThreshold        float32 `json:"TSThreshold"`
	PreCThreshold      float32 `json:"PreCThreshold"`
	ClsMaxSubcenter    int32   `json:"ClsMaxSubcenter"`
	RetrievalTime      int32   `json:"RetrievalTime"`
	SpacetimeThreshold float32 `json:"SpacetimeThreshold"`
	MaxSubcenterNum    int     `json:"MaxSubcenterNum"`
	MinSubcenterNum    int     `json:"MinSubcenterNum"`
	MaxIndexSubsNum    int     `json:"MaxIndexSubsNum"`
	Alpha              float32 `json:"Alpha"`
	Density            int     `json:"Density"`
	Iteration          int     `json:"Iteration"`
	SimThreshold       float32 `json:"SimThreshold"`
	MaxSubcenterSize   int     `json:"MaxSubcenterSize"`
}

// Snapshot相关配置
type SnapshotConfig struct {
	RootPath            string `json:"RootPath"`
	IntervalTimeSecs    int32  `json:"IntervalTimeSecs"`
	WorkerId            string `json:"WorkerId"`
	RetainSnapshotCount int    `json:"RetainSnapshotCount"`
}

// 离线数据仓库配置
type OfflineDataRepositoryConfig struct {
	SavePath string `json:"SavePath"`
}

type SpacetimeFilterConfig struct {
	CameraNums int32     `json:"CameraNums"`
	RegionIds  []int32   `json:"RegionIds"`
	CameraIds  []int32   `json:"CameraIds"`
	ReachTimes [][]int32 `json:"ReachTimes"`
}
