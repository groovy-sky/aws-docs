This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::DataSource S3DataSourceConfiguration

The configuration information to connect to Amazon S3 as your data source.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BucketArn" : String,
  "BucketOwnerAccountId" : String,
  "InclusionPrefixes" : [ String, ... ]
}

```

### YAML

```yaml

  BucketArn: String
  BucketOwnerAccountId: String
  InclusionPrefixes:
    - String

```

## Properties

`BucketArn`

The Amazon Resource Name (ARN) of the S3 bucket that contains your data.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws(-cn|-us-gov|-eusc|-iso(-[b-f])?)?:s3:::[a-z0-9][a-z0-9.-]{1,61}[a-z0-9]$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BucketOwnerAccountId`

The account ID for the owner of the S3 bucket.

_Required_: No

_Type_: String

_Pattern_: `^[0-9]{12}$`

_Minimum_: `12`

_Maximum_: `12`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InclusionPrefixes`

A list of S3 prefixes to include certain files or content. This field is an array with
a maximum of one item, which can contain a string that has a maximum length of 300
characters. For more information, see [Organizing objects using\
prefixes](../../../s3/latest/userguide/using-prefixes.md).

_Required_: No

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `300 | 1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PatternObjectFilterConfiguration

S3Location

All content copied from https://docs.aws.amazon.com/.
