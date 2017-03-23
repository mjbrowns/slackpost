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
  "errors"
	"github.com/spf13/cobra"
)

// addfieldCmd represents the addfield command
var addfieldCmd = &cobra.Command{
	Use:   "addfield <title> <value>",
	Short: "Add fields to the current attachment",
	Long: `Add fields to the current attachment.

	Message atachments can have multiple fields (see slack message documentation).
	This function adds a few field to the most recently added attachment.`,
	RunE: do_addfield,
}

func do_addfield (cmd *cobra.Command, args []string) error {
	// TODO: Work your own magic here
	Log("\nStarting function addfield")
	if len(args) <2 {
		return errors.New("ERROR: title and value arguments are required")
	}
	var msg slackMessage
	ReadMessageFile(&msg,MessageFile)

  var myField slackAttachmentFields
	myField.Title=args[0]
	myField.Value=args[1]
  if short,_:=cmd.Flags().GetBool("short"); short {
		myField.Short="true"
	} else {
		myField.Short="false"
	}

	lasta := len(msg.Payload.Attachments)-1

  acnt := len(msg.Payload.Attachments[lasta].Fields)

	slice := make([]slackAttachmentFields,acnt+1,acnt+1)
	if acnt > 0 {
		for i := range msg.Payload.Attachments[lasta].Fields {
			slice[i] = msg.Payload.Attachments[lasta].Fields[i]
		}
	}
	slice[acnt]=myField
	msg.Payload.Attachments[lasta].Fields = slice

	DumpMessage(msg)
	WriteMessageFile(&msg,MessageFile)
  return nil
}

func init() {
	RootCmd.AddCommand(addfieldCmd)
	addfieldCmd.Flags().BoolP("short","s",false,"If set, indicates that the title/value pair is short enough to be displayed vertically with other fields")
}
