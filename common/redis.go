package common


type RedisMapConf struct {
	List map[string]*RedisConf `mapstructure:"list"`
}

type RedisConf struct {
	ProxyList []string `mapstructure:"proxy_list"`
	MaxActive int      `mapstructure:"max_active"`
	MaxIdle   int      `mapstructure:"max_idle"`
	Downgrade bool     `mapstructure:"down_grade"`
}

func InitRedis(env *Env) error {
	ConfRedis := &RedisMapConf{}

	err := env.ParseConfig(env.GetConfPath("redis_map"), ConfRedis)
	if err != nil {
		return err
	}
	//初始化redis 配置


	return nil
}