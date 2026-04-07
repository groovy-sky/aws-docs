This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3Tables::TablePolicy

Creates a new maintenance configuration or replaces an existing table policy for a table. For more information, see [Adding a table policy](../../../s3/latest/userguide/s3-tables-table-policy.md#table-policy-add) in the _Amazon Simple Storage Service User Guide_.

Permissions

You must have the `s3tables:PutTablePolicy` permission to use this operation.

AWS Cloud Development Kit (AWS CDK)

To use S3 Tables AWS CDK constructs, add the `@aws-cdk/aws-s3tables-alpha` dependency with one of the following options:

- NPM: `npm i @aws-cdk/aws-s3tables-alpha`

- Yarn: `yarn add @aws-cdk/aws-s3tables-alpha`

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::S3Tables::TablePolicy",
  "Properties" : {
      "ResourcePolicy" : Json,
      "TableARN" : String
    }
}

```

### YAML

```yaml

Type: AWS::S3Tables::TablePolicy
Properties:
  ResourcePolicy: Json
  TableARN: String

```

## Properties

`ResourcePolicy`

The `JSON` that defines the policy.

_Required_: Yes

_Type_: Json

_Minimum_: `1`

_Maximum_: `20480`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TableARN`

The Amazon Resource Name (ARN) of the table.

_Required_: Yes

_Type_: String

_Pattern_: `(arn:aws[-a-z0-9]*:[a-z0-9]+:[-a-z0-9]*:[0-9]{12}:bucket/[a-z0-9_-]{3,63}/table/[a-zA-Z0-9-_]{1,255})`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns some information about your table policy.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Namespace`

The namespace to associated with the table.

`TableBucketARN`

The Amazon Resource Name (ARN) of the table bucket that contains the table.

`TableName`

The name of the table.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::S3Tables::TableBucketPolicy

Next
