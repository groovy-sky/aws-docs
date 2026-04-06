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

_Type_: [CustomConnectorProfileProperties](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appflow-connectorprofile-customconnectorprofileproperties.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Datadog`

The connector-specific properties required by Datadog.

_Required_: No

_Type_: [DatadogConnectorProfileProperties](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appflow-connectorprofile-datadogconnectorprofileproperties.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Dynatrace`

The connector-specific properties required by Dynatrace.

_Required_: No

_Type_: [DynatraceConnectorProfileProperties](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appflow-connectorprofile-dynatraceconnectorprofileproperties.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InforNexus`

The connector-specific properties required by Infor Nexus.

_Required_: No

_Type_: [InforNexusConnectorProfileProperties](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appflow-connectorprofile-infornexusconnectorprofileproperties.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Marketo`

The connector-specific properties required by Marketo.

_Required_: No

_Type_: [MarketoConnectorProfileProperties](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appflow-connectorprofile-marketoconnectorprofileproperties.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Pardot`

The connector-specific properties required by Salesforce Pardot.

_Required_: No

_Type_: [PardotConnectorProfileProperties](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appflow-connectorprofile-pardotconnectorprofileproperties.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Redshift`

The connector-specific properties required by Amazon Redshift.

_Required_: No

_Type_: [RedshiftConnectorProfileProperties](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appflow-connectorprofile-redshiftconnectorprofileproperties.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Salesforce`

The connector-specific properties required by Salesforce.

_Required_: No

_Type_: [SalesforceConnectorProfileProperties](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appflow-connectorprofile-salesforceconnectorprofileproperties.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SAPOData`

The connector-specific profile properties required when using SAPOData.

_Required_: No

_Type_: [SAPODataConnectorProfileProperties](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appflow-connectorprofile-sapodataconnectorprofileproperties.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServiceNow`

The connector-specific properties required by serviceNow.

_Required_: No

_Type_: [ServiceNowConnectorProfileProperties](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appflow-connectorprofile-servicenowconnectorprofileproperties.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Slack`

The connector-specific properties required by Slack.

_Required_: No

_Type_: [SlackConnectorProfileProperties](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appflow-connectorprofile-slackconnectorprofileproperties.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Snowflake`

The connector-specific properties required by Snowflake.

_Required_: No

_Type_: [SnowflakeConnectorProfileProperties](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appflow-connectorprofile-snowflakeconnectorprofileproperties.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Veeva`

The connector-specific properties required by Veeva.

_Required_: No

_Type_: [VeevaConnectorProfileProperties](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appflow-connectorprofile-veevaconnectorprofileproperties.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Zendesk`

The connector-specific properties required by Zendesk.

_Required_: No

_Type_: [ZendeskConnectorProfileProperties](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appflow-connectorprofile-zendeskconnectorprofileproperties.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [ConnectorProfileProperties](https://docs.aws.amazon.com/appflow/1.0/APIReference/API_ConnectorProfileProperties.html) in the _Amazon AppFlow API_
_Reference_.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ConnectorProfileCredentials

CustomAuthCredentials
