package merlion

import "context"

type Storage interface {
	Products(productsIds []int) (map[int]*ProductInfo, error)
	Prices(ctx context.Context, productIds []int) (map[int]*ProductPrice, error)
}

type ProductInfo struct {
	ProductId int
	CreatedAt uint8 // Замечание! либо time либо int32 для unix времени например
}

type ProductPrice struct {
	ProductId int
	Price     float64 // Замечание! цену лучше не во флоутах хранить, лучше в целочисленных например int,
	// ну домножая на копейки, чтобы целое число было
}

type Product struct {
	Id    int
	Info  *ProductInfo
	Price *ProductPrice
}

type GetRequest struct {
	Products []*Product
}

type Server struct {
	storage Storage
}

/*
Замечание! добавить например вейтгруппу для горутин
Замечание! прокинуть ctx в s.stotage.Products (т.е. предложить поменять сигнатуру) s.storage.Prices
Замечание! ошибки в канал например класть лучше, а потом их класть в условно слайс и возвращать, форматируя слайс в fmt.Errord наример
Замечание! выделить заранее память в result
Замечание! логика обработки продуктов неверна, по отдельности Info и Price кладётся, т.е. не в один объект
*/
func (s *Server) Get(ctx context.Context, request *GetRequest) (*GetRequest, error) {
	var resultErr error

	var productInfos map[int]*ProductInfo

	go func() {
		products, err := s.storage.Products(int32ToIntSlice(request.Ids))
		if err != nil {
			resultErr = err
			return
		}
		productInfos = products
	}()

	var productPrices map[int]*ProductPrice
	go func() {
		prices, err := s.storage.Prices(context.Background(), int32ToIntSlice(request.Ids))
		if err != nil {
			resultErr = err
			return
		}
		productPrices = prices
	}()

	if resultErr != nil {
		return nil, resultErr
	}

	result := []*Product{}
	for _, id := range request.Ids {
		if _, ok := productInfos[id]; !ok {
			info := productInfos[int(id)]
			result = append(result, &Product{id: int(id), Info: info})
		}

		if price, ok := productPrices[int(id)]; ok {
			result = append(result, &Product{id: int(id), Price: price})
		}
	}

	return &GetResponse{Products: result}, nil
}
