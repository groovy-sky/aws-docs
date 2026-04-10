This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::UsageProfile Tag

The `Tag` object represents a label that you can assign to an AWS resource. Each tag consists of a key and an optional value, both of which you define.

For more information about tags, and controlling access to resources in AWS Glue, see
[AWS Tags in AWS Glue](../../../glue/latest/dg/monitor-tags.md) and [Specifying AWS Glue Resource\
ARNs](../../../glue/latest/dg/glue-specifying-resource-arns.md) in the developer guide.

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

The tag key. The key is required when you create a tag on an object. The key is case-sensitive, and must not contain the prefix aws.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The tag value. The value is optional when you create a tag on an object. The value is case-sensitive, and must not contain the prefix aws.

_Required_: Yes

_Type_: String

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ProfileConfiguration

AWS::Glue::Workflow

All content copied from https://docs.aws.amazon.com/.
