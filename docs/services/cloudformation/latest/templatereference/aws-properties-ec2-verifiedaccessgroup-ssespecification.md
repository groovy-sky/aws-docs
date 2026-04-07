This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::VerifiedAccessGroup SseSpecification

AWS Verified Access provides server side encryption by default to data at rest using AWS-owned KMS keys. You also have the option of using customer managed KMS keys, which can be specified using the options below.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CustomerManagedKeyEnabled" : Boolean,
  "KmsKeyArn" : String
}

```

### YAML

```yaml

  CustomerManagedKeyEnabled: Boolean
  KmsKeyArn: String

```

## Properties

`CustomerManagedKeyEnabled`

Enable or disable the use of customer managed KMS keys for server side encryption.

Valid values: `True` \| `False`

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsKeyArn`

The ARN of the KMS key.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::EC2::VerifiedAccessGroup

Tag
