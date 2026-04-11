This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Wisdom::MessageTemplate AgentAttributes

Information about an agent.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FirstName" : String,
  "LastName" : String
}

```

### YAML

```yaml

  FirstName: String
  LastName: String

```

## Properties

`FirstName`

The agent’s first name as entered in their Amazon Connect user account.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `32767`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LastName`

The agent’s last name as entered in their Amazon Connect user account.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `32767`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Wisdom::MessageTemplate

Content

All content copied from https://docs.aws.amazon.com/.
