This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::Bucket DefaultRetention

The container element for optionally specifying the default Object Lock retention settings for new
objects placed in the specified bucket.

###### Note

- The `DefaultRetention` settings require both a mode and a period.

- The `DefaultRetention` period can be either `Days` or `Years`
but you must select one. You cannot specify `Days` and `Years` at the same
time.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Days" : Integer,
  "Mode" : String,
  "Years" : Integer
}

```

### YAML

```yaml

  Days: Integer
  Mode: String
  Years: Integer

```

## Properties

`Days`

The number of days that you want to specify for the default retention period. If Object
Lock is turned on, you must specify `Mode` and specify either `Days` or
`Years`.

_Required_: Conditional

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Mode`

The default Object Lock retention mode you want to apply to new objects placed in the
specified bucket. If Object Lock is turned on, you must specify `Mode` and specify
either `Days` or `Years`.

_Required_: Conditional

_Type_: String

_Allowed values_: `COMPLIANCE | GOVERNANCE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Years`

The number of years that you want to specify for the default retention period. If Object
Lock is turned on, you must specify `Mode` and specify either `Days` or
`Years`.

_Required_: Conditional

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- AWS::S3::Bucket [Examples](../userguide/aws-properties-s3-bucket.md#aws-properties-s3-bucket--examples)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DataExport

DeleteMarkerReplication

All content copied from https://docs.aws.amazon.com/.
