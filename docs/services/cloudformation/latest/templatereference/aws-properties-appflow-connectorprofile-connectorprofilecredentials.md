This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppFlow::ConnectorProfile ConnectorProfileCredentials

The connector-specific credentials required by a connector.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Amplitude" : AmplitudeConnectorProfileCredentials,
  "CustomConnector" : CustomConnectorProfileCredentials,
  "Datadog" : DatadogConnectorProfileCredentials,
  "Dynatrace" : DynatraceConnectorProfileCredentials,
  "GoogleAnalytics" : GoogleAnalyticsConnectorProfileCredentials,
  "InforNexus" : InforNexusConnectorProfileCredentials,
  "Marketo" : MarketoConnectorProfileCredentials,
  "Pardot" : PardotConnectorProfileCredentials,
  "Redshift" : RedshiftConnectorProfileCredentials,
  "Salesforce" : SalesforceConnectorProfileCredentials,
  "SAPOData" : SAPODataConnectorProfileCredentials,
  "ServiceNow" : ServiceNowConnectorProfileCredentials,
  "Singular" : SingularConnectorProfileCredentials,
  "Slack" : SlackConnectorProfileCredentials,
  "Snowflake" : SnowflakeConnectorProfileCredentials,
  "Trendmicro" : TrendmicroConnectorProfileCredentials,
  "Veeva" : VeevaConnectorProfileCredentials,
  "Zendesk" : ZendeskConnectorProfileCredentials
}

```

### YAML

```yaml

  Amplitude:
    AmplitudeConnectorProfileCredentials
  CustomConnector:
    CustomConnectorProfileCredentials
  Datadog:
    DatadogConnectorProfileCredentials
  Dynatrace:
    DynatraceConnectorProfileCredentials
  GoogleAnalytics:
    GoogleAnalyticsConnectorProfileCredentials
  InforNexus:
    InforNexusConnectorProfileCredentials
  Marketo:
    MarketoConnectorProfileCredentials
  Pardot:
    PardotConnectorProfileCredentials
  Redshift:
    RedshiftConnectorProfileCredentials
  Salesforce:
    SalesforceConnectorProfileCredentials
  SAPOData:
    SAPODataConnectorProfileCredentials
  ServiceNow:
    ServiceNowConnectorProfileCredentials
  Singular:
    SingularConnectorProfileCredentials
  Slack:
    SlackConnectorProfileCredentials
  Snowflake:
    SnowflakeConnectorProfileCredentials
  Trendmicro:
    TrendmicroConnectorProfileCredentials
  Veeva:
    VeevaConnectorProfileCredentials
  Zendesk:
    ZendeskConnectorProfileCredentials

```

## Properties

`Amplitude`

The connector-specific credentials required when using Amplitude.

_Required_: No

_Type_: [AmplitudeConnectorProfileCredentials](aws-properties-appflow-connectorprofile-amplitudeconnectorprofilecredentials.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomConnector`

The connector-specific profile credentials that are required when using the custom
connector.

_Required_: No

_Type_: [CustomConnectorProfileCredentials](aws-properties-appflow-connectorprofile-customconnectorprofilecredentials.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Datadog`

The connector-specific credentials required when using Datadog.

_Required_: No

_Type_: [DatadogConnectorProfileCredentials](aws-properties-appflow-connectorprofile-datadogconnectorprofilecredentials.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Dynatrace`

The connector-specific credentials required when using Dynatrace.

_Required_: No

_Type_: [DynatraceConnectorProfileCredentials](aws-properties-appflow-connectorprofile-dynatraceconnectorprofilecredentials.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GoogleAnalytics`

The connector-specific credentials required when using Google Analytics.

_Required_: No

_Type_: [GoogleAnalyticsConnectorProfileCredentials](aws-properties-appflow-connectorprofile-googleanalyticsconnectorprofilecredentials.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InforNexus`

The connector-specific credentials required when using Infor Nexus.

_Required_: No

_Type_: [InforNexusConnectorProfileCredentials](aws-properties-appflow-connectorprofile-infornexusconnectorprofilecredentials.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Marketo`

The connector-specific credentials required when using Marketo.

_Required_: No

_Type_: [MarketoConnectorProfileCredentials](aws-properties-appflow-connectorprofile-marketoconnectorprofilecredentials.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Pardot`

The connector-specific credentials required when using Salesforce Pardot.

_Required_: No

_Type_: [PardotConnectorProfileCredentials](aws-properties-appflow-connectorprofile-pardotconnectorprofilecredentials.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Redshift`

The connector-specific credentials required when using Amazon Redshift.

_Required_: No

_Type_: [RedshiftConnectorProfileCredentials](aws-properties-appflow-connectorprofile-redshiftconnectorprofilecredentials.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Salesforce`

The connector-specific credentials required when using Salesforce.

_Required_: No

_Type_: [SalesforceConnectorProfileCredentials](aws-properties-appflow-connectorprofile-salesforceconnectorprofilecredentials.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SAPOData`

The connector-specific profile credentials required when using SAPOData.

_Required_: No

_Type_: [SAPODataConnectorProfileCredentials](aws-properties-appflow-connectorprofile-sapodataconnectorprofilecredentials.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServiceNow`

The connector-specific credentials required when using ServiceNow.

_Required_: No

_Type_: [ServiceNowConnectorProfileCredentials](aws-properties-appflow-connectorprofile-servicenowconnectorprofilecredentials.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Singular`

The connector-specific credentials required when using Singular.

_Required_: No

_Type_: [SingularConnectorProfileCredentials](aws-properties-appflow-connectorprofile-singularconnectorprofilecredentials.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Slack`

The connector-specific credentials required when using Slack.

_Required_: No

_Type_: [SlackConnectorProfileCredentials](aws-properties-appflow-connectorprofile-slackconnectorprofilecredentials.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Snowflake`

The connector-specific credentials required when using Snowflake.

_Required_: No

_Type_: [SnowflakeConnectorProfileCredentials](aws-properties-appflow-connectorprofile-snowflakeconnectorprofilecredentials.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Trendmicro`

The connector-specific credentials required when using Trend Micro.

_Required_: No

_Type_: [TrendmicroConnectorProfileCredentials](aws-properties-appflow-connectorprofile-trendmicroconnectorprofilecredentials.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Veeva`

The connector-specific credentials required when using Veeva.

_Required_: No

_Type_: [VeevaConnectorProfileCredentials](aws-properties-appflow-connectorprofile-veevaconnectorprofilecredentials.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Zendesk`

The connector-specific credentials required when using Zendesk.

_Required_: No

_Type_: [ZendeskConnectorProfileCredentials](aws-properties-appflow-connectorprofile-zendeskconnectorprofilecredentials.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [ConnectorProfileCredentials](../../../../reference/appflow/1-0/apireference/api-connectorprofilecredentials.md) in the _Amazon AppFlow API_
_Reference_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ConnectorProfileConfig

ConnectorProfileProperties

All content copied from https://docs.aws.amazon.com/.
