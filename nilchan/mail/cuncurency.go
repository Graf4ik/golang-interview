package mail

//type User struct {
//	Name string
//}
//
//
//// main менять нельзя
//func main() {
//	fmt.Println(Do(context.Background(), []User{{"aaa"}, {"bbb"}, {"ccc"}, {"ddd"}, {"eeee"}}))
//}
//
//func fetchByName(ctx context.Context, userName string) (int, error) {
//	// Тут происходит сетевой подход, который по userName возвращает userID
//
//	time.Sleep(10 * time.Millisecond) // имитация сетевого подхода
//	return rand.Int() & 100000, nil
//}
//
//
//// Все изменения долдеы производится в данной функции
//func Do(ctx context.Context, users []User) (map[string]int, error) {
//	collected := make(map[string]int)
//
//	// TODO необходимо реализовать конкурентные запросы для каждого юзера и вернуть результат
//}
