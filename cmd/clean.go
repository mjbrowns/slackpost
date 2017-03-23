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
	"io/ioutil"
	"regexp"
	"github.com/spf13/cobra"
)

// cleanCmd represents the clean command
var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Cleans up the message object",
	Long: `Cleans up the temporary file(s) used to create the message object.
	This function is normally performed automatically after sending the object,
	but in some cases you may need to clean up the message manually.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		Log("\nStarting function clean")
		if all,_:=cmd.Flags().GetBool("all"); all {
      files, err := ioutil.ReadDir(WorkDir)
			if err != nil { Die(err.Error()) }
			r,_ := regexp.Compile(fmt.Sprintf("^%s(.*)",WorkPrefix))
			for _,file := range files {
				if m:=r.FindStringSubmatch(file.Name()); m !=nil {
					Log(fmt.Sprintf("Deleting Message ID: %s (%s)",m[1],m[0]))
					if force,_:=cmd.Flags().GetBool("force"); force {
						RemoveFile(fmt.Sprintf("%s/%s",WorkDir,m[0]))
					}
				}
			}
		} else {
			Log(fmt.Sprintf("Deleting Message ID: %s (%s)",MessageID,MessageFile))
			RemoveFile(MessageFile)
		}
	},
}

func init() {
	RootCmd.AddCommand(cleanCmd)
	cleanCmd.Flags().BoolP("all","a",false,"If set, will clean up any detected messages")
	cleanCmd.Flags().Bool("force",false,"When using -a|--all this is required to actually delete files")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cleanCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cleanCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
