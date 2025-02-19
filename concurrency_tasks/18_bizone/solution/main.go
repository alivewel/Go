package techtask

import (
	"fmt"
)

// Manager отвечает за чтение данных, их обработку и запись
type manager struct {
	reader    Reader
	processors []Processor
	writer    Writer
}

// NewManager создает новый экземпляр Manager
func NewManager(reader Reader, processors []Processor, writer Writer) Manager {
	return &manager{
		reader:    reader,
		processors: processors,
		writer:    writer,
	}
}

// Manage запускает процесс обработки данных
func (m *manager) Manage() {
	for {
		dataBatch := m.reader.Read()
		if len(dataBatch) == 0 {
			continue
		}

		var processedData []*Data
		for _, data := range dataBatch {
			tempData := []*Data{data}
			for _, processor := range m.processors {
				var err error
				tempData, err = processor.Process(*data)
				if err != nil {
					fmt.Println("Error processing data:", err)
					tempData = nil
					break
				}
			}

			if tempData != nil {
				processedData = append(processedData, tempData...)
			}
		}

		if len(processedData) > 0 {
			m.writer.Write(processedData)
		}
	}
}
