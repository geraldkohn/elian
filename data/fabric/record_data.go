package fabric

import "sync"

/**
存放record数据
注: 下一版需要引进链码, 这种存放数据的方式会被废弃
*/

type tempDataRepo struct {
	data map[string]string //k, v对
	mu   sync.RWMutex      //防止高并发造成恐慌, 需要读写锁
}

func newTempData() tempDataRepo {
	return tempDataRepo{
		data: make(map[string]string),
	}
}