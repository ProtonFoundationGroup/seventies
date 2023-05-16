package config

type Config struct {
	Serial string        `yaml:"serial"`
	Chat   *ChatConfig   `yaml:"chat"`
	Wallet *WalletConfig `yaml:"wallet"`
	Net    *NetConfig    `yaml:"net"`
}

func Load(filename string) (*Config, error) {

	// 解析 filename 指定的 yml文件
	config := &Config{}
	err := parseYaml(filename, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
