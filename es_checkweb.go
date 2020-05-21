package main

import (
  "golang.org/x/net/html"
  "log"
  "strings"
  "fmt"
  "os"
  "time"
  "io/ioutil"
	"net/http"
  "strconv"
  "errors"

  "github.com/aws/aws-sdk-go/aws"
//  "github.com/aws/aws-sdk-go/aws/credentials"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/ses"

)

const txt_url = "http://wdc.nict.go.jp/IONO/fxEs/latest-fxEs.html"    // Target URL

const i_mhz = 9.1   // mininum of fxEs value by NICT

const con_baseDir = "/home/ubuntu/go/src/github.com/Tsukiyama-do/es_notice/"
// const con_baseDir = "/home/yuichi/github.com/Tsukiyama-do/laserbeam/src/es_notice/"
// const con_baseDir = "/home/yuichi01/go/src/es_notice/"

var hlog *log.Logger    // custom logger object


func init() {
// To generate log file and put log


}

func sendSESEmail(from string, to string, title string, body string) error {

  // Specify profile for config and region for requests
  awsSession := session.Must(session.NewSessionWithOptions(session.Options{
  	 Config: aws.Config{Region: aws.String("us-east-1")},
  	 Profile: "yuichi01",
  }))

/*
  awsSession := session.New(&aws.Config{
    Region:      aws.String("us-east-1"),
    Credentials: credentials.NewStaticCredentials("AWS_ACCESS_KEY_ID", "AWS_SECRET_KEY", ""),
  })
*/

  svc := ses.New(awsSession)
  input := &ses.SendEmailInput{
    Destination: &ses.Destination{
      ToAddresses: []*string{
        aws.String(to),
      },
    },
    Message: &ses.Message{
      Body: &ses.Body{
        Text: &ses.Content{
          Charset: aws.String("UTF-8"),
          Data:    aws.String(body),
        },
      },
      Subject: &ses.Content{
        Charset: aws.String("UTF-8"),
        Data:    aws.String(title),
      },
    },
    Source: aws.String(from),
  }
  _, err := svc.SendEmail(input)
  if err != nil {
    return errors.New(err.Error())
  }
  return nil
}

func espo_text () ([]byte, error ) {

  hlog.Println("espo_text started!")

	res, err := http.Get(txt_url)    // http request : GET

	if err != nil {
    hlog.Printf("%v \n",err)
    return nil, err
  }
  defer res.Body.Close()

  hlog.Println("espo_text Get ended successfully!")


	robots, err := ioutil.ReadAll(res.Body)
  if err != nil {
    hlog.Printf("%v \n",err)
    return nil, err
  }

  hlog.Println("espo_text ReadAll ended successfully!")

	return robots, nil

}

type espo_data struct  {    //  espo status struct
  s_datetime string
  f_fxes float64
}


