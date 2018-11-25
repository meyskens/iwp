package immoweb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/alecthomas/template"
)

type immowebSearchResponse struct {
	Adims   []immowebAdims        `json:"adims"`
	Results []immowebSearchResult `json:"results"`
}

type immowebAdims struct {
	AdimID            int    `json:"adimId"`
	PublisherID       int    `json:"publisherId"`
	Title             string `json:"title"`
	Image             string `json:"image"`
	Phone             string `json:"phone"`
	PhoneCallingTime  string `json:"phoneCallingTime"`
	MobileCallingTime string `json:"mobileCallingTime"`
	PostalAddress     struct {
		Country    string `json:"country"`
		Province   string `json:"province"`
		District   string `json:"district"`
		Locality   string `json:"locality"`
		Postalcode string `json:"postalcode"`
		Street     string `json:"street"`
		Number     string `json:"number"`
		Box        string `json:"box"`
		Region     string `json:"region"`
	} `json:"postalAddress"`
	IsCertifiedPartner bool `json:"isCertifiedPartner"`
}

type immowebSearchResult struct {
	ID          int       `json:"id"`
	PublisherID int       `json:"publisherId"`
	PubDate     time.Time `json:"pubDate,omitempty"`
	Photos      []string  `json:"photos"`
	Flags       struct {
		New                  bool      `json:"new"`
		DefaultOrderingScore int       `json:"defaultOrderingScore"`
		PriceOrderingScore   float64   `json:"priceOrderingScore"`
		DateOrderingScore    time.Time `json:"dateOrderingScore"`
		Pricem2OrderScore    float64   `json:"pricem2orderScore"`
		AdQualityScore       float64   `json:"adQualityScore"`
	} `json:"flags"`
	AdType string `json:"adType"`
	Energy struct {
		EnergyConsumption  interface{} `json:"energyConsumption"`
		EPCScore           string      `json:"EPCScore"`
		CO2Emission        interface{} `json:"CO2emission"`
		EPCreferenceNumber interface{} `json:"EPCreferenceNumber"`
	} `json:"energy,omitempty"`
	AgencyLogo string `json:"agencyLogo,omitempty"`
	Info       struct {
		Idestate int    `json:"idestate"`
		BuyRent  string `json:"buyRent"`
		MainType string `json:"mainType"`
		Subtype  string `json:"subtype"`
		Location struct {
			Postal struct {
				Country    string `json:"country"`
				Province   string `json:"province"`
				District   string `json:"district"`
				Locality   string `json:"locality"`
				Postalcode string `json:"postalcode"`
				Region     string `json:"region"`
			} `json:"postal"`
			PostalAddr struct {
				Country    string `json:"country"`
				Province   string `json:"province"`
				District   string `json:"district"`
				Locality   string `json:"locality"`
				Postalcode string `json:"postalcode"`
				Region     string `json:"region"`
			} `json:"postalAddr"`
		} `json:"location"`
		Price struct {
			VatProfile      string `json:"vatProfile"`
			SalePrice       int    `json:"salePrice"`
			CadastralIncome int    `json:"cadastralIncome"`
			PriceBasedOn    bool   `json:"priceBasedOn"`
			IsMinimumBid    bool   `json:"isMinimumBid"`
			IsPriceBased    bool   `json:"isPriceBased"`
		} `json:"price"`
		Rooms int `json:"rooms"`
		M2    int `json:"m2"`
	} `json:"info"`
	Sold            bool `json:"sold,omitempty"`
	EstateGroupInfo struct {
		ProjectTag string `json:"projectTag"`
		Prices     struct {
			Min float64 `json:"min"`
			Max float64 `json:"max"`
		} `json:"prices"`
		Area struct {
			Min int `json:"min"`
			Max int `json:"max"`
		} `json:"area"`
		RoomCount struct {
			Min int `json:"min"`
			Max int `json:"max"`
		} `json:"roomCount"`
		EstateGroupItems []struct {
			ID          int       `json:"id"`
			PublisherID int       `json:"publisherId"`
			PubDate     time.Time `json:"pubDate"`
			Flags       struct {
				DefaultOrderingScore int       `json:"defaultOrderingScore"`
				PriceOrderingScore   float64   `json:"priceOrderingScore"`
				DateOrderingScore    time.Time `json:"dateOrderingScore"`
				Pricem2OrderScore    float64   `json:"pricem2orderScore"`
				AdQualityScore       float64   `json:"adQualityScore"`
			} `json:"flags"`
			AdType     string  `json:"adType"`
			AgencyLogo string  `json:"agencyLogo"`
			Mnemonic   string  `json:"mnemonic"`
			Area       int     `json:"area"`
			RoomCount  int     `json:"roomCount"`
			Floor      int     `json:"floor"`
			Price      float64 `json:"price"`
			Sold       string  `json:"sold"`
		} `json:"estateGroupItems"`
		UnitCount        int    `json:"unitCount"`
		UnitsDisplayMode string `json:"unitsDisplayMode"`
	} `json:"estateGroupInfo,omitempty"`
	ExpiryDate time.Time `json:"expiryDate,omitempty"`
}

