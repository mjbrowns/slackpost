package cmd

import (
  "encoding/json"
  "io/ioutil"
  "os"
  "fmt"
  "net/http"
  "net/url"
)

func Log(txt string) {
  if ! Verbose { return }
  fmt.Println(txt)
}

func Die(txt string) {
  fmt.Println(txt)
  os.Exit(255)
}

func Check(err error,txt string) {
	if err == nil { return }
	fmt.Println(txt)
	Die(err.Error())
}

func RemoveFile(fn string) {
  if _, err := os.Stat(fn); os.IsNotExist(err) {
		return
	}
  if err := os.Remove(fn); err != nil {
    Log(fmt.Sprintf("Error deleting message file %s",fn))
    }
}

func WriteMessageFile(msg *slackMessage,fn string) {
	b,err:=json.Marshal(msg)
	Check(err,"Error decoding JSON message")
	d := []byte(b)
	err = ioutil.WriteFile(fn,d,0600)
	Check(err,"Error writing JSON object")
	if Verbose { fmt.Printf("JSON Object written to %s\n",fn)}
	return
}

func ReadMessageFile(msg *slackMessage,fn string) {
  if _, err := os.Stat(fn); os.IsNotExist(err) {
		Die("Error: Message has not yet been initialized")
	}
	dat, err  := ioutil.ReadFile(fn)
  Check (err,"Error reading JSON object")
	err = json.Unmarshal(dat,msg)
  Check (err,"Error decoding JSON object")
	return
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil { return true, nil }
	if os.IsNotExist(err) { return false, nil }
	return true, err
}

func DumpMessage(msg slackMessage) {
  if ! Verbose { return }
  fmt.Println("\nMessage Data   : START")
  fmt.Printf("  Webhook      : %s\n",msg.Hook)
  fmt.Printf("  Message      : %s\n",msg.Payload.Text)
  fmt.Printf("  User         : %s\n",msg.Payload.User)
  fmt.Printf("  Icon         : %s\n",msg.Payload.Icon)
  fmt.Printf("  Icon-url     : %s\n",msg.Payload.Iconlink)
  fmt.Printf("  Channel      : %s\n",msg.Payload.Channel)
  acnt := len(msg.Payload.Attachments)
  if acnt > 0 {
    for i := range msg.Payload.Attachments {
    fmt.Printf("  Attachment   # %d\n",i)
    DumpAttach(msg.Payload.Attachments[i])
    }
  } else {
    fmt.Println("  Attachments  : (none)")
  }
  fmt.Println("Message Data   : END\n")
}

func DumpAttach(a slackAttachment) {
  if ! Verbose { return }
  fmt.Printf("    Fallback   : %s\n",a.Fallback)
  fmt.Printf("    Color      : %s\n",a.Color)
  fmt.Printf("    PreText    : %s\n",a.PreText)
  fmt.Printf("    AuthorName : %s\n",a.AuthorName)
  fmt.Printf("    AuthorLink : %s\n",a.AuthorLink)
  fmt.Printf("    AuthorIcon : %s\n",a.AuthorIcon)
  fmt.Printf("    Title      : %s\n",a.Title)
  fmt.Printf("    TitleLink  : %s\n",a.TitleLink)
  fmt.Printf("    Text       : %s\n",a.Text)
  fmt.Printf("    ImageUrl   : %s\n",a.ImageUrl)
  fmt.Printf("    ThumbUrl   : %s\n",a.ThumbUrl)
  fmt.Printf("    Footer     : %s\n",a.Footer)
  fmt.Printf("    FooterIcon : %s\n",a.FooterIcon)
  fmt.Printf("    TimeStamp  : %s\n",a.TS)
  fcnt := len(a.Fields)
  if fcnt > 0 {
    for i := range a.Fields {
      fmt.Printf("    Field      # %d\n",i)
      DumpFields(a.Fields[i])
    }
  } else {
    fmt.Println("    Fields     : (none)")
  }
}

func DumpFields(f slackAttachmentFields) {
  if ! Verbose { return }
  fmt.Printf("      Title    : %s\n",f.Title)
  fmt.Printf("      Value    : %s\n",f.Value)
  fmt.Printf("      Short    : %s\n",f.Short)
}

func SlackSend(msg *slackMessage) {
	params, _ := json.Marshal(msg.Payload)

	resp, _ := http.PostForm(
		msg.Hook,
		url.Values{"payload": {string(params)}},
	)

	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
  if string(body) != "ok" {Die(fmt.Sprintf("Error: Slack post returned error: %s",string(body)))}
	return
}
