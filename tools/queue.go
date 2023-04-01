package tools

import (
	"sync"
)

type node struct {
	val   any
	id    int64
	next  *node
	front *node
}

type Queue struct {
	data  *node //数据头指针
	end   *node
	mutex sync.RWMutex //并发控制锁
}

func NewQueue() *Queue {
	n := new(node)
	return &Queue{
		data:  n,
		end:   n,
		mutex: sync.RWMutex{},
	}
}

func (q *Queue) Add(message any, id int64) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	q.end.val = message
	q.end.id = id
	q.end.next = new(node)
	q.end.next.front = q.end
	q.end = q.end.next
}

func (q *Queue) Delete(id int64) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	for idx := q.data; idx != q.end; idx = idx.next {
		if idx.id == id {
			if idx.front != nil {
				idx.front.next = idx.next
				idx.next.front = idx.front
			} else {
				idx.next.front = nil
				q.data = idx.next
			}
			break
		}
	}
}

func (q *Queue) GetUnreadSlice(id int64) []any {
	q.mutex.RLock()
	defer q.mutex.RUnlock()
	var s []any
	for idx := q.data; idx != q.end; idx = idx.next {
		if idx.id > id {
			s = append(s, idx.val)
		}
	}
	return s
}
