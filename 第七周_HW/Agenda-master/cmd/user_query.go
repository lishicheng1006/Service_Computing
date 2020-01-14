// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"agenda-go-cli/service"
	"github.com/spf13/cobra"
)

// queryuserCmd represents the queryuser command
var queryuserCmd = &cobra.Command{
	Use:   "toqueryuser",
	Short: "query user by name",
	Run: func(cmd *cobra.Command, args []string) {
		errLog.Println("Query User called")
		if _, flag := service.GetCurUser(); flag != true {
			fmt.Println("Please Log in firstly")
		}
		ru := service.ListAllUser()
		for _, u := range ru {
			fmt.Println("----------------")
			fmt.Println("Username: ", u.Name)
			fmt.Println("Phone: ",u.Phone)
			fmt.Println("Email: ",u.Email)
			fmt.Println("----------------")
		}
	},
}

func init() {
	RootCmd.AddCommand(queryuserCmd)

	// Here you will define your flags and configuration settings.
	// queryuserCmd.Flags().StringP("username", "u", "", "specify user")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// queryuserCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// queryuserCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
