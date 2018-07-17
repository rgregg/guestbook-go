# guestbook-go
A simple guest book web app written in Go.

## Setup builds

```shell
$ kubectl apply -f setup/kaniko.yaml

$ kubectl apply -f setup/build_secret.yaml

$ kubectl edit serviceaccount default
```

Add `build-secret` to the list of credentials in the default service account.

## Setup data store

Configure to allow off-cluster communication, which involves editing the network
map and including the local IP addresses.

```shell
$ gcloud container clusters describe your-cluster-id --zone=us-west1-c | grep -e clusterIpv4Cidr -e servicesIpv4Cidr

$ kubectl edit configmap config-network -n knative-serving
```

Change the value to look something like this:

```yaml
data:
  istio.sidecar.includeOutboundIPRanges: '10.16.0.0/14,10.19.240.0/20'
```

[See more details here](https://github.com/knative/docs/blob/master/serving/outbound-network-access.md).

## Running the demo

```shell
kubectl apply -f service.yaml
```

