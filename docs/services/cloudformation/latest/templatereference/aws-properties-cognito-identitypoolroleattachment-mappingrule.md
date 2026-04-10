This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cognito::IdentityPoolRoleAttachment MappingRule

Defines how to map a claim to a role ARN.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Claim" : String,
  "MatchType" : String,
  "RoleARN" : String,
  "Value" : String
}

```

### YAML

```yaml

  Claim: String
  MatchType: String
  RoleARN: String
  Value: String

```

## Properties

`Claim`

The claim name that must be present in the token. For example: "isAdmin" or
"paid".

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MatchType`

The match condition that specifies how closely the claim value in the IdP token must
match `Value`.

Valid values are: `Equals`, `Contains`, `StartsWith`,
and `NotEqual`.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleARN`

The Amazon Resource Name (ARN) of the role.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

A brief string that the claim must match. For example, "paid" or "yes".

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Cognito::IdentityPoolRoleAttachment

RoleMapping

All content copied from https://docs.aws.amazon.com/.
