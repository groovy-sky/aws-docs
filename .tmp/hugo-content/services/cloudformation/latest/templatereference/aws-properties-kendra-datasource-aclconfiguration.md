This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Kendra::DataSource AclConfiguration

Provides information about the column that should be used for filtering the query
response by groups.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AllowedGroupsColumnName" : String
}

```

### YAML

```yaml

  AllowedGroupsColumnName: String

```

## Properties

`AllowedGroupsColumnName`

A list of groups, separated by semi-colons, that filters a query response based on
user context. The document is only returned to users that are in one of the groups
specified in the `UserContext` field of the [Query](../../../kendra/latest/dg/api-query.md) operation.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AccessControlListConfiguration

ColumnConfiguration

All content copied from https://docs.aws.amazon.com/.
