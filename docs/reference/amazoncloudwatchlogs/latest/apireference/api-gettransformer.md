# GetTransformer

Returns the information about the log transformer associated with this log group.

This operation returns data only for transformers created at the log group level. To get
information for an account-level transformer, use [DescribeAccountPolicies](api-describeaccountpolicies.md).

## Request Syntax

```nohighlight

{
   "logGroupIdentifier": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[logGroupIdentifier](#API_GetTransformer_RequestSyntax)**

Specify either the name or ARN of the log group to return transformer information for. If
the log group is in a source account and you are using a monitoring account, you must use the
log group ARN.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Pattern: `[\w#+=/:,.@-]*`

Required: Yes

## Response Syntax

```nohighlight

{
   "creationTime": number,
   "lastModifiedTime": number,
   "logGroupIdentifier": "string",
   "transformerConfig": [
      {
         "addKeys": {
            "entries": [
               {
                  "key": "string",
                  "overwriteIfExists": boolean,
                  "value": "string"
               }
            ]
         },
         "copyValue": {
            "entries": [
               {
                  "overwriteIfExists": boolean,
                  "source": "string",
                  "target": "string"
               }
            ]
         },
         "csv": {
            "columns": [ "string" ],
            "delimiter": "string",
            "destination": "string",
            "quoteCharacter": "string",
            "source": "string"
         },
         "dateTimeConverter": {
            "locale": "string",
            "matchPatterns": [ "string" ],
            "source": "string",
            "sourceTimezone": "string",
            "target": "string",
            "targetFormat": "string",
            "targetTimezone": "string"
         },
         "deleteKeys": {
            "withKeys": [ "string" ]
         },
         "grok": {
            "match": "string",
            "source": "string"
         },
         "listToMap": {
            "flatten": boolean,
            "flattenedElement": "string",
            "key": "string",
            "source": "string",
            "target": "string",
            "valueKey": "string"
         },
         "lowerCaseString": {
            "withKeys": [ "string" ]
         },
         "moveKeys": {
            "entries": [
               {
                  "overwriteIfExists": boolean,
                  "source": "string",
                  "target": "string"
               }
            ]
         },
         "parseCloudfront": {
            "source": "string"
         },
         "parseJSON": {
            "destination": "string",
            "source": "string"
         },
         "parseKeyValue": {
            "destination": "string",
            "fieldDelimiter": "string",
            "keyPrefix": "string",
            "keyValueDelimiter": "string",
            "nonMatchValue": "string",
            "overwriteIfExists": boolean,
            "source": "string"
         },
         "parsePostgres": {
            "source": "string"
         },
         "parseRoute53": {
            "source": "string"
         },
         "parseToOCSF": {
            "eventSource": "string",
            "mappingVersion": "string",
            "ocsfVersion": "string",
            "source": "string"
         },
         "parseVPC": {
            "source": "string"
         },
         "parseWAF": {
            "source": "string"
         },
         "renameKeys": {
            "entries": [
               {
                  "key": "string",
                  "overwriteIfExists": boolean,
                  "renameTo": "string"
               }
            ]
         },
         "splitString": {
            "entries": [
               {
                  "delimiter": "string",
                  "source": "string"
               }
            ]
         },
         "substituteString": {
            "entries": [
               {
                  "from": "string",
                  "source": "string",
                  "to": "string"
               }
            ]
         },
         "trimString": {
            "withKeys": [ "string" ]
         },
         "typeConverter": {
            "entries": [
               {
                  "key": "string",
                  "type": "string"
               }
            ]
         },
         "upperCaseString": {
            "withKeys": [ "string" ]
         }
      }
   ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[creationTime](#API_GetTransformer_ResponseSyntax)**

The creation time of the transformer, expressed as the number of milliseconds after Jan
1, 1970 00:00:00 UTC.

Type: Long

Valid Range: Minimum value of 0.

**[lastModifiedTime](#API_GetTransformer_ResponseSyntax)**

The date and time when this transformer was most recently modified, expressed as the
number of milliseconds after Jan 1, 1970 00:00:00 UTC.

Type: Long

Valid Range: Minimum value of 0.

**[logGroupIdentifier](#API_GetTransformer_ResponseSyntax)**

The ARN of the log group that you specified in your request.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Pattern: `[\w#+=/:,.@-]*`

**[transformerConfig](#API_GetTransformer_ResponseSyntax)**

This sructure contains the configuration of the requested transformer.

Type: Array of [Processor](api-processor.md) objects

Array Members: Minimum number of 1 item. Maximum number of 20 items.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidOperationException**

The operation is not valid on the specified resource.

HTTP Status Code: 400

**InvalidParameterException**

A parameter is specified incorrectly.

HTTP Status Code: 400

**ResourceNotFoundException**

The specified resource does not exist.

HTTP Status Code: 400

**ServiceUnavailableException**

The service cannot complete the request.

HTTP Status Code: 500

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/logs-2014-03-28/gettransformer.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/logs-2014-03-28/gettransformer.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/gettransformer.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/logs-2014-03-28/gettransformer.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/gettransformer.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/logs-2014-03-28/gettransformer.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/logs-2014-03-28/gettransformer.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/logs-2014-03-28/gettransformer.md)

- [AWS SDK for Python](../../../../services/goto/boto3/logs-2014-03-28/gettransformer.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/gettransformer.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetScheduledQueryHistory

ListAggregateLogGroupSummaries

All content copied from https://docs.aws.amazon.com/.
