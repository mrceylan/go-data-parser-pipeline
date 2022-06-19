package main

type Drug struct {
	ProductNdc        string `json:"product_ndc"`
	GenericName       string `json:"generic_name"`
	LabelerName       string `json:"labeler_name"`
	BrandName         string `json:"brand_name"`
	ActiveIngredients []struct {
		Name     string `json:"name"`
		Strength string `json:"strength"`
	} `json:"active_ingredients"`
	Finished  bool `json:"finished"`
	Packaging []struct {
		PackageNdc         string `json:"package_ndc"`
		Description        string `json:"description"`
		MarketingStartDate string `json:"marketing_start_date"`
		Sample             bool   `json:"sample"`
	} `json:"packaging"`
	ListingExpirationDate string `json:"listing_expiration_date"`
	Openfda               struct {
		ManufacturerName   []string `json:"manufacturer_name"`
		Rxcui              []string `json:"rxcui"`
		SplSetID           []string `json:"spl_set_id"`
		IsOriginalPackager []bool   `json:"is_original_packager"`
		Upc                []string `json:"upc"`
		Unii               []string `json:"unii"`
	} `json:"openfda"`
	MarketingCategory  string   `json:"marketing_category"`
	DosageForm         string   `json:"dosage_form"`
	SplID              string   `json:"spl_id"`
	ProductType        string   `json:"product_type"`
	Route              []string `json:"route"`
	MarketingStartDate string   `json:"marketing_start_date"`
	ProductID          string   `json:"product_id"`
	ApplicationNumber  string   `json:"application_number"`
	BrandNameBase      string   `json:"brand_name_base"`
}
