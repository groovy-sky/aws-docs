This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QBusiness::Application QAppsConfiguration

Configuration information about Amazon Q Apps.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "QAppsControlMode" : String
}

```

### YAML

```yaml

  QAppsControlMode: String

```

## Properties

`QAppsControlMode`

Status information about whether end users can create and use Amazon Q Apps in the web
experience.

_Required_: Yes

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

PersonalizationConfiguration

QuickSightConfiguration
