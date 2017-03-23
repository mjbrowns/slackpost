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
	"os"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var MessageID string
var Verbose bool
var WorkDir string
var MessageFile string
var WorkPrefix string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   ProgName,
	Short: "Post a message to a slack channel",
	Long: `Sends a message to a slack channel by constructing a mesage object
	containing all available slack API capabilities`,
// Uncomment the following line if your bare application
// has an action associated with it:
//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := RootCmd.Execute()
  Check(err,"Error running root command")
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", fmt.Sprintf("config file (default is $HOME/.%s.json)",ProgName))
	RootCmd.PersistentFlags().StringVarP(&MessageID,"message-id","m","","message id number.  Can be anything.  Defaults to PID of caller")
	RootCmd.PersistentFlags().BoolVarP(&Verbose,"verbose","v",false,"Enable verbose messages")
	RootCmd.PersistentFlags().StringVarP(&WorkDir,"workdir","w",WorkDir,"working directory to use to store message objects")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	Log("\nStarting function root")
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}
	viper.SetConfigName(fmt.Sprintf(".%s",ProgName)) // name of config file (without extension)
	viper.AddConfigPath("$HOME")  // adding home directory as first search path
	viper.SetEnvPrefix("SLACKPOST")
	viper.AutomaticEnv()          // read in environment variables that match
	viper.SetDefault("workdir","/tmp")
	viper.SetDefault("workprefix",fmt.Sprintf("%s_",ProgName))
	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		if Verbose { fmt.Println("Using config file:", viper.ConfigFileUsed()) }
	}
	if WorkDir == "" { WorkDir = viper.GetString("workdir")	}
	WorkPrefix = viper.GetString("workprefix")
	if MessageID == "" { MessageID = viper.GetString("id") }
	if MessageID == "" { MessageID = fmt.Sprintf("%d",os.Getppid())	}
	found,err := PathExists(WorkDir)
	Check(err,fmt.Sprintf("Couldn't stat %s\n",WorkDir))
	if ! found { Die(fmt.Sprintf("ERROR: WorkDir %s does not exist!",WorkDir)) }
	MessageFile=fmt.Sprintf("%s/%s%s",WorkDir,WorkPrefix,MessageID)
	Log(fmt.Sprintf("Using WorkDir: %s",WorkDir))
	Log(fmt.Sprintf("Using Message-ID: %s",MessageID))
	Log(fmt.Sprintf("Using MessageFile: %s",MessageFile))
}
