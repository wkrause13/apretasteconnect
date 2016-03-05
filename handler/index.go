package handler
import(
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	files,_ := ioutil.ReadDir("/root/go/src/apretasteconnect/public/app/active/")
	var newestDir os.FileInfo
	for _,file := range files {
		if newestDir == nil{
                       newestDir = file
		       continue
		}
		diff :=  newestDir.ModTime().Sub(file.ModTime())
		if diff.Seconds() < float64 (0){
			newestDir = file
		}
	}

	appIndex := "/root/go/src/apretasteconnect/public/app/revisions/" + newestDir.Name() + "/index.html"
	body, _ := ioutil.ReadFile(appIndex)
	fmt.Fprint(w, string(body))
}
