package ostrovok

// нужно пофиксить код
//func main() {
//	rates, err := readConversionRates()
//
//	if err != nil {
//		log.Fatalf("readConversionRates() failed: %s", err)
//	}
//	// background task update conversion rates
//	go func() {
//		for {
//			time.Sleep(time.Minute)
//			rates, err = readConversionRates()
//			if err != nil {
//				log.Printf("readConversionRates() failed: %s", err)
//			}
//			rates = newrates
//		}
//	}()
//
//	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//		// parse 'from' and 'value' from query params
//		from := "RUB"
//		val := 140.0
//
//		rate, ok := rates[from]
//		if !ok {
//			http.NotFound(w, r)
//			return
//		}
//
//		convVal := val / rate
//
//		fmt.Fprint(w, convVal)
//	})
//
//	if err := http.ListenAndServe(":8080", nil); err != nil {
//		log.Fatal(err)
//	}
//}
//
//// readConversionRates reads rates from a file or an external service (relatively long running function).
//func readConversionRates() (map[string]float64, error) {
//	// resp, err := http.Get("https://example.org/conv-rates")
//	time.Sleep(100 * time.Millisecond)
//
//	return map[string]float64{
//		"USD": 1.0,
//		"RUB": 70.0,
//		"EUR": 95.0,
//	}, nil
//}
