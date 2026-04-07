This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KMS::Alias

The `AWS::KMS::Alias` resource specifies a display name for a [KMS key](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#kms_keys).
You can use an alias to identify a KMS key in the AWS KMS console, in the [DescribeKey](https://docs.aws.amazon.com/kms/latest/APIReference/API_DescribeKey.html)
operation, and in [cryptographic\
operations](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#cryptographic-operations), such as [Decrypt](https://docs.aws.amazon.com/kms/latest/APIReference/API_Decrypt.html) and [GenerateDataKey](https://docs.aws.amazon.com/kms/latest/APIReference/API_GenerateDataKey.html).

###### Note

Adding, deleting, or updating an alias can allow or deny permission to the KMS key. For
details, see [ABAC for AWS KMS](https://docs.aws.amazon.com/kms/latest/developerguide/abac.html) in the _AWS Key Management Service Developer_
_Guide_.

Using an alias to refer to a KMS key can help you simplify key management. For example, an
alias in your code can be associated with different KMS keys in different AWS Regions. For more information, see [Using aliases](https://docs.aws.amazon.com/kms/latest/developerguide/kms-alias.html) in the
_AWS Key Management Service Developer Guide_.

When specifying an alias, observe the following rules.

- Each alias is associated with one KMS key, but multiple aliases can be associated with the
same KMS key.

- The alias and its associated KMS key must be in the same AWS account and
Region.

- The alias name must be unique in the AWS account and Region. However,
you can create aliases with the same name in different AWS Regions. For
example, you can have an `alias/projectKey` in multiple Regions, each of which
is associated with a KMS key in its Region.

- Each alias name must begin with `alias/` followed by a name, such as
`alias/exampleKey`. The alias name can contain only alphanumeric characters,
forward slashes (/), underscores (\_), and dashes (-). Alias names cannot begin with
`alias/aws/`. That alias name prefix is reserved for [AWS managed keys](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-managed-cmk).

**Regions**

AWS KMS CloudFormation resources are available in all AWS Regions in which AWS KMS and
CloudFormation are supported.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::KMS::Alias",
  "Properties" : {
      "AliasName" : String,
      "TargetKeyId" : String
    }
}

```

### YAML

```yaml

Type: AWS::KMS::Alias
Properties:
  AliasName: String
  TargetKeyId: String

```

## Properties

`AliasName`

Specifies the alias name. This value must begin with `alias/` followed by a
name, such as `alias/ExampleAlias`.

###### Note

If you change the value of the `AliasName` property, the existing alias is
deleted and a new alias is created for the specified KMS key. This change can disrupt
applications that use the alias. It can also allow or deny access to a KMS key affected by
attribute-based access control (ABAC).

The alias must be string of 1-256 characters. It can contain only alphanumeric characters,
forward slashes (/), underscores (\_), and dashes (-). The alias name cannot begin with
`alias/aws/`. The `alias/aws/` prefix is reserved for [AWS managed keys](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-managed-cmk).

_Required_: Yes

_Type_: String

_Pattern_: `^(alias/)[a-zA-Z0-9:/_-]+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TargetKeyId`

Associates the alias with the specified [customer managed key](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#customer-cmk). The
KMS key must be in the same AWS account and Region.

A valid key ID is required. If you supply a null or empty string value, this operation
returns an error.

For help finding the key ID and ARN, see [Finding the key ID and\
ARN](https://docs.aws.amazon.com/kms/latest/developerguide/viewing-keys.html#find-cmk-id-arn) in the _AWS Key Management Service Developer Guide_.

Specify the key ID or the key ARN of the KMS key.

For example:

- Key ID: `1234abcd-12ab-34cd-56ef-1234567890ab`

- Key ARN:
`arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab`

To get the key ID and key ARN for a KMS key, use [ListKeys](https://docs.aws.amazon.com/kms/latest/APIReference/API_ListKeys.html) or [DescribeKey](https://docs.aws.amazon.com/kms/latest/APIReference/API_DescribeKey.html).

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the alias name, such as `alias/exampleAlias`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## Examples

### Create an alias

The following examples create the `alias/exampleAlias` alias for a KMS key. The
KMS key is identified by a reference to its CloudFormation resource name. Before using these
examples, replace the example target key ID and example alias with valid values.

#### JSON

```json

{
    "myAlias": {
        "Type": "AWS::KMS::Alias",
        "Properties": {
            "AliasName": "alias/exampleAlias",
            "TargetKeyId": {
                "Ref": "myKey"
            }
        }
    }
}
```

#### YAML

```yaml

myAlias:
  Type: 'AWS::KMS::Alias'
  Properties:
    AliasName: alias/exampleAlias
    TargetKeyId: !Ref myKey
```

## See also

- [CreateAlias](https://docs.aws.amazon.com/kms/latest/APIReference/API_CreateAlias.html) in the _AWS Key Management Service API_
_Reference_.

- [Using\
aliases](https://docs.aws.amazon.com/kms/latest/developerguide/kms-alias.html) in the _AWS Key Management Service Developer_
_Guide_.

- [ABAC for\
AWS KMS](https://docs.aws.amazon.com/kms/latest/developerguide/abac.html) in the _AWS Key Management Service Developer_
_Guide_.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS Key Management Service

AWS::KMS::Key
