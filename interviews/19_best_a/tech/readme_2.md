**Мерлион техсобес**

**Вопросы:**
- Как выглядел флоу в команде? Когда задача решается завершённой?
- Как фичлидство происходило? Насколько сильно декомпозировал задачу? Были ли Фичи на несколько спринтов?
- Какая была самая большая фича, которая была? Были ли задачи более чем не спринт?
- Допустим приходит Тим лид и говорит нужна такая фича, но я понимаю что он для нее нужен новый микросервис или новое взаимодействие микросервисное или новая технология. Что делал?
- Кто деплоил?
- Зачем девопс, если CI настроен, если мы сами можем кнопочку нажимать?
- Деплоили сразу по факту?
- Были ли релизные даты?
- Ответственность разработчика на какой момент по задаче заканчивалась?
- В prod мы деплоили или девопс?
- Прошел реквест, код ревью прошел, кто мерджил в релизную ветку?
- Технические метрики: кол-во нагрузок. Кто за ними следил?
- Девопс смотрит и видит что в CPU упирается. Код норм написан. Нужно масштабироваться, что делать?

Написан некий код. Толи работает, толи нет, надо отрефакторить:

``` Go
type Storage interface {
  Products (productIds []string) ([]*Products, error)
  Compilations (ctx context.Context, productIds []string) ([]*Compilations, error)
}

/* domain package */
/
type Products struct {
  Id int
  CreatedAt uint8
}

type Compilations struct {
  ProductId int
  MinPrice float64
  MaxPrice float64
}

type ProductWithCompilations struct {
  Id CreatedAt
  int uint8
  Compilations []*Compilations
}


/* transport package */
type GetRequest struct {
  Ids []string
}

type GetResponse struct {
  ProductWithCompilations []*ProductWithCompilations
}

type Server struct {
  storage Storage
}

func (s *Server) SearchBywords (ctx context.Context, request *GetRequest) (*GetResponse, error) { 
  var errChan chan error
  var products []*Products
  go func() {
    result, err := s.storage.Products(request.Ids)
    if err != nil {
      errChan <- err
      return
    }
    
  products = result
  }()

  var compilations []*Compilations
  go func() {
    result, err := s.storage.Compilations (context.Background(), request.Ids)
    if err != nil {
      errChan<- err
      return
    }
  comlitations = result
  }()
  
  for err := range errChan {
    if err != nil {
      return nil, err
    }
  }
  
  result := []*ProductWithComplitation{}
  for _, product := range products {
    for  _, complitation := range complitation {
      if complitation.ProductId == product.id {
        result = append(result, &ProductWithComplitation{Id: product.Id, createdAt: product.CreatedAt, Complitations: []*Complitations{complitation}})
        
      }
    }
  }
  return &GetResponse{ProductWithComplitations: result}, nil
}
```
