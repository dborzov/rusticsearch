package main

type SearchItem struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type SearchPageItem struct {
	Id     string   `json:"id"`
	Name   string   `json:"name"`
	Price  string   `json:"price"`
	Images []string `json:"images"`
}

type SearchResult struct {
	Results []SearchItem `json:"products"`
}

type SearchPageResult struct {
	Results []SearchPageItem `json:"products"`
}

func (this SearchItem) Convert2SearchPageItem() SearchPageItem {
	output := SearchPageItem{}
	output.Id = this.Id
	output.Name = this.Name
	output.Price = "10.00"
	output.Images = []string{"https://www.filepicker.io/api/file/vVk61RqQqaD1cDvHLp1w"}
	return output
}
