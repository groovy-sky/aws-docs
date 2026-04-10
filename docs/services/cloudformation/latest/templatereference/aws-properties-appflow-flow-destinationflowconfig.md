This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppFlow::Flow DestinationFlowConfig

Contains information about the configuration of destination connectors present in the
flow.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ApiVersion" : String,
  "ConnectorProfileName" : String,
  "ConnectorType" : String,
  "DestinationConnectorProperties" : DestinationConnectorProperties
}

```

### YAML

```yaml

  ApiVersion: String
  ConnectorProfileName: String
  ConnectorType: String
  DestinationConnectorProperties:
    DestinationConnectorProperties

```

## Properties

`ApiVersion`

The API version that the destination connector uses.

_Required_: No

_Type_: String

_Pattern_: `\S+`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConnectorProfileName`

The name of the connector profile. This name must be unique for each connector profile in
the AWS account.

_Required_: No

_Type_: String

_Pattern_: `[\w/!@#+=.-]+`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConnectorType`

The type of destination connector, such as Sales force, Amazon S3, and so on.

_Required_: Yes

_Type_: String

_Allowed values_: `SAPOData | Salesforce | Pardot | Singular | Slack | Redshift | S3 | Marketo | Googleanalytics | Zendesk | Servicenow | Datadog | Trendmicro | Snowflake | Dynatrace | Infornexus | Amplitude | Veeva | CustomConnector | EventBridge | Upsolver | LookoutMetrics`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DestinationConnectorProperties`

This stores the information that is required to query a particular connector.

_Required_: Yes

_Type_: [DestinationConnectorProperties](aws-properties-appflow-flow-destinationconnectorproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [DestinationConnectorProperties](../../../../reference/appflow/1-0/apireference/api-destinationconnectorproperties.md) in the _Amazon AppFlow API_
_Reference_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DestinationConnectorProperties

DynatraceSourceProperties

All content copied from https://docs.aws.amazon.com/.
