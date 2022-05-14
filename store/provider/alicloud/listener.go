package alicloud

import (
	"fmt"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/schollz/progressbar/v3"
)

type listener struct {
	bar     *progressbar.ProgressBar
	startAt time.Time
}

func NewListener() oss.ProgressListener {
	return &listener{}
}

func (l *listener) ProgressChanged(event *oss.ProgressEvent) {
	switch event.EventType {
	case oss.TransferStartedEvent: //开始上传事件
		l.bar = progressbar.DefaultBytes(
			event.TotalBytes,
			"文件上传中",
		)
		l.startAt = time.Now()
	case oss.TransferDataEvent: //正在上传事件
		l.bar.Add64(event.RwBytes)
	case oss.TransferCompletedEvent: //上传完成事件
		fmt.Printf("\n上传完成: 耗时%d秒\n", int(time.Since(l.startAt).Seconds()))
	case oss.TransferFailedEvent: //上传失败事件
		fmt.Printf("\n上传失败\n")
	}
}
