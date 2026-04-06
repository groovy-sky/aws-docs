This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Athena::NamedQuery

The `AWS::Athena::NamedQuery` resource specifies an Amazon Athena saved query, where `QueryString` contains the SQL query statements that
make up the query.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Athena::NamedQuery",
  "Properties" : {
      "Database" : String,
      "Description" : String,
      "Name" : String,
      "QueryString" : String,
      "WorkGroup" : String
    }
}

```

### YAML

```yaml

Type: AWS::Athena::NamedQuery
Properties:
  Database: String
  Description: String
  Name: String
  QueryString:
    String
  WorkGroup: String

```

## Properties

`Database`

The database to which the query belongs.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

The query description.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The query name.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`QueryString`

The SQL statements that make up the query.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `262144`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`WorkGroup`

The name of the workgroup that contains the named query.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource name.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`NamedQueryId`

The unique ID of the query.

## Examples

### Create a named query

The following example creates a named query.

#### JSON

```json

{
    "Resources": {
        "AthenaNamedQuery": {
            "Type": "AWS::Athena::NamedQuery",
            "Properties": {
                "Database": "swfmetadata",
                "Description": "A query that selects all aggregated data",
                "Name": "MostExpensiveWorkflow",
                "QueryString": "SELECT workflowname, AVG(activitytaskstarted) AS AverageWorkflow FROM swfmetadata WHERE year='17' AND GROUP BY workflowname ORDER BY AverageWorkflow DESC LIMIT 10"
            }
        }
    }
}
```

#### YAML

```yaml

Resources:
  AthenaNamedQuery:
    Type: AWS::Athena::NamedQuery
    Properties:
      Database: "swfmetadata"
      Description: "A query that selects all aggregated data"
      Name: "MostExpensiveWorkflow"
      QueryString: >
                    SELECT workflowname, AVG(activitytaskstarted) AS AverageWorkflow
                    FROM swfmetadata
                    WHERE year='17' AND GROUP BY workflowname
                    ORDER BY AverageWorkflow DESC LIMIT 10
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AWS::Athena::PreparedStatement
