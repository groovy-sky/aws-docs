This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppFlow::Flow SourceFlowConfig

Contains information about the configuration of the source connector used in the flow.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ApiVersion" : String,
  "ConnectorProfileName" : String,
  "ConnectorType" : String,
  "IncrementalPullConfig" : IncrementalPullConfig,
  "SourceConnectorProperties" : SourceConnectorProperties
}

```

### YAML

```yaml

  ApiVersion: String
  ConnectorProfileName: String
  ConnectorType: String
  IncrementalPullConfig:
    IncrementalPullConfig
  SourceConnectorProperties:
    SourceConnectorProperties

```

## Properties

`ApiVersion`

The API version of the connector when it's used as a source in the flow.

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

The type of connector, such as Salesforce, Amplitude, and so on.

_Required_: Yes

_Type_: String

_Allowed values_: `SAPOData | Salesforce | Pardot | Singular | Slack | Redshift | S3 | Marketo | Googleanalytics | Zendesk | Servicenow | Datadog | Trendmicro | Snowflake | Dynatrace | Infornexus | Amplitude | Veeva | CustomConnector | EventBridge | Upsolver | LookoutMetrics`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IncrementalPullConfig`

Defines the configuration for a scheduled incremental data pull. If a valid configuration
is provided, the fields specified in the configuration are used when querying for the
incremental data pull.

_Required_: No

_Type_: [IncrementalPullConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appflow-flow-incrementalpullconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceConnectorProperties`

Specifies the information that is required to query a particular source connector.

_Required_: Yes

_Type_: [SourceConnectorProperties](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appflow-flow-sourceconnectorproperties.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [SourceFlowConfig](https://docs.aws.amazon.com/appflow/1.0/APIReference/API_SourceFlowConfig.html) in the _Amazon AppFlow API_
_Reference_.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

SourceConnectorProperties

SuccessResponseHandlingConfig
