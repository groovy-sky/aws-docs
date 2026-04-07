# GetTrafficPolicy

Gets information about a specific traffic policy version.

For information about how of deleting a traffic policy affects the response from
`GetTrafficPolicy`, see [DeleteTrafficPolicy](api-deletetrafficpolicy.md).

## Request Syntax

```nohighlight

GET /2013-04-01/trafficpolicy/Id/Version HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[Id](#API_GetTrafficPolicy_RequestSyntax)**

The ID of the traffic policy that you want to get information about.

Length Constraints: Minimum length of 1. Maximum length of 36.

Required: Yes

**[Version](#API_GetTrafficPolicy_RequestSyntax)**

The version number of the traffic policy that you want to get information
about.

Valid Range: Minimum value of 1. Maximum value of 1000.

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
<?xml version="1.0" encoding="UTF-8"?>
<GetTrafficPolicyResponse>
   <TrafficPolicy>
      <Comment>string</Comment>
      <Document>string</Document>
      <Id>string</Id>
      <Name>string</Name>
      <Type>string</Type>
      <Version>integer</Version>
   </TrafficPolicy>
</GetTrafficPolicyResponse>
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in XML format by the service.

**[GetTrafficPolicyResponse](#API_GetTrafficPolicy_ResponseSyntax)**

Root level tag for the GetTrafficPolicyResponse parameters.

Required: Yes

**[TrafficPolicy](#API_GetTrafficPolicy_ResponseSyntax)**

A complex type that contains settings for the specified traffic policy.

Type: [TrafficPolicy](https://docs.aws.amazon.com/Route53/latest/APIReference/API_TrafficPolicy.html) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidInput**

The input is not valid.

**message**

HTTP Status Code: 400

**NoSuchTrafficPolicy**

No traffic policy exists with the specified ID.

**message**

HTTP Status Code: 404

## Examples

### Example Request

This example illustrates one usage of GetTrafficPolicy.

```

GET /2013-04-01/trafficpolicy/12345678-abcd-9876-fedc-1a2b3c4de5f6/2
```

### Example Response

This example illustrates one usage of GetTrafficPolicy.

```

HTTP/1.1 200 OK
<?xml version="1.0" encoding="UTF-8"?>
<GetTrafficPolicyResponse xmlns="https://route53.amazonaws.com/doc/2013-04-01/">
   <TrafficPolicy>
      <Id>12345678-abcd-9876-fedc-1a2b3c4de5f6</Id>
      <Version>2</Version>
      <Name>MyTrafficPolicy</Name>
      <Type>A</Type>
      <Document>traffic policy definition in JSON format</Document>
      <Comment>New traffic policy version</Comment>
   </TrafficPolicy>
</GetTrafficPolicyResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/route53-2013-04-01/GetTrafficPolicy)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/route53-2013-04-01/GetTrafficPolicy)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/route53-2013-04-01/GetTrafficPolicy)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/route53-2013-04-01/GetTrafficPolicy)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/route53-2013-04-01/GetTrafficPolicy)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/route53-2013-04-01/GetTrafficPolicy)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/route53-2013-04-01/GetTrafficPolicy)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/route53-2013-04-01/GetTrafficPolicy)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/route53-2013-04-01/GetTrafficPolicy)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/route53-2013-04-01/GetTrafficPolicy)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

GetReusableDelegationSetLimit

GetTrafficPolicyInstance
