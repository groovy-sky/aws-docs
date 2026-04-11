# Processor

This structure contains the information about one processor in a log transformer.

## Contents

**addKeys**

Use this parameter to include the [addKeys](../../../../services/amazoncloudwatch/latest/logs/cloudwatch-logs-transformation.md#CloudWatch-Logs-Transformation-addKeys) processor in your transformer.

Type: [AddKeys](api-addkeys.md) object

Required: No

**copyValue**

Use this parameter to include the [copyValue](../../../../services/amazoncloudwatch/latest/logs/cloudwatch-logs-transformation.md#CloudWatch-Logs-Transformation-copyValue) processor in your transformer.

Type: [CopyValue](api-copyvalue.md) object

Required: No

**csv**

Use this parameter to include the [CSV](../../../../services/amazoncloudwatch/latest/logs/cloudwatch-logs-transformation.md#CloudWatch-Logs-Transformation-CSV) processor in your transformer.

Type: [CSV](api-csv.md) object

Required: No

**dateTimeConverter**

Use this parameter to include the [datetimeConverter](../../../../services/amazoncloudwatch/latest/logs/cloudwatch-logs-transformation.md#CloudWatch-Logs-Transformation-datetimeConverter) processor in your transformer.

Type: [DateTimeConverter](api-datetimeconverter.md) object

Required: No

**deleteKeys**

Use this parameter to include the [deleteKeys](../../../../services/amazoncloudwatch/latest/logs/cloudwatch-logs-transformation.md#CloudWatch-Logs-Transformation-deleteKeys) processor in your transformer.

Type: [DeleteKeys](api-deletekeys.md) object

Required: No

**grok**

Use this parameter to include the [grok](../../../../services/amazoncloudwatch/latest/logs/cloudwatch-logs-transformation.md#CloudWatch-Logs-Transformation-grok) processor in your transformer.

Type: [Grok](api-grok.md) object

Required: No

**listToMap**

Use this parameter to include the [listToMap](../../../../services/amazoncloudwatch/latest/logs/cloudwatch-logs-transformation.md#CloudWatch-Logs-Transformation-listToMap) processor in your transformer.

Type: [ListToMap](api-listtomap.md) object

Required: No

**lowerCaseString**

Use this parameter to include the [lowerCaseString](../../../../services/amazoncloudwatch/latest/logs/cloudwatch-logs-transformation.md#CloudWatch-Logs-Transformation-lowerCaseString) processor in your transformer.

Type: [LowerCaseString](api-lowercasestring.md) object

Required: No

**moveKeys**

Use this parameter to include the [moveKeys](../../../../services/amazoncloudwatch/latest/logs/cloudwatch-logs-transformation.md#CloudWatch-Logs-Transformation-moveKeys) processor in your transformer.

Type: [MoveKeys](api-movekeys.md) object

Required: No

**parseCloudfront**

Use this parameter to include the [parseCloudfront](../../../../services/amazoncloudwatch/latest/logs/cloudwatch-logs-transformation.md#CloudWatch-Logs-Transformation-parseCloudfront) processor in your transformer.

If you use this processor, it must be the first processor in your transformer.

Type: [ParseCloudfront](api-parsecloudfront.md) object

Required: No

**parseJSON**

Use this parameter to include the [parseJSON](../../../../services/amazoncloudwatch/latest/logs/cloudwatch-logs-transformation.md#CloudWatch-Logs-Transformation-parseJSON) processor in your transformer.

Type: [ParseJSON](api-parsejson.md) object

Required: No

**parseKeyValue**

Use this parameter to include the [parseKeyValue](../../../../services/amazoncloudwatch/latest/logs/cloudwatch-logs-transformation.md#CloudWatch-Logs-Transformation-parseKeyValue) processor in your transformer.

Type: [ParseKeyValue](api-parsekeyvalue.md) object

Required: No

**parsePostgres**

Use this parameter to include the [parsePostGres](../../../../services/amazoncloudwatch/latest/logs/cloudwatch-logs-transformation.md#CloudWatch-Logs-Transformation-parsePostGres) processor in your transformer.

If you use this processor, it must be the first processor in your transformer.

Type: [ParsePostgres](api-parsepostgres.md) object

Required: No

**parseRoute53**

Use this parameter to include the [parseRoute53](../../../../services/amazoncloudwatch/latest/logs/cloudwatch-logs-transformation.md#CloudWatch-Logs-Transformation-parseRoute53) processor in your transformer.

If you use this processor, it must be the first processor in your transformer.

Type: [ParseRoute53](api-parseroute53.md) object

Required: No

**parseToOCSF**

Use this parameter to convert logs into Open Cybersecurity Schema (OCSF) format.

Type: [ParseToOCSF](api-parsetoocsf.md) object

Required: No

**parseVPC**

Use this parameter to include the [parseVPC](../../../../services/amazoncloudwatch/latest/logs/cloudwatch-logs-transformation.md#CloudWatch-Logs-Transformation-parseVPC) processor in your transformer.

If you use this processor, it must be the first processor in your transformer.

Type: [ParseVPC](api-parsevpc.md) object

Required: No

**parseWAF**

Use this parameter to include the [parseWAF](../../../../services/amazoncloudwatch/latest/logs/cloudwatch-logs-transformation.md#CloudWatch-Logs-Transformation-parseWAF) processor in your transformer.

If you use this processor, it must be the first processor in your transformer.

Type: [ParseWAF](api-parsewaf.md) object

Required: No

**renameKeys**

Use this parameter to include the [renameKeys](../../../../services/amazoncloudwatch/latest/logs/cloudwatch-logs-transformation.md#CloudWatch-Logs-Transformation-renameKeys) processor in your transformer.

Type: [RenameKeys](api-renamekeys.md) object

Required: No

**splitString**

Use this parameter to include the [splitString](../../../../services/amazoncloudwatch/latest/logs/cloudwatch-logs-transformation.md#CloudWatch-Logs-Transformation-splitString) processor in your transformer.

Type: [SplitString](api-splitstring.md) object

Required: No

**substituteString**

Use this parameter to include the [substituteString](../../../../services/amazoncloudwatch/latest/logs/cloudwatch-logs-transformation.md#CloudWatch-Logs-Transformation-substituteString) processor in your transformer.

Type: [SubstituteString](api-substitutestring.md) object

Required: No

**trimString**

Use this parameter to include the [trimString](../../../../services/amazoncloudwatch/latest/logs/cloudwatch-logs-transformation.md#CloudWatch-Logs-Transformation-trimString) processor in your transformer.

Type: [TrimString](api-trimstring.md) object

Required: No

**typeConverter**

Use this parameter to include the [typeConverter](../../../../services/amazoncloudwatch/latest/logs/cloudwatch-logs-transformation.md#CloudWatch-Logs-Transformation-typeConverter) processor in your transformer.

Type: [TypeConverter](api-typeconverter.md) object

Required: No

**upperCaseString**

Use this parameter to include the [upperCaseString](../../../../services/amazoncloudwatch/latest/logs/cloudwatch-logs-transformation.md#CloudWatch-Logs-Transformation-upperCaseString) processor in your transformer.

Type: [UpperCaseString](api-uppercasestring.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/processor.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/processor.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/processor.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Policy

QueryCompileError

All content copied from https://docs.aws.amazon.com/.
