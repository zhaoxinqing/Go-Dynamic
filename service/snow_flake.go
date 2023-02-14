// SnowFlake：雪花算法
package service

import (
	"errors"
	"sync"
	"time"
)

const (
	workerBits  uint8 = 10 //机器码位数
	numberBits  uint8 = 12 //序列号位数
	workerMax   int64 = -1 ^ (-1 << workerBits)
	numberMax   int64 = -1 ^ (-1 << numberBits)
	timeShift   uint8 = workerBits + numberBits //时间戳偏移量
	workerShift uint8 = numberBits              //机器码偏移量
	startTime   int64 = 1650793569000           // 起始时间戳(单位:ms)
)

type Worker struct {
	mu        sync.Mutex
	timestamp int64
	workerId  int64
	number    int64
}

// NewWorker 生成一个新节点
func NewWorker(workerId int64) (*Worker, error) {
	if workerId < 0 || workerId > workerMax {
		return nil, errors.New("Worker ID excess of quantity")
	}
	return &Worker{
		timestamp: 0,
		workerId:  workerId,
		number:    0,
	}, nil
}

func (w *Worker) GetId() int64 {
	w.mu.Lock()
	defer w.mu.Unlock()
	now := time.Now().UnixNano() / 1e6
	if w.timestamp == now {
		w.number++
		if w.number > numberMax {
			for now <= w.timestamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		w.number = 0
		w.timestamp = now
	}
	ID := int64((now-startTime)<<timeShift | (w.workerId << workerShift) | (w.number))
	return ID
}
