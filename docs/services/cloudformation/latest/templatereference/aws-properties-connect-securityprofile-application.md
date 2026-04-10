This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::SecurityProfile Application

This API is in preview release for Amazon Connect and is subject to change.

A third-party application's metadata.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ApplicationPermissions" : [ String, ... ],
  "Namespace" : String,
  "Type" : String
}

```

### YAML

```yaml

  ApplicationPermissions:
    - String
  Namespace: String
  Type: String

```

## Properties

`ApplicationPermissions`

The permissions that the agent is granted on the application. For third-party applications, only the
`ACCESS` permission is supported. For MCP Servers, the permissions are tool Identifiers accepted by MCP
Server.

_Required_: Yes

_Type_: Array of String

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Namespace`

Namespace of the application that you want to give access to.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

Type of Application.

_Required_: No

_Type_: String

_Allowed values_: `MCP | THIRD_PARTY_APPLICATION`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Connect::SecurityProfile

DataTableAccessControlConfiguration

All content copied from https://docs.aws.amazon.com/.
