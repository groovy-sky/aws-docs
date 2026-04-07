This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3Tables::Namespace

Creates a namespace. A namespace is a logical grouping of tables within your table bucket, which you can use to organize tables. For more information, see [Create a namespace](../../../s3/latest/userguide/s3-tables-namespace-create.md) in the _Amazon Simple Storage Service User Guide_.

Permissions

You must have the `s3tables:CreateNamespace` permission to use this operation.

AWS Cloud Development Kit (AWS CDK)

To use S3 Tables AWS CDK constructs, add the `@aws-cdk/aws-s3tables-alpha` dependency with one of the following options:

- NPM: `npm i @aws-cdk/aws-s3tables-alpha`

- Yarn: `yarn add @aws-cdk/aws-s3tables-alpha`

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::S3Tables::Namespace",
  "Properties" : {
      "Namespace" : String,
      "TableBucketARN" : String
    }
}

```

### YAML

```yaml

Type: AWS::S3Tables::Namespace
Properties:
  Namespace: String
  TableBucketARN: String

```

## Properties

`Namespace`

The name of the namespace.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TableBucketARN`

The Amazon Resource Name (ARN) of the table bucket to create the namespace in.

_Required_: Yes

_Type_: String

_Pattern_: `(arn:aws[-a-z0-9]*:[a-z0-9]+:[-a-z0-9]*:[0-9]{12}:bucket/[a-z0-9_-]{3,63})`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the namespace name.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Amazon S3 Tables

AWS::S3Tables::Table
