package webkit

import (
	"fmt"
	"runtime"
	"time"

	"github.com/gotk3/gotk3/gtk"
	"github.com/sourcegraph/webloop"
)

// Render a webpage and resturn static html
func Render(url string) ([]byte, error) {
	gtk.Init(nil)
	go func() {
		runtime.LockOSThread()
		gtk.Main()
	}()

	ctx := webloop.New()
	view := ctx.NewView()
	//defer view.Close()
	view.Open(url)
	//err := view.Wait()
	//if err != nil {
	//	return nil, err
	//}

	time.Sleep(20 * time.Second) // to do fix me for better
	fmt.Println("wait done")
	result, err := view.EvaluateJavaScript("document.documentElement.outerHTML")
	if err != nil {
		return nil, err
	}

	html := result.(string)

	fmt.Println(html)

	return []byte(html), nil
}
