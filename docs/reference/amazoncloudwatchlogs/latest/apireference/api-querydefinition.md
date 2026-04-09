# QueryDefinition

This structure contains details about a saved CloudWatch Logs Insights query
definition.

## Contents

**lastModified**

The date that the query definition was most recently modified.

Type: Long

Valid Range: Minimum value of 0.

Required: No

**logGroupNames**

If this query definition contains a list of log groups that it is limited to, that list
appears here.

Type: Array of strings

Length Constraints: Minimum length of 1. Maximum length of 512.

Pattern: `[\.\-_/#A-Za-z0-9]+`

Required: No

**name**

The name of the query definition.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Required: No

**parameters**

If this query definition contains a list of query parameters that define placeholder
variables for the query string, that list appears here.

Type: Array of [QueryParameter](api-queryparameter.md) objects

Array Members: Maximum number of 20 items.

Required: No

**queryDefinitionId**

The unique ID of the query definition.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Required: No

**queryLanguage**

The query language used for this query. For more information about the query languages
that CloudWatch Logs supports, see [Supported query\
languages](../../../../services/amazoncloudwatch/latest/logs/cwl-analyzelogdata-languages.md).

Type: String

Valid Values: `CWLI | SQL | PPL`

Required: No

**queryString**

The query string to use for this definition. For more information, see [CloudWatch Logs\
Insights Query Syntax](../../../../services/amazoncloudwatch/latest/logs/cwl-querysyntax.md).

Type: String

Length Constraints: Minimum length of 1. Maximum length of 10000.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/querydefinition.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/querydefinition.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/querydefinition.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

QueryCompileErrorLocation

QueryInfo

All content copied from https://docs.aws.amazon.com/.
