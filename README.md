# Example Kubernetes Restarter from FireHydrant

This repository contains an example Go HTTP service that can handle Slack command extensions from FireHydrant.com. To learn more about command extensions, checkout https://firehydrant.com/docs/integration-guides/slack-command-extensions/

## Running this service

Make sure you've exported your `KUBECONFIG` environment variable. This service does not support in-cluster configuration for Kubernetes clientsets. Please adopt it as you need!

Once you've done that, you can run in your terminal:

```
export KUBECONIG=~/.kube/config
go get
go run *.go
```

Then, you should be able to send an example request to it from another terminal window/tab:

```
$ curl -X POST -H "Content-Type: application/json" -d @examples/restart-deployment.json https://localhost:8080/webhooks
```

You'll see this service do a quick restart of the pods in the `default` namespace for the deployment `nginx-deployment`.

Take a look at the [example payload](examples/restart-deployment.json) FireHydrant sends when extensions are executed.
