This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataZone::SubscriptionTarget SubscriptionTargetForm

The details of the subscription target configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Content" : String,
  "FormName" : String
}

```

### YAML

```yaml

  Content: String
  FormName: String

```

## Properties

`Content`

The content of the subscription target configuration.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FormName`

The form name included in the subscription target configuration.

_Required_: Yes

_Type_: String

_Pattern_: `^(?![0-9_])\w+$|^_\w*[a-zA-Z0-9]\w*$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::DataZone::SubscriptionTarget

AWS::DataZone::UserProfile
