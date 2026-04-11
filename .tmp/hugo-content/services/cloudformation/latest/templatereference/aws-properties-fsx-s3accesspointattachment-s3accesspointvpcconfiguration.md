This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FSx::S3AccessPointAttachment S3AccessPointVpcConfiguration

If included, Amazon S3 restricts access to this access point to requests from the specified virtual private cloud (VPC).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "VpcId" : String
}

```

### YAML

```yaml

  VpcId: String

```

## Properties

`VpcId`

Specifies the virtual private cloud (VPC) for the S3 access point VPC configuration, if one exists.

_Required_: Yes

_Type_: String

_Pattern_: `^(vpc-[0-9a-f]{8,})$`

_Minimum_: `12`

_Maximum_: `21`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

S3AccessPointOpenZFSConfiguration

AWS::FSx::Snapshot

All content copied from https://docs.aws.amazon.com/.
