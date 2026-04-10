This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataZone::PolicyGrant UserPolicyGrantPrincipal

The user policy grant principal.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AllUsersGrantFilter" : Json,
  "UserIdentifier" : String
}

```

### YAML

```yaml

  AllUsersGrantFilter: Json
  UserIdentifier: String

```

## Properties

`AllUsersGrantFilter`

The all users grant filter of the user policy grant principal.

_Required_: No

_Type_: Json

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`UserIdentifier`

The user ID of the user policy grant principal.

_Required_: No

_Type_: String

_Pattern_: `(^([0-9a-f]{10}-|)[A-Fa-f0-9]{8}-[A-Fa-f0-9]{4}-[A-Fa-f0-9]{4}-[A-Fa-f0-9]{4}-[A-Fa-f0-9]{12}$|^[a-zA-Z_0-9+=,.@-]+$|^arn:aws[^:]*:iam::\d{12}:.+$)`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ProjectPolicyGrantPrincipal

AWS::DataZone::Project

All content copied from https://docs.aws.amazon.com/.
