This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppFlow::ConnectorProfile PardotConnectorProfileProperties

The connector-specific profile properties required when using Salesforce Pardot.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BusinessUnitId" : String,
  "InstanceUrl" : String,
  "IsSandboxEnvironment" : Boolean
}

```

### YAML

```yaml

  BusinessUnitId: String
  InstanceUrl: String
  IsSandboxEnvironment: Boolean

```

## Properties

`BusinessUnitId`

The business unit id of Salesforce Pardot instance.

_Required_: Yes

_Type_: String

_Pattern_: `\S+`

_Maximum_: `18`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceUrl`

The location of the Salesforce Pardot resource.

_Required_: No

_Type_: String

_Pattern_: `\S+`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IsSandboxEnvironment`

Indicates whether the connector profile applies to a sandbox or production
environment.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

PardotConnectorProfileCredentials

RedshiftConnectorProfileCredentials
