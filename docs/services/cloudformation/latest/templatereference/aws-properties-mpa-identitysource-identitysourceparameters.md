This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MPA::IdentitySource IdentitySourceParameters

Contains details for the resource that provides identities to the identity source. For example, an IAM Identity Center instance.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IamIdentityCenter" : IamIdentityCenter
}

```

### YAML

```yaml

  IamIdentityCenter:
    IamIdentityCenter

```

## Properties

`IamIdentityCenter`

AWS IAM Identity Center credentials.

_Required_: Yes

_Type_: [IamIdentityCenter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-mpa-identitysource-iamidentitycenter.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

IamIdentityCenter

Tag
