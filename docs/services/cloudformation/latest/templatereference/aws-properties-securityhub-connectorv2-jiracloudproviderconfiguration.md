This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecurityHub::ConnectorV2 JiraCloudProviderConfiguration

The initial configuration settings required to establish an integration between Security Hub CSPM and Jira Cloud.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ProjectKey" : String
}

```

### YAML

```yaml

  ProjectKey: String

```

## Properties

`ProjectKey`

The project key for a JiraCloud instance.

_Required_: Yes

_Type_: String

_Minimum_: `2`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::SecurityHub::ConnectorV2

Provider

All content copied from https://docs.aws.amazon.com/.
