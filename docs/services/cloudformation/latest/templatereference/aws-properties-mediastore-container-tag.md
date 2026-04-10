This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaStore::Container Tag

A collection of tags associated with a container. Each tag consists of a key:value pair, which can be anything you define. Typically, the tag key
represents a category (such as "environment") and the tag value represents a specific value within that category (such as "test," "development," or
"production"). You can add up to 50
tags to each container. For more information about tagging, including naming and usage conventions, see [Tagging Resources in MediaStore](../../../mediastore/latest/ug/tagging.md).

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

Part of the key:value pair that defines a tag. You can use a tag key to describe a category of information, such as "customer." Tag keys are
case-sensitive.

_Required_: Yes

_Type_: String

_Pattern_: `[\p{L}\p{Z}\p{N}_.:/=+\-@]*`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

Part of the key:value pair that defines a tag. You can use a tag value to describe a specific value within a category, such as "companyA" or
"companyB." Tag values are case-sensitive.

_Required_: Yes

_Type_: String

_Pattern_: `[\p{L}\p{Z}\p{N}_.:/=+\-@]*`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MetricPolicyRule

Next

All content copied from https://docs.aws.amazon.com/.
