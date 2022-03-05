// package runtime2

// import (
// 	"context"
// 	"fmt"

// 	"github.com/Shopify/sarama"
// 	"github.com/pkg/errors"
// 	zfield "github.com/tkeel-io/core/pkg/logger"
// 	"github.com/tkeel-io/kit/log"
// 	"go.uber.org/zap"
// )

// type RuntimeConfig struct {
// 	Source SourceConf
// }

// type Runtime struct {
// 	containers map[string]*Container
// 	dispatch   Dispatch
// 	dao        Dao

// 	ctx    context.Context
// 	cancel context.CancelFunc
// }

// func NewRuntime(ctx context.Context, d Dao, dispatcher Dispatch) *Runtime {
// 	ctx, cacel := context.WithCancel(ctx)
// 	return &Runtime{
// 		dao:        d,
// 		ctx:        ctx,
// 		cancel:     cacel,
// 		dispatch:   dispatcher,
// 		containers: make(map[string]*Container),
// 	}
// }

// func (r *Runtime) Start(cfg RuntimeConfig) error {
// 	config := sarama.NewConfig()
// 	config.Version = sarama.V2_3_0_0
// 	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange
// 	config.Consumer.Offsets.Initial = sarama.OffsetNewest

// 	gconsumer, err := sarama.NewConsumerGroup(cfg.Source.Brokers, cfg.Source.GroupName, config)
// 	if err != nil {
// 		log.Error("create consumer", zfield.Endpoints(cfg.Source.Brokers), zap.Error(err))
// 		return errors.Wrap(err, "create consumer")
// 	}

// 	if err = gconsumer.Consume(r.ctx, cfg.Source.Topics, r); nil != err {
// 		log.Error("consume source", zfield.Endpoints(cfg.Source.Brokers), zap.Error(err))
// 		return errors.Wrap(err, "consume source")
// 	}

// 	return nil
// }

// func (r *Runtime) GetContainer(id string) Container {
// 	if _, has := r.containers[id]; !has {
// 		log.Info("create container", zfield.ID(id))
// 		r.containers[id] = NewContainer(r.ctx, id)
// 	}
// }

// type runtimeConsumer struct {
// 	runtime *Runtime
// }

// // Setup is run at the beginning of a new session, before ConsumeClaim.
// func (r *runtimeConsumer) Setup(sarama.ConsumerGroupSession) error {
// 	return nil
// }

// // Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
// // but before the offsets are committed for the very last time.
// func (r *runtimeConsumer) Cleanup(sarama.ConsumerGroupSession) error {
// 	return nil
// }

// // ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
// // Once the Messages() channel is closed, the Handler must finish its processing
// // loop and exit.
// func (r *runtimeConsumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {

// 	for {
// 		select {
// 		case <-r.ctx.Done():
// 			return nil
// 		case msg := <-claim.Messages():
// 			log.Debug("consume message", zfield.Offset(msg.Offset),
// 				zfield.Partition(msg.Partition), zfield.Topic(msg.Topic), zap.Any("header", msg.Headers))

// 			containerID := fmt.Sprintf("%s-%d", msg.Topic, msg.Partition)
// 			if _, has := r.containers[containerID]; !has {
// 				log.Info("create container", zfield.ID(containerID),
// 					zfield.Partition(msg.Partition), zfield.Topic(msg.Topic))
// 				r.containers[containerID] = NewContainer(r.ctx, msg.Partition)
// 			}

// 			container := r.containers[containerID]
// 			container.DeliveredEvent(context.Background(), msg)
// 		}
// 	}
// }