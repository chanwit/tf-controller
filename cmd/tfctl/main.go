package main

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/weaveworks/tf-controller/tfctl"
)

// BuildSHA is the tfctl version
var BuildSHA string

func main() {
	cmd := run()
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func run() *cobra.Command {
	app := tfctl.New(BuildSHA)

	rootCmd := &cobra.Command{
		Use:           "tfctl",
		SilenceErrors: false,
		SilenceUsage:  true,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return app.Init(viper.GetString("kubeconfig"), viper.GetString("namespace"), viper.GetString("terraform"))
		},
	}

	rootCmd.PersistentFlags().String("kubeconfig", "", "Path to the kubeconfig file to use for CLI requests.")
	rootCmd.PersistentFlags().StringP("namespace", "n", "flux-system", "The kubernetes namespace to use for CLI requests.")
	rootCmd.PersistentFlags().String("terraform", "/usr/bin/terraform", "The location of the terraform binary.")

	viper.BindPFlag("kubeconfig", rootCmd.PersistentFlags().Lookup("kubeconfig"))
	viper.BindPFlag("namespace", rootCmd.PersistentFlags().Lookup("namespace"))
	viper.BindPFlag("terraform", rootCmd.PersistentFlags().Lookup("terraform"))

	viper.BindEnv("kubeconfig")

	rootCmd.AddCommand(buildVersionCmd(app))
	rootCmd.AddCommand(buildPlanCmd(app))
	rootCmd.AddCommand(buildInstallCmd(app))
	rootCmd.AddCommand(buildUninstallCmd(app))
	rootCmd.AddCommand(buildReconcileCmd(app))
	rootCmd.AddCommand(buildSuspendCmd(app))
	rootCmd.AddCommand(buildResumeCmd(app))
	rootCmd.AddCommand(buildGetCmd(app))

	return rootCmd
}

func buildVersionCmd(app *tfctl.CLI) *cobra.Command {
	install := &cobra.Command{
		Use:   "version",
		Short: "Prints tf-controller and tfctl version information",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			return app.Version(os.Stdout)
		},
	}
	install.Flags().String("version", "", "The version of tf-controller to install.")
	viper.BindPFlag("version", install.Flags().Lookup("version"))
	return install
}
func buildInstallCmd(app *tfctl.CLI) *cobra.Command {
	install := &cobra.Command{
		Use:   "install",
		Short: "Install the tf-controller",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			return app.Install(viper.GetString("version"))
		},
	}
	install.Flags().String("version", "", "The version of tf-controller to install.")
	viper.BindPFlag("version", install.Flags().Lookup("version"))
	return install
}

func buildUninstallCmd(app *tfctl.CLI) *cobra.Command {
	return &cobra.Command{
		Use:   "uninstall",
		Short: "Uninstall the tf-controller",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			return app.Uninstall()
		},
	}
}

func buildReconcileCmd(app *tfctl.CLI) *cobra.Command {
	return &cobra.Command{
		Use:   "reconcile NAME",
		Short: "Trigger a reconcile of the provided resource",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return app.Reconcile(args[0])
		},
	}
}

func buildSuspendCmd(app *tfctl.CLI) *cobra.Command {
	return &cobra.Command{
		Use:   "suspend NAME",
		Short: "Suspend reconciliation for the provided resource",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return app.Suspend(args[0])
		},
	}
}

func buildResumeCmd(app *tfctl.CLI) *cobra.Command {
	return &cobra.Command{
		Use:   "resume NAME",
		Short: "Resume reconciliation for the provided resource",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return app.Resume(args[0])
		},
	}
}

func buildPlanCmd(app *tfctl.CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "plan",
		Short: "Plan a Terraform configuration",
	}

	cmd.AddCommand(buildPlanShowCmd(app))
	cmd.AddCommand(buildPlanApproveCmd(app))

	return cmd
}

func buildPlanShowCmd(app *tfctl.CLI) *cobra.Command {
	return &cobra.Command{
		Use:   "show NAME",
		Short: "Show pending Terraform plan",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return app.ShowPlan(os.Stdout, args[0])
		},
	}
}

func buildPlanApproveCmd(app *tfctl.CLI) *cobra.Command {
	return &cobra.Command{
		Use:   "approve NAME",
		Short: "Approve pending Terraform plan",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return app.ApprovePlan(os.Stdout, args[0])
		},
	}
}

func buildGetCmd(app *tfctl.CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get Terraform resources",
		RunE: func(cmd *cobra.Command, args []string) error {
			return app.Get(os.Stdout)
		},
	}

	cmd.AddCommand(buildGetTerraformCmd(app))

	return cmd
}

func buildGetTerraformCmd(app *tfctl.CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get NAME",
		Short: "Get a Terraform resource",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return app.GetTerraform(os.Stdout, args[0])
		},
	}

	return cmd
}