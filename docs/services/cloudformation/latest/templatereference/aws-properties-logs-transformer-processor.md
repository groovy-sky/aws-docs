This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Logs::Transformer Processor

This structure contains the information about one processor in a log
transformer.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AddKeys" : AddKeys,
  "CopyValue" : CopyValue,
  "Csv" : Csv,
  "DateTimeConverter" : DateTimeConverter,
  "DeleteKeys" : DeleteKeys,
  "Grok" : Grok,
  "ListToMap" : ListToMap,
  "LowerCaseString" : LowerCaseString,
  "MoveKeys" : MoveKeys,
  "ParseCloudfront" : ParseCloudfront,
  "ParseJSON" : ParseJSON,
  "ParseKeyValue" : ParseKeyValue,
  "ParsePostgres" : ParsePostgres,
  "ParseRoute53" : ParseRoute53,
  "ParseToOCSF" : ParseToOCSF,
  "ParseVPC" : ParseVPC,
  "ParseWAF" : ParseWAF,
  "RenameKeys" : RenameKeys,
  "SplitString" : SplitString,
  "SubstituteString" : SubstituteString,
  "TrimString" : TrimString,
  "TypeConverter" : TypeConverter,
  "UpperCaseString" : UpperCaseString
}

```

### YAML

```yaml

  AddKeys:
    AddKeys
  CopyValue:
    CopyValue
  Csv:
    Csv
  DateTimeConverter:
    DateTimeConverter
  DeleteKeys:
    DeleteKeys
  Grok:
    Grok
  ListToMap:
    ListToMap
  LowerCaseString:
    LowerCaseString
  MoveKeys:
    MoveKeys
  ParseCloudfront:
    ParseCloudfront
  ParseJSON:
    ParseJSON
  ParseKeyValue:
    ParseKeyValue
  ParsePostgres:
    ParsePostgres
  ParseRoute53:
    ParseRoute53
  ParseToOCSF:
    ParseToOCSF
  ParseVPC:
    ParseVPC
  ParseWAF:
    ParseWAF
  RenameKeys:
    RenameKeys
  SplitString:
    SplitString
  SubstituteString:
    SubstituteString
  TrimString:
    TrimString
  TypeConverter:
    TypeConverter
  UpperCaseString:
    UpperCaseString

```

## Properties

`AddKeys`

Use this parameter to include the [addKeys](../../../amazoncloudwatch/latest/logs/cloudwatch-logs-transformation.md#CloudWatch-Logs-Transformation-addKeys) processor in your transformer.

_Required_: No

_Type_: [AddKeys](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-logs-transformer-addkeys.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CopyValue`

