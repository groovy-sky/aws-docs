This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFv2::LoggingConfiguration LabelNameCondition

A single label name condition for a condition in a logging filter.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LabelName" : String
}

```

### YAML

```yaml

  LabelName: String

```

## Properties

`LabelName`

The label name that a log record must contain in order to meet the condition. This must
be a fully qualified label name. Fully qualified labels have a prefix, optional namespaces, and label name. The prefix identifies the rule group or web ACL context of the rule that added the label.

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9A-Za-z_\-:]+$`

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Filter

LoggingFilter

All content copied from https://docs.aws.amazon.com/.
