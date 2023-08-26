package commands

import (
	"fmt"
	"os"

	"github.com/kehiy/taar/utils"
	cobra "github.com/spf13/cobra"
)

func BuildDNSCommand(parentCmd *cobra.Command) {
	dnsCmd := &cobra.Command{
		Use:   "dns",
		Short: "change and manage dns",
	}
	buildSetCommand(dnsCmd)
	buildShowCommand(dnsCmd)

	parentCmd.AddCommand(dnsCmd)
}

func buildShowCommand(parentCmd *cobra.Command) {
	showCmd := &cobra.Command{
		Use:   "show",
		Short: "show dns setting",
	}
	parentCmd.AddCommand(showCmd)

	showCmd.Run = func(cmd *cobra.Command, args []string) {
		cmd.Println(utils.ShowResolve())
	}
}

func buildSetCommand(parentCmd *cobra.Command) {
	setCmd := &cobra.Command{
		Use:   "set",
		Short: "set new dns",
	}
	parentCmd.AddCommand(setCmd)

	setCmd.Run = func(cmd *cobra.Command, args []string) {
		err := changeDNS(args)
		if err != nil {
			cmd.PrintErrf("can't change dns: error:\n%v\n", err)
		} else {
			cmd.Printf("dns successfully changed, new config:\n%s\n", args)
		}
	}
}

// ! not CMDs.
func changeDNS(DNSs []string) error {
	path := "/etc/resolv.conf"

	newContent := `
	# DO NOT CHANGE!
	# managed by taar network manager.
	`
	for _, dns := range DNSs {
		newContent += fmt.Sprintf("nameserver %s\n", dns)
	}

	err := os.WriteFile(path, []byte(newContent), 0o600)
	if err != nil {
		return err
	}

	return nil
}