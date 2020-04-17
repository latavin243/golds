package golds

/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-04-17 22:01:57
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-04-17 22:09:42
 */

type RequestEncoder struct{}

func NewRequestEncoder() *RequestEncoder {
	return new(RequestEncoder)
}

func (this *RequestEncoder) Encode() {
	// TODO(monitor1379)
}
