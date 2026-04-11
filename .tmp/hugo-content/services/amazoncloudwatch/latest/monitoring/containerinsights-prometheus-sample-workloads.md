# (Optional) Set up sample containerized Amazon EKS workloads for Prometheus metric testing

To test the Prometheus metric support in CloudWatch Container Insights, you can set up
one or more of the following containerized workloads. The CloudWatch agent with Prometheus
support automatically collects metrics from each of these workloads. To see the
metrics that are collected by default, see [Prometheus metrics collected by the CloudWatch agent](containerinsights-prometheus-metrics.md).

Before you can install any of these workloads, you must install Helm 3.x by
entering the following commands:

```

brew install helm
```

For more information, see [Helm](https://helm.sh/).

###### Topics

- [Set up AWS App Mesh sample workload for Amazon EKS and Kubernetes](containerinsights-prometheus-sample-workloads-appmesh.md)

- [Set up NGINX with sample traffic on Amazon EKS and Kubernetes](containerinsights-prometheus-sample-workloads-nginx.md)

- [Set up memcached with a metric exporter on Amazon EKS and Kubernetes](containerinsights-prometheus-sample-workloads-memcached.md)

- [Set up Java/JMX sample workload on Amazon EKS and Kubernetes](containerinsights-prometheus-sample-workloads-javajmx.md)

- [Set up HAProxy with a metric exporter on Amazon EKS and Kubernetes](containerinsights-prometheus-sample-workloads-haproxy.md)

- [Tutorial for adding a new Prometheus scrape target: Redis OSS on Amazon EKS and Kubernetes clusters](containerinsights-prometheus-setup-redis-eks.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Scraping additional Prometheus sources and importing those metrics

Set up AWS App Mesh sample workload for Amazon EKS and Kubernetes

All content copied from https://docs.aws.amazon.com/.
