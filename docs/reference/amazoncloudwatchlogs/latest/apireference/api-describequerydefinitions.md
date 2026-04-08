# DescribeQueryDefinitions

This operation returns a paginated list of your saved CloudWatch Logs Insights query
definitions. You can retrieve query definitions from the current account or from a source
account that is linked to the current account.

You can use the `queryDefinitionNamePrefix` parameter to limit the results to
only the query definitions that have names that start with a certain string.

## Request Syntax

```nohighlight

{
   "maxResults": number,
   "nextToken": "string",
   "queryDefinitionNamePrefix": "string",
   "queryLanguage": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[maxResults](#API_DescribeQueryDefinitions_RequestSyntax)**

Limits the number of returned query definitions to the specified number.

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 1000.

Required: No

**[nextToken](#API_DescribeQueryDefinitions_RequestSyntax)**

The token for the next set of items to return. The token expires after 24
hours.

Type: String

Length Constraints: Minimum length of 1.

Required: No

**[queryDefinitionNamePrefix](#API_DescribeQueryDefinitions_RequestSyntax)**

Use this parameter to filter your results to only the query definitions that have names
that start with the prefix you specify.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Required: No

**[queryLanguage](#API_DescribeQueryDefinitions_RequestSyntax)**

The query language used for this query. For more information about the query languages
that CloudWatch Logs supports, see [Supported query\
languages](../../../../services/amazoncloudwatch/latest/logs/cwl-analyzelogdata-languages.md).

Type: String

Valid Values: `CWLI | SQL | PPL`

Required: No

## Response Syntax

```nohighlight

{
   "nextToken": "string",
   "queryDefinitions": [
      {
         "lastModified": number,
         "logGroupNames": [ "string" ],
         "name": "string",
         "parameters": [
            {
               "defaultValue": "string",
               "description": "string",
               "name": "string"
            }
         ],
         "queryDefinitionId": "string",
         "queryLanguage": "string",
         "queryString": "string"
      }
   ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[nextToken](#API_DescribeQueryDefinitions_ResponseSyntax)**

The token for the next set of items to return. The token expires after 24
hours.

Type: String

Length Constraints: Minimum length of 1.

**[queryDefinitions](#API_DescribeQueryDefinitions_ResponseSyntax)**

The list of query definitions that match your request.

Type: Array of [QueryDefinition](api-querydefinition.md) objects

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidParameterException**

A parameter is specified incorrectly.

HTTP Status Code: 400

**ServiceUnavailableException**

The service cannot complete the request.

HTTP Status Code: 500

## Examples

### Example

This example retrieves a list of query definitions that have names that begin with
`lambda`.

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
X-Amz-Target: Logs_20140328.DescribeQueryDefinitions
{
   "queryDefinitionNamePrefix": "lambda",
   "maxResults": 2
}
```

#### Sample Response

```

HTTP/1.1 200 OK
x-amzn-RequestId: <RequestId>
Content-Type: application/x-amz-json-1.1
Content-Length: <PayloadSizeBytes>
Date: <Date>
{
   "nextToken": "abdcefghijlkmn",
   "queryDefinitions": [
      {
         "lastModified": 1549321515,
         "logGroupNames": [ "VPC_Flow_Log1", "VPC_Flow_Log2" ],
         "name": "VPC-top15-packet-transfers",
         "queryDefinitionId": "123456ab-12ab-123a-789e-1234567890ab",
         "queryString": "stats sum(packets) as packetsTransferred by srcAddr, dstAddr | sort packetsTransferred  desc | limit 15",
         "parameters": []
      },
      {
         "lastModified": 1557321299,
         "name": "ErrorsByLevel",
         "queryDefinitionId": "456789ab-abcd-1234-789e-0987654321ab",
         "queryString": "fields @timestamp, @message | filter level = {{logLevel}}",
         "parameters": [
            {
               "name": "logLevel",
               "defaultValue": "ERROR",
               "description": "Log level to filter (ERROR, WARN, INFO, DEBUG)"
            }
         ]
      }
   ]
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/logs-2014-03-28/describequerydefinitions.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/logs-2014-03-28/describequerydefinitions.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/describequerydefinitions.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/logs-2014-03-28/describequerydefinitions.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/describequerydefinitions.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/logs-2014-03-28/describequerydefinitions.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/logs-2014-03-28/describequerydefinitions.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/logs-2014-03-28/describequerydefinitions.md)

- [AWS SDK for Python](../../../../services/goto/boto3/logs-2014-03-28/describequerydefinitions.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/describequerydefinitions.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeQueries

DescribeResourcePolicies
