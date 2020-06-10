package monitor

// FootLockerCalendar refer to the footlocker calendar
type FootLockerCalendar struct {
	ID        string `json:"id"`
	Image     string `json:"image"`
	DeepLinks []struct {
		Locale string `json:"locale"`
		Link   string `json:"link"`
	} `json:"deepLinks"`
	ReleaseDatetime string `json:"releaseDatetime"`
	Colors          string `json:"colors"`
	Name            string `json:"name"`
	Brand           string `json:"brand"`
	BrandImage      string `json:"brandImage"`
	BaseID          string `json:"baseID"`
	HasStock        string `json:"hasStock"`
	HasImage        string `json:"hasImage"`
	Gender          string `json:"gender"`
	ManufacturerSku string `json:"manufacturerSku"`
}

type FootLockerSize struct {
	Content string `json:"content"`
}

// Start Discord Webhook struct
type Thumbnail struct {
	URL string `json:"url"`
}

type Fields struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}

type Footer struct {
	Text    string `json:"text"`
	IconURL string `json:"icon_url"`
}

type Embeds struct {
	Title     string    `json:"title"`
	URL       string    `json:"url"`
	Color     int       `json:"color"`
	Timestamp string    `json:"timestamp"`
	Thumbnail Thumbnail `json:"thumbnail"`
	Fields    []Fields  `json:"fields"`
	Footer    Footer    `json:"footer"`
}

type Webhook struct {
	Username  string   `json:"username"`
	AvatarURL string   `json:"avatar_url"`
	Embeds    []Embeds `json:"embeds"`
}

// End Discord Webhook struct

type ShopifyProducts struct {
	Products []struct {
		ID          int64       `json:"id"`
		Title       string      `json:"title"`
		Handle      string      `json:"handle"`
		BodyHTML    interface{} `json:"body_html"`
		PublishedAt string      `json:"published_at"`
		CreatedAt   string      `json:"created_at"`
		UpdatedAt   string      `json:"updated_at"`
		Vendor      string      `json:"vendor"`
		ProductType string      `json:"product_type"`
		Tags        []string    `json:"tags"`
		Variants    []struct {
			ID               int64       `json:"id"`
			Title            string      `json:"title"`
			Option1          string      `json:"option1"`
			Option2          interface{} `json:"option2"`
			Option3          interface{} `json:"option3"`
			Sku              string      `json:"sku"`
			RequiresShipping bool        `json:"requires_shipping"`
			Taxable          bool        `json:"taxable"`
			FeaturedImage    interface{} `json:"featured_image"`
			Available        bool        `json:"available"`
			Price            string      `json:"price"`
			Grams            int         `json:"grams"`
			CompareAtPrice   interface{} `json:"compare_at_price"`
			Position         int         `json:"position"`
			ProductID        int64       `json:"product_id"`
			CreatedAt        string      `json:"created_at"`
			UpdatedAt        string      `json:"updated_at"`
		} `json:"variants"`
		Images []struct {
			ID         int64         `json:"id"`
			CreatedAt  string        `json:"created_at"`
			Position   int           `json:"position"`
			UpdatedAt  string        `json:"updated_at"`
			ProductID  int64         `json:"product_id"`
			VariantIds []interface{} `json:"variant_ids"`
			Src        string        `json:"src"`
			Width      int           `json:"width"`
			Height     int           `json:"height"`
		} `json:"images"`
		Options []struct {
			Name     string   `json:"name"`
			Position int      `json:"position"`
			Values   []string `json:"values"`
		} `json:"options"`
	} `json:"products"`
}

