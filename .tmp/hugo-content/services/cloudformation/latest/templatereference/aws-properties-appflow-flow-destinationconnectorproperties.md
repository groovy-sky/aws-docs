This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppFlow::Flow DestinationConnectorProperties

This stores the information that is required to query a particular connector.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CustomConnector" : CustomConnectorDestinationProperties,
  "EventBridge" : EventBridgeDestinationProperties,
  "LookoutMetrics" : LookoutMetricsDestinationProperties,
  "Marketo" : MarketoDestinationProperties,
  "Redshift" : RedshiftDestinationProperties,
  "S3" : S3DestinationProperties,
  "Salesforce" : SalesforceDestinationProperties,
  "SAPOData" : SAPODataDestinationProperties,
  "Snowflake" : SnowflakeDestinationProperties,
  "Upsolver" : UpsolverDestinationProperties,
  "Zendesk" : ZendeskDestinationProperties
}

```

### YAML

```yaml

  CustomConnector:
    CustomConnectorDestinationProperties
  EventBridge:
    EventBridgeDestinationProperties
  LookoutMetrics:
    LookoutMetricsDestinationProperties
  Marketo:
    MarketoDestinationProperties
  Redshift:
    RedshiftDestinationProperties
  S3:
    S3DestinationProperties
  Salesforce:
    SalesforceDestinationProperties
  SAPOData:
    SAPODataDestinationProperties
  Snowflake:
    SnowflakeDestinationProperties
  Upsolver:
    UpsolverDestinationProperties
  Zendesk:
    ZendeskDestinationProperties

```

## Properties

`CustomConnector`

The properties that are required to query the custom Connector.

_Required_: No

_Type_: [CustomConnectorDestinationProperties](aws-properties-appflow-flow-customconnectordestinationproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EventBridge`

The properties required to query Amazon EventBridge.

_Required_: No

_Type_: [EventBridgeDestinationProperties](aws-properties-appflow-flow-eventbridgedestinationproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LookoutMetrics`

The properties required to query Amazon Lookout for Metrics.

_Required_: No

_Type_: [LookoutMetricsDestinationProperties](aws-properties-appflow-flow-lookoutmetricsdestinationproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Marketo`

The properties required to query Marketo.

_Required_: No

_Type_: [MarketoDestinationProperties](aws-properties-appflow-flow-marketodestinationproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Redshift`

The properties required to query Amazon Redshift.

_Required_: No

_Type_: [RedshiftDestinationProperties](aws-properties-appflow-flow-redshiftdestinationproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3`

The properties required to query Amazon S3.

_Required_: No

_Type_: [S3DestinationProperties](aws-properties-appflow-flow-s3destinationproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Salesforce`

The properties required to query Salesforce.

_Required_: No

_Type_: [SalesforceDestinationProperties](aws-properties-appflow-flow-salesforcedestinationproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SAPOData`

The properties required to query SAPOData.

_Required_: No

_Type_: [SAPODataDestinationProperties](aws-properties-appflow-flow-sapodatadestinationproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Snowflake`

The properties required to query Snowflake.

_Required_: No

_Type_: [SnowflakeDestinationProperties](aws-properties-appflow-flow-snowflakedestinationproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Upsolver`

The properties required to query Upsolver.

_Required_: No

_Type_: [UpsolverDestinationProperties](aws-properties-appflow-flow-upsolverdestinationproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Zendesk`

The properties required to query Zendesk.

_Required_: No

_Type_: [ZendeskDestinationProperties](aws-properties-appflow-flow-zendeskdestinationproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [DestinationConnectorProperties](../../../../reference/appflow/1-0/apireference/api-destinationconnectorproperties.md) in the _Amazon AppFlow API_
_Reference_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DataTransferApi

DestinationFlowConfig

All content copied from https://docs.aws.amazon.com/.
