This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppFlow::Flow SourceConnectorProperties

Specifies the information that is required to query a particular connector.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Amplitude" : AmplitudeSourceProperties,
  "CustomConnector" : CustomConnectorSourceProperties,
  "Datadog" : DatadogSourceProperties,
  "Dynatrace" : DynatraceSourceProperties,
  "GoogleAnalytics" : GoogleAnalyticsSourceProperties,
  "InforNexus" : InforNexusSourceProperties,
  "Marketo" : MarketoSourceProperties,
  "Pardot" : PardotSourceProperties,
  "S3" : S3SourceProperties,
  "Salesforce" : SalesforceSourceProperties,
  "SAPOData" : SAPODataSourceProperties,
  "ServiceNow" : ServiceNowSourceProperties,
  "Singular" : SingularSourceProperties,
  "Slack" : SlackSourceProperties,
  "Trendmicro" : TrendmicroSourceProperties,
  "Veeva" : VeevaSourceProperties,
  "Zendesk" : ZendeskSourceProperties
}

```

### YAML

```yaml

  Amplitude:
    AmplitudeSourceProperties
  CustomConnector:
    CustomConnectorSourceProperties
  Datadog:
    DatadogSourceProperties
  Dynatrace:
    DynatraceSourceProperties
  GoogleAnalytics:
    GoogleAnalyticsSourceProperties
  InforNexus:
    InforNexusSourceProperties
  Marketo:
    MarketoSourceProperties
  Pardot:
    PardotSourceProperties
  S3:
    S3SourceProperties
  Salesforce:
    SalesforceSourceProperties
  SAPOData:
    SAPODataSourceProperties
  ServiceNow:
    ServiceNowSourceProperties
  Singular:
    SingularSourceProperties
  Slack:
    SlackSourceProperties
  Trendmicro:
    TrendmicroSourceProperties
  Veeva:
    VeevaSourceProperties
  Zendesk:
    ZendeskSourceProperties

```

## Properties

`Amplitude`

Specifies the information that is required for querying Amplitude.

_Required_: No

_Type_: [AmplitudeSourceProperties](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appflow-flow-amplitudesourceproperties.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomConnector`

The properties that are applied when the custom connector is being used as a
source.

_Required_: No

_Type_: [CustomConnectorSourceProperties](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appflow-flow-customconnectorsourceproperties.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Datadog`

Specifies the information that is required for querying Datadog.

_Required_: No

_Type_: [DatadogSourceProperties](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appflow-flow-datadogsourceproperties.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Dynatrace`

Specifies the information that is required for querying Dynatrace.

_Required_: No

_Type_: [DynatraceSourceProperties](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appflow-flow-dynatracesourceproperties.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GoogleAnalytics`

Specifies the information that is required for querying Google Analytics.

_Required_: No

_Type_: [GoogleAnalyticsSourceProperties](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appflow-flow-googleanalyticssourceproperties.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InforNexus`

Specifies the information that is required for querying Infor Nexus.

_Required_: No

_Type_: [InforNexusSourceProperties](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appflow-flow-infornexussourceproperties.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Marketo`

Specifies the information that is required for querying Marketo.

_Required_: No

_Type_: [MarketoSourceProperties](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appflow-flow-marketosourceproperties.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Pardot`

Specifies the information that is required for querying Salesforce Pardot.

_Required_: No

_Type_: [PardotSourceProperties](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appflow-flow-pardotsourceproperties.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3`

Specifies the information that is required for querying Amazon S3.

_Required_: No

_Type_: [S3SourceProperties](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appflow-flow-s3sourceproperties.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Salesforce`

Specifies the information that is required for querying Salesforce.

_Required_: No

_Type_: [SalesforceSourceProperties](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appflow-flow-salesforcesourceproperties.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SAPOData`

The properties that are applied when using SAPOData as a flow source.

_Required_: No

_Type_: [SAPODataSourceProperties](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appflow-flow-sapodatasourceproperties.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServiceNow`

Specifies the information that is required for querying ServiceNow.

_Required_: No

_Type_: [ServiceNowSourceProperties](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appflow-flow-servicenowsourceproperties.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Singular`

Specifies the information that is required for querying Singular.

_Required_: No

_Type_: [SingularSourceProperties](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appflow-flow-singularsourceproperties.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Slack`

Specifies the information that is required for querying Slack.

_Required_: No

_Type_: [SlackSourceProperties](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appflow-flow-slacksourceproperties.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Trendmicro`

Specifies the information that is required for querying Trend Micro.

_Required_: No

_Type_: [TrendmicroSourceProperties](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appflow-flow-trendmicrosourceproperties.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Veeva`

Specifies the information that is required for querying Veeva.

_Required_: No

_Type_: [VeevaSourceProperties](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appflow-flow-veevasourceproperties.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Zendesk`

Specifies the information that is required for querying Zendesk.

_Required_: No

_Type_: [ZendeskSourceProperties](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appflow-flow-zendesksourceproperties.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [SourceConnectorProperties](https://docs.aws.amazon.com/appflow/1.0/APIReference/API_SourceConnectorProperties.html) in the _Amazon AppFlow API_
_Reference_.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

SnowflakeDestinationProperties

SourceFlowConfig
