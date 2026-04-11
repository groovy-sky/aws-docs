This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Logs::QueryDefinition

Creates a query definition for CloudWatch Logs Insights. For more information,
see [Analyzing Log Data with CloudWatch Logs Insights](../../../amazoncloudwatch/latest/logs/analyzinglogdata.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Logs::QueryDefinition",
  "Properties" : {
      "LogGroupNames" : [ String, ... ],
      "Name" : String,
      "Parameters" : [ QueryParameter, ... ],
      "QueryLanguage" : String,
      "QueryString" : String
    }
}

```

### YAML

```yaml

Type: AWS::Logs::QueryDefinition
Properties:
  LogGroupNames:
    - String
  Name: String
  Parameters:
    - QueryParameter
  QueryLanguage: String
  QueryString:
    String

```

## Properties

`LogGroupNames`

Use this parameter if you want the query to query only certain log groups.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

A name for the query definition.

###### Note

You can use the name to create a folder structure for your queries. To create a
folder, use a forward slash (/) to prefix your desired query name with your desired
folder name. For example,
`folder-name/query-name`.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Parameters`

If this query definition contains a list of query parameters that define placeholder
variables for the query string, that list appears here.

_Required_: No

_Type_: Array of [QueryParameter](aws-properties-logs-querydefinition-queryparameter.md)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`QueryLanguage`

The query language used for this query. For more information about the query languages
that CloudWatch Logs supports, see [Supported query\
languages](../../../amazoncloudwatch/latest/logs/cwl-analyzelogdata-languages.md).

_Required_: No

_Type_: String

_Allowed values_: `CWLI | SQL | PPL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`QueryString`

The query string to use for this query definition. For more information, see [CloudWatch Logs Insights Query Syntax](../../../amazoncloudwatch/latest/logs/cwl-querysyntax.md).

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `10000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the query definition ID.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`QueryDefinitionId`

The ID of the query definition.

## Examples

### Query definition example

The following example creates a query definition.

#### JSON

```json

"myQueryDefinition": {
  "Type": "AWS::Logs::QueryDefinition",
  "Properties": {
    "Name": "myQueryName",
    "QueryString": "fields @timestamp, @message | sort @timestamp desc | limit 20"
    }
}
```

#### YAML

```yaml

myQueryDefinition:
  Type: AWS::Logs::QueryDefinition
  Properties:
    Name: "myQueryName"
    QueryString: “fields @timestamp, @message | sort @timestamp desc | limit 20"
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MetricTransformation

QueryParameter

All content copied from https://docs.aws.amazon.com/.
