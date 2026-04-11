This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::XRay::SamplingRule Tag

A map that contains tag keys and tag values to attach to an AWS X-Ray group or sampling
rule. For more information about ways to use tags, see [Tagging AWS resources](../../../../general/latest/gr/aws-tagging.md)
in the _AWS General Reference_.

The following restrictions apply to tags:

- Maximum number of user-applied tags per resource: 50

- Tag keys and values are case sensitive.

- Don't use `aws:` as a prefix for keys; it's reserved for AWS use. You
cannot edit or delete system tags.

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

A tag key, such as `Stage` or `Name`. A tag key cannot be empty. The
key can be a maximum of 128 characters, and can contain only Unicode letters, numbers, or separators,
or the following special characters: `+ - = . _ : /`

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

An optional tag value, such as `Production` or `test-only`. The value can be
a maximum of 255 characters, and contain only Unicode letters, numbers, or separators, or the following
special characters: `+ - = . _ : /`

_Required_: Yes

_Type_: String

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SamplingRule

AWS::XRay::TransactionSearchConfig

All content copied from https://docs.aws.amazon.com/.