type SupremeProducts struct {
	UniqueImageURLPrefixes []interface{} `json:"unique_image_url_prefixes"`
	ProductsAndCategories  struct {
		Bags []struct {
			Name          string `json:"name"`
			ID            int    `json:"id"`
			ImageURL      string `json:"image_url"`
			ImageURLHi    string `json:"image_url_hi"`
			Price         int    `json:"price"`
			SalePrice     int    `json:"sale_price"`
			NewItem       bool   `json:"new_item"`
			Position      int    `json:"position"`
			CategoryName  string `json:"category_name"`
			PriceEuro     int    `json:"price_euro"`
			SalePriceEuro int    `json:"sale_price_euro"`
		} `json:"Bags"`
		Pants []struct {
			Name          string `json:"name"`
			ID            int    `json:"id"`
			ImageURL      string `json:"image_url"`
			ImageURLHi    string `json:"image_url_hi"`
			Price         int    `json:"price"`
			SalePrice     int    `json:"sale_price"`
			NewItem       bool   `json:"new_item"`
			Position      int    `json:"position"`
			CategoryName  string `json:"category_name"`
			PriceEuro     int    `json:"price_euro"`
			SalePriceEuro int    `json:"sale_price_euro"`
		} `json:"Pants"`
		Accessories []struct {
			Name          string `json:"name"`
			ID            int    `json:"id"`
			ImageURL      string `json:"image_url"`
			ImageURLHi    string `json:"image_url_hi"`
			Price         int    `json:"price"`
			SalePrice     int    `json:"sale_price"`
			NewItem       bool   `json:"new_item"`
			Position      int    `json:"position"`
			CategoryName  string `json:"category_name"`
			PriceEuro     int    `json:"price_euro"`
			SalePriceEuro int    `json:"sale_price_euro"`
		} `json:"Accessories"`
		Skate []struct {
			Name          string `json:"name"`
			ID            int    `json:"id"`
			ImageURL      string `json:"image_url"`
			ImageURLHi    string `json:"image_url_hi"`
			Price         int    `json:"price"`
			SalePrice     int    `json:"sale_price"`
			NewItem       bool   `json:"new_item"`
			Position      int    `json:"position"`
			CategoryName  string `json:"category_name"`
			PriceEuro     int    `json:"price_euro"`
			SalePriceEuro int    `json:"sale_price_euro"`
		} `json:"Skate"`
		Shoes []struct {
			Name          string `json:"name"`
			ID            int    `json:"id"`
			ImageURL      string `json:"image_url"`
			ImageURLHi    string `json:"image_url_hi"`
			Price         int    `json:"price"`
			SalePrice     int    `json:"sale_price"`
			NewItem       bool   `json:"new_item"`
			Position      int    `json:"position"`
			CategoryName  string `json:"category_name"`
			PriceEuro     int    `json:"price_euro"`
			SalePriceEuro int    `json:"sale_price_euro"`
		} `json:"Shoes"`
		Hats []struct {
			Name          string `json:"name"`
			ID            int    `json:"id"`
			ImageURL      string `json:"image_url"`
			ImageURLHi    string `json:"image_url_hi"`
			Price         int    `json:"price"`
			SalePrice     int    `json:"sale_price"`
			NewItem       bool   `json:"new_item"`
			Position      int    `json:"position"`
			CategoryName  string `json:"category_name"`
			PriceEuro     int    `json:"price_euro"`
			SalePriceEuro int    `json:"sale_price_euro"`
		} `json:"Hats"`
		TopsSweaters []struct {
			Name          string `json:"name"`
			ID            int    `json:"id"`
			ImageURL      string `json:"image_url"`
			ImageURLHi    string `json:"image_url_hi"`
			Price         int    `json:"price"`
			SalePrice     int    `json:"sale_price"`
			NewItem       bool   `json:"new_item"`
			Position      int    `json:"position"`
			CategoryName  string `json:"category_name"`
			PriceEuro     int    `json:"price_euro"`
			SalePriceEuro int    `json:"sale_price_euro"`
		} `json:"Tops/Sweaters"`
		Jackets []struct {
			Name          string `json:"name"`
			ID            int    `json:"id"`
			ImageURL      string `json:"image_url"`
			ImageURLHi    string `json:"image_url_hi"`
			Price         int    `json:"price"`
			SalePrice     int    `json:"sale_price"`
			NewItem       bool   `json:"new_item"`
			Position      int    `json:"position"`
			CategoryName  string `json:"category_name"`
			PriceEuro     int    `json:"price_euro"`
			SalePriceEuro int    `json:"sale_price_euro"`
		} `json:"Jackets"`
		Sweatshirts []struct {
			Name          string `json:"name"`
			ID            int    `json:"id"`
			ImageURL      string `json:"image_url"`
			ImageURLHi    string `json:"image_url_hi"`
			Price         int    `json:"price"`
			SalePrice     int    `json:"sale_price"`
			NewItem       bool   `json:"new_item"`
			Position      int    `json:"position"`
			CategoryName  string `json:"category_name"`
			PriceEuro     int    `json:"price_euro"`
			SalePriceEuro int    `json:"sale_price_euro"`
		} `json:"Sweatshirts"`
		Shirts []struct {
			Name          string `json:"name"`
			ID            int    `json:"id"`
			ImageURL      string `json:"image_url"`
			ImageURLHi    string `json:"image_url_hi"`
			Price         int    `json:"price"`
			SalePrice     int    `json:"sale_price"`
			NewItem       bool   `json:"new_item"`
			Position      int    `json:"position"`
			CategoryName  string `json:"category_name"`
			PriceEuro     int    `json:"price_euro"`
			SalePriceEuro int    `json:"sale_price_euro"`
		} `json:"Shirts"`
		TShirts []struct {
			Name          string `json:"name"`
			ID            int    `json:"id"`
			ImageURL      string `json:"image_url"`
			ImageURLHi    string `json:"image_url_hi"`
			Price         int    `json:"price"`
			SalePrice     int    `json:"sale_price"`
			NewItem       bool   `json:"new_item"`
			Position      int    `json:"position"`
			CategoryName  string `json:"category_name"`
			PriceEuro     int    `json:"price_euro"`
			SalePriceEuro int    `json:"sale_price_euro"`
		} `json:"T-Shirts"`
		Shorts []struct {
			Name          string `json:"name"`
			ID            int    `json:"id"`
			ImageURL      string `json:"image_url"`
			ImageURLHi    string `json:"image_url_hi"`
			Price         int    `json:"price"`
			SalePrice     int    `json:"sale_price"`
			NewItem       bool   `json:"new_item"`
			Position      int    `json:"position"`
			CategoryName  string `json:"category_name"`
			PriceEuro     int    `json:"price_euro"`
			SalePriceEuro int    `json:"sale_price_euro"`
		} `json:"Shorts"`
		New []struct {
			Name          string `json:"name"`
			ID            int    `json:"id"`
			ImageURL      string `json:"image_url"`
			ImageURLHi    string `json:"image_url_hi"`
			Price         int    `json:"price"`
			SalePrice     int    `json:"sale_price"`
			NewItem       bool   `json:"new_item"`
			Position      int    `json:"position"`
			CategoryName  string `json:"category_name"`
			PriceEuro     int    `json:"price_euro"`
			SalePriceEuro int    `json:"sale_price_euro"`
		} `json:"new"`
	} `json:"products_and_categories"`
	ReleaseDate string `json:"release_date"`
	ReleaseWeek string `json:"release_week"`
}

