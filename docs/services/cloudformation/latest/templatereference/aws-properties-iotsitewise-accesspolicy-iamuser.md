This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTSiteWise::AccessPolicy IamUser

Contains information about an AWS Identity and Access Management user.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "arn" : String
}

```

### YAML

```yaml

  arn: String

```

## Properties

`arn`

The ARN of the IAM user. For more information, see [IAM ARNs](../../../iam/latest/userguide/reference-identifiers.md) in the
_IAM User Guide_.

###### Note

If you delete the IAM user, access policies that contain this identity include an
empty `arn`. You can delete the access policy for the IAM user that no longer
exists.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IamRole

Portal

All content copied from https://docs.aws.amazon.com/.
