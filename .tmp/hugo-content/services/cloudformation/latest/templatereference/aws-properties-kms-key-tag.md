This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KMS::Key Tag

A key-value pair. A tag consists of a tag key and a tag value. Tag keys and tag values are
both required, but tag values can be empty (null) strings.

###### Important

Do not include confidential or sensitive information in this field. This field may be displayed in plaintext in CloudTrail logs and other output.

For information about the rules that apply to tag keys and tag values, see [User-Defined Tag Restrictions](../../../awsaccountbilling/latest/aboutv2/allocation-tag-restrictions.md) in the _AWS Billing and Cost Management_
_User Guide_.

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

The key name of the tag. You can specify a value that's 1 to 128 Unicode characters in length and can't be prefixed with `aws:`. digits, whitespace, `_`, `.`, `:`, `/`, `=`, `+`, `@`, `-`, and `"`.

For more information, see [Tag](../userguide/aws-properties-resource-tags.md).

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The value for the tag. You can specify a value that's 1 to 256 characters in length. You can use any of the following characters: the set of Unicode letters, digits, whitespace, `_`, `.`, `/`, `=`, `+`, and `-`.

For more information, see [Tag](../userguide/aws-properties-resource-tags.md).

_Required_: Yes

_Type_: String

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::KMS::Key

AWS::KMS::ReplicaKey

All content copied from https://docs.aws.amazon.com/.
