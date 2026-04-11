This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::VerifiedPermissions::Policy StaticPolicyDefinition

A structure that defines a static policy.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Description" : String,
  "Statement" : String
}

```

### YAML

```yaml

  Description: String
  Statement: String

```

## Properties

`Description`

The description of the static policy.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `150`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Statement`

The policy content of the static policy, written in the Cedar policy language.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `10000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PolicyDefinition

TemplateLinkedPolicyDefinition

All content copied from https://docs.aws.amazon.com/.
