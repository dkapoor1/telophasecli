package cmd

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/spf13/cobra"

	"github.com/santiago-labs/telophasecli/cmd/runner"
	"github.com/santiago-labs/telophasecli/lib/awsorgs"
	"github.com/santiago-labs/telophasecli/lib/awssess"
	"github.com/santiago-labs/telophasecli/lib/ymlparser"
	"github.com/santiago-labs/telophasecli/resource"
	"github.com/santiago-labs/telophasecli/resourceoperation"
)

var orgFile string

func init() {
	rootCmd.AddCommand(accountProvision)
	accountProvision.Flags().StringVar(&orgFile, "org", "organization.yml", "Path to the organization.yml file")
}

func isValidAccountArg(arg string) bool {
	switch arg {
	case "import":
		return true
	case "diff":
		return true
	case "deploy":
		return true
	default:
		return false
	}
}

var accountProvision = &cobra.Command{
	Use:   "account",
	Short: "account - Provision and import AWS accounts",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires at least one arg")
		}
		if isValidAccountArg(args[0]) {
			return nil
		}
		return fmt.Errorf("invalid color specified: %s", args[0])
	},
	Run: func(cmd *cobra.Command, args []string) {
		orgClient := awsorgs.New()

		var consoleUI runner.ConsoleUI
		if useTUI {
			consoleUI = runner.NewTUI()
		} else {
			consoleUI = runner.NewSTDOut()
		}

		ctx := context.Background()
		if args[0] == "import" {
			if err := importOrgV2(orgClient); err != nil {
				panic(fmt.Sprintf("error importing organization: %s", err))
			}
		}

		rootAWSGroup, err := ymlparser.ParseOrganizationV2(orgFile)
		if err != nil {
			panic(fmt.Sprintf("error: %s", err))
		}
		if args[0] == "diff" {
			orgV2Diff(ctx, consoleUI, orgClient, rootAWSGroup, resourceoperation.Diff)
		}

		if args[0] == "deploy" {
			operations := orgV2Diff(ctx, consoleUI, orgClient, rootAWSGroup, resourceoperation.Deploy)

			for _, op := range operations {
				err := op.Call(ctx)
				if err != nil {
					panic(fmt.Sprintf("error: %s", err))
				}
			}
		}
	},
}

func orgV2Diff(
	ctx context.Context,
	outputUI runner.ConsoleUI,
	orgClient awsorgs.Client,
	rootAWSGroup *resource.AccountGroup,
	operation int,
) []resourceoperation.ResourceOperation {

	var operations []resourceoperation.ResourceOperation
	if rootAWSGroup != nil {
		mgmtAcct, err := orgClient.FetchManagementAccount(ctx)
		if err != nil {
			panic(err)
		}
		operations = append(operations, resourceoperation.CollectOrganizationUnitOps(
			ctx, outputUI, orgClient, rootAWSGroup, operation,
		)...)
		for _, op := range resourceoperation.FlattenOperations(operations) {
			outputUI.Print(op.ToString(), *mgmtAcct)
		}
		if len(operations) == 0 {
			outputUI.Print("\033[32m No changes to AWS Organization. \033[0m", *mgmtAcct)
		}
	}

	return operations
}

func currentAccountID() (string, error) {
	stsClient := sts.New(session.Must(awssess.DefaultSession()))
	caller, err := stsClient.GetCallerIdentity(&sts.GetCallerIdentityInput{})
	if err != nil {
		return "", err
	}

	return *caller.Account, nil
}

func importOrgV2(orgClient awsorgs.Client) error {
	managingAccountID, err := currentAccountID()
	if err != nil {
		return err
	}

	rootId, err := orgClient.GetRootId()
	if err != nil {
		return err
	}
	if rootId == "" {
		return fmt.Errorf("no root ID found")
	}

	rootGroup, err := orgClient.FetchGroupAndDescendents(context.TODO(), rootId, managingAccountID)
	if err != nil {
		return err
	}
	org := resource.AccountGroup{
		Name:        rootGroup.Name,
		ChildGroups: rootGroup.ChildGroups,
		Accounts:    rootGroup.Accounts,
	}

	if err := ymlparser.WriteOrgV2File(orgFile, &org); err != nil {
		return err
	}

	return nil
}
