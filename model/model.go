package model

import (
	"strings"
)

var GlobalCfg OpenclashCfg

//var ProxyMap map[string]Proxy // key-》 name
//var AddrMap map[string]Proxy  // key-》 server‘s doamin or ip
type DNS struct {
	Enable         bool     `yaml:"enable,omitempty" json:"enable"`
	Ipv6           bool     `yaml:"ipv6,omitempty" json:"ipv6"` // 是否启用 IPv6
	UseHosts       bool     `yaml:"use-hosts,omitempty" json:"use-hosts"`
	Nameserver     []string `yaml:"nameserver,omitempty" json:"nameserver"`
	Fallback       []string `yaml:"fallback,omitempty" json:"fallback"`
	FallbackFilter struct {
		GeoIP     bool     `yaml:"geoip,omitempty" json:"geoip"`
		GeoIPCode string   `yaml:"geoip-code,omitempty" json:"geoip-code"`
		IPCIDR    []string `yaml:"ipcidr,omitempty" json:"ipcidr"`
		Domain    []string `yaml:"domain,omitempty" json:"domain"`
	} `yaml:"fallback-filter,omitempty" json:"fallback-filter"`
	Listen            string            `yaml:"listen,omitempty" json:"listen"` // 是否使用 hosts 文件
	EnhancedMode      interface{}       `yaml:"enhanced-mode,omitempty"`
	FakeIpRange       string            `yaml:"fake-ip-range,omitempty" json:"fake-ip-range"`   // 伪造 IP 的范围
	FakeIPFilter      []string          `yaml:"fake-ip-filter,omitempty" json:"fake-ip-filter"` // 是否使用 hosts 文件
	DefaultNameserver []string          `yaml:"default-nameserver,omitempty" json:"default-nameserver"`
	NameServerPolicy  map[string]string `yaml:"nameserver-policy,omitempty" json:"nameserver-policy"`
	SearchDomains     []string          `yaml:"search-domains,omitempty" json:"search-domains"`
}

type Proxy struct {
	Name       string      `yaml:"name" json:"name"`                                   // 代理名称，用于标识不同的代理配置
	Type       string      `yaml:"type" json:"type"`                                   // 代理类型，如 ss、vmess、trojan 等
	Server     string      `yaml:"server" json:"server"`                               // 服务器地址，代理服务器的 IP 地址或域名
	Port       interface{} `yaml:"port" json:"port"`                                   // 服务器端口，代理服务器的端口号
	Cipher     string      `yaml:"cipher,omitempty" json:"cipher,omitempty"`           // 加密方式，如 aes-256-gcm，根据代理类型指定
	Password   string      `yaml:"password,omitempty" json:"password,omitempty"`       // 密码，用于 ss、trojan 等代理的认证
	Plugin     string      `yaml:"plugin,omitempty" json:"plugin,omitempty"`           // 插件名称，如 v2ray-plugin，用于增强代理功能
	PluginOpts *PluginOpts `yaml:"plugin-opts,omitempty" json:"plugin-opts,omitempty"` // 插件选项，定义插件具体的配置参数

	UUID    string `yaml:"uuid,omitempty" json:"uuid,omitempty"` // UUID，用于 vmess 类型的唯一识别码
	AlterID int    `yaml:"alterId" json:"alterId"`               // 额外ID，用于 vmess 的额外身份标识

	TLS            bool    `yaml:"tls,omitempty" json:"tls,omitempty"`                           // 是否启用 TLS 加密，通常用于 HTTPS 连接
	SkipCertVerify bool    `yaml:"skip-cert-verify,omitempty" json:"skip-cert-verify,omitempty"` // 是否跳过证书验证，用于自签名证书的情况
	Network        string  `yaml:"network,omitempty" json:"network,omitempty"`                   // 网络类型，如 tcp、ws（WebSocket）
	WSOpts         *WSOpts `yaml:"ws-optgos,omitempty" json:"ws-opts,omitempty"`                 // WebSocket 选项，当 network 为 ws 时使用

	Obfs      string `yaml:"obfs,omitempty" json:"obfs,omitempty"`             // 混淆方式，如 http，用于 SSR 的混淆设置
	ObfsParam string `yaml:"obfs-param,omitempty" json:"obfs-param,omitempty"` // 混淆参数，用于 SSR 混淆时的额外参数设置

	Protocol      string `yaml:"protocol,omitempty" json:"protocol,omitempty"`             // 协议类型，如 auth_sha1_v4，用于 SSR
	ProtocolParam string `yaml:"protocol-param,omitempty" json:"protocol-param,omitempty"` // 协议参数，为 SSR 的协议提供额外参数
	SNI           string `yaml:"sni,omitempty" json:"sni,omitempty"`                       // 服务名称指示，用于 HTTPS 连接时指定服务器名称
	Username      string `yaml:"username,omitempty" json:"username,omitempty"`             // 用户名，用于需要认证的 SOCKS5/HTTP 代理

	PSK      string    `yaml:"psk,omitempty" json:"psk,omitempty"`             // 预共享密钥，用于 Snell 代理的加密
	ObfsOpts *ObfsOpts `yaml:"obfs-opts,omitempty" json:"obfs-opts,omitempty"` // 混淆选项，用于 Snell 的混淆配置

	IP         string `yaml:"ip,omitempty" json:"ip,omitempty"`                   // IP 地址，用于 WireGuard 类型代理的本地 IP 地址
	Address    string `yaml:"address,omitempty" json:"address,omitempty"`         // 地址，用于 WireGuard 配置的网络地址
	PrivateKey string `yaml:"private-key,omitempty" json:"private-key,omitempty"` // 私钥，用于 WireGuard 类型代理的私钥
	DNSServer  string `yaml:"dns-server,omitempty" json:"dns-server,omitempty"`   // DNS 服务器，用于 WireGuard 配置的 DNS 解析
	MTU        int    `yaml:"mtu,omitempty" json:"mtu,omitempty"`                 // 最大传输单元，用于 WireGuard 的网络配置
	Peers      []Peer `yaml:"peers,omitempty" json:"peers,omitempty"`             // 对等点，用于 WireGuard 配置中的远程对等点列表
	SubFrom    string `yaml:"-" json:"subFrom,omitempty"`                         //自定义字段，订阅来源
	InnerName  string `yaml:"innerName,omitempty" json:"innerName,omitempty"`
}

