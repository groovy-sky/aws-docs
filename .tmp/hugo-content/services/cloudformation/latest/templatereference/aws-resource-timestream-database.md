This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Timestream::Database

Creates a new Timestream database. If the AWS KMS key is not
specified, the database will be encrypted with a Timestream managed AWS KMS key located in your account. Refer to [AWS managed\
AWS KMS keys](../../../kms/latest/developerguide/concepts.md#aws-managed-cmk) for more info. [Service quotas apply](../../../timestream/latest/developerguide/ts-limits.md). See
[code sample](../../../timestream/latest/developerguide/code-samples-create-db.md) for
details.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Timestream::Database",
  "Properties" : {
      "DatabaseName" : String,
      "KmsKeyId" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Timestream::Database
Properties:
  DatabaseName: String
  KmsKeyId: String
  Tags:
    - Tag

```

## Properties

`DatabaseName`

The name of the Timestream database.

_Length Constraints_: Minimum length of 3 bytes. Maximum length of 256
bytes.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9_.-]{3,256}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`KmsKeyId`

The identifier of the AWS KMS key used to encrypt the data stored in the
database.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags to add to the database.

_Required_: No

_Type_: Array of [Tag](aws-properties-timestream-database-tag.md)

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the database name `DATABASE_NAME`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` returns a value for the specified attribute of this type. The
following are the available attributes:

`Arn`

The `arn` of the database.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon Timestream

Tag

All content copied from https://docs.aws.amazon.com/.