func main () {

  // To generate log file and put log


    t := time.Now()    // to get current time
    st_log := con_baseDir + "log_" + t.Format("20060102-150405") + ".log"   // log's path and file
    outfile, err := os.Create(st_log)      // create file pointer
    if err != nil {      log.Fatalf("error opening file: %v", err)  }

    defer outfile.Close()

    hlog = log.New(outfile, "[Espo]", log.Ltime|log.Lshortfile)   //  define logger format.

    // Program started
    hlog.Println("main() started!")

//  Below is test data
//  s := `<p>Links:</p><ul><li><a href="foo">Foo</a><li><a href="/bar/baz">BarBaz</a></ul>`

/*
  s_html := `<!-- saved from url=(0048)http://wdc.nict.go.jp/IONO/fxEs/latest-fxEs.html -->
  <html><head><meta http-equiv="Content-Type" content="text/html; charset=UTF-8"><title>NICT Latest fxEs</title><meta http-equiv="refresh" content="900"></head><body bgcolor="c3c3c3"><font color="black">
  <table border=""><tbody><tr><th colspan="5" bgcolor="pink">NICT 'fxEs' values for the past two hours
  </th></tr><tr><td bgcolor="e3e3e3">Undefined</td><td bgcolor="aqua">Under7[MHz]</td><td bgcolor="yellow">Over7[MHz]</td><td bgcolor="hotpink">Over8[MHz]</td><td bgcolor="red">Over9[MHz]</td></tr></tbody></table>
  <br><table border=""><tbody><tr><td bgcolor="pink"><p align="right">Date Time (JST)</p></td><td bgcolor="pink">Okinawa</td><td bgcolor="pink">Yamagawa</td><td bgcolor="pink">Kokubunji</td><td bgcolor="pink">Wakkanai</td>
  </tr><tr><td bgcolor="pink">2019/06/01 18:15</td>
  <td bgcolor="aqua"><p align="right"> 3.2</p></td>
  <td bgcolor="red"><p align="right">11.4</p></td>
  <td bgcolor="red"><p align="right">14.7</p></td>
  <td bgcolor="aqua"><p align="right"> 4.2</p></td>
  </tr><tr><td bgcolor="pink">2019/06/01 18:30</td>
  <td bgcolor="aqua"><p align="right"> 3.0</p></td>
  <td bgcolor="red"><p align="right">16.2</p></td>
  <td bgcolor="aqua"><p align="right">3.3</p></td>
  <td bgcolor="aqua"><p align="right"> 3.9</p></td>
  </tr><tr><td bgcolor="pink">2019/06/01 18:45</td>
  <td bgcolor="aqua"><p align="right"> 2.5</p></td>
  <td bgcolor="red"><p align="right">10.4</p></td>
  <td bgcolor="red"><p align="right">13.6</p></td>
  <td bgcolor="aqua"><p align="right"> 4.9</p></td>
  </tr><tr><td bgcolor="pink">2019/06/01 19:00</td>
  <td bgcolor="aqua"><p align="right"> 2.7</p></td>
  <td bgcolor="aqua"><p align="right"> 6.0</p></td>
  <td bgcolor="red"><p align="right">15.1</p></td>
  <td bgcolor="aqua"><p align="right"> 5.9</p></td>
  </tr><tr><td bgcolor="pink">2019/06/01 19:15</td>
  <td bgcolor="aqua"><p align="right"> 2.7</p></td>
  <td bgcolor="yellow"><p align="right"> 7.7</p></td>
  <td bgcolor="yellow"><p align="right">7.1</p></td>
  <td bgcolor="e3e3e3"><p align="right">----</p></td>
  </tr><tr><td bgcolor="pink">2019/06/01 19:30</td>
  <td bgcolor="aqua"><p align="right"> 2.8</p></td>
  <td bgcolor="aqua"><p align="right"> 6.7</p></td>
  <td bgcolor="red"><p align="right">14.5</p></td>
  <td bgcolor="hotpink"><p align="right"> 8.8</p></td>
  </tr><tr><td bgcolor="pink">2019/06/01 19:45</td>
  <td bgcolor="aqua"><p align="right"> 3.1</p></td>
  <td bgcolor="yellow"><p align="right"> 7.0</p></td>
  <td bgcolor="red"><p align="right">14.9</p></td>
  <td bgcolor="red"><p align="right"> 9.2</p></td>
  </tr><tr><td bgcolor="pink">2019/06/16 08:25</td>
  <td bgcolor="e3e3e3"><p align="right">----</p></td>
  <td bgcolor="aqua"><p align="right"> 5.9</p></td>
  <td bgcolor="aqua"><p align="right">7.1</p></td>
  <td bgcolor="red"><p align="right">10.6</p></td>
  </tr></tbody></table>
  URL for cell phone ... <a href="http://wdc.nict.go.jp/IONO/fxEs/latest-fxEs-k.html">(Type a)</a>
  <a href="http://wdc.nict.go.jp/IONO/fxEs/latest-fxEs-i.html">(Type i)</a>

  </font></body></html>`

*/


  // Get html from internet
  hlog.Println("Get html started!")

  s_html, err := espo_text()
  if err != nil {   hlog.Fatal(err)  }   // identify error

  // Parse html
  hlog.Println("Parse html started!")
  doc, err := html.Parse(strings.NewReader(string(s_html)))
  if err != nil {   hlog.Fatal(err)  }   // identify error

  // Pick Kokubunji value of fxes of the last check
  hlog.Println("Picking the Kokubunji value of fxes just started!")
  var f func(*html.Node)
  var  i_col int    // what row?
  var f_last bool = false   // to indicate the last row or not
  var tmp_data espo_data

  f = func(n *html.Node) {
    if n.Type == html.TextNode  {
                byear := []byte(n.Data)
                if string(byear[:4]) == "2019" || string(byear[:4]) == "2020" ||string(byear[:4]) == "2021"  {

                    layout := "2006/01/02 15:04"
                    loc, _ := time.LoadLocation("Asia/Tokyo")

                    t, _ := time.ParseInLocation(layout, n.Data,loc)
                    t1 := time.Now()
                    minutes := 19
                  if t1.Sub(t) <  time.Duration(minutes)*time.Minute  {  //  to identify last line or not.
                      hlog.Println("Recent check datetime : ", t)   // => "2003-04-18 00:00:00 +0000 UTC"
                      tmp_data.s_datetime = t.Format(layout)     // set time with string format
                      f_last = true    // flag => identify as last line.
                  }
                }
                if f_last == true && i_col == 3 {

                      s_wk := string(byear)
                      hlog.Println("Recent fxes of Kokubunji : ",s_wk)
                      s_fxes_tmp := strings.TrimSpace(fmt.Sprintf("%s",s_wk))
                      f_fxes_tmp, err := strconv.ParseFloat(s_fxes_tmp,32)
                      if err != nil  {   hlog.Fatal(err)  }   // identify error
                      if f_fxes_tmp >= i_mhz {      // fxes over i_mhz will be set.
                        tmp_data.f_fxes = f_fxes_tmp
                      }
                      f_last = false  // reset the flag
                }
    }
    if f_last == true {
        if n.Type == html.ElementNode && n.Data == "td" {
            i_col++    // used to identify what row it is.
        }
    }

    for c := n.FirstChild; c != nil; c = c.NextSibling {
        f(c)
    }
  }
  f(doc)

  s_fxes := fmt.Sprintf("%f",tmp_data.f_fxes)
  hlog.Println("Result is " + s_fxes + " value on " + tmp_data.s_datetime + ". " )


  // Start aws ese process
  hlog.Println("Start aws ese process. " )

  if tmp_data.f_fxes > 0 {  // if value of fxes is set.

    from := "noreply@www.jj1pow.com"
    to := "yuichisierra_kusc_fm@willcom.com"
    title := "Espo notice from AWS"
    body := "Espo occurred with fxes value of " + fmt.Sprintf("%f",tmp_data.f_fxes) + " on " + tmp_data.s_datetime
    err := sendSESEmail(from, to, title, body)
    if err != nil {
  	   hlog.Println("Error at sending the message.")
       hlog.Fatal(err)
    }
    hlog.Println("The message was successfully sent to ", to)

  }

  hlog.Println("End aws ese process. " )
  // End aws ese process


}
