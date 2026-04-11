This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Athena::PreparedStatement

Specifies a prepared statement for use with SQL queries in Athena.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Athena::PreparedStatement",
  "Properties" : {
      "Description" : String,
      "QueryStatement" : String,
      "StatementName" : String,
      "WorkGroup" : String
    }
}

```

### YAML

```yaml

Type: AWS::Athena::PreparedStatement
Properties:
  Description: String
  QueryStatement: String
  StatementName: String
  WorkGroup: String

```

## Properties

`Description`

The description of the prepared statement.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`QueryStatement`

The query string for the prepared statement.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `262144`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StatementName`

The name of the prepared statement.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`WorkGroup`

The workgroup to which the prepared statement belongs.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the name of the prepared statement.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Athena::NamedQuery

AWS::Athena::WorkGroup

All content copied from https://docs.aws.amazon.com/.
