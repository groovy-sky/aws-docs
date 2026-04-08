# GetQueryLoggingConfig

Gets information about a specified configuration for DNS query logging.

For more information about DNS query logs, see [CreateQueryLoggingConfig](api-createqueryloggingconfig.md) and [Logging DNS\
Queries](../../../../services/route53/latest/developerguide/query-logs.md).

## Request Syntax

```nohighlight

GET /2013-04-01/queryloggingconfig/Id HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[Id](#API_GetQueryLoggingConfig_RequestSyntax)**

The ID of the configuration for DNS query logging that you want to get information
about.

Length Constraints: Minimum length of 1. Maximum length of 36.

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
<?xml version="1.0" encoding="UTF-8"?>
<GetQueryLoggingConfigResponse>
   <QueryLoggingConfig>
      <CloudWatchLogsLogGroupArn>string</CloudWatchLogsLogGroupArn>
      <HostedZoneId>string</HostedZoneId>
      <Id>string</Id>
   </QueryLoggingConfig>
</GetQueryLoggingConfigResponse>
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in XML format by the service.

**[GetQueryLoggingConfigResponse](#API_GetQueryLoggingConfig_ResponseSyntax)**

Root level tag for the GetQueryLoggingConfigResponse parameters.

Required: Yes

**[QueryLoggingConfig](#API_GetQueryLoggingConfig_ResponseSyntax)**

A complex type that contains information about the query logging configuration that
you specified in a [GetQueryLoggingConfig](api-getqueryloggingconfig.md) request.

Type: [QueryLoggingConfig](api-queryloggingconfig.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidInput**

The input is not valid.

**message**

HTTP Status Code: 400

**NoSuchQueryLoggingConfig**

There is no DNS query logging configuration with the specified ID.

HTTP Status Code: 404

## Examples

### Example Request

The following request gets information about the configuration with the ID
`87654321-dcba-1234-abcd-1a2b3c4d5e6f`.

```

GET /2013-04-01/queryloggingconfig HTTP/1.1
<?xml version="1.0" encoding="UTF-8"?>
<GetQueryLoggingConfigRequest xmlns="https://route53.amazonaws.com/doc/2013-04-01/">
   <Id>87654321-dcba-1234-abcd-1a2b3c4d5e6f</Id>
</GetQueryLoggingConfigRequest>
```

### Example Response

This example illustrates one usage of GetQueryLoggingConfig.

```

HTTP/1.1 200 OK
<?xml version="1.0" encoding="UTF-8"?>
<GetQueryLoggingConfigResponse xmlns="https://route53.amazonaws.com/doc/2013-04-01/">
   <QueryLoggingConfig>
      <Id>87654321-dcba-1234-abcd-1a2b3c4d5e6f</Id>
      <HostedZoneId>Z1D633PJN98FT9</HostedZoneId>
      <CloudWatchLogsLogGroupArn>arn:aws:logs:us-east-1:111111111111:log-group:example.com:*</CloudWatchLogsLogGroupArn>
   </QueryLoggingConfig>
</GetQueryLoggingConfigResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/route53-2013-04-01/getqueryloggingconfig.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/route53-2013-04-01/getqueryloggingconfig.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/route53-2013-04-01/getqueryloggingconfig.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/route53-2013-04-01/getqueryloggingconfig.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/route53-2013-04-01/getqueryloggingconfig.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/route53-2013-04-01/getqueryloggingconfig.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/route53-2013-04-01/getqueryloggingconfig.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/route53-2013-04-01/getqueryloggingconfig.md)

- [AWS SDK for Python](../../../../services/goto/boto3/route53-2013-04-01/getqueryloggingconfig.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/route53-2013-04-01/getqueryloggingconfig.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

GetHostedZoneLimit

GetReusableDelegationSet
