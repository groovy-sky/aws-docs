This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::APS::Scraper

A scraper is a fully-managed agentless collector that discovers and pulls metrics
automatically. A scraper pulls metrics from Prometheus-compatible sources within an
Amazon EKS cluster, and sends them to your Amazon Managed Service for Prometheus
workspace. Scrapers are flexible. You can configure the scraper to control what metrics
are collected, the frequency of collection, what transformations are applied to the
metrics, and more.

An IAM role will be created for you that Amazon Managed Service for Prometheus uses to
access the metrics in your cluster. You must configure this role with a policy that
allows it to scrape metrics from your cluster. For more information, see [Configuring your Amazon EKS cluster](https://docs.aws.amazon.com/prometheus/latest/userguide/AMP-collector-how-to.html#AMP-collector-eks-setup) in the _Amazon Managed Service_
_for Prometheus User Guide_.

The `scrapeConfiguration` parameter contains the YAML configuration for the
scraper.

###### Note

For more information about collectors, including what metrics are collected, and
how to configure the scraper, see [Using an AWS managed collector](https://docs.aws.amazon.com/prometheus/latest/userguide/AMP-collector-how-to.html) in the _Amazon Managed Service_
_for Prometheus User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::APS::Scraper",
  "Properties" : {
      "Alias" : String,
      "Destination" : Destination,
      "RoleConfiguration" : RoleConfiguration,
      "ScrapeConfiguration" : ScrapeConfiguration,
      "ScraperLoggingConfiguration" : ScraperLoggingConfiguration,
      "Source" : Source,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::APS::Scraper
Properties:
  Alias: String
  Destination:
    Destination
  RoleConfiguration:
    RoleConfiguration
  ScrapeConfiguration:
    ScrapeConfiguration
  ScraperLoggingConfiguration:
    ScraperLoggingConfiguration
  Source:
    Source
  Tags:
    - Tag

```

## Properties

`Alias`

An optional user-assigned scraper alias.

_Required_: No

_Type_: String

_Pattern_: `^[0-9A-Za-z][-.0-9A-Z_a-z]*$`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Destination`

The Amazon Managed Service for Prometheus workspace the scraper sends metrics
to.

_Required_: Yes

_Type_: [Destination](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-aps-scraper-destination.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleConfiguration`

The role configuration in an Amazon Managed Service for Prometheus scraper.

_Required_: No

_Type_: [RoleConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-aps-scraper-roleconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScrapeConfiguration`

The configuration in use by the scraper.

_Required_: Yes

_Type_: [ScrapeConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-aps-scraper-scrapeconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScraperLoggingConfiguration`

The definition of logging configuration in an Amazon Managed Service for Prometheus workspace.

_Required_: No

_Type_: [ScraperLoggingConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-aps-scraper-scraperloggingconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Source`

The Amazon EKS cluster from which the scraper collects metrics.

_Required_: Yes

_Type_: [Source](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-aps-scraper-source.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

(Optional) The list of tag keys and values associated with the scraper.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-aps-scraper-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource name. For example:

`{ "Ref": "Id" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the scraper. For example,
`arn:aws:aps:<region>:123456798012:scraper/s-example1-1234-abcd-5678-ef9012abcd34`.

`RoleArn`

The Amazon Resource Name (ARN) of the IAM role that provides permissions for the
scraper to discover and collect metrics on your behalf.

For example,
`arn:aws:iam::123456789012:role/service-role/AmazonGrafanaServiceRole-12example`.

`ScraperId`

The ID of the scraper. For example,
`s-example1-1234-abcd-5678-ef9012abcd34`.

## Examples

### Amazon Managed Service for Prometheus scraper example

The following example creates a scraper, showing how to pass in configuration.
You must replace values in braces ( `{}`) with values that match your
system.

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09 Resources: Scraper: Type:
                AWS::APS::Scraper UpdateReplacePolicy: Retain DeletionPolicy: Delete Properties:
                Alias: "Scraper-Test" Source: EksConfiguration: ClusterArn:
                "arn:aws:eks:{region}:{account-id}:cluster/{cluster-name}" SecurityGroupIds: -
                "sg-{security-group-id}" SubnetIds: - "subnet-{subnet-id-1}" -
                "subnet-{subnet-id-2}" Destination: AmpConfiguration: WorkspaceArn:
                "arn:aws:aps:{region}:{account-id}:workspace/ws-{workspace-id}" ScrapeConfiguration:
                ConfigurationBlob: "global:\n scrape_interval: 30s\n # external_labels:\n #
                clusterArn: {cluster-arn}\nscrape_configs:\n # pod metrics\n - job_name:
                pod_exporter\n kubernetes_sd_configs:\n - role: pod\n # container metrics\n -
                job_name: cadvisor\n scheme: https\n authorization:\n credentials_file:
                /var/run/secrets/kubernetes.io/serviceaccount/token\n kubernetes_sd_configs:\n -
                role: node\n relabel_configs:\n - action: labelmap\n regex:
                __meta_kubernetes_node_label_(.+)\n - replacement: kubernetes.default.svc:443\n
                target_label: __address__\n - source_labels: [__meta_kubernetes_node_name]\n regex:
                (.+)\n target_label: __metrics_path__\n replacement:
                /api/v1/nodes/$1/proxy/metrics/cadvisor\n # apiserver metrics\n - bearer_token_file:
                /var/run/secrets/kubernetes.io/serviceaccount/token\n job_name:
                kubernetes-apiservers\n kubernetes_sd_configs:\n - role: endpoints\n
                relabel_configs:\n - action: keep\n regex: default;kubernetes;https\n
                source_labels:\n - __meta_kubernetes_namespace\n - __meta_kubernetes_service_name\n
                - __meta_kubernetes_endpoint_port_name\n scheme: https\n # kube proxy metrics\n -
                job_name: kube-proxy\n honor_labels: true\n kubernetes_sd_configs:\n - role: pod\n
                relabel_configs:\n - action: keep\n source_labels:\n - __meta_kubernetes_namespace\n
                - __meta_kubernetes_pod_name\n separator: '/'\n regex: 'kube-system/kube-proxy.+'\n
                - source_labels:\n - __address__\n action: replace\n target_label: __address__\n
                regex: (.+?)(\\:\\d+)?\n replacement: $1:10249" Tags: - Key: "BusinessPurpose"
                Value: "Testing"
```

#### JSON

```json

{ "CreateInputs": { "Alias": "Scraper-Test",
                "ScrapeConfiguration": { "ConfigurationBlob": "global:\n scrape_interval: 30s\n #
                external_labels:\n # clusterArn: {cluster-arn}\nscrape_configs:\n # pod metrics\n -
                job_name: pod_exporter\n kubernetes_sd_configs:\n - role: pod\n # container
                metrics\n - job_name: cadvisor\n scheme: https\n authorization:\n credentials_file:
                /var/run/secrets/kubernetes.io/serviceaccount/token\n kubernetes_sd_configs:\n -
                role: node\n relabel_configs:\n - action: labelmap\n regex:
                __meta_kubernetes_node_label_(.+)\n - replacement: kubernetes.default.svc:443\n
                target_label: __address__\n - source_labels: [__meta_kubernetes_node_name]\n regex:
                (.+)\n target_label: __metrics_path__\n replacement:
                /api/v1/nodes/$1/proxy/metrics/cadvisor\n # apiserver metrics\n - bearer_token_file:
                /var/run/secrets/kubernetes.io/serviceaccount/token\n job_name:
                kubernetes-apiservers\n kubernetes_sd_configs:\n - role: endpoints\n
                relabel_configs:\n - action: keep\n regex: default;kubernetes;https\n
                source_labels:\n - __meta_kubernetes_namespace\n - __meta_kubernetes_service_name\n
                - __meta_kubernetes_endpoint_port_name\n scheme: https\n # kube proxy metrics\n -
                job_name: kube-proxy\n honor_labels: true\n kubernetes_sd_configs:\n - role: pod\n
                relabel_configs:\n - action: keep\n source_labels:\n - __meta_kubernetes_namespace\n
                - __meta_kubernetes_pod_name\n separator: '/'\n regex: 'kube-system/kube-proxy.+'\n
                - source_labels:\n - __address__\n action: replace\n target_label: __address__\n
                regex: (.+?)(\\:\\d+)?\n replacement: $1:10249" }, "Source": { "EksConfiguration": {
                "ClusterArn": "arn:aws:eks:{region}:{account-id}:cluster/{cluster-name}",
                "SecurityGroupIds": [ "sg-{security-group-id}" ], "SubnetIds": [
                "subnet-{subnet-id-1}", "subnet-{subnet-id-2}" ] } }, "Destination": {
                "AmpConfiguration": { "WorkspaceArn":
                "arn:aws:aps:{region}:{account-id}:workspace/ws-{workspace-id}" } }, "Tags": [ {
                "Key": "BusinessPurpose", "Value": "Testing" } ] }, "PatchInputs": [] }
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AmpConfiguration
