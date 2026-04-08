# PutTransformer

Creates or updates a _log transformer_ for a single log group. You use
log transformers to transform log events into a different format, making them easier for you
to process and analyze. You can also transform logs from different sources into standardized
formats that contains relevant, source-specific information.

After you have created a transformer, CloudWatch Logs performs the transformations at
the time of log ingestion. You can then refer to the transformed versions of the logs during
operations such as querying with CloudWatch Logs Insights or creating metric filters or
subscription filers.

You can also use a transformer to copy metadata from metadata keys into the log events
themselves. This metadata can include log group name, log stream name, account ID and
Region.

A transformer for a log group is a series of processors, where each processor applies one
type of transformation to the log events ingested into this log group. The processors work one
after another, in the order that you list them, like a pipeline. For more information about
the available processors to use in a transformer, see [Processors that you can use](../../../../services/amazoncloudwatch/latest/logs/cloudwatch-logs-transformation.md#CloudWatch-Logs-Transformation-Processors).

Having log events in standardized format enables visibility across your applications for
your log analysis, reporting, and alarming needs. CloudWatch Logs provides transformation
for common log types with out-of-the-box transformation templates for major AWS
log sources such as VPC flow logs, Lambda, and Amazon RDS. You can use
pre-built transformation templates or create custom transformation policies.

You can create transformers only for the log groups in the Standard log class.

You can also set up a transformer at the account level. For more information, see [PutAccountPolicy](api-putaccountpolicy.md). If there is both a log-group level transformer created with
`PutTransformer` and an account-level transformer that could apply to the same
log group, the log group uses only the log-group level transformer. It ignores the
account-level transformer.

## Request Syntax

```nohighlight

{
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

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[logGroupIdentifier](#API_PutTransformer_RequestSyntax)**

Specify either the name or ARN of the log group to create the transformer for.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Pattern: `[\w#+=/:,.@-]*`

Required: Yes

**[transformerConfig](#API_PutTransformer_RequestSyntax)**

This structure contains the configuration of this log transformer. A log transformer is an
array of processors, where each processor applies one type of transformation to the log events
that are ingested.

Type: Array of [Processor](api-processor.md) objects

Array Members: Minimum number of 1 item. Maximum number of 20 items.

Required: Yes

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidOperationException**

The operation is not valid on the specified resource.

HTTP Status Code: 400

**InvalidParameterException**

A parameter is specified incorrectly.

HTTP Status Code: 400

**LimitExceededException**

You have reached the maximum number of resources that can be created.

HTTP Status Code: 400

**OperationAbortedException**

Multiple concurrent requests to update the same resource were in conflict.

HTTP Status Code: 400

**ResourceNotFoundException**

The specified resource does not exist.

HTTP Status Code: 400

**ServiceUnavailableException**

The service cannot complete the request.

HTTP Status Code: 500

## Examples

### To create a log transformer

The following example creates a log transformer for the specified log
group.

#### Sample Request

```

POST / HTTP/1.1
Host: logs.<region>.<domain>
X-Amz-Date: <DATE>
Authorization: AWS4-HMAC-SHA256 Credential=<Credential>, SignedHeaders=content-type;date;host;user-agent;x-amz-date;x-amz-target;x-amzn-requestid, Signature=<Signature>
User-Agent: <UserAgentString>
Accept: application/json
Content-Type: application/x-amz-json-1.1
Content-Length: <PayloadSizeBytes>
Connection: Keep-Alive
X-Amz-Target: Logs_20140328.PutTransformer
{
    "logGroupIdentifier": "my-log-group-name",
    "transformerConfig": [
        {
            "parseJSON": {}
        },
        {
            "addKeys": {
                "entries": [
                    {
                        "key": "metadata.transformed_in",
                        "value": "CloudWatchLogs"
                    }
                ]
            }
        },
        {
            "trimString": {
                "withKeys": [
                    "status"
                ]
            }
        },
        {
            "lowerCaseString": {
                "withKeys": [
                    "status"
                ]
            }
        }
    ]
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/logs-2014-03-28/puttransformer.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/logs-2014-03-28/puttransformer.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/puttransformer.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/logs-2014-03-28/puttransformer.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/puttransformer.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/logs-2014-03-28/puttransformer.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/logs-2014-03-28/puttransformer.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/logs-2014-03-28/puttransformer.md)

- [AWS SDK for Python](../../../../services/goto/boto3/logs-2014-03-28/puttransformer.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/puttransformer.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

PutSubscriptionFilter

StartLiveTail
