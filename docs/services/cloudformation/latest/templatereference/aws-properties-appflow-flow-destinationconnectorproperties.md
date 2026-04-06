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

_Type_: [CustomConnectorDestinationProperties](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appflow-flow-customconnectordestinationproperties.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EventBridge`

The properties required to query Amazon EventBridge.

_Required_: No

_Type_: [EventBridgeDestinationProperties](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appflow-flow-eventbridgedestinationproperties.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LookoutMetrics`

The properties required to query Amazon Lookout for Metrics.

_Required_: No

_Type_: [LookoutMetricsDestinationProperties](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appflow-flow-lookoutmetricsdestinationproperties.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Marketo`

The properties required to query Marketo.

_Required_: No

_Type_: [MarketoDestinationProperties](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appflow-flow-marketodestinationproperties.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Redshift`

The properties required to query Amazon Redshift.

_Required_: No

_Type_: [RedshiftDestinationProperties](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appflow-flow-redshiftdestinationproperties.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3`

The properties required to query Amazon S3.

_Required_: No

_Type_: [S3DestinationProperties](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appflow-flow-s3destinationproperties.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Salesforce`

The properties required to query Salesforce.

_Required_: No

_Type_: [SalesforceDestinationProperties](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appflow-flow-salesforcedestinationproperties.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SAPOData`

The properties required to query SAPOData.

_Required_: No

_Type_: [SAPODataDestinationProperties](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appflow-flow-sapodatadestinationproperties.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Snowflake`

The properties required to query Snowflake.

_Required_: No

_Type_: [SnowflakeDestinationProperties](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appflow-flow-snowflakedestinationproperties.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Upsolver`

The properties required to query Upsolver.

_Required_: No

_Type_: [UpsolverDestinationProperties](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appflow-flow-upsolverdestinationproperties.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Zendesk`

The properties required to query Zendesk.

_Required_: No

_Type_: [ZendeskDestinationProperties](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appflow-flow-zendeskdestinationproperties.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [DestinationConnectorProperties](https://docs.aws.amazon.com/appflow/1.0/APIReference/API_DestinationConnectorProperties.html) in the _Amazon AppFlow API_
_Reference_.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DataTransferApi

DestinationFlowConfig
