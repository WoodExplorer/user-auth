/* ******************************************************************************
* 2019 - present Contributed by Apulis Technology (Shenzhen) Co. LTD
*
* This program and the accompanying materials are made available under the
* terms of the MIT License, which is available at
* https://www.opensource.org/licenses/MIT
*
* See the NOTICE file distributed with this work for additional
* information regarding copyright ownership.
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
* WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
* License for the specific language governing permissions and limitations
* under the License.
*
* SPDX-License-Identifier: MIT
******************************************************************************/
package cmd

import (
	roleRepo "github.com/WoodExplorer/user-auth/internal/repository/role"
	userRepo "github.com/WoodExplorer/user-auth/internal/repository/user"
	"github.com/WoodExplorer/user-auth/internal/router"
	"github.com/WoodExplorer/user-auth/internal/services/role"
	"github.com/WoodExplorer/user-auth/internal/services/user"
	"github.com/WoodExplorer/user-auth/internal/stores/memory"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run datahub-bff service",
	Run: func(cmd *cobra.Command, args []string) {

		go func() {
			warn := http.ListenAndServe(":6060", nil)
			if warn != nil {
				log.Warn().Msgf("starting metric server error: %+v", warn)
			} else {
				log.Info().Msg("metric server started")
			}
		}()

		store := memory.NewStore()
		store.Start()

		rr := roleRepo.NewRepo(store)
		ur := userRepo.NewRepo(store)

		roleSvc := role.NewService(rr)
		userSvc := user.NewService(ur)

		r := router.InitRouter(roleSvc, userSvc)
		go func() {
			err := r.Start()
			if err != nil {
				log.Fatal().Msgf("starting server error: %+v", err)
			}
		}()

		// wait for interrupt signal to gracefully shut down the server
		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit

		store.Stop()

		log.Info().Msg("server exited")
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
