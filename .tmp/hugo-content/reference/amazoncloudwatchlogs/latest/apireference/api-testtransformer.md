# TestTransformer

Use this operation to test a log transformer. You enter the transformer configuration and
a set of log events to test with. The operation responds with an array that includes the
original log events and the transformed versions.

## Request Syntax

```nohighlight

{
   "logEventMessages": [ "string" ],
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

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[logEventMessages](#API_TestTransformer_RequestSyntax)**

An array of the raw log events that you want to use to test this transformer.

Type: Array of strings

Array Members: Minimum number of 1 item. Maximum number of 50 items.

Length Constraints: Minimum length of 1.

Required: Yes

**[transformerConfig](#API_TestTransformer_RequestSyntax)**

This structure contains the configuration of this log transformer that you want to test. A
log transformer is an array of processors, where each processor applies one type of
transformation to the log events that are ingested.

Type: Array of [Processor](api-processor.md) objects

Array Members: Minimum number of 1 item. Maximum number of 20 items.

Required: Yes

## Response Syntax

```nohighlight

{
   "transformedLogs": [
      {
         "eventMessage": "string",
         "eventNumber": number,
         "transformedEventMessage": "string"
      }
   ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[transformedLogs](#API_TestTransformer_ResponseSyntax)**

An array where each member of the array includes both the original version and the
transformed version of one of the log events that you input.

Type: Array of [TransformedLogRecord](api-transformedlogrecord.md) objects

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidOperationException**

The operation is not valid on the specified resource.

HTTP Status Code: 400

**InvalidParameterException**

A parameter is specified incorrectly.

HTTP Status Code: 400

**ServiceUnavailableException**

The service cannot complete the request.

HTTP Status Code: 500

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/logs-2014-03-28/testtransformer.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/logs-2014-03-28/testtransformer.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/testtransformer.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/logs-2014-03-28/testtransformer.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/testtransformer.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/logs-2014-03-28/testtransformer.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/logs-2014-03-28/testtransformer.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/logs-2014-03-28/testtransformer.md)

- [AWS SDK for Python](../../../../services/goto/boto3/logs-2014-03-28/testtransformer.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/testtransformer.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TestMetricFilter

UntagLogGroup

All content copied from https://docs.aws.amazon.com/.
