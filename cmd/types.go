// this file holds datatypes used by the package

package cmd

type slackAttachmentFields struct {
  Title       string `json:"title"`
  Value       string `json:"value"`
  Short       string `json:"short"`
}

type slackAttachment struct {
  Fallback    string `json:"fallback"`
  Color       string `json:"color"`
  PreText     string `json:"pretext"`
  AuthorName  string `json:"author_name"`
  AuthorLink  string `json:"author_link"`
  AuthorIcon  string `json:"author_icon"`
  Title       string `json:"title"`
  TitleLink   string `json:"title_link"`
  Text        string `json:"text"`
  Fields      []slackAttachmentFields `json:"fields"`
  ImageUrl    string `json:"image_url"`
  ThumbUrl    string `json:"thumb_url"`
  Footer      string `json:"footer"`
  FooterIcon  string `json:"footer_icon"`
  TS          string `json:"ts"`
}


type slackPayload struct {
	Text string 									`json:"text"`
	User string 									`json:"username"`
	Icon string 									`json:"icon_emoji"`
	Iconlink string 							`json:"icon_url"`
	Channel string                `json:"channel"`
	Attachments []slackAttachment `json:"attachments"`
}

type slackMessage struct {
  Hook string 									`json:"hook"`
  Payload slackPayload          `json:"payload"`
}
