package commands

import (
	"flag"
	"strconv"

	"github.com/cloudfoundry-incubator/multiapps-cli-plugin/util"
	"github.com/cloudfoundry/cli/plugin"
)

const noConfirmOpt = "no-confirm"

// BlueGreenDeployCommand is a command for blue green deployment of MTAs.
type BlueGreenDeployCommand struct {
	*DeployCommand
}

// NewBlueGreenDeployCommand creates a new BlueGreenDeployCommand.
func NewBlueGreenDeployCommand() *BlueGreenDeployCommand {
	baseCmd := &BaseCommand{flagsParser: deployCommandLineArgumentsParser{}, flagsValidator: deployCommandFlagsValidator{}}
	deployCmd := &DeployCommand{baseCmd, blueGreenDeployProcessParametersSetter(), &blueGreenDeployCommandProcessTypeProvider{}}
	bgDeployCmd := &BlueGreenDeployCommand{deployCmd}
	baseCmd.Command = bgDeployCmd
	return bgDeployCmd
}

// GetPluginCommand returns more information for the blue green deploy command.
func (c *BlueGreenDeployCommand) GetPluginCommand() plugin.Command {
	return plugin.Command{
		Name:     "bg-deploy",
		HelpText: "Deploy a multi-target app using blue-green deployment",
		UsageDetails: plugin.Usage{
			Usage: `Deploy a multi-target app using blue-green deployment
   cf bg-deploy MTA [-e EXT_DESCRIPTOR[,...]] [-t TIMEOUT] [--version-rule VERSION_RULE] [-u URL] [-f]  [--retries RETRIES] [--no-start]  [--namespace NAMESPACE] [--delete-services] [--delete-service-keys] [--delete-service-brokers] [--keep-files] [--no-restart-subscribed-apps]  [--no-confirm] [--do-not-fail-on-missing-permissions] [--abort-on-error] [--verify-archive-signature]

   Perform action on an active deploy operation
   cf deploy -i OPERATION_ID -a ACTION [-u URL]`,
			Options: map[string]string{
				extDescriptorsOpt:                                  "Extension descriptors",
				deployServiceURLOpt:                                "Deploy service URL, by default 'deploy-service.<system-domain>'",
				timeoutOpt:                                         "Start timeout in seconds",
				versionRuleOpt:                                     "Version rule (HIGHER, SAME_HIGHER, ALL)",
				operationIDOpt:                                     "Active deploy operation ID",
				actionOpt:                                          "Action to perform on active deploy operation (abort, retry, resume, monitor)",
				forceOpt:                                           "Force deploy without confirmation for aborting conflicting processes",
				util.GetShortOption(noStartOpt):                    "Do not start apps",
				util.GetShortOption(namespaceOpt):                  "(EXPERIMENTAL) Namespace for the mta, applied to app and service names as well",
				util.GetShortOption(deleteServicesOpt):             "Recreate changed services / delete discontinued services",
				util.GetShortOption(deleteServiceKeysOpt):          "Delete existing service keys and apply the new ones",
				util.GetShortOption(deleteServiceBrokersOpt):       "Delete discontinued service brokers",
				util.GetShortOption(keepFilesOpt):                  "Keep files used for deployment",
				util.GetShortOption(noRestartSubscribedAppsOpt):    "Do not restart subscribed apps, updated during the deployment",
				util.GetShortOption(noConfirmOpt):                  "Do not require confirmation for deleting the previously deployed MTA apps",
				util.GetShortOption(noFailOnMissingPermissionsOpt): "Do not fail on missing permissions for admin operations",
				util.GetShortOption(abortOnErrorOpt):               "Auto-abort the process on any errors",
				util.GetShortOption(verifyArchiveSignatureOpt):     "Verify the archive is correctly signed",
				util.GetShortOption(retriesOpt):                    "Retry the operation N times in case a non-content error occurs (default 3)",
			},
		},
	}
}

func (c *BlueGreenDeployCommand) defineCommandOptions(flags *flag.FlagSet) {
	c.DeployCommand.defineCommandOptions(flags)
	flags.Bool(noConfirmOpt, false, "")
}

// blueGreenDeployProcessParametersSetter returns a new ProcessParametersSetter.
func blueGreenDeployProcessParametersSetter() ProcessParametersSetter {
	return func(flags *flag.FlagSet, processBuilder *util.ProcessBuilder) {
		deployProcessParametersSetter()(flags, processBuilder)
		processBuilder.Parameter("noConfirm", strconv.FormatBool(GetBoolOpt(noConfirmOpt, flags)))
		processBuilder.Parameter("keepOriginalAppNamesAfterDeploy", strconv.FormatBool(false))
	}
}

type blueGreenDeployCommandProcessTypeProvider struct{}

func (bg blueGreenDeployCommandProcessTypeProvider) GetProcessType() string {
	return "BLUE_GREEN_DEPLOY"
}
