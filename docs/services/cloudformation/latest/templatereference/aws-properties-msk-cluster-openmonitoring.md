This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MSK::Cluster OpenMonitoring

JMX and Node monitoring for the MSK cluster.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Prometheus" : Prometheus
}

```

### YAML

```yaml

  Prometheus:
    Prometheus

```

## Properties

`Prometheus`

Prometheus exporter settings.

_Required_: Yes

_Type_: [Prometheus](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-msk-cluster-prometheus.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

NodeExporter

Prometheus
