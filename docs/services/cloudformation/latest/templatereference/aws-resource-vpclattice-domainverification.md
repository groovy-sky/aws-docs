This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::VpcLattice::DomainVerification

A domain name verification is an entity that allows you to prove your ownership of a given domain. When you create a domain verification using CloudFormation, use a waiter to make sure the domain verification is complete before you create a service network resource association, a VPC endpoint, or a service network VPC association with this domain.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::VpcLattice::DomainVerification",
  "Properties" : {
      "DomainName" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::VpcLattice::DomainVerification
Properties:
  DomainName: String
  Tags:
    - Tag

```

## Properties

`DomainName`

The domain name being verified.

_Required_: Yes

_Type_: String

_Minimum_: `3`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags associated with the domain verification.

_Required_: No

_Type_: Array of [Tag](aws-properties-vpclattice-domainverification-tag.md)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`Arn`

The Amazon Resource Name (ARN) of the domain verification.

`Id`

The ID of the domain verification.

`Status`

The current status of the domain verification process.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::VpcLattice::AuthPolicy

Tag

All content copied from https://docs.aws.amazon.com/.
