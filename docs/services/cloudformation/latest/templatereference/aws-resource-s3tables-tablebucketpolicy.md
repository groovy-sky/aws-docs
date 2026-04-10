This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3Tables::TableBucketPolicy

Creates a new maintenance configuration or replaces an existing table bucket policy for a
table bucket. For more information, see [Adding a table bucket policy](../../../s3/latest/userguide/s3-tables-bucket-policy.md#table-bucket-policy-add) in the _Amazon Simple Storage Service User Guide_.

Permissions

You must have the `s3tables:PutTableBucketPolicy` permission to use this operation.

AWS Cloud Development Kit (AWS CDK)

To use S3 Tables AWS CDK constructs, add the `@aws-cdk/aws-s3tables-alpha` dependency with one of the following options:

- NPM: `npm i @aws-cdk/aws-s3tables-alpha`

- Yarn: `yarn add @aws-cdk/aws-s3tables-alpha`

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::S3Tables::TableBucketPolicy",
  "Properties" : {
      "ResourcePolicy" : Json,
      "TableBucketARN" : String
    }
}

```

### YAML

```yaml

Type: AWS::S3Tables::TableBucketPolicy
Properties:
  ResourcePolicy: Json
  TableBucketARN: String

```

## Properties

`ResourcePolicy`

The bucket policy JSON for the table bucket.

_Required_: Yes

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TableBucketARN`

The Amazon Resource Name (ARN) of the table bucket.

_Required_: Yes

_Type_: String

_Pattern_: `(arn:aws[-a-z0-9]*:[a-z0-9]+:[-a-z0-9]*:[0-9]{12}:bucket/[a-z0-9_-]{3,63})`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns some information about your table bucket policy.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UnreferencedFileRemoval

AWS::S3Tables::TablePolicy

All content copied from https://docs.aws.amazon.com/.
