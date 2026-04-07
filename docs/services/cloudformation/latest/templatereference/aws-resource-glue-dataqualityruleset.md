This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::DataQualityRuleset

The `AWS::Glue::DataQualityRuleset` resource specifies a data quality ruleset with DQDL rules applied to a specified AWS Glue table. For more information, see AWS Glue Data Quality in the AWS Glue Developer Guide.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Glue::DataQualityRuleset",
  "Properties" : {
      "ClientToken" : String,
      "Description" : String,
      "Name" : String,
      "Ruleset" : String,
      "Tags" : [ Tag, ... ],
      "TargetTable" : DataQualityTargetTable
    }
}

```

### YAML

```yaml

Type: AWS::Glue::DataQualityRuleset
Properties:
  ClientToken: String
  Description: String
  Name: String
  Ruleset: String
  Tags:
    - Tag
  TargetTable:
    DataQualityTargetTable

```

## Properties

`ClientToken`

Used for idempotency and is recommended to be set to a random ID (such as a UUID) to avoid creating or starting multiple instances of the same resource.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A description of the data quality ruleset.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the data quality ruleset.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Ruleset`

A Data Quality Definition Language (DQDL) ruleset. For more information see the AWS Glue Developer Guide.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A list of tags applied to the data quality ruleset.

_Required_: No

_Type_: Array of [`Tag`](aws-properties-resource-tags.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetTable`

An object representing an AWS Glue table.

_Required_: No

_Type_: [DataQualityTargetTable](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-glue-dataqualityruleset-dataqualitytargettable.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

EncryptionAtRest

DataQualityTargetTable
