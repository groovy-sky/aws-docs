This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::DataQualityRuleset DataQualityTargetTable

An object representing an AWS Glue table.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DatabaseName" : String,
  "TableName" : String
}

```

### YAML

```yaml

  DatabaseName: String
  TableName: String

```

## Properties

`DatabaseName`

The name of the database where the AWS Glue table exists.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TableName`

The name of the AWS Glue table.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Glue::DataQualityRuleset

AWS::Glue::DevEndpoint
