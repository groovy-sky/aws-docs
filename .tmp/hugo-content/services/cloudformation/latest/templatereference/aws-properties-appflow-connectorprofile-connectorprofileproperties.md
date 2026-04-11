This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppFlow::ConnectorProfile ConnectorProfileProperties

The connector-specific profile properties required by each connector.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CustomConnector" : CustomConnectorProfileProperties,
  "Datadog" : DatadogConnectorProfileProperties,
  "Dynatrace" : DynatraceConnectorProfileProperties,
  "InforNexus" : InforNexusConnectorProfileProperties,
  "Marketo" : MarketoConnectorProfileProperties,
  "Pardot" : PardotConnectorProfileProperties,
  "Redshift" : RedshiftConnectorProfileProperties,
  "Salesforce" : SalesforceConnectorProfileProperties,
  "SAPOData" : SAPODataConnectorProfileProperties,
  "ServiceNow" : ServiceNowConnectorProfileProperties,
  "Slack" : SlackConnectorProfileProperties,
  "Snowflake" : SnowflakeConnectorProfileProperties,
  "Veeva" : VeevaConnectorProfileProperties,
  "Zendesk" : ZendeskConnectorProfileProperties
}

```

### YAML

```yaml

  CustomConnector:
    CustomConnectorProfileProperties
  Datadog:
    DatadogConnectorProfileProperties
  Dynatrace:
    DynatraceConnectorProfileProperties
  InforNexus:
    InforNexusConnectorProfileProperties
  Marketo:
    MarketoConnectorProfileProperties
  Pardot:
    PardotConnectorProfileProperties
  Redshift:
    RedshiftConnectorProfileProperties
  Salesforce:
    SalesforceConnectorProfileProperties
  SAPOData:
    SAPODataConnectorProfileProperties
  ServiceNow:
    ServiceNowConnectorProfileProperties
  Slack:
    SlackConnectorProfileProperties
  Snowflake:
    SnowflakeConnectorProfileProperties
  Veeva:
    VeevaConnectorProfileProperties
  Zendesk:
    ZendeskConnectorProfileProperties

```

## Properties

`CustomConnector`

The properties required by the custom connector.

_Required_: No

_Type_: [CustomConnectorProfileProperties](aws-properties-appflow-connectorprofile-customconnectorprofileproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Datadog`

The connector-specific properties required by Datadog.

_Required_: No

_Type_: [DatadogConnectorProfileProperties](aws-properties-appflow-connectorprofile-datadogconnectorprofileproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Dynatrace`

The connector-specific properties required by Dynatrace.

_Required_: No

_Type_: [DynatraceConnectorProfileProperties](aws-properties-appflow-connectorprofile-dynatraceconnectorprofileproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InforNexus`

The connector-specific properties required by Infor Nexus.

_Required_: No

_Type_: [InforNexusConnectorProfileProperties](aws-properties-appflow-connectorprofile-infornexusconnectorprofileproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Marketo`

The connector-specific properties required by Marketo.

_Required_: No

_Type_: [MarketoConnectorProfileProperties](aws-properties-appflow-connectorprofile-marketoconnectorprofileproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Pardot`

The connector-specific properties required by Salesforce Pardot.

_Required_: No

_Type_: [PardotConnectorProfileProperties](aws-properties-appflow-connectorprofile-pardotconnectorprofileproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Redshift`

The connector-specific properties required by Amazon Redshift.

_Required_: No

_Type_: [RedshiftConnectorProfileProperties](aws-properties-appflow-connectorprofile-redshiftconnectorprofileproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Salesforce`

The connector-specific properties required by Salesforce.

_Required_: No

_Type_: [SalesforceConnectorProfileProperties](aws-properties-appflow-connectorprofile-salesforceconnectorprofileproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SAPOData`

The connector-specific profile properties required when using SAPOData.

_Required_: No

_Type_: [SAPODataConnectorProfileProperties](aws-properties-appflow-connectorprofile-sapodataconnectorprofileproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServiceNow`

The connector-specific properties required by serviceNow.

_Required_: No

_Type_: [ServiceNowConnectorProfileProperties](aws-properties-appflow-connectorprofile-servicenowconnectorprofileproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Slack`

The connector-specific properties required by Slack.

_Required_: No

_Type_: [SlackConnectorProfileProperties](aws-properties-appflow-connectorprofile-slackconnectorprofileproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Snowflake`

The connector-specific properties required by Snowflake.

_Required_: No

_Type_: [SnowflakeConnectorProfileProperties](aws-properties-appflow-connectorprofile-snowflakeconnectorprofileproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Veeva`

The connector-specific properties required by Veeva.

_Required_: No

_Type_: [VeevaConnectorProfileProperties](aws-properties-appflow-connectorprofile-veevaconnectorprofileproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Zendesk`

The connector-specific properties required by Zendesk.

_Required_: No

_Type_: [ZendeskConnectorProfileProperties](aws-properties-appflow-connectorprofile-zendeskconnectorprofileproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [ConnectorProfileProperties](../../../../reference/appflow/1-0/apireference/api-connectorprofileproperties.md) in the _Amazon AppFlow API_
_Reference_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ConnectorProfileCredentials

CustomAuthCredentials

All content copied from https://docs.aws.amazon.com/.
