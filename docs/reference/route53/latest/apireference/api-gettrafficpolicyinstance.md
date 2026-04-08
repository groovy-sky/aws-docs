# GetTrafficPolicyInstance

Gets information about a specified traffic policy instance.

###### Note

Use `GetTrafficPolicyInstance` with the `id` of new traffic policy instance to confirm that the
`CreateTrafficPolicyInstance` or an `UpdateTrafficPolicyInstance` request completed successfully.
For more information, see the `State` response
element.

###### Note

In the Route 53 console, traffic policy instances are known as policy
records.

## Request Syntax

```nohighlight

GET /2013-04-01/trafficpolicyinstance/Id HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[Id](#API_GetTrafficPolicyInstance_RequestSyntax)**

The ID of the traffic policy instance that you want to get information about.

Length Constraints: Minimum length of 1. Maximum length of 36.

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
<?xml version="1.0" encoding="UTF-8"?>
<GetTrafficPolicyInstanceResponse>
   <TrafficPolicyInstance>
      <HostedZoneId>string</HostedZoneId>
      <Id>string</Id>
      <Message>string</Message>
      <Name>string</Name>
      <State>string</State>
      <TrafficPolicyId>string</TrafficPolicyId>
      <TrafficPolicyType>string</TrafficPolicyType>
      <TrafficPolicyVersion>integer</TrafficPolicyVersion>
      <TTL>long</TTL>
   </TrafficPolicyInstance>
</GetTrafficPolicyInstanceResponse>
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in XML format by the service.

**[GetTrafficPolicyInstanceResponse](#API_GetTrafficPolicyInstance_ResponseSyntax)**

Root level tag for the GetTrafficPolicyInstanceResponse parameters.

Required: Yes

**[TrafficPolicyInstance](#API_GetTrafficPolicyInstance_ResponseSyntax)**

A complex type that contains settings for the traffic policy instance.

Type: [TrafficPolicyInstance](api-trafficpolicyinstance.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidInput**

The input is not valid.

**message**

HTTP Status Code: 400

**NoSuchTrafficPolicyInstance**

No traffic policy instance exists with the specified ID.

**message**

HTTP Status Code: 404

## Examples

### Example Request

This example illustrates one usage of GetTrafficPolicyInstance.

```

GET /2013-04-01/trafficpolicyinstance/12131415-abac-5432-caba-6f5e4d3c2b1a
```

### Example Response

This example illustrates one usage of GetTrafficPolicyInstance.

```

HTTP/1.1 200 OK
<?xml version="1.0" encoding="UTF-8"?>
<GetTrafficPolicyInstanceResponse xmlns="https://route53.amazonaws.com/doc/2013-04-01/">
   <TrafficPolicyInstance>
      <Id>12131415-abac-5432-caba-6f5e4d3c2b1a</Id>
      <HostedZoneId>Z1D633PJN98FT9</HostedZoneId>
      <Name>www.example.com</Name>
      <TTL>300</TTL>
      <State>Applied</State>
      <Message/>
      <TrafficPolicyId>12345678-abcd-9876-fedc-1a2b3c4de5f6</TrafficPolicyId>
      <TrafficPolicyVersion>7</TrafficPolicyVersion>
      <TrafficPolicyType>A</TrafficPolicyType>
   </TrafficPolicyInstance>
</GetTrafficPolicyInstanceResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/route53-2013-04-01/gettrafficpolicyinstance.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/route53-2013-04-01/gettrafficpolicyinstance.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/route53-2013-04-01/gettrafficpolicyinstance.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/route53-2013-04-01/gettrafficpolicyinstance.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/route53-2013-04-01/gettrafficpolicyinstance.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/route53-2013-04-01/gettrafficpolicyinstance.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/route53-2013-04-01/gettrafficpolicyinstance.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/route53-2013-04-01/gettrafficpolicyinstance.md)

- [AWS SDK for Python](../../../../services/goto/boto3/route53-2013-04-01/gettrafficpolicyinstance.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/route53-2013-04-01/gettrafficpolicyinstance.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

GetTrafficPolicy

GetTrafficPolicyInstanceCount
