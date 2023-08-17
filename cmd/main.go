package cmd

import (
	role_repo "github.com/WoodExplorer/user-auth/internal/repository/role"
	token_blacklist_repo "github.com/WoodExplorer/user-auth/internal/repository/token_blacklist"
	user_repo "github.com/WoodExplorer/user-auth/internal/repository/user"
	user_role_repo "github.com/WoodExplorer/user-auth/internal/repository/user_role"
	"github.com/WoodExplorer/user-auth/internal/router"
	"github.com/WoodExplorer/user-auth/internal/services/authn"
	"github.com/WoodExplorer/user-auth/internal/services/authz"
	"github.com/WoodExplorer/user-auth/internal/services/role"
	"github.com/WoodExplorer/user-auth/internal/services/user"
	"github.com/WoodExplorer/user-auth/internal/services/user_role"
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
	Short: "Run user-auth service",
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

		rr := role_repo.NewRepo(store)
		ur := user_repo.NewRepo(store)
		urr := user_role_repo.NewRepo(store)
		tbr := token_blacklist_repo.NewRepo(store)

		roleSvc := role.NewService(rr)
		userSvc := user.NewService(ur, urr)
		userRoleSvc := user_role.NewService(urr)
		authnSvc := authn.NewService(ur, tbr)
		authzSvc := authz.NewService(urr, tbr)

		_, r := router.InitRouter(roleSvc, userSvc, userRoleSvc, authnSvc, authzSvc)
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

		r.Stop()
		store.Stop()
		log.Info().Msg("server exited")
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
