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

_Type_: [AddKeys](aws-properties-logs-transformer-addkeys.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CopyValue`

Use this parameter to include the [copyValue](../../../amazoncloudwatch/latest/logs/cloudwatch-logs-transformation-processors.md#CloudWatch-Logs-Transformation-copyValue) processor in your transformer.

_Required_: No

_Type_: [CopyValue](aws-properties-logs-transformer-copyvalue.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Csv`

Use this parameter to include the [CSV](../../../amazoncloudwatch/latest/logs/cloudwatch-logs-transformation.md#CloudWatch-Logs-Transformation-CSV) processor in your transformer.

_Required_: No

_Type_: [Csv](aws-properties-logs-transformer-csv.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DateTimeConverter`

Use this parameter to include the [datetimeConverter](../../../amazoncloudwatch/latest/logs/cloudwatch-logs-transformation-processors.md#CloudWatch-Logs-Transformation-datetimeConverter) processor in your transformer.

_Required_: No

_Type_: [DateTimeConverter](aws-properties-logs-transformer-datetimeconverter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeleteKeys`

Use this parameter to include the [deleteKeys](../../../amazoncloudwatch/latest/logs/cloudwatch-logs-transformation.md#CloudWatch-Logs-Transformation-deleteKeys) processor in your transformer.

_Required_: No

_Type_: [DeleteKeys](aws-properties-logs-transformer-deletekeys.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Grok`

Use this parameter to include the [grok](../../../amazoncloudwatch/latest/logs/cloudwatch-logs-transformation-processors.md#CloudWatch-Logs-Transformation-grok) processor in your transformer.

_Required_: No

_Type_: [Grok](aws-properties-logs-transformer-grok.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ListToMap`

Use this parameter to include the [listToMap](../../../amazoncloudwatch/latest/logs/cloudwatch-logs-transformation.md#CloudWatch-Logs-Transformation-listToMap) processor in your transformer.

_Required_: No

_Type_: [ListToMap](aws-properties-logs-transformer-listtomap.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LowerCaseString`

Use this parameter to include the [lowerCaseString](../../../amazoncloudwatch/latest/logs/cloudwatch-logs-transformation-processors.md#CloudWatch-Logs-Transformation-lowerCaseString) processor in your transformer.

_Required_: No

_Type_: [LowerCaseString](aws-properties-logs-transformer-lowercasestring.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MoveKeys`

Use this parameter to include the [moveKeys](../../../amazoncloudwatch/latest/logs/cloudwatch-logs-transformation-processors.md#CloudWatch-Logs-Transformation-moveKeys) processor in your transformer.

_Required_: No

_Type_: [MoveKeys](aws-properties-logs-transformer-movekeys.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ParseCloudfront`

Use this parameter to include the [parseCloudfront](../../../amazoncloudwatch/latest/logs/cloudwatch-logs-transformation-processors.md#CloudWatch-Logs-Transformation-parseCloudfront) processor in your transformer.

If you use this processor, it must be the first processor in your transformer.

_Required_: No

_Type_: [ParseCloudfront](aws-properties-logs-transformer-parsecloudfront.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ParseJSON`

Use this parameter to include the [parseJSON](../../../amazoncloudwatch/latest/logs/cloudwatch-logs-transformation-processors.md#CloudWatch-Logs-Transformation-parseJSON) processor in your transformer.

_Required_: No

_Type_: [ParseJSON](aws-properties-logs-transformer-parsejson.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ParseKeyValue`

Use this parameter to include the [parseKeyValue](../../../amazoncloudwatch/latest/logs/cloudwatch-logs-transformation-processors.md#CloudWatch-Logs-Transformation-parseKeyValue) processor in your transformer.

_Required_: No

_Type_: [ParseKeyValue](aws-properties-logs-transformer-parsekeyvalue.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ParsePostgres`

Use this parameter to include the [parsePostGres](../../../amazoncloudwatch/latest/logs/cloudwatch-logs-transformation.md#CloudWatch-Logs-Transformation-parsePostGres) processor in your transformer.

If you use this processor, it must be the first processor in your transformer.

_Required_: No

_Type_: [ParsePostgres](aws-properties-logs-transformer-parsepostgres.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ParseRoute53`

Use this parameter to include the [parseRoute53](../../../amazoncloudwatch/latest/logs/cloudwatch-logs-transformation-processors.md#CloudWatch-Logs-Transformation-parseRoute53) processor in your transformer.

If you use this processor, it must be the first processor in your transformer.

_Required_: No

_Type_: [ParseRoute53](aws-properties-logs-transformer-parseroute53.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ParseToOCSF`

Use this parameter to convert logs into Open Cybersecurity Schema (OCSF) format.

_Required_: No

_Type_: [ParseToOCSF](aws-properties-logs-transformer-parsetoocsf.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ParseVPC`

Use this parameter to include the [parseVPC](../../../amazoncloudwatch/latest/logs/cloudwatch-logs-transformation-processors.md#CloudWatch-Logs-Transformation-parseVPC) processor in your transformer.

If you use this processor, it must be the first processor in your transformer.

_Required_: No

_Type_: [ParseVPC](aws-properties-logs-transformer-parsevpc.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ParseWAF`

Use this parameter to include the [parseWAF](../../../amazoncloudwatch/latest/logs/cloudwatch-logs-transformation.md#CloudWatch-Logs-Transformation-parseWAF) processor in your transformer.

If you use this processor, it must be the first processor in your transformer.

_Required_: No

_Type_: [ParseWAF](aws-properties-logs-transformer-parsewaf.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RenameKeys`

Use this parameter to include the [renameKeys](../../../amazoncloudwatch/latest/logs/cloudwatch-logs-transformation.md#CloudWatch-Logs-Transformation-renameKeys) processor in your transformer.

_Required_: No

_Type_: [RenameKeys](aws-properties-logs-transformer-renamekeys.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SplitString`

Use this parameter to include the [splitString](../../../amazoncloudwatch/latest/logs/cloudwatch-logs-transformation-processors.md#CloudWatch-Logs-Transformation-splitString) processor in your transformer.

_Required_: No

_Type_: [SplitString](aws-properties-logs-transformer-splitstring.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubstituteString`

Use this parameter to include the [substituteString](../../../amazoncloudwatch/latest/logs/cloudwatch-logs-transformation-processors.md#CloudWatch-Logs-Transformation-substituteString) processor in your transformer.

_Required_: No

_Type_: [SubstituteString](aws-properties-logs-transformer-substitutestring.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TrimString`

Use this parameter to include the [trimString](../../../amazoncloudwatch/latest/logs/cloudwatch-logs-transformation-processors.md#CloudWatch-Logs-Transformation-trimString) processor in your transformer.

_Required_: No

_Type_: [TrimString](aws-properties-logs-transformer-trimstring.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TypeConverter`

Use this parameter to include the [typeConverter](../../../amazoncloudwatch/latest/logs/cloudwatch-logs-transformation-processors.md#CloudWatch-Logs-Transformation-typeConverter) processor in your transformer.

_Required_: No

_Type_: [TypeConverter](aws-properties-logs-transformer-typeconverter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UpperCaseString`

Use this parameter to include the [upperCaseString](../../../amazoncloudwatch/latest/logs/cloudwatch-logs-transformation-processors.md#CloudWatch-Logs-Transformation-upperCaseString) processor in your transformer.

_Required_: No

_Type_: [UpperCaseString](aws-properties-logs-transformer-uppercasestring.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ParseWAF

RenameKeyEntry

All content copied from https://docs.aws.amazon.com/.
