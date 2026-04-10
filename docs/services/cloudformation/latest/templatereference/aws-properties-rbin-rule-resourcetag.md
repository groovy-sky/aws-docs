This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Rbin::Rule ResourceTag

\[Tag-level retention rules only\] Information about the resource tags used to identify resources that are retained by the retention
rule.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ResourceTagKey" : String,
  "ResourceTagValue" : String
}

```

### YAML

```yaml

  ResourceTagKey: String
  ResourceTagValue: String

```

## Properties

`ResourceTagKey`

The tag key.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceTagValue`

The tag value.

_Required_: Yes

_Type_: String

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Rbin::Rule

RetentionPeriod

All content copied from https://docs.aws.amazon.com/.
