package crawler

// PageData stores the crawled information
type PageData struct {
    URL         string   `json:"url"`
    Title       string   `json:"title"`
    Description string   `json:"description"`
    Slokas      []string `json:"slokas"`  // NEW field for storing shlokas
}

