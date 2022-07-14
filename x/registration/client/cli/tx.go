package cli

import (
	"bufio"
	//"fmt"
	"io/ioutil"
	//"strconv"

	"github.com/spf13/cobra"
	//"github.com/spf13/viper"

	"github.com/enigmampc/cosmos-sdk/client"
	"github.com/enigmampc/cosmos-sdk/client/context"
	"github.com/enigmampc/cosmos-sdk/client/flags"
	"github.com/enigmampc/cosmos-sdk/codec"
	sdk "github.com/enigmampc/cosmos-sdk/types"
	"github.com/enigmampc/cosmos-sdk/x/auth"
	"github.com/enigmampc/cosmos-sdk/x/auth/client/utils"

	"github.com/enigmampc/SecretNetwork/x/registration/internal/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Registration transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	txCmd.AddCommand(flags.PostCommands(
		AuthenticateNodeCmd(cdc),
	)...)
	return txCmd
}

// AuthenticateNodeCmd will upload code to be reused.
func AuthenticateNodeCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "auth [cert file]",
		Short: "Upload a certificate to authenticate the node",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			// parse coins trying to be sent
			cert, err := ioutil.ReadFile(args[0])
			if err != nil {
				return err
			}

			// build and sign the transaction, then broadcast to Tendermint
			msg := types.RaAuthenticate{
				Sender:      cliCtx.GetFromAddress(),
				Certificate: cert,
			}
			err = msg.ValidateBasic()

			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return cmd
}
