- Теория та же, что есть в навигации.
- Задача вот чуть другая, но по сути в некоторых местах похожа на уже имеющующуюся (замечания свои пометил "Замечание" - интервьюрер соглашался с ними)

``` Go
package some_package

import (
  "context"
)

type Storage interface {
  Products(productsIds []int) (map[int]*ProductInfo, error)
  Prices(ctx context.Context, productIds []int) (map[int]*ProductPrice, error)
}

/* domain package */

type ProductInfo struct {
  ProductId int
  CreatedAt uint8 // ЗАМЕЧАНИЕ!. либо time либо int32 для unix времени например
}

type ProductPrice struct {
  ProductId int
  Price     float64 // ЗАМЕЧАНИЕ!. цену лучше не во флоутах хранить, лучше в целочисленных например int, ну домножая там копейки, чтобы целое число было
}

type Product struct {
  Id    int
  Info  *ProductInfo
  Price *ProductPrice
}

/* transport package */

type GetRequest struct {
  Products []*Product
}

type Server struct {
  storage Storage
}

// ЗАМЕЧАНИЕ! добавить например вейтгруппу для горутин
// ЗАМЕЧАНИЕ! прокинуть лучше ctx в s.storage.Products (т.е. сигнатуру поменять предложить) и s.storage.Prices
// ЗАМЕЧАНИЕ! ошибки в канал например класть лучше, а потом их класть в условно слайс и возвращать, форматируя слайс в fmt.Errorf например
// ЗАМЕЧАНИЕ! выделить заранее память в result
// ЗАМЕЧАНИЕ! логика обработки продуктов неверна, по отдельности Info и Price кладётся, т.е. не в один объект
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
  
  result := []*Product()
  for _, id := range request.Ids {
    if _, ok := productInfoS[int(id)]; ok {
      info := productInfoS[int(id)]
      result = append(result, &Product{id: int(id), Info: info})
    }
  
    if price, ok := productPrices[int(id)]; ok {
      result = append(result, &Product{id: int(id), Price: price})
    }
  
  return &GetResponse{Products: result}, nil
}

func int32ToIntSlice(slice []int32) []int {
  /* Функция переводит []int32 в []int */
}
```