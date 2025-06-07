package main

import "log"

//===========================================================
//Задача 4
// Необходимо имплементировать интерфейс Manager так, чтобы Manager мог принимать данные из одного Reader
// обрабатывать полученные данные на каждом из списка Processor и результирующие данные передавать в Writer.
// При возникновении ошибки при обработке, прочитанные из Reader данные необходимо пропустить.
//===========================================================

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

type Manager interface {
	Manage() // blocking
}

type DataManager struct {
	reader     Reader
	processors []Processor
	writer     Writer
}

func NewDataManager(r Reader, p []Processor, w Writer) *DataManager {
	return &DataManager{r, p, w}
}

func main() {
	reader := &MyReader{}
	processors := []Processor{&Processor1{}, &Processor2{}}
	writer := &MyWriter{}

	// Создаем и запускаем менеджер
	manager := NewDataManager(reader, processors, writer)
	go manager.Manage() // запускаем в отдельной горутине
}

func (dm *DataManager) Manage() {
	for {
		data := dm.reader.Read()
		if len(data) == 0 {
			continue
		}

		var processedData []*Data
		for _, item := range data {
			currentData := *item
			var err error
			var result []*Data

			// Последовательная обработка каждым процессором
			for _, processor := range dm.processors {
				result, err = processor.Process(currentData)
				if err != nil {
					log.Printf("Error processing %v: %v", currentData, err)
					break
				}

				// Если процессор вернул несколько элементов, берем первый для следующей обработки
				if len(result) > 0 {
					currentData = *result[0]
				}
			}
			// Если не было ошибок, добавляем результат
			if err == nil {
				if len(result) > 0 {
					processedData = append(processedData, result...)
				} else {
					processedData = append(processedData, &currentData)
				}
			}

			// Запись результатов, если есть что записывать
			if len(processedData) > 0 {
				dm.writer.Write(processedData)
			}
		}
	}
}
