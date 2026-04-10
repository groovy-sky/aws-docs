This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTAnalytics::Dataset OutputFileUriValue

The value of the variable as a structure that specifies an output file URI.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FileName" : String
}

```

### YAML

```yaml

  FileName: String

```

## Properties

`FileName`

The URI of the location where dataset contents are stored, usually the URI of a file in an
S3 bucket.

_Required_: Yes

_Type_: String

_Pattern_: `^[\w\.-]{1,255}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LateDataRuleConfiguration

QueryAction

All content copied from https://docs.aws.amazon.com/.
