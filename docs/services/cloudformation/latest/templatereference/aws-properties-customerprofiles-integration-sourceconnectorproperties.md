This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CustomerProfiles::Integration SourceConnectorProperties

Specifies the information that is required to query a particular Amazon AppFlow
connector. Customer Profiles supports Salesforce, Zendesk, Marketo, ServiceNow and
Amazon S3.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Marketo" : MarketoSourceProperties,
  "S3" : S3SourceProperties,
  "Salesforce" : SalesforceSourceProperties,
  "ServiceNow" : ServiceNowSourceProperties,
  "Zendesk" : ZendeskSourceProperties
}

```

### YAML

```yaml

  Marketo:
    MarketoSourceProperties
  S3:
    S3SourceProperties
  Salesforce:
    SalesforceSourceProperties
  ServiceNow:
    ServiceNowSourceProperties
  Zendesk:
    ZendeskSourceProperties

```

## Properties

`Marketo`

The properties that are applied when Marketo is being used as a source.

_Required_: No

_Type_: [MarketoSourceProperties](aws-properties-customerprofiles-integration-marketosourceproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3`

The properties that are applied when Amazon S3 is being used as the flow
source.

_Required_: No

_Type_: [S3SourceProperties](aws-properties-customerprofiles-integration-s3sourceproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Salesforce`

The properties that are applied when Salesforce is being used as a source.

_Required_: No

_Type_: [SalesforceSourceProperties](aws-properties-customerprofiles-integration-salesforcesourceproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServiceNow`

The properties that are applied when ServiceNow is being used as a source.

_Required_: No

_Type_: [ServiceNowSourceProperties](aws-properties-customerprofiles-integration-servicenowsourceproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Zendesk`

The properties that are applied when using Zendesk as a flow source.

_Required_: No

_Type_: [ZendeskSourceProperties](aws-properties-customerprofiles-integration-zendesksourceproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ServiceNowSourceProperties

SourceFlowConfig

All content copied from https://docs.aws.amazon.com/.