type SupremeProduct struct {
	Styles []struct {
		ID                int         `json:"id"`
		Name              string      `json:"name"`
		Currency          string      `json:"currency"`
		Description       interface{} `json:"description"`
		ImageURL          string      `json:"image_url"`
		ImageURLHi        string      `json:"image_url_hi"`
		SwatchURL         string      `json:"swatch_url"`
		SwatchURLHi       string      `json:"swatch_url_hi"`
		MobileZoomedURL   string      `json:"mobile_zoomed_url"`
		MobileZoomedURLHi string      `json:"mobile_zoomed_url_hi"`
		BiggerZoomedURL   string      `json:"bigger_zoomed_url"`
		Sizes             []struct {
			Name       string `json:"name"`
			ID         int    `json:"id"`
			StockLevel int    `json:"stock_level"`
		} `json:"sizes"`
		Additional []struct {
			SwatchURL       string `json:"swatch_url"`
			SwatchURLHi     string `json:"swatch_url_hi"`
			ImageURL        string `json:"image_url"`
			ImageURLHi      string `json:"image_url_hi"`
			ZoomedURL       string `json:"zoomed_url"`
			ZoomedURLHi     string `json:"zoomed_url_hi"`
			BiggerZoomedURL string `json:"bigger_zoomed_url"`
		} `json:"additional"`
	} `json:"styles"`
	Description             string `json:"description"`
	CanAddStyles            bool   `json:"can_add_styles"`
	CanBuyMultiple          bool   `json:"can_buy_multiple"`
	Ino                     string `json:"ino"`
	CodBlocked              bool   `json:"cod_blocked"`
	CanadaBlocked           bool   `json:"canada_blocked"`
	PurchasableQty          int    `json:"purchasable_qty"`
	NewItem                 bool   `json:"new_item"`
	Apparel                 bool   `json:"apparel"`
	Handling                int    `json:"handling"`
	NoFreeShipping          bool   `json:"no_free_shipping"`
	CanBuyMultipleWithLimit int    `json:"can_buy_multiple_with_limit"`
	NonEuBlocked            *bool  `json:"non_eu_blocked"`
	RussiaBlocked           *bool  `json:"russia_blocked"`
}
