# ListQueryLoggingConfigs

Lists the configurations for DNS query logging that are associated with the current
AWS account or the configuration that is associated with a specified
hosted zone.

For more information about DNS query logs, see [CreateQueryLoggingConfig](api-createqueryloggingconfig.md). Additional information, including the format of
DNS query logs, appears in [Logging DNS Queries](../../../../services/route53/latest/developerguide/query-logs.md) in
the _Amazon Route 53 Developer Guide_.

## Request Syntax

```nohighlight

GET /2013-04-01/queryloggingconfig?hostedzoneid=HostedZoneId&maxresults=MaxResults&nexttoken=NextToken HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[hostedzoneid](#API_ListQueryLoggingConfigs_RequestSyntax)**

(Optional) If you want to list the query logging configuration that is associated with
a hosted zone, specify the ID in `HostedZoneId`.

If you don't specify a hosted zone ID, `ListQueryLoggingConfigs` returns
all of the configurations that are associated with the current AWS account.

Length Constraints: Maximum length of 32.

**[maxresults](#API_ListQueryLoggingConfigs_RequestSyntax)**

(Optional) The maximum number of query logging configurations that you want Amazon
Route 53 to return in response to the current request. If the current AWS account has more than `MaxResults` configurations, use the
value of [NextToken](api-listqueryloggingconfigs-api-listqueryloggingconfigs-requestsyntax.md) in the response to get the next page of results.

If you don't specify a value for `MaxResults`, Route 53 returns up to 100
configurations.

**[nexttoken](#API_ListQueryLoggingConfigs_RequestSyntax)**

(Optional) If the current AWS account has more than
`MaxResults` query logging configurations, use `NextToken` to
get the second and subsequent pages of results.

For the first `ListQueryLoggingConfigs` request, omit this value.

For the second and subsequent requests, get the value of `NextToken` from
the previous response and specify that value for `NextToken` in the
request.

Length Constraints: Maximum length of 1024.

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
<?xml version="1.0" encoding="UTF-8"?>
<ListQueryLoggingConfigsResponse>
   <NextToken>string</NextToken>
   <QueryLoggingConfigs>
      <QueryLoggingConfig>
         <CloudWatchLogsLogGroupArn>string</CloudWatchLogsLogGroupArn>
         <HostedZoneId>string</HostedZoneId>
         <Id>string</Id>
      </QueryLoggingConfig>
   </QueryLoggingConfigs>
</ListQueryLoggingConfigsResponse>
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in XML format by the service.

**[ListQueryLoggingConfigsResponse](#API_ListQueryLoggingConfigs_ResponseSyntax)**

Root level tag for the ListQueryLoggingConfigsResponse parameters.

Required: Yes

**[NextToken](#API_ListQueryLoggingConfigs_ResponseSyntax)**

If a response includes the last of the query logging configurations that are
associated with the current AWS account, `NextToken` doesn't
appear in the response.

If a response doesn't include the last of the configurations, you can get more
configurations by submitting another [ListQueryLoggingConfigs](api-listqueryloggingconfigs.md) request. Get the value of `NextToken`
that Amazon Route 53 returned in the previous response and include it in
`NextToken` in the next request.

Type: String

Length Constraints: Maximum length of 1024.

**[QueryLoggingConfigs](#API_ListQueryLoggingConfigs_ResponseSyntax)**

An array that contains one [QueryLoggingConfig](api-queryloggingconfig.md) element for each configuration for DNS query logging
that is associated with the current AWS account.

Type: Array of [QueryLoggingConfig](api-queryloggingconfig.md) objects

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidInput**

The input is not valid.

**message**

HTTP Status Code: 400

**InvalidPaginationToken**

The value that you specified to get the second or subsequent page of results is
invalid.

HTTP Status Code: 400

**NoSuchHostedZone**

No hosted zone exists with the ID that you specified.

**message**

HTTP Status Code: 404

## Examples

### Example Request

The following request gets the configuration that is associated with the
hosted zone `Z1D633PJN98FT9`.

```

GET /2013-04-01/queryloggingconfig HTTP/1.1
<?xml version="1.0" encoding="UTF-8"?>
<ListQueryLoggingConfigsRequest xmlns="https://route53.amazonaws.com/doc/2013-04-01/">
   <HostedZoneId>Z1D633PJN98FT9</HostedZoneId>
   <MaxResults>1</MaxResults>
</ListQueryLoggingConfigsRequest>
```

### Example Response

This example illustrates one usage of ListQueryLoggingConfigs.

```

HTTP/1.1 200 OK
<?xml version="1.0" encoding="UTF-8"?>
<ListQueryLoggingConfigsResponse xmlns="https://route53.amazonaws.com/doc/2013-04-01/">
   <NextToken>87654321-dcba-1234-abcd-1a2b3c4d5e70</NextToken>
   <QueryLoggingConfigs>
      <QueryLoggingConfig>
         <Id>87654321-dcba-1234-abcd-1a2b3c4d5e6f</Id>
         <HostedZoneId>Z1D633PJN98FT9</HostedZoneId>
         <CloudWatchLogsLogGroupArn>arn:aws:logs:us-east-1:111111111111:log-group:example.com:*</CloudWatchLogsLogGroupArn>
      </QueryLoggingConfig>
   </QueryLoggingConfigs>
</ListQueryLoggingConfigsResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/route53-2013-04-01/listqueryloggingconfigs.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/route53-2013-04-01/listqueryloggingconfigs.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/route53-2013-04-01/listqueryloggingconfigs.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/route53-2013-04-01/listqueryloggingconfigs.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/route53-2013-04-01/listqueryloggingconfigs.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/route53-2013-04-01/listqueryloggingconfigs.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/route53-2013-04-01/listqueryloggingconfigs.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/route53-2013-04-01/listqueryloggingconfigs.md)

- [AWS SDK for Python](../../../../services/goto/boto3/route53-2013-04-01/listqueryloggingconfigs.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/route53-2013-04-01/listqueryloggingconfigs.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ListHostedZonesByVPC

ListResourceRecordSets
