package kubeapi

import (
	"context"
	"errors"

	"github.com/miekg/dns"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/plugin"
)

// KubeAPI implements a plugin that establishes a Kubernetes API client.
type KubeAPI struct {
	Next plugin.Handler

	APIServer     string
	APICertAuth   string
	APIClientCert string
	APIClientKey  string
	ClientConfig  clientcmd.ClientConfig

	Client kubernetes.Interface
}

// Name implements the Handler interface.
func (k KubeAPI) Name() string { return pluginName }

// ServeDNS implements the plugin.Handler interface.
func (k KubeAPI) ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {
	return plugin.NextOrFailure(k.Name(), k.Next, ctx, w, r)
}

// Client returns the Kubernetes API client defined by kubeapi plugin in the server config.
// This should only be called after all plugins have loaded, e.g. from a caddy Controller OnStartup function.
func Client(config *dnsserver.Config) (kubernetes.Interface, error) {
	for _, h := range config.Handlers() {
		k, ok := h.(*KubeAPI)
		if !ok {
			continue
		}
		return k.Client, nil
	}
	// Either the kubeapi plugin was not defined in the serverblock, or Client() was called
	// before it had a chance to register itself.
	return nil, errors.New("kubeapi plugin not registered in server block")
}
