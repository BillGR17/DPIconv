package main

import (
  "github.com/andlabs/ui"
  "strconv"
  "fmt"
)

func main() {
  err := ui.Main(func() {
    input := ui.NewEntry()
    result := ui.NewLabel("")
    box := ui.NewVerticalBox()
    box.Append(ui.NewLabel("Enter Font 96DPI Size:"), false)
    box.Append(input, false)
    box.Append(result, false)
    window := ui.NewWindow("WebFontFix", 150, 80, false)
    window.SetMargined(true)
    window.SetChild(box)
    input.OnChanged(func(*ui.Entry) {
      if len(input.Text())!=0{
        ti,e:=strconv.ParseInt(input.Text(),0,64)
        fix:=(float64(ti)/72)*96
        res:=fmt.Sprintf("%f",fix)
        if e!=nil{
          result.SetText("Please Enter A Number");
        }else{
          result.SetText(res);
        }
      }else{
        result.SetText("Please Enter A Number");
      }
    })
    window.OnClosing(func(*ui.Window) bool {
     ui.Quit()
     return true
    })
    window.Show()
  })
  if err != nil {
    panic(err)
  }
}
