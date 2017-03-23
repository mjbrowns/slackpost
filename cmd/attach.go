// Copyright © 2017 Mitch Brown <mitch@mjbrowns.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"github.com/spf13/cobra"
  "fmt"
	"errors"
)

var myAttachment slackAttachment

// attachCmd represents the attach command
var attachCmd = &cobra.Command{
	Use:   "attach <text>",
	Short: "Attach a text string to the message",
	Long: `Attach a text string to the message.
	This function will add a string as an attachment to the message.
	There can be unlimited attachments to a message, and each attachment has many
	options available to customize how the attachment will appear.

	Example:  attach "$(<logfile.txt)"`,
	RunE: do_attach,
}

func do_attach (cmd *cobra.Command, args []string) error {
	// TODO: Work your own magic here
	Log("\nAttach function started")
	if len(args) <1 {
		return errors.New("ERROR: text string argument is required")
	}
	var msg slackMessage
	ReadMessageFile(&msg,MessageFile)

	myAttachment.Text=args[0]

	acnt := len(msg.Payload.Attachments)
	Log(fmt.Sprintf("Old Attachment Count: %d",acnt))
	slice := make([]slackAttachment,acnt+1,acnt+1)
	if acnt > 0 {
		for i := range msg.Payload.Attachments {
			slice[i] = msg.Payload.Attachments[i]
		}
	}
	slice[acnt]=myAttachment
	msg.Payload.Attachments = slice
	DumpMessage(msg)
	WriteMessageFile(&msg,MessageFile)
	return nil
}

func init() {
	RootCmd.AddCommand(attachCmd)
	attachCmd.Flags().StringVarP(&myAttachment.Fallback,"fallback","B","","Fallback string to use for clients that can't display the message")
	attachCmd.Flags().StringVarP(&myAttachment.Color,"color","c","","Name of color to use with this attachment")
	attachCmd.Flags().StringVarP(&myAttachment.PreText,"pretext","p","","String to display above the message attachment block")
	attachCmd.Flags().StringVarP(&myAttachment.AuthorName,"author","a","","Name of author")
	attachCmd.Flags().StringVarP(&myAttachment.AuthorLink,"author-link","A","","URL to hyperlink to the author's name")
	attachCmd.Flags().StringVarP(&myAttachment.AuthorIcon,"author-icon","I","","URL to a 16x16px icon to display by the author's name")
	attachCmd.Flags().StringVarP(&myAttachment.Title,"title","t","","String to display as the title of the attachment.  Displays as larger, bolder text")
	attachCmd.Flags().StringVarP(&myAttachment.TitleLink,"title-link","T","","URL to hyperlink to the attachment title.")
	attachCmd.Flags().StringVarP(&myAttachment.ImageUrl,"image","i","","URL of image to display inside the attachment.")
	attachCmd.Flags().StringVarP(&myAttachment.ThumbUrl,"thumb","b","","URL of a thumbnail to display to the top-right of the attachment.")
	attachCmd.Flags().StringVarP(&myAttachment.Footer,"footer","f","","Text that appears as a footer to the attachment.")
	attachCmd.Flags().StringVarP(&myAttachment.FooterIcon,"footer-icon","F","","URL of an icon to display with the footer.")
	attachCmd.Flags().StringVarP(&myAttachment.TS,"timestamp","d","","Epoch-time date/time stamp for the attachment.")
}
