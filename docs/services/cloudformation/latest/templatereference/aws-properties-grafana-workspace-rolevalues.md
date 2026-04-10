This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Grafana::Workspace RoleValues

This structure defines which groups defined in the SAML assertion attribute are to be
mapped to the Grafana `Admin` and `Editor` roles in the workspace.
SAML authenticated users not part of `Admin` or `Editor` role
groups have `Viewer` permission over the workspace.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Admin" : [ String, ... ],
  "Editor" : [ String, ... ]
}

```

### YAML

```yaml

  Admin:
    - String
  Editor:
    - String

```

## Properties

`Admin`

A list of groups from the SAML assertion attribute to grant the Grafana
`Admin` role to.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Editor`

A list of groups from the SAML assertion attribute to grant the Grafana
`Editor` role to.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NetworkAccessControl

SamlConfiguration

All content copied from https://docs.aws.amazon.com/.
