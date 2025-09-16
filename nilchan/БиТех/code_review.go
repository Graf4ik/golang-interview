package БиТех

import "context"

// нужно сделать код-ревью
type Service struct {
	db     repository.Repository
	logger *zap.Logger
	cache  cache.Cache
}

func New(db repository.Repository, logger *zap.Logger, cache cache.Cache) *Service {
	return &Service{
		db:     db,
		logger: logger,
		cache:  cache,
	}
}

// GetProduct возвращает сущность по ее ID
func (s *Service) GetProduct(_ context.Context, id int64) (*models.Product, error) {
	product, err := s.cache.GetProduct(id)
	if err != nil {
		return nil, err
	}
	// Что если cache.GetProduct возвращает err, но это ErrCacheMiss? Сейчас ты прерываешь выполнение, хотя можно пойти в БД.
	/*
		✅ Решение: проверять ошибку более точно:
		product, err := s.cache.GetProduct(id)
		if err != nil && !errors.Is(err, cache.ErrCacheMiss) {
		    return nil, err
		}
	*/

	// Также желательно после чтения из БД — прокешировать:
	if product != nil {
		return product, nil //  ignore error or log
	}

	product, err := s.db.GetProduct(context.Background(), id) // Использование context.Background() в обработке запроса — антипаттерн.
	// Ты теряешь таймауты, дедлайны, отмену и trace/span из вызова GetProduct.
	if err != nil {
		//  Хорошая практика — логировать ошибки хотя бы уровня Debug или Warn.
		//    s.logger.Warn("failed to get product from cache", zap.Error(err))
		return nil, err
	}
	return product, nil
}

// Update обновляет информацию о сущности
// Нейминг: Update != CreateProduct
func (s Service) CreateProduct(c context.Context, p models.Product) error {
	_, err := s.db.CreateProduct(c, p)
	if err != nil {
		return err
	}
	return nil
}

/*
Метод получает значение структуры, а не указатель.
Это означает, что ты создаешь копию Service, и если в будущем тебе нужно будет мутировать поля (например, счётчики, мьютексы и т.п.) — это не сработает.
*/