type PluginOpts struct {
	Mode string `yaml:"mode" json:"mode"` // 插件模式，如 websocket
	Host string `yaml:"host" json:"host"` // 插件主机名
}

type WSOpts struct {
	Path    string            `yaml:"path" json:"path"`       // WebSocket 路径
	Headers map[string]string `yaml:"headers" json:"headers"` // WebSocket 头部信息
}

type ObfsOpts struct {
	Mode string `yaml:"mode" json:"mode"` // 混淆模式，如 http
	Host string `yaml:"host" json:"host"` // 混淆主机名
}

type Peer struct {
	PublicKey  string   `yaml:"public-key" json:"public-key"`   // 对等点的公钥
	Endpoint   string   `yaml:"endpoint" json:"endpoint"`       // 对等点的端点地址
	AllowedIPs []string `yaml:"allowed-ips" json:"allowed-ips"` // 允许的 IP 地址段
}

type ProxyGroup struct {
	Name     string   `yaml:"name" json:"name"`         // 代理组名称
	Type     string   `yaml:"type" json:"type"`         // 代理组类型，如 select, url-test
	Proxies  []string `yaml:"proxies" json:"proxies"`   // 包含的代理名称列表
	URL      string   `yaml:"url" json:"url"`           // 包含的代理名称列表
	Interval int      `yaml:"interval" json:"interval"` // 包含的代理名称列表
}

type OpenclashCfg struct {
	Port               int      `yaml:"port,omitempty"`
	SocksPort          int      `yaml:"socks-port,omitempty"`
	RedirPort          int      `yaml:"redir-port,omitempty"`
	TProxyPort         int      `yaml:"tproxy-port,omitempty"`
	MixedPort          int      `yaml:"mixed-port,omitempty"`
	Authentication     []string `yaml:"authentication,omitempty"`
	AllowLAN           bool     `yaml:"allow-lan,omitempty"`
	BindAddress        string   `yaml:"bind-address,omitempty"`
	LogLevel           string   `yaml:"log-level,omitempty"`
	Mode               string   `yaml:"mode,omitempty"`
	ExternalController string   `yaml:"external-controller,omitempty"`
	ExternalUI         string   `yaml:"external-ui,omitempty"`
	Secret             string   `yaml:"secret,omitempty"`
	Interface          string   `yaml:"interface-name,omitempty"`
	RoutingMark        int      `yaml:"routing-mark,omitempty"`
	Tunnels            []Tunnel `yaml:"tunnels,omitempty"`

	ProxyProvider map[string]map[string]any `yaml:"proxy-providers,omitempty"`
	Hosts         map[string]string         `yaml:"hosts,omitempty"`
	Inbounds      []Inbound                 `yaml:"inbounds,omitempty"`
	DNS           DNS                       `yaml:"dns,omitempty"`
	Proxies       []Proxy                   `yaml:"proxies,omitempty"`
	ProxyGroups   []ProxyGroup              `yaml:"proxy-groups,omitempty"` // Assuming it can be a slice even if empty
	Rules         []string                  `yaml:"rules,omitempty"`
}

// Profile config
type Profile struct {
	StoreSelected bool `yaml:"store-selected"`
	StoreFakeIP   bool `yaml:"store-fake-ip"`
}

type tunnel struct {
	Network []string `yaml:"network"`
	Address string   `yaml:"address"`
	Target  string   `yaml:"target"`
	Proxy   string   `yaml:"proxy"`
}
type InboundType string

type inbound struct {
	Type          InboundType `json:"type" yaml:"type"`
	BindAddress   string      `json:"bind-address" yaml:"bind-address"`
	IsFromPortCfg bool        `json:"-" yaml:"-"`
}

// Inbound
type Inbound inbound
type Tunnel tunnel

// 更新ProxyMap
func UpdateProxyMap(allProxy []Proxy, wordMap map[string]struct{}) map[string]Proxy {
	// 去重, server是每个节点唯一的
	addrMap := make(map[string]Proxy, 0)
	proxyMap := make(map[string]Proxy, 0)

	for _, val := range allProxy {
		if nameFilter(val.Name, wordMap) {
			continue
		}
		newName := val.Name
		if strings.Contains(val.Name, ":") {
			newName = strings.ReplaceAll(val.Name, ":", "-")
			//fmt.Println("************ newName", newName)
		}

		key := val.Server + ":" + newName
		addrMap[key] = val
	}

	for _, proxy := range addrMap {
		newName := proxy.Name
		if strings.Contains(proxy.Name, ":") {
			newName = strings.ReplaceAll(proxy.Name, ":", "-")
			//fmt.Println("************ newName", newName)
		}
		key := proxy.Server + ":" + newName
		proxyMap[key] = proxy
	}
	return proxyMap
}
func nameFilter(name string, wordMap map[string]struct{}) bool {
	for word := range wordMap {
		if strings.Contains(name, word) {
			return true
		}
	}
	return false
}
