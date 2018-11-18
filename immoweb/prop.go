package immoweb

import "time"

type immowebPropertyInfo struct {
	ID              int       `json:"id"`
	PublisherID     int       `json:"publisherId"`
	PubDate         time.Time `json:"pubDate"`
	Photos          []string  `json:"photos"`
	PicturesResized []struct {
		LastModificationDate time.Time `json:"lastModificationDate"`
		Ordering             int       `json:"ordering"`
		PictureFormatLarge   string    `json:"pictureFormatLarge"`
		PictureFormatMedium  string    `json:"pictureFormatMedium"`
		PictureFormatSmall   string    `json:"pictureFormatSmall"`
		PictureOrientation   string    `json:"pictureOrientation"`
	} `json:"picturesResized"`
	Flags struct {
		DefaultOrderingScore int       `json:"defaultOrderingScore"`
		PriceOrderingScore   float64   `json:"priceOrderingScore"`
		DateOrderingScore    time.Time `json:"dateOrderingScore"`
		Pricem2OrderScore    float64   `json:"pricem2orderScore"`
		AdQualityScore       float64   `json:"adQualityScore"`
	} `json:"flags"`
	AdType string `json:"adType"`
	Energy struct {
		EnergyConsumption                           float64     `json:"energyConsumption"`
		EPCScore                                    string      `json:"EPCScore"`
		CO2Emission                                 interface{} `json:"CO2emission"`
		EPCreferenceNumber                          string      `json:"EPCreferenceNumber"`
		Heating                                     string      `json:"Heating"`
		DoubleGlazing                               bool        `json:"doubleGlazing"`
		InspectionReportOfTheElectricalInstallation string      `json:"inspectionReportOfTheElectricalInstallation"`
		EpcreferenceNumber                          string      `json:"epcreferenceNumber"`
	} `json:"energy"`
	Media []struct {
		Type    string `json:"type"`
		Content struct {
			Value string `json:"value"`
		} `json:"content"`
	} `json:"media"`
	MainType     string `json:"mainType"`
	Subtype      string `json:"subtype"`
	Geographical struct {
		POIs []struct {
			PoiName  string  `json:"poiName"`
			Distance float64 `json:"distance"`
		} `json:"POIs"`
		Geo struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"geo"`
		GeoPoint struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"geoPoint"`
		HasSeaView bool `json:"hasSeaView"`
		Postal     struct {
			Country      string `json:"country"`
			Province     string `json:"province"`
			District     string `json:"district"`
			Locality     string `json:"locality"`
			Postalcode   string `json:"postalcode"`
			Street       string `json:"street"`
			Number       string `json:"number"`
			Box          string `json:"box"`
			BuildingName string `json:"buildingName"`
			Region       string `json:"region"`
		} `json:"postal"`
		PostalAddr struct {
			Country      string `json:"country"`
			Province     string `json:"province"`
			District     string `json:"district"`
			Locality     string `json:"locality"`
			Postalcode   string `json:"postalcode"`
			Street       string `json:"street"`
			Number       string `json:"number"`
			Box          string `json:"box"`
			BuildingName string `json:"buildingName"`
			Region       string `json:"region"`
		} `json:"postalAddr"`
	} `json:"geographical"`
	BuyRent            string `json:"buyRent"`
	TransactionSubtype string `json:"transactionSubtype"`
	ActivationStatus   string `json:"activationStatus"`
	General            struct {
		NetFloorArea           int    `json:"netFloorArea"`
		FrontageNumber         int    `json:"frontageNumber"`
		InsideParkingNumber    int    `json:"insideParkingNumber"`
		OutsideParkingNumber   int    `json:"outsideParkingNumber"`
		BuildingCondition      string `json:"buildingCondition"`
		SurroundingEnvironment string `json:"surroundingEnvironment"`
		YearOfConstruction     int    `json:"yearOfConstruction"`
		Availability           string `json:"availability"`
	} `json:"general"`
	Interior struct {
		Bedrooms       int    `json:"bedrooms"`
		BathroomNumber int    `json:"bathroomNumber"`
		ToiletsNumber  int    `json:"toiletsNumber"`
		KitchenSetup   string `json:"kitchenSetup"`
		CellarProperty bool   `json:"cellarProperty"`
	} `json:"interior"`
	BuildingRegulation struct {
		PlanningPermissionObtained                string `json:"planningPermissionObtained"`
		AsBuiltPlan                               string `json:"asBuiltPlan"`
		SubdivisionPermit                         string `json:"subdivisionPermit"`
		PossiblePriorityPurchaseRight             string `json:"possiblePriorityPurchaseRight"`
		ProceedingsForBreachOfPlanningRegulations string `json:"proceedingsForBreachOfPlanningRegulations"`
		LatestLandUseDesignation                  string `json:"latestLandUseDesignation"`
		FloodZoneInfo                             string `json:"floodZoneInfo"`
	} `json:"buildingRegulation"`
	Financial struct {
		VatProfile      string `json:"vatProfile"`
		SalePrice       int    `json:"salePrice"`
		CadastralIncome int    `json:"cadastralIncome"`
		PriceBasedOn    bool   `json:"priceBasedOn"`
		IsMinimumBid    bool   `json:"isMinimumBid"`
		IsPriceBased    bool   `json:"isPriceBased"`
	} `json:"financial"`
	SoldInfo struct {
		Sold      bool   `json:"sold"`
		SoldLabel string `json:"soldLabel"`
	} `json:"soldInfo"`
	ContactInfo struct {
		Phone      string `json:"phone"`
		Email      string `json:"email"`
		Mobile     string `json:"mobile"`
		Idclient   int    `json:"idclient"`
		ClientType string `json:"clientType"`
		Title      string `json:"title"`
		Location   struct {
			Geo struct {
				Latitude  float64 `json:"latitude"`
				Longitude float64 `json:"longitude"`
			} `json:"geo"`
			GeoPoint struct {
				Latitude  float64 `json:"latitude"`
				Longitude float64 `json:"longitude"`
			} `json:"geoPoint"`
			Postal struct {
				Country    string `json:"country"`
				Province   string `json:"province"`
				District   string `json:"district"`
				Locality   string `json:"locality"`
				Postalcode string `json:"postalcode"`
				Street     string `json:"street"`
				Number     string `json:"number"`
				Box        string `json:"box"`
				Region     string `json:"region"`
			} `json:"postal"`
			PostalAddr struct {
				Country    string `json:"country"`
				Province   string `json:"province"`
				District   string `json:"district"`
				Locality   string `json:"locality"`
				Postalcode string `json:"postalcode"`
				Street     string `json:"street"`
				Number     string `json:"number"`
				Box        string `json:"box"`
				Region     string `json:"region"`
			} `json:"postalAddr"`
		} `json:"location"`
		Logo                    string `json:"logo"`
		Www                     string `json:"www"`
		IpiNo                   string `json:"ipiNo"`
		ReferenceWithinSoftware string `json:"referenceWithinSoftware"`
		ReferenceWithinAgency   string `json:"referenceWithinAgency"`
		Responsible             string `json:"responsible"`
	} `json:"contactInfo"`
	BookmarksCount int `json:"bookmarksCount"`
	ViewsCount     int `json:"viewsCount"`
	Land           struct {
		SurfaceAreaOfPlot int  `json:"surfaceAreaOfPlot"`
		Terrace           bool `json:"terrace"`
	} `json:"land"`
}
