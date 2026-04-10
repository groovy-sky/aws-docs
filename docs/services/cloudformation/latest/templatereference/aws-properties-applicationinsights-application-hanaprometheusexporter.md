This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApplicationInsights::Application HANAPrometheusExporter

The `AWS::ApplicationInsights::Application HANAPrometheusExporter` property
type defines the HANA DB Prometheus Exporter settings. For more information, see the
[component configuration](../../../amazoncloudwatch/latest/monitoring/component-config-sections.md#component-configuration-prometheus) in the CloudWatch Application Insights
documentation.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AgreeToInstallHANADBClient" : Boolean,
  "HANAPort" : String,
  "HANASecretName" : String,
  "HANASID" : String,
  "PrometheusPort" : String
}

```

### YAML

```yaml

  AgreeToInstallHANADBClient: Boolean
  HANAPort: String
  HANASecretName: String
  HANASID: String
  PrometheusPort: String

```

## Properties

`AgreeToInstallHANADBClient`

Designates whether you agree to install the HANA DB client.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HANAPort`

The HANA database port by which the exporter will query HANA metrics.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HANASecretName`

The AWS Secrets Manager secret that stores HANA monitoring user credentials. The HANA Prometheus exporter uses these credentials to connect to the database and query HANA metrics.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HANASID`

The three-character SAP system ID (SID) of the SAP HANA system.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PrometheusPort`

The target port to which Prometheus sends metrics. If not specified, the default port 9668 is used.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HAClusterPrometheusExporter

JMXPrometheusExporter

All content copied from https://docs.aws.amazon.com/.
