# kubeapi

## Name

*kubeapi* - configures a Kubernetes API client.

## Description

*kubeapi* configures a Kubernetes API client and makes it available to other plugins.

Plugins that wish to use the Kubernetes API client can call `kubeapi.Client()` after the server starts,
e.g. via `caddy.OnStartup`.

## Syntax

```
kubeapi {
    endpoint URL
    tls CERT KEY CACERT
    kubeconfig KUBECONFIG [CONTEXT]
}
```

* `endpoint` specifies the **URL** for a remote k8s API endpoint.
  If omitted, it will connect to k8s in-cluster using the cluster service account.
* `tls` **CERT** **KEY** **CACERT** are the TLS cert, key and the CA cert file names for remote k8s connection.
  This option is ignored if connecting in-cluster (i.e. endpoint is not specified).
* `kubeconfig` **KUBECONFIG [CONTEXT]** authenticates the connection to a remote k8s cluster using a kubeconfig file.
  **[CONTEXT]** is optional, if not set, then the current context specified in kubeconfig will be used.
  It supports TLS, username and password, or token-based authentication.
  This option is ignored if connecting in-cluster (i.e., the endpoint is not specified).


## External Plugin

To use this plugin, compile CoreDNS with this plugin added to `plugin.cfg`. It should be placed at/near the end of the
chain of plugins.

## Examples

The _kubenodes_ plugin (https://github.com/infobloxopen/kubenodes) uses _kubeapi_ to connect to the Kubernetes API.

```
. {
  kubeapi
  kubenodes node.cluster.local in-addr.arpa ip6.arpa
}
```
