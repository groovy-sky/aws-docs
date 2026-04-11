This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MSK::Cluster Prometheus

Prometheus settings for open monitoring.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "JmxExporter" : JmxExporter,
  "NodeExporter" : NodeExporter
}

```

### YAML

```yaml

  JmxExporter:
    JmxExporter
  NodeExporter:
    NodeExporter

```

## Properties

`JmxExporter`

Indicates whether you want to enable or disable the JMX Exporter.

_Required_: No

_Type_: [JmxExporter](aws-properties-msk-cluster-jmxexporter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NodeExporter`

Indicates whether you want to enable or disable the Node Exporter.

_Required_: No

_Type_: [NodeExporter](aws-properties-msk-cluster-nodeexporter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OpenMonitoring

ProvisionedThroughput

All content copied from https://docs.aws.amazon.com/.
