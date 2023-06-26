package main

import (
	"checklistCGREG/src/functions"
	"fmt"
	"io"
	"os"
	"net/http"
	"github.com/gin-gonic/gin"
	"os/exec"
	"runtime"
)

func main() {
	router := gin.Default()
	router.Static("/styles","./styles")
	router.Static("/assets","./src/assets")
	router.GET("/", Index)
	router.GET("/path", Path)
	
	url := "http://localhost:9090/"// Replace with your desired URL
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "linux":
		cmd = exec.Command("xdg-open", url)
		case "windows":
			cmd = exec.Command("cmd", "/c", "start", url)
			case "darwin":
				cmd = exec.Command("open", url)
				default:
					// Unsupported operating system
					return
	}
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		panic(err)
	}
	
	http.ListenAndServe(":9090", router)
	

	path := "/home/jhquihuiri7/Downloads"
	data := functions.WalkFunc(path)
	functions.GenerateExcel(data)
	//tmp := functions.ParseTemplate()
	//tmp.Execute(os.Stdout, nil)

}
func Index(c *gin.Context) {
	tmp := functions.ParseTemplate()
	tmp.Execute(c.Writer, nil)
	//c.Writer.WriteString("HOLA")
}
func Path (c *gin.Context){
	err := c.Request.ParseForm()
	if err != nil {
		// Handle error
		http.Error(c.Writer, "Failed to parse form", http.StatusBadRequest)
		return
	}
	// Access form values
	path := c.Request.Form.Get("path")
	fmt.Println(path)
	data := functions.WalkFunc(path)
	tempFile := functions.GenerateExcel(data)
	
	c.Writer.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Writer.Header().Set("Content-Disposition", "attachment; filename=example.xlsx")
	file, err := os.Open(tempFile)
	if err != nil {
		// Handle error
		http.Error(c.Writer, "Failed to open Excel file", http.StatusInternalServerError)
		fmt.Println(tempFile)
		return
		
	}
	defer file.Close()

	// Stream the file contents to the response writer
	_, err = file.Seek(0, 0) // Reset file pointer to the beginning
	if err != nil {
		// Handle error
		http.Error(c.Writer, "Failed to read Excel file", http.StatusInternalServerError)
		return
	}
	
	_, err = io.Copy(c.Writer, file)
	if err != nil {
		// Handle error
		http.Error(c.Writer, "Failed to stream Excel file", http.StatusInternalServerError)
		return
	}	
	//c.Redirect(303,"/")
}
