package servicediscovery

type IServiceDiscovery interface {
	GetServiceIp(serviceName string) string
}
