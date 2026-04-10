This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::Bucket IntelligentTieringConfiguration

Specifies the S3 Intelligent-Tiering configuration for an Amazon S3 bucket.

For information about the S3 Intelligent-Tiering storage class, see [Storage class for\
automatically optimizing frequently and infrequently accessed objects](../../../s3/latest/dev/storage-class-intro.md#sc-dynamic-data-access).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Id" : String,
  "Prefix" : String,
  "Status" : String,
  "TagFilters" : [ TagFilter, ... ],
  "Tierings" : [ Tiering, ... ]
}

```

### YAML

```yaml

  Id: String
  Prefix: String
  Status: String
  TagFilters:
    - TagFilter
  Tierings:
    - Tiering

```

## Properties

`Id`

The ID used to identify the S3 Intelligent-Tiering configuration.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Prefix`

An object key name prefix that identifies the subset of objects to which the rule
applies.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Status`

Specifies the status of the configuration.

_Required_: Yes

_Type_: String

_Allowed values_: `Disabled | Enabled`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TagFilters`

A container for a key-value pair.

_Required_: No

_Type_: Array of [TagFilter](aws-properties-s3-bucket-tagfilter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tierings`

Specifies a list of S3 Intelligent-Tiering storage class tiers in the configuration. At
least one tier must be defined in the list. At most, you can specify two tiers in the list,
one for each available AccessTier: `ARCHIVE_ACCESS` and
`DEEP_ARCHIVE_ACCESS`.

###### Note

You only need Intelligent Tiering Configuration enabled on a bucket if you want to
automatically move objects stored in the Intelligent-Tiering storage class to Archive Access
or Deep Archive Access tiers.

_Required_: Yes

_Type_: Array of [Tiering](aws-properties-s3-bucket-tiering.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FilterRule

InventoryConfiguration

All content copied from https://docs.aws.amazon.com/.
