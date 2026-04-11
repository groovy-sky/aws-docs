This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Grafana::Workspace AssertionAttributes

A structure that defines which attributes in the IdP assertion are to be used to
define information about the users authenticated by the IdP to use the workspace.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Email" : String,
  "Groups" : String,
  "Login" : String,
  "Name" : String,
  "Org" : String,
  "Role" : String
}

```

### YAML

```yaml

  Email: String
  Groups: String
  Login: String
  Name: String
  Org: String
  Role: String

```

## Properties

`Email`

The name of the attribute within the SAML assertion to use as the email names for SAML
users.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Groups`

The name of the attribute within the SAML assertion to use as the user full "friendly"
names for user groups.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Login`

The name of the attribute within the SAML assertion to use as the login names for SAML
users.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the attribute within the SAML assertion to use as the user full "friendly"
names for SAML users.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Org`

The name of the attribute within the SAML assertion to use as the user full "friendly"
names for the users' organizations.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Role`

The name of the attribute within the SAML assertion to use as the user roles.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Grafana::Workspace

IdpMetadata

All content copied from https://docs.aws.amazon.com/.
