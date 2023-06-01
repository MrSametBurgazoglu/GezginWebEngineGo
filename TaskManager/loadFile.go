package TaskManager

import "gezgin_web_engine/FileManager"

type LoadFileTask struct {
	fileUrl string
	data    []byte
}

func (receiver *LoadFileTask) ExecuteTask() {
	receiver.data = FileManager.LoadFile(receiver.fileUrl)
}
