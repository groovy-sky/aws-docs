This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CustomerProfiles::Integration SourceFlowConfig

The configuration that controls how Customer Profiles retrieves data from the
source.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ConnectorProfileName" : String,
  "ConnectorType" : String,
  "IncrementalPullConfig" : IncrementalPullConfig,
  "SourceConnectorProperties" : SourceConnectorProperties
}

```

### YAML

```yaml

  ConnectorProfileName: String
  ConnectorType: String
  IncrementalPullConfig:
    IncrementalPullConfig
  SourceConnectorProperties:
    SourceConnectorProperties

```

## Properties

`ConnectorProfileName`

The name of the Amazon AppFlow connector profile. This name must be unique for each
connector profile in the AWS account.

_Required_: No

_Type_: String

_Pattern_: `[\w/!@#+=.-]+`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConnectorType`

The type of connector, such as Salesforce, Marketo, and so on.

_Required_: Yes

_Type_: String

_Allowed values_: `Salesforce | Marketo | ServiceNow | Zendesk | S3`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IncrementalPullConfig`

Defines the configuration for a scheduled incremental data pull. If a valid
configuration is provided, the fields specified in the configuration are used when
querying for the incremental data pull.

_Required_: No

_Type_: [IncrementalPullConfig](aws-properties-customerprofiles-integration-incrementalpullconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceConnectorProperties`

Specifies the information that is required to query a particular source
connector.

_Required_: Yes

_Type_: [SourceConnectorProperties](aws-properties-customerprofiles-integration-sourceconnectorproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SourceConnectorProperties

Tag

All content copied from https://docs.aws.amazon.com/.
