This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMRServerless::Application Tag

A map of key-value pairs to help you manage EMR Serverless resources. One resource can have a maximum
of 50 tags. For more information, see [Tagging resources](../../../emr/latest/emr-serverless-userguide/tagging.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Key" : String,
  "Value" : String
}

```

### YAML

```yaml

  Key: String
  Value: String

```

## Properties

`Key`

The key to use in the tag.

_Required_: Yes

_Type_: String

_Pattern_: `^[A-Za-z0-9 /_.:=+@-]+$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The value of the tag.

_Required_: Yes

_Type_: String

_Pattern_: `^[A-Za-z0-9 /_.:=+@-]*$`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SchedulerConfiguration

WorkerConfiguration

All content copied from https://docs.aws.amazon.com/.
