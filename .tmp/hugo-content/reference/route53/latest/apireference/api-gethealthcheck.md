# GetHealthCheck

Gets information about a specified health check.

## Request Syntax

```nohighlight

GET /2013-04-01/healthcheck/HealthCheckId HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[HealthCheckId](#API_GetHealthCheck_RequestSyntax)**

The identifier that Amazon Route 53 assigned to the health check when you created it.
When you add or update a resource record set, you use this value to specify which health
check to use. The value can be up to 64 characters long.

Length Constraints: Maximum length of 64.

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
<?xml version="1.0" encoding="UTF-8"?>
<GetHealthCheckResponse>
   <HealthCheck>
      <CallerReference>string</CallerReference>
      <CloudWatchAlarmConfiguration>
         <ComparisonOperator>string</ComparisonOperator>
         <Dimensions>
            <Dimension>
               <Name>string</Name>
               <Value>string</Value>
            </Dimension>
         </Dimensions>
         <EvaluationPeriods>integer</EvaluationPeriods>
         <MetricName>string</MetricName>
         <Namespace>string</Namespace>
         <Period>integer</Period>
         <Statistic>string</Statistic>
         <Threshold>double</Threshold>
      </CloudWatchAlarmConfiguration>
      <HealthCheckConfig>
         <AlarmIdentifier>
            <Name>string</Name>
            <Region>string</Region>
         </AlarmIdentifier>
         <ChildHealthChecks>
            <ChildHealthCheck>string</ChildHealthCheck>
         </ChildHealthChecks>
         <Disabled>boolean</Disabled>
         <EnableSNI>boolean</EnableSNI>
         <FailureThreshold>integer</FailureThreshold>
         <FullyQualifiedDomainName>string</FullyQualifiedDomainName>
         <HealthThreshold>integer</HealthThreshold>
         <InsufficientDataHealthStatus>string</InsufficientDataHealthStatus>
         <Inverted>boolean</Inverted>
         <IPAddress>string</IPAddress>
         <MeasureLatency>boolean</MeasureLatency>
         <Port>integer</Port>
         <Regions>
            <Region>string</Region>
         </Regions>
         <RequestInterval>integer</RequestInterval>
         <ResourcePath>string</ResourcePath>
         <RoutingControlArn>string</RoutingControlArn>
         <SearchString>string</SearchString>
         <Type>string</Type>
      </HealthCheckConfig>
      <HealthCheckVersion>long</HealthCheckVersion>
      <Id>string</Id>
      <LinkedService>
         <Description>string</Description>
         <ServicePrincipal>string</ServicePrincipal>
      </LinkedService>
   </HealthCheck>
</GetHealthCheckResponse>
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in XML format by the service.

**[GetHealthCheckResponse](#API_GetHealthCheck_ResponseSyntax)**

Root level tag for the GetHealthCheckResponse parameters.

Required: Yes

**[HealthCheck](#API_GetHealthCheck_ResponseSyntax)**

A complex type that contains information about one health check that is associated
with the current AWS account.

Type: [HealthCheck](api-healthcheck.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**IncompatibleVersion**

The resource you're trying to access is unsupported on this Amazon Route 53
endpoint.

HTTP Status Code: 400

**InvalidInput**

The input is not valid.

**message**

HTTP Status Code: 400

**NoSuchHealthCheck**

No health check exists with the specified ID.

**message**

HTTP Status Code: 404

## Examples

### Example Request

This example illustrates one usage of GetHealthCheck.

```

GET /2013-04-01/healthcheck/018927304987
```

### Example Response

This example illustrates one usage of GetHealthCheck.

```

HTTP/1.1 200 OK
<?xml version="1.0" encoding="UTF-8"?>
<GetHealthCheckResponse xmlns="https://route53.amazonaws.com/doc/2013-04-01/">
   <HealthCheck>
      <Id>abcdef11-2222-3333-4444-555555fedcba</Id>
      <CallerReference>example.com 192.0.2.17</CallerReference>
      <HealthCheckConfig>
         <IPAddress>192.0.2.17</IPAddress>
         <Port>80</Port>
         <Type>HTTP</Type>
         <ResourcePath>/docs/route-53-health-check.html</ResourcePath>
         <FullyQualifiedDomainName>example.com</FullyQualifiedDomainName>
         <RequestInterval>30</RequestInterval>
         <FailureThreshold>3</FailureThreshold>
         <MeasureLatency>true</MeasureLatency>
         <EnableSNI>true</EnableSNI>
         <Regions>
            <Region>ap-southeast-1</Region>
            <Region>ap-southeast-2</Region>
            <Region>ap-northeast-1</Region>
         </Regions>
         <Inverted>false</Inverted>
      </HealthCheckConfig>
      <HealthCheckVersion>2<HealthCheckVersion>
   </HealthCheck>
</GetHealthCheckResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/route53-2013-04-01/gethealthcheck.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/route53-2013-04-01/gethealthcheck.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/route53-2013-04-01/gethealthcheck.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/route53-2013-04-01/gethealthcheck.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/route53-2013-04-01/gethealthcheck.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/route53-2013-04-01/gethealthcheck.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/route53-2013-04-01/gethealthcheck.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/route53-2013-04-01/gethealthcheck.md)

- [AWS SDK for Python](../../../../services/goto/boto3/route53-2013-04-01/gethealthcheck.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/route53-2013-04-01/gethealthcheck.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

GetGeoLocation

GetHealthCheckCount
