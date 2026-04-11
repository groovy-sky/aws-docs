This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WorkSpaces::ConnectionAlias

The `AWS::WorkSpaces::ConnectionAlias` resource specifies a connection alias.
Connection aliases are used for cross-Region redirection. For more information, see [Cross-Region Redirection for Amazon WorkSpaces](../../../workspaces/latest/adminguide/cross-region-redirection.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::WorkSpaces::ConnectionAlias",
  "Properties" : {
      "ConnectionString" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::WorkSpaces::ConnectionAlias
Properties:
  ConnectionString:
    String
  Tags:
    - Tag

```

## Properties

`ConnectionString`

The connection string specified for the connection alias. The connection string must be
in the form of a fully qualified domain name (FQDN), such as
`www.example.com`.

_Required_: Yes

_Type_: String

_Pattern_: `^[.0-9a-zA-Z\-]{1,255}$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags to associate with the connection alias.

_Required_: No

_Type_: Array of [Tag](aws-properties-workspaces-connectionalias-tag.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource name.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AliasId`

The identifier of the connection alias, returned as a string.

`Associations`

The association status of the connection alias.

`ConnectionAliasState`

The current state of the connection alias, returned as a string.

## See also

- [ConnectionAlias](../../../workspaces/latest/api/api-connectionalias.md) in the _Amazon WorkSpaces API_
_Reference_

- [Cross-Region\
Redirection for Amazon WorkSpaces](../../../workspaces/latest/adminguide/cross-region-redirection.md) in the _Amazon WorkSpaces_
_Administration Guide_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon WorkSpaces

ConnectionAliasAssociation

All content copied from https://docs.aws.amazon.com/.
