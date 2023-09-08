package tools

import (
	"errors"
	"sync"
	"time"
)

/*
* Snowflake
*
* 1                                               42           52             64
* +-----------------------------------------------+------------+---------------+
* | timestamp(ms)                                 | workerId   | sequence      |
* +-----------------------------------------------+------------+---------------+
* | 0000000000 0000000000 0000000000 0000000000 0 | 0000000000 | 0000000000 00 |
* +-----------------------------------------------+------------+---------------+
* -1 << sequenceBits 就是将所有位都为1的数左移12位，得到 111111111111000000000000。 000000000000111111111111，即十进制的 4095。
* 1. 41位时间截(毫秒级)，注意这是时间截的差值（当前时间截 - 开始时间截)。可以使用约70年: (1L << 41) / (1000L * 60 * 60 * 24 * 365) = 69
* 2. 10位数据机器位，可以部署在1024个节点
* 3. 12位序列，毫秒内的计数，同一机器，同一时间截并发4096个序号
 */

const (
	startTimeStamp = int64(1483228800000)             //开始时间截 (2017-01-01),设置为一个固定的起始时间。
	workerIdBits   = uint(10)                         //机器id所占的位数，用于标识不同的机器。
	sequenceBits   = uint(12)                         //序列所占的位数，用于保证同一毫秒内生成的ID是唯一的。
	workerIdMax    = int64(-1 ^ (-1 << workerIdBits)) //支持的最大机器id数量
	sequenceMask   = int64(-1 ^ (-1 << sequenceBits)) //序列掩码，用于限制序列的范围。
	workerIdShift  = sequenceBits                     //机器id左移位数
	timestampShift = sequenceBits + workerIdBits      //时间戳左移位数
)

// A Snowflake struct holds the basic information needed for a snowflake generator worker
type Snowflake struct {
	sync.Mutex
	timestamp int64
	workerId  int64
	sequence  int64
}

// NewSnowflake 创建一个新的雪花工作节点（Snowflake Worker）
func NewSnowflake(workerId int64) (*Snowflake, error) {
	if workerId < 0 || workerId > workerIdMax {
		return nil, errors.New("workerId must be between 0 and 1023")
	}
	return &Snowflake{
		timestamp: 0,
		workerId:  workerId,
		sequence:  0,
	}, nil
}

// Generate creates and returns a unique snowflake ID
func (s *Snowflake) Generate() int64 {
	s.Lock()
	// 毫秒级别时间戳
	now := time.Now().UnixNano() / 1000000
	// 检查当前时间戳 now 是否与 Snowflake 实例中记录的时间戳 s.timestamp 相等。
	// 如果相等，表示在同一毫秒内多次调用生成ID的操作。
	if s.timestamp == now {
		// 位与运算 & 与 sequenceMask 进行按位截断，以保证序列号不超出位数限制。 1 & 111111111111 的结果是 000000000001
		s.sequence = (s.sequence + 1) & sequenceMask
		// 序列号为0，表示当前毫秒内的ID已经用完了。在这种情况下，需要等待下一毫秒再次生成ID，以避免序列号重复。
		if s.sequence == 0 {
			// 直到now>s.timestamp跳出
			for now <= s.timestamp {
				now = time.Now().UnixNano() / 1000000
			}
		}
	} else {
		s.sequence = 0
	}

	s.timestamp = now
	// 这一行代码使用位运算将时间戳、机器ID和序列号组合成一个64位的雪花ID。
	r := int64((now-startTimeStamp)<<timestampShift | (s.workerId << workerIdShift) | (s.sequence))

	s.Unlock()
	return r
}
