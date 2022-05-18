package kafka_checker

import (
	"flag"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/golang/protobuf/proto" // nolint
	log "github.com/sirupsen/logrus"
	oplog_api "gitlab.sz.sensetime.com/rtc/cluster-searching/api/oplog"
	"gitlab.sz.sensetime.com/rtc/realtime-clustering/kafka"
	"os"
	"os/signal"
	"reflect"
	"syscall"
	"time"
	"utils"
)

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

var (
	configPath = flag.String("configPath", "./config.json", "config path")
)

type ServiceConfig struct {
	KafkaConfig KafkaConfig
}

type EventConsumer struct {
	AddCount    int
	DeleteCount int
	UpdateCount int
	NilCount    int
	OthersCount int

	AddFaceIdCount    int
	AddPedIdCount     int
	AddFacePedIdCount int

	UpdateFaceIdCount    int
	UpdatePedIdCount     int
	UpdateFacePedIdCount int
}

func NewConsumer(cfgPath string) {
	config := &ServiceConfig{}
	LoadConfigFileOrPanic(cfgPath, config)
	consumer := &EventConsumer{}
	go consumer.Run()

	cg := kafka.NewConsumerGroup(config.KafkaConfig, consumer.Process, "fake_consumer", 1)
	cg.Run()

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-sigc
}

func (t *EventConsumer) Process(msg *sarama.ConsumerMessage) error {
	oplog := &oplog_api.OpLog{}
	if err := proto.Unmarshal(msg.Value, oplog); err != nil {
		log.WithError(err).Error("bad oplog message")
		return err
	}
	offset := msg.Offset
	switch oplog.GetLog().(type) {
	case *oplog_api.OpLog_OpAddDocument:
		t.AddCount++
		docs := oplog.GetLog().(*oplog_api.OpLog_OpAddDocument).OpAddDocument.Docs
		if docs == nil || len(docs) == 0 {
			fmt.Println("add (nil or empty) docs ")
			return nil
		}
		for _, doc := range docs {
			if doc.FaceObjectIds != nil && len(doc.FaceObjectIds) > 0 &&
				doc.PedObjectIds != nil && len(doc.PedObjectIds) > 0 {
				t.AddFacePedIdCount++
				//fmt.Printf("add docId: %d ; faceids: %+v ; pedids: %+v\n", doc.Id, doc.FaceObjectIds, doc.PedObjectIds)
			} else if doc.FaceObjectIds != nil && len(doc.FaceObjectIds) > 0 {
				t.AddFaceIdCount++
			} else if doc.PedObjectIds != nil && len(doc.PedObjectIds) > 0 {
				t.AddPedIdCount++
			}

			fmt.Printf("add docid:%d, offset:%d\n", doc.Id, offset)
		}
	case *oplog_api.OpLog_OpDeleteDocument:
		t.DeleteCount++
		dbId := oplog.GetLog().(*oplog_api.OpLog_OpDeleteDocument).OpDeleteDocument.DocIds
		fmt.Printf("delete docId: %d\n", dbId)
	case *oplog_api.OpLog_OpUpdateDocument:
		t.UpdateCount++
		doc := oplog.GetLog().(*oplog_api.OpLog_OpUpdateDocument).OpUpdateDocument.Doc
		if doc.FaceObjectIds != nil && len(doc.FaceObjectIds) > 0 &&
			doc.PedObjectIds != nil && len(doc.PedObjectIds) > 0 {
			t.UpdateFacePedIdCount++
			//fmt.Printf("update docId: %d ; faceids: %+v ; pedids: %+v\n", doc.Id, doc.FaceObjectIds, doc.PedObjectIds)
		} else if doc.FaceObjectIds != nil && len(doc.FaceObjectIds) > 0 {
			t.UpdateFaceIdCount++
		} else if doc.PedObjectIds != nil && len(doc.PedObjectIds) > 0 {
			t.UpdatePedIdCount++
		}
		//time.Sleep(time.Second * 2)
	default:
		log.Infof("oplog: %v", oplog)
		if oplog.GetLog() == nil {
			//fmt.Println("nil oplog")
			t.NilCount++
		} else {
			fmt.Println(reflect.TypeOf(oplog.GetLog()).String())
			t.OthersCount++
		}
	}

	return nil
}

func (t *EventConsumer) Run() {
	ticker := time.Tick(time.Second * 2)
	for {
		select {
		case <-ticker:
			fmt.Printf("add:%d delete:%d update:%d nil:%d other:%d |||| [add] face: %d ped: %d faceped: %d |||| [update] face: %d ped: %d faceped: %d \n", t.AddCount,
				t.DeleteCount, t.UpdateCount, t.NilCount, t.OthersCount, t.AddFaceIdCount, t.AddPedIdCount, t.AddFacePedIdCount,
				t.UpdateFaceIdCount, t.UpdatePedIdCount, t.UpdateFacePedIdCount)
		}
	}
}

func main() {
	flag.Parse()
	NewConsumer(*configPath)
}
