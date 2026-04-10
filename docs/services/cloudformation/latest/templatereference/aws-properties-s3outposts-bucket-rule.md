This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3Outposts::Bucket Rule

A container for an Amazon S3 on Outposts bucket lifecycle rule.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AbortIncompleteMultipartUpload" : AbortIncompleteMultipartUpload,
  "ExpirationDate" : String,
  "ExpirationInDays" : Integer,
  "Filter" : Filter,
  "Id" : String,
  "Status" : String
}

```

### YAML

```yaml

  AbortIncompleteMultipartUpload:
    AbortIncompleteMultipartUpload
  ExpirationDate: String
  ExpirationInDays: Integer
  Filter:
    Filter
  Id: String
  Status: String

```

## Properties

`AbortIncompleteMultipartUpload`

The container for the abort incomplete multipart upload rule.

_Required_: No

_Type_: [AbortIncompleteMultipartUpload](aws-properties-s3outposts-bucket-abortincompletemultipartupload.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExpirationDate`

Specifies the expiration for the lifecycle of the object by specifying an expiry date.

_Required_: No

_Type_: String

_Pattern_: `^([0-2]\d{3})-(0[0-9]|1[0-2])-([0-2]\d|3[01])T([01]\d|2[0-4]):([0-5]\d):([0-6]\d)((\.\d{3})?)Z$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExpirationInDays`

Specifies the expiration for the lifecycle of the object in the form of days that the object has been in the S3 on Outposts bucket.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Filter`

The container for the filter of the lifecycle rule.

_Required_: No

_Type_: [Filter](aws-properties-s3outposts-bucket-filter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Id`

Property description not available.

_Required_: No

_Type_: String

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Status`

If `Enabled`, the rule is currently being applied. If `Disabled`,
the rule is not currently being applied.

_Required_: No

_Type_: String

_Allowed values_: `Enabled | Disabled`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LifecycleConfiguration

Tag

All content copied from https://docs.aws.amazon.com/.
