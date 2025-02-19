package techtask

// Вводные:

type Data struct {
	ID      int
	Payload map[string]interface{}
}

type Reader interface {
	Read() []*Data
}

type Processor interface {
	Process(Data) ([]*Data, error)
}

type Writer interface {
	Write([]*Data)
}

// Необходимо имплементировать интерфейс Manager так, чтобы он:
// - перманентно и синхронно принимал данные из Reader
// - обрабатывал все полученные данные на каждом из списка Processor
// - в случае отсутствия ошибок при обработке, полученные в процессе обработки данные передавал в Writer
// - при возникновении ошибки при обработке, ничего не отправлял во Writer

type Manager interface {
	Manage() // blocking
}

// Реализация:
