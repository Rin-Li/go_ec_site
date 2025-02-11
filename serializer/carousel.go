package serializer

import "gin-mall-tmp/model"

type Carousel struct {
    Id        uint   `json:"id"`
    ImgPath   string `json:"img_path"`
    ProductId uint   `json:"product_id"`
    CreateAt  int64  `json:"create_at"`
}

func BuildCarousel(item *model.Carousel) Carousel {
    return Carousel{
        Id:        item.ID,
        ImgPath:   item.ImgPath,
        ProductId: item.ProductId,
        CreateAt:  item.CreatedAt.Unix(),
    }
}

func BuildCarousels(items []model.Carousel) (Carousels []Carousel){
	for _, item := range items{
		carousel := BuildCarousel(&item)
		Carousels = append(Carousels, carousel)
	}
	return Carousels
}


