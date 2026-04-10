This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MPA::IdentitySource IamIdentityCenter

AWS IAM Identity Center credentials. For more information see, [AWS IAM Identity Center](https://aws.amazon.com/identity-center) .

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ApprovalPortalUrl" : String,
  "InstanceArn" : String,
  "Region" : String
}

```

### YAML

```yaml

  ApprovalPortalUrl: String
  InstanceArn: String
  Region: String

```

## Properties

`ApprovalPortalUrl`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`InstanceArn`

Amazon Resource Name (ARN) for the IAM Identity Center instance.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:.+:sso:::instance/(?:sso)?ins-[a-zA-Z0-9-.]{16}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Region`

AWS Region where the IAM Identity Center instance is located.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::MPA::IdentitySource

IdentitySourceParameters

All content copied from https://docs.aws.amazon.com/.
