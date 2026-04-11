This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppTest::TestCase M2NonManagedApplicationAction

Specifies the AWS Mainframe Modernization non-managed application action.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ActionType" : String,
  "Resource" : String
}

```

### YAML

```yaml

  ActionType: String
  Resource: String

```

## Properties

`ActionType`

The action type of the AWS Mainframe Modernization non-managed application action.

_Required_: Yes

_Type_: String

_Allowed values_: `Configure | Deconfigure`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Resource`

The resource of the AWS Mainframe Modernization non-managed application action.

_Required_: Yes

_Type_: String

_Pattern_: `^\S{1,1000}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

M2ManagedApplicationAction

MainframeAction

All content copied from https://docs.aws.amazon.com/.
