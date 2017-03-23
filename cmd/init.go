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
	"fmt"
  "errors"
	"github.com/spf13/cobra"
)

var MyMessage slackMessage

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init <webhook> <message>",
	Short: "Initializes the message object",
	Long: `Initializes the message object`,
	RunE: do_init,
}

func do_init(cmd *cobra.Command, args []string) error {
	Log ("\nInit function started")
	if len(args) <2 {
		return errors.New("ERROR: webhook and message arguments are required")
	}
	if cmd.Flags().Changed("icon") && cmd.Flags().Changed("icon-url") {
		return errors.New("ERROR: icon and icon-url options are mutually exclusive")
	}

	MyMessage.Hook=args[0]
	MyMessage.Payload.Text=args[1]

	if MyMessage.Payload.Icon != "" && string(MyMessage.Payload.Icon[0]) != ":" {
		MyMessage.Payload.Icon = fmt.Sprintf(":%s:",MyMessage.Payload.Icon)
	}
	if MyMessage.Payload.Iconlink != "" { MyMessage.Payload.Icon=""}

  DumpMessage(MyMessage)
	WriteMessageFile(&MyMessage,MessageFile)
	return nil
}

func init() {
	RootCmd.AddCommand(initCmd)
	initCmd.Flags().StringVarP(&MyMessage.Payload.User,"username","u","","User Name to associate with the message")
	initCmd.Flags().StringVarP(&MyMessage.Payload.Icon,"icon","i","","Icon name to post with the message")
	initCmd.Flags().StringVarP(&MyMessage.Payload.Iconlink,"icon-url","I","","URL of Icon to post with the message")
	initCmd.Flags().StringVarP(&MyMessage.Payload.Channel,"channel","c","","Name of channel to send to (ignored if passing a channel hook)")

}
