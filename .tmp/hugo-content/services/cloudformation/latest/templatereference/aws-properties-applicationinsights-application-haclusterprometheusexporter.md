This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApplicationInsights::Application HAClusterPrometheusExporter

The `AWS::ApplicationInsights::Application HAClusterPrometheusExporter`
property type defines the HA cluster Prometheus Exporter settings. For more information,
see the [component configuration](../../../amazoncloudwatch/latest/monitoring/component-config-sections.md#component-configuration-prometheus) in the CloudWatch Application Insights
documentation.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PrometheusPort" : String
}

```

### YAML

```yaml

  PrometheusPort: String

```

## Properties

`PrometheusPort`

The target port to which Prometheus sends metrics. If not specified, the default port 9668 is used.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CustomComponent

HANAPrometheusExporter

All content copied from https://docs.aws.amazon.com/.