// Scraper is the scraper interface implementation for immoweb
type Scraper struct {
}

// GetProperties gets the link to the property with the given parameters
func (s *Scraper) GetProperties(saleType, propertyType, zip string, sellers map[string]bool) ([]string, error) {
	buyRent := "BUY"
	if saleType == "rent" {
		buyRent = "RENT"
	}
	houseAppertment := "HOUSE"
	if propertyType == "appartment" {
		houseAppertment = "APARTMENT"
	}

	client := &http.Client{}
	q := `{"mainType":"` + houseAppertment + `","location":{"postalcodes":["` + zip + `"]},"land":{"min":0},"subtypes":["HOUSE","APARTMENT_BLOCK","VILLA","CASTLE","TOWN_HOUSE","MIXED_USE_BUILDING","FARMHOUSE","BUNGALOW","MANSION","PAVILION","COUNTRY_COTTAGE","CHALET","EXCEPTIONAL_PROPERTY","OTHER_PROPERTY","MANOR_HOUSE","DUPLEX","APARTMENT","PENTHOUSE","FLAT_STUDIO","TRIPLEX","KOT","LOFT","GROUND_FLOOR","SERVICE_FLAT"],"buyRent":"` + buyRent + `","parkingPlaces":{"min":0},"frontageNumber":{"min":0},"rooms":{"min":0},"price":{},"buildingConditions":["AS_NEW","JUST_RENOVATED","GOOD","TO_REFURBISH","TO_RENOVATE","TO_RESTORE"],"area":{"min":0}}`
	req, err := http.NewRequest("GET", "https://api.immoweb.be/rest/estate?q="+template.URLQueryEscaper(q), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "immoweb/3.6.2.1 (iPhone; iOS 12.0; scale/2.0)")
	req.Header.Set("Accept", `application/vnd.immoweb.estateresults+json; profile="/schema/estateresults-r1#"`)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	data := immowebSearchResponse{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	out := []string{}

	for _, prop := range data.Results {
		if prop.Info.Location.PostalAddr.Postalcode == zip {
			info, err := s.getPropertyDetail(q, strconv.FormatInt(int64(prop.ID), 10))
			if err != nil {
				continue
			}
			//fmt.Println(info.MainType)

			if info.ContactInfo.ClientType == "AGE" && sellers["agency"] {
				out = append(out, fmt.Sprintf("https://www.immoweb.be/nl/zoekertje/huis/te-koop/undefined/0000/id%d", prop.ID))
			}
			if info.ContactInfo.ClientType == "NOT" && sellers["notary"] {
				out = append(out, fmt.Sprintf("https://www.immoweb.be/nl/zoekertje/huis/te-koop/undefined/0000/id%d", prop.ID))
			}
			if info.ContactInfo.ClientType == "PAR" && sellers["individual"] {
				out = append(out, fmt.Sprintf("https://www.immoweb.be/nl/zoekertje/huis/te-koop/undefined/0000/id%d", prop.ID))
			}
		}
	}

	return out, nil
}

func (s *Scraper) getPropertyDetail(q, id string) (*immowebPropertyInfo, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.immoweb.be/rest/estate/"+id+"?q="+template.URLQueryEscaper(q), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "immoweb/3.6.2.1 (iPhone; iOS 12.0; scale/2.0)")
	req.Header.Set("Accept", `application/vnd.immoweb.estate+json; profile="/schema/estate-r1#"`)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	data := immowebPropertyInfo{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
