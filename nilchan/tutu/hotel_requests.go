package tutu

// Есть поток данных, в виде идентификаторов отелей, для каждого отеля нужно сделать поисковый запрос (запрос выполняется
// минимум 500ms) и отправить результаты в другой поток.

//type SearchResult struct {
//	HotelID int
//}
//
//func main() {
//	dataCh := make(chan int)
//
//	go func() {
//		for i := 0; i <= 10; i++ {
//			dataCh <- i
//		}
//		defer close(dataCh)
//	}()
//}
//
//func search(hotelID int) SearchResult {
//	time.Sleep(time.Millisecond * 500)
//	return SearchResult{
//		HotelID: hotelID,
//	}
//}
