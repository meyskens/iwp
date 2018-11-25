package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/meyskens/iwp/immoweb"
	"github.com/skratchdot/open-golang/open"
	"github.com/zserge/webview"
)

// Scraper is an interface to an immo site scraper
type Scraper interface {
	GetProperties(saleType, propertyType, zip string, sellers map[string]bool) ([]string, error)
}

var sites = []Scraper{
	//&heylen.Scraper{},
	&immoweb.Scraper{},
	//&zimmo.Scraper{},
	//&immovlan.Scraper{},
}

var w webview.WebView

func main() {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()
	go func() {
		// load in bindata
		http.Handle("/api/lookup", http.HandlerFunc(getForZip))
		http.Handle("/api/lookup/csv", http.HandlerFunc(getForZipToCSV))
		http.Handle("/api/open", http.HandlerFunc(openURL))
		http.Handle("/", http.FileServer(&assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: AssetInfo, Prefix: "/frontend/"}))
		log.Fatal(http.Serve(ln, nil))
	}()
	w = webview.New(webview.Settings{Debug: true, Title: "IWP", Width: 800, Height: 600, Resizable: true, URL: "http://" + ln.Addr().String()})
	w.Dispatch(func() {
		// Inject controller
		w.Bind("goapp", &App{})
	})
	w.Run()
}

func getForZip(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	saleType := r.URL.Query().Get("saletype")
	zip := r.URL.Query().Get("postcode")
	agency := r.URL.Query().Get("agency")
	individual := r.URL.Query().Get("individual")
	notary := r.URL.Query().Get("notary")

	sellers := map[string]bool{}
	if agency != "" {
		sellers["agency"] = true
	}
	if individual != "" {
		sellers["individual"] = true
	}
	if notary != "" {
		sellers["notary"] = true
	}

	all := []string{}
	for _, site := range sites {
		urls, _ := site.GetProperties(saleType, "house", zip, sellers)
		all = append(all, urls...)

		urls, _ = site.GetProperties(saleType, "appartment", zip, sellers)
		all = append(all, urls...)
	}

	out, _ := json.Marshal(all)
	w.Write(out)
}

func getForZipToCSV(w http.ResponseWriter, r *http.Request) {
	saleType := r.URL.Query().Get("saletype")
	zip := r.URL.Query().Get("postcode")
	agency := r.URL.Query().Get("agency")
	individual := r.URL.Query().Get("individual")
	notary := r.URL.Query().Get("notary")

	sellers := map[string]bool{}
	if agency != "" {
		sellers["agency"] = true
	}
	if individual != "" {
		sellers["individual"] = true
	}
	if notary != "" {
		sellers["notary"] = true
	}

	all := []string{}
	for _, site := range sites {
		urls, _ := site.GetProperties(saleType, "house", zip, sellers)
		all = append(all, urls...)

		urls, _ = site.GetProperties(saleType, "appartment", zip, sellers)
		all = append(all, urls...)
	}

	f, err := os.Create(zip + "-" + time.Now().Format("2006-01-02T15-04") + ".csv")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	f.WriteString(strings.Join(all, "\n"))
	f.Close()
	w.WriteHeader(http.StatusOK)
}

func openURL(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	open.Run(r.Form.Get("url"))
	w.Write([]byte("ok"))
}

type App struct{}

func (a *App) OpenURL(url string) {
	fmt.Println(open.Run(url))
}