Use this parameter to include the [copyValue](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Processors.html#CloudWatch-Logs-Transformation-copyValue) processor in your transformer.

_Required_: No

_Type_: [CopyValue](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-logs-transformer-copyvalue.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Csv`

Use this parameter to include the [CSV](../../../amazoncloudwatch/latest/logs/cloudwatch-logs-transformation.md#CloudWatch-Logs-Transformation-CSV) processor in your transformer.

_Required_: No

_Type_: [Csv](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-logs-transformer-csv.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DateTimeConverter`

Use this parameter to include the [datetimeConverter](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Processors.html#CloudWatch-Logs-Transformation-datetimeConverter) processor in your transformer.

_Required_: No

_Type_: [DateTimeConverter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-logs-transformer-datetimeconverter.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeleteKeys`

Use this parameter to include the [deleteKeys](../../../amazoncloudwatch/latest/logs/cloudwatch-logs-transformation.md#CloudWatch-Logs-Transformation-deleteKeys) processor in your transformer.

_Required_: No

_Type_: [DeleteKeys](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-logs-transformer-deletekeys.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Grok`

Use this parameter to include the [grok](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Processors.html#CloudWatch-Logs-Transformation-grok) processor in your transformer.

_Required_: No

_Type_: [Grok](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-logs-transformer-grok.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ListToMap`

Use this parameter to include the [listToMap](../../../amazoncloudwatch/latest/logs/cloudwatch-logs-transformation.md#CloudWatch-Logs-Transformation-listToMap) processor in your transformer.

_Required_: No

_Type_: [ListToMap](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-logs-transformer-listtomap.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LowerCaseString`

Use this parameter to include the [lowerCaseString](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Processors.html#CloudWatch-Logs-Transformation-lowerCaseString) processor in your transformer.

_Required_: No

_Type_: [LowerCaseString](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-logs-transformer-lowercasestring.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MoveKeys`

Use this parameter to include the [moveKeys](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Processors.html#CloudWatch-Logs-Transformation-moveKeys) processor in your transformer.

_Required_: No

_Type_: [MoveKeys](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-logs-transformer-movekeys.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ParseCloudfront`

Use this parameter to include the [parseCloudfront](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Processors.html#CloudWatch-Logs-Transformation-parseCloudfront) processor in your transformer.

If you use this processor, it must be the first processor in your transformer.

_Required_: No

_Type_: [ParseCloudfront](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-logs-transformer-parsecloudfront.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ParseJSON`

Use this parameter to include the [parseJSON](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Processors.html#CloudWatch-Logs-Transformation-parseJSON) processor in your transformer.

_Required_: No

_Type_: [ParseJSON](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-logs-transformer-parsejson.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ParseKeyValue`

Use this parameter to include the [parseKeyValue](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Processors.html#CloudWatch-Logs-Transformation-parseKeyValue) processor in your transformer.

_Required_: No

_Type_: [ParseKeyValue](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-logs-transformer-parsekeyvalue.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ParsePostgres`

Use this parameter to include the [parsePostGres](../../../amazoncloudwatch/latest/logs/cloudwatch-logs-transformation.md#CloudWatch-Logs-Transformation-parsePostGres) processor in your transformer.

If you use this processor, it must be the first processor in your transformer.

_Required_: No

_Type_: [ParsePostgres](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-logs-transformer-parsepostgres.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ParseRoute53`

Use this parameter to include the [parseRoute53](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Processors.html#CloudWatch-Logs-Transformation-parseRoute53) processor in your transformer.

If you use this processor, it must be the first processor in your transformer.

_Required_: No

_Type_: [ParseRoute53](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-logs-transformer-parseroute53.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ParseToOCSF`

Use this parameter to convert logs into Open Cybersecurity Schema (OCSF) format.

_Required_: No

_Type_: [ParseToOCSF](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-logs-transformer-parsetoocsf.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ParseVPC`

Use this parameter to include the [parseVPC](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Processors.html#CloudWatch-Logs-Transformation-parseVPC) processor in your transformer.

If you use this processor, it must be the first processor in your transformer.

_Required_: No

_Type_: [ParseVPC](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-logs-transformer-parsevpc.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ParseWAF`

Use this parameter to include the [parseWAF](../../../amazoncloudwatch/latest/logs/cloudwatch-logs-transformation.md#CloudWatch-Logs-Transformation-parseWAF) processor in your transformer.

If you use this processor, it must be the first processor in your transformer.

_Required_: No

_Type_: [ParseWAF](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-logs-transformer-parsewaf.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RenameKeys`

Use this parameter to include the [renameKeys](../../../amazoncloudwatch/latest/logs/cloudwatch-logs-transformation.md#CloudWatch-Logs-Transformation-renameKeys) processor in your transformer.

_Required_: No

_Type_: [RenameKeys](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-logs-transformer-renamekeys.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SplitString`

Use this parameter to include the [splitString](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Processors.html#CloudWatch-Logs-Transformation-splitString) processor in your transformer.

_Required_: No

_Type_: [SplitString](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-logs-transformer-splitstring.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubstituteString`

Use this parameter to include the [substituteString](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Processors.html#CloudWatch-Logs-Transformation-substituteString) processor in your transformer.

_Required_: No

_Type_: [SubstituteString](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-logs-transformer-substitutestring.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TrimString`

Use this parameter to include the [trimString](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Processors.html#CloudWatch-Logs-Transformation-trimString) processor in your transformer.

_Required_: No

_Type_: [TrimString](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-logs-transformer-trimstring.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TypeConverter`

Use this parameter to include the [typeConverter](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Processors.html#CloudWatch-Logs-Transformation-typeConverter) processor in your transformer.

_Required_: No

_Type_: [TypeConverter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-logs-transformer-typeconverter.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UpperCaseString`

Use this parameter to include the [upperCaseString](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Processors.html#CloudWatch-Logs-Transformation-upperCaseString) processor in your transformer.

_Required_: No

_Type_: [UpperCaseString](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-logs-transformer-uppercasestring.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ParseWAF

RenameKeyEntry
