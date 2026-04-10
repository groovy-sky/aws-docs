This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSet RefreshFailureEmailAlert

The configuration settings for the email alerts that are sent when a dataset refresh fails.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AlertStatus" : String
}

```

### YAML

```yaml

  AlertStatus: String

```

## Properties

`AlertStatus`

The status value that determines if email alerts are sent.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RefreshFailureConfiguration

RelationalTable

All content copied from https://docs.aws.amazon.com/.
