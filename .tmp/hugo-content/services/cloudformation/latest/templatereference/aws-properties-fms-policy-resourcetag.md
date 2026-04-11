This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FMS::Policy ResourceTag

The resource tags that AWS Firewall Manager uses to determine if a particular resource
should be included or excluded from the AWS Firewall Manager policy. Tags enable you to
categorize your AWS resources in different ways, for example, by purpose, owner, or
environment. Each tag consists of a key and an optional value. Firewall Manager combines the
tags with "AND" so that, if you add more than one tag to a policy scope, a resource must have
all the specified tags to be included or excluded. For more information, see
[Working with Tag Editor](../../../awsconsolehelpdocs/latest/gsg/tag-editor.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Key" : String,
  "Value" : String
}

```

### YAML

```yaml

  Key: String
  Value: String

```

## Properties

`Key`

The resource tag key.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The resource tag value.

_Required_: No

_Type_: String

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PortRange

SecurityServicePolicyData

All content copied from https://docs.aws.amazon.com/.
