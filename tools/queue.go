package tools

import (
	"FIM-Chat/errors"
	"sync"
)

type Queue struct {
	data  []any        //泛型切片
	begin uint         //首节点下标
	end   uint         //尾节点下标
	cap   uint         //容量
	len   uint         // 队列元素大小
	mutex sync.RWMutex //并发控制锁
}

const minQueueLen = 16

func New() (q *Queue) {
	return &Queue{
		data:  make([]any, minQueueLen),
		begin: 0,
		end:   0,
		cap:   minQueueLen,
		len:   0,
		mutex: sync.RWMutex{},
	}
}

func (q *Queue) growCap() {
	oldCap := q.cap
	doubleCap := oldCap + oldCap
	newCap := oldCap
	const threshold = 256
	if oldCap < threshold {
		newCap = doubleCap
	} else {
		newCap = oldCap + (oldCap+3*threshold)/4
	}
	tmp := make([]any, newCap)
	l, r := q.begin, q.end
	for l != r {
		tmp = append(tmp, q.data[l])
		l = (l + 1) % oldCap
	}
	q.begin = 0
	q.end = oldCap
	q.cap = newCap
	q.data = tmp
}

func (q *Queue) reduceCap() {
	oldCap := q.cap
	newCap := Max(minQueueLen, int(oldCap)/2)
	tmp := make([]any, newCap)
	l, r := q.begin, q.end
	for l != r {
		tmp = append(tmp, q.data[l])
		l = (l + 1) % q.cap
	}
	q.data = tmp
	q.begin = 0
	q.end = q.len
	q.cap = uint(newCap)
}
func (q *Queue) Size() uint {
	q.mutex.RLock()
	defer q.mutex.RUnlock()
	return q.len
}

func (q *Queue) Clear() {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	q.begin, q.end, q.cap, q.len = 0, 0, minQueueLen, 0
	q.data = make([]any, minQueueLen)
}

func (q *Queue) Empty() bool {
	q.mutex.RLock()
	defer q.mutex.RUnlock()
	return q.Size() <= 0
}

func (q *Queue) Push(v any) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	if q.len == q.cap {
		q.growCap()
	}
	q.data[q.end] = v
	q.end = (q.end + 1) % q.cap
	q.len++
}

func (q *Queue) Pop() error {
	q.mutex.Lock()
	q.mutex.Unlock()
	if q.Empty() {
		return errors.QueueEmptyError
	}
	q.begin = (q.begin + 1) % q.cap
	q.len--
	if q.cap > minQueueLen && q.len+q.len < q.cap {
		q.reduceCap()
	}
	return nil
}

func (q *Queue) Front() (any, error) {
	q.mutex.RLock()
	defer q.mutex.RUnlock()
	if q.Empty() {
		return nil, errors.QueueEmptyError
	}
	return q.data[q.begin], nil
}
