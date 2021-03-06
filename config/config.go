package config

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/go-yaml/yaml"
	"github.com/spf13/viper"
)

type ConfigViper struct {
	*viper.Viper
	Account  *Account
	Accounts map[string]*Account
}

type Account struct {
	Host         string
	ID           int
	RefreshToken string `mapstructure:"refresh_token" yaml:"refresh_token"`
}

var Config ConfigViper

func init() {
	Config.Viper = viper.New()
	Config.SetDefault("login", map[string]interface{}{"accounts": make(map[string]interface{})})
	Config.SetDefault("update", map[string]interface{}{"check": true})
	Config.SetEnvPrefix("fpt")
	Config.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	Config.AutomaticEnv()
}

func ReadConfig(configFile, account string) error {
	Config.SetConfigFile(configFile)
	err := Config.ReadInConfig()
	if err != nil {
		if _, ok := err.(*os.PathError); !(ok &&
			Config.IsSet("login.account.id") &&
			Config.IsSet("login.account.host") &&
			Config.IsSet("login.account.refresh_token")) {
			return err
		}
	}

	err = Config.UnmarshalKey("login.accounts", &Config.Accounts)
	if err != nil {
		return fmt.Errorf("%s: %s", configFile, err)
	}

	if Config.IsSet("login.account.id") &&
		Config.IsSet("login.account.host") &&
		Config.IsSet("login.account.refresh_token") {
		Config.Account = &Account{
			ID:           Config.GetInt("login.account.id"),
			Host:         Config.GetString("login.account.host"),
			RefreshToken: Config.GetString("login.account.refresh_token"),
		}
	} else {
		var ok bool
		if account == "" {
			defaultAccount := Config.GetString("login.default_account")
			Config.Account, ok = Config.Accounts[strings.ToLower(defaultAccount)]
			if !ok {
				return fmt.Errorf("%s: could not find default account: %s", configFile, defaultAccount)
			}
		} else {
			Config.Account, ok = Config.Accounts[strings.ToLower(account)]
			if !ok {
				return fmt.Errorf("%s: could not find account: %s", configFile, account)
			}
		}
	}

	return nil
}

func (config *ConfigViper) GetAccount(id int, host string) (*Account, error) {
	for _, account := range config.Accounts {
		if account.ID == id && account.Host == host {
			return account, nil
		}
	}

	return nil, fmt.Errorf("Could not find account for account/host: %d %s", id, host)
}

// Obtain input via STDIN then print out to config file
// Example of config file
// login:
//   default_account: acct1
//   accounts:
//     acct1:
//       id: 67972
//       host: us-3.rightscale.com
//       refresh_token: abc123abc123abc123abc123abc123abc123abc1
//     acct2:
//       id: 60073
//       host: us-4.rightscale.com
//       refresh_token: zxy987zxy987zxy987zxy987xzy987zxy987xzy9
func (config *ConfigViper) SetAccount(name string, setDefault bool, input io.Reader, output io.Writer) error {
	// if the default account isn't set we should set it to the account we are setting
	if !config.IsSet("login.default_account") {
		setDefault = true
	}

	// get the settings and specifically the login settings into a map we can manipulate and marshal to YAML unhindered
	// by the meddling of the Viper
	settings := config.AllSettings()
	if _, ok := settings["login"]; !ok {
		settings["login"] = map[string]interface{}{"accounts": make(map[string]interface{})}
	}
	loginSettings := settings["login"].(map[string]interface{})

	// set the default account if we want or need to
	if setDefault {
		loginSettings["default_account"] = name
	}

	// get the previous value for the named account if it exists and construct a new account to populate
	oldAccount, ok := config.Accounts[name]
	newAccount := &Account{}

	// prompt for the account ID and use the old value if nothing is entered
	fmt.Fprint(output, "Account ID")
	if ok {
		fmt.Fprintf(output, " (%d)", oldAccount.ID)
	}
	fmt.Fprint(output, ": ")
	fmt.Fscanln(input, &newAccount.ID)
	if ok && newAccount.ID == 0 {
		newAccount.ID = oldAccount.ID
	}

	// prompt for the API endpoint host and use the old value if nothing is entered
	fmt.Fprint(output, "API endpoint host")
	if ok {
		fmt.Fprintf(output, " (%s)", oldAccount.Host)
	}
	fmt.Fprint(output, ": ")
	fmt.Fscanln(input, &newAccount.Host)
	if ok && newAccount.Host == "" {
		newAccount.Host = oldAccount.Host
	}

	// prompt for the refresh token and use the old value if nothing is entered
	fmt.Fprint(output, "Refresh token")
	if ok {
		fmt.Fprintf(output, " (%s)", oldAccount.RefreshToken)
	}
	fmt.Fprint(output, ": ")
	fmt.Fscanln(input, &newAccount.RefreshToken)
	if ok && newAccount.RefreshToken == "" {
		newAccount.RefreshToken = oldAccount.RefreshToken
	}

	// add the new account to the map of accounts overwriting any old value
	accounts := loginSettings["accounts"].(map[string]interface{})
	accounts[name] = newAccount

	// render the settings map as YAML
	yml, err := yaml.Marshal(settings)
	if err != nil {
		return err
	}

	configPath := config.ConfigFileUsed()
	// back up the current config file before writing a new one or if one does not exist, make sure the directory exists
	if err := os.Rename(configPath, configPath+".bak"); err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(filepath.Dir(configPath), 0700); err != nil {
				return err
			}
		} else {
			return err
		}
	}

	// create a new config file which only the current user can read or write
	configFile, err := os.OpenFile(configPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer configFile.Close()

	// write the YAML into the config file
	if _, err := configFile.Write(yml); err != nil {
		return err
	}

	return nil
}

func (config *ConfigViper) ShowConfiguration(output io.Writer) error {
	// Check if config file exists
	if _, err := os.Stat(config.ConfigFileUsed()); err != nil {
		return err
	}

	yml, err := yaml.Marshal(config.AllSettings())
	if err != nil {
		return err
	}
	output.Write(yml)

	return nil
}

var hostRe = regexp.MustCompile(`governance-(\d+)\.(test.)?rightscale\.com`)

func (a *Account) Validate() error {
	if !hostRe.MatchString(a.Host) {
		return fmt.Errorf("invalid host: must be of form governance-<shard number>.rightscale.com")
	}
	return nil
}

func (a *Account) AuthHost() string {
	matches := hostRe.FindStringSubmatch(a.Host)
	if len(matches) == 0 {
		return ""
	}
	shardNum := matches[1]
	testHost := matches[2]
	prefix := "us"
	if shardNum == "10" {
		prefix = "telstra"
	}
	if testHost != "" {
		prefix = "moo"
	}
	return fmt.Sprintf("%s-%s.%srightscale.com", prefix, shardNum, testHost)
}
