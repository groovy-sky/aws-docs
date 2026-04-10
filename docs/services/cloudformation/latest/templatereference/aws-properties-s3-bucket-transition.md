This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::Bucket Transition

Specifies when an object transitions to a specified storage class. For more information about Amazon S3
lifecycle configuration rules, see [Transitioning Objects Using\
Amazon S3 Lifecycle](../../../s3/latest/dev/lifecycle-transition-general-considerations.md) in the _Amazon S3 User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "StorageClass" : String,
  "TransitionDate" : String,
  "TransitionInDays" : Integer
}

```

### YAML

```yaml

  StorageClass: String
  TransitionDate: String
  TransitionInDays: Integer

```

## Properties

`StorageClass`

The storage class to which you want the object to transition.

_Required_: Yes

_Type_: String

_Allowed values_: `DEEP_ARCHIVE | GLACIER | Glacier | GLACIER_IR | INTELLIGENT_TIERING | ONEZONE_IA | STANDARD_IA`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TransitionDate`

Indicates when objects are transitioned to the specified storage class. The date value must be in
ISO 8601 format. The time is always midnight UTC.

_Required_: Conditional

_Type_: String

_Pattern_: `^(\d{4})-(0[0-9]|1[0-2])-([0-2]\d|3[01])T([01]\d|2[0-4]):([0-5]\d):([0-6]\d)((\.\d{3})?)Z$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TransitionInDays`

Indicates the number of days after creation when objects are transitioned to the specified storage
class. If the specified storage class is `INTELLIGENT_TIERING`, `GLACIER_IR`,
`GLACIER`, or `DEEP_ARCHIVE`, valid values are `0` or positive
integers. If the specified storage class is `STANDARD_IA` or `ONEZONE_IA`, valid
values are positive integers greater than `30`. Be aware that some storage classes have a
minimum storage duration and that you're charged for transitioning objects before their minimum storage
duration. For more information, see [Constraints and considerations for transitions](../../../s3/latest/userguide/lifecycle-transition-general-considerations.md#lifecycle-configuration-constraints) in the _Amazon S3 User_
_Guide_.

_Required_: Conditional

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- AWS::S3::Bucket [Examples](../userguide/aws-properties-s3-bucket.md#aws-properties-s3-bucket--examples)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TopicConfiguration

VersioningConfiguration

All content copied from https://docs.aws.amazon.com/.
