This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::DistributionTenant WebAclCustomization

The AWS WAF web ACL customization specified for the distribution tenant.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Action" : String,
  "Arn" : String
}

```

### YAML

```yaml

  Action: String
  Arn: String

```

## Properties

`Action`

The action for the AWS WAF web ACL customization. You can specify `override` to specify a separate AWS WAF web ACL for the distribution tenant. If you specify `disable`, the distribution tenant won't have AWS WAF web ACL protections and won't inherit from the multi-tenant distribution.

_Required_: No

_Type_: String

_Allowed values_: `override | disable`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Arn`

The Amazon Resource Name (ARN) of the AWS WAF web ACL.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AWS::CloudFront::Function

All content copied from https://docs.aws.amazon.com/.
