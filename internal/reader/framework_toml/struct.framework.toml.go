package framework_toml

// FrameworkConfig framework.toml配置文件结构
type FrameworkConfig struct {
	Title string `toml:"title"`
	Framework struct{
		FrameworkHash string `toml:"framework_hash"`
		Timezone string `toml:"timezone"`
		GlVersion string `toml:"gv_version"`
		FrameworkPath string `toml:"framework_path"`
		Desc string `toml:"desc"`
		Author string `toml:"author"`
		Email string `toml:"email"`
		Github string `toml:"github"`
		Gitee string `toml:"gitee"`
	} `toml:"framework"`
	HttpServer struct{
		Host string `toml:"host"`
		Port string `toml:"port"`
	} `toml:"http_server"`
	GRPCServer struct{
		Host string `toml:"host"`
		Port string `toml:"port"`
	} `toml:"grpc_server"`
	GRPCClient struct{
		Host string `toml:"host"`
		Port string `toml:"port"`
	} `toml:"grpc_client"`
	Mysql1 struct{
		Host string `toml:"host"`
		Port string `toml:"port"`
		Database string `toml:"database"`
		User string `toml:"user"`
		Pwd string `toml:"pwd"`
		Encode string `toml:"encode"`
	} `toml:"mysql_1"`
	Mysql2 struct{
		Host string `toml:"host"`
		Port string `toml:"port"`
		Database string `toml:"database"`
		User string `toml:"user"`
		Pwd string `toml:"pwd"`
		Encode string `toml:"encode"`
	} `toml:"mysql_2"`
	Redis1 struct{
		Host string `toml:"host"`
		Port string `toml:"port"`
		Pwd string `toml:"pwd"`
		DataBase string `toml:"database"`
	} `toml:"redis_1"`
	Redis2 struct{
		Host string `toml:"host"`
		Port string `toml:"port"`
		Pwd string `toml:"pwd"`
		DataBase string `toml:"database"`
	} `toml:"redis_2"`
	ElasticSearch struct{
		HttpHost string `toml:"http_host"`
		Port string `toml:"port"`
	} `toml:"elastic_search"`
	View struct{
		ViewHTML string `toml:"view_html"`
		ViewStatic string `toml:"view_static"`
		ViewTPL string `toml:"view_tpl"`
	} `toml:"view"`
	VersionBlock struct{
		Api string `toml:"api"`
		Web string `toml:"web"`
	} `toml:"version_block"`
	Common struct{
		Env string `toml:"env"`
		LogHost string `toml:"log_host"`
		LogPort string `toml:"log_port"`
	} `toml:"common"`
}