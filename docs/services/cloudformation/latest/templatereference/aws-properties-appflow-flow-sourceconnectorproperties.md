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

_Type_: [AmplitudeSourceProperties](aws-properties-appflow-flow-amplitudesourceproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomConnector`

The properties that are applied when the custom connector is being used as a
source.

_Required_: No

_Type_: [CustomConnectorSourceProperties](aws-properties-appflow-flow-customconnectorsourceproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Datadog`

Specifies the information that is required for querying Datadog.

_Required_: No

_Type_: [DatadogSourceProperties](aws-properties-appflow-flow-datadogsourceproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Dynatrace`

Specifies the information that is required for querying Dynatrace.

_Required_: No

_Type_: [DynatraceSourceProperties](aws-properties-appflow-flow-dynatracesourceproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GoogleAnalytics`

Specifies the information that is required for querying Google Analytics.

_Required_: No

_Type_: [GoogleAnalyticsSourceProperties](aws-properties-appflow-flow-googleanalyticssourceproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InforNexus`

Specifies the information that is required for querying Infor Nexus.

_Required_: No

_Type_: [InforNexusSourceProperties](aws-properties-appflow-flow-infornexussourceproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Marketo`

Specifies the information that is required for querying Marketo.

_Required_: No

_Type_: [MarketoSourceProperties](aws-properties-appflow-flow-marketosourceproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Pardot`

Specifies the information that is required for querying Salesforce Pardot.

_Required_: No

_Type_: [PardotSourceProperties](aws-properties-appflow-flow-pardotsourceproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3`

Specifies the information that is required for querying Amazon S3.

_Required_: No

_Type_: [S3SourceProperties](aws-properties-appflow-flow-s3sourceproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Salesforce`

Specifies the information that is required for querying Salesforce.

_Required_: No

_Type_: [SalesforceSourceProperties](aws-properties-appflow-flow-salesforcesourceproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SAPOData`

The properties that are applied when using SAPOData as a flow source.

_Required_: No

_Type_: [SAPODataSourceProperties](aws-properties-appflow-flow-sapodatasourceproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServiceNow`

Specifies the information that is required for querying ServiceNow.

_Required_: No

_Type_: [ServiceNowSourceProperties](aws-properties-appflow-flow-servicenowsourceproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Singular`

Specifies the information that is required for querying Singular.

_Required_: No

_Type_: [SingularSourceProperties](aws-properties-appflow-flow-singularsourceproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Slack`

Specifies the information that is required for querying Slack.

_Required_: No

_Type_: [SlackSourceProperties](aws-properties-appflow-flow-slacksourceproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Trendmicro`

Specifies the information that is required for querying Trend Micro.

_Required_: No

_Type_: [TrendmicroSourceProperties](aws-properties-appflow-flow-trendmicrosourceproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Veeva`

Specifies the information that is required for querying Veeva.

_Required_: No

_Type_: [VeevaSourceProperties](aws-properties-appflow-flow-veevasourceproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Zendesk`

Specifies the information that is required for querying Zendesk.

_Required_: No

_Type_: [ZendeskSourceProperties](aws-properties-appflow-flow-zendesksourceproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [SourceConnectorProperties](../../../../reference/appflow/1-0/apireference/api-sourceconnectorproperties.md) in the _Amazon AppFlow API_
_Reference_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SnowflakeDestinationProperties

SourceFlowConfig

All content copied from https://docs.aws.amazon.com/.
