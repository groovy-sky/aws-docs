# UpdateTrafficPolicyComment

Updates the comment for a specified traffic policy version.

## Request Syntax

```nohighlight

POST /2013-04-01/trafficpolicy/Id/Version HTTP/1.1
<?xml version="1.0" encoding="UTF-8"?>
<UpdateTrafficPolicyCommentRequest xmlns="https://route53.amazonaws.com/doc/2013-04-01/">
   <Comment>string</Comment>
</UpdateTrafficPolicyCommentRequest>
```

## URI Request Parameters

The request uses the following URI parameters.

**[Id](#API_UpdateTrafficPolicyComment_RequestSyntax)**

The value of `Id` for the traffic policy that you want to update the
comment for.

Length Constraints: Minimum length of 1. Maximum length of 36.

Required: Yes

**[Version](#API_UpdateTrafficPolicyComment_RequestSyntax)**

The value of `Version` for the traffic policy that you want to update the
comment for.

Valid Range: Minimum value of 1. Maximum value of 1000.

Required: Yes

## Request Body

The request accepts the following data in XML format.

**[UpdateTrafficPolicyCommentRequest](#API_UpdateTrafficPolicyComment_RequestSyntax)**

Root level tag for the UpdateTrafficPolicyCommentRequest parameters.

Required: Yes

**[Comment](#API_UpdateTrafficPolicyComment_RequestSyntax)**

The new comment for the specified traffic policy and version.

Type: String

Length Constraints: Maximum length of 1024.

Required: Yes

## Response Syntax

```nohighlight

HTTP/1.1 200
<?xml version="1.0" encoding="UTF-8"?>
<UpdateTrafficPolicyCommentResponse>
   <TrafficPolicy>
      <Comment>string</Comment>
      <Document>string</Document>
      <Id>string</Id>
      <Name>string</Name>
      <Type>string</Type>
      <Version>integer</Version>
   </TrafficPolicy>
</UpdateTrafficPolicyCommentResponse>
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in XML format by the service.

**[UpdateTrafficPolicyCommentResponse](#API_UpdateTrafficPolicyComment_ResponseSyntax)**

Root level tag for the UpdateTrafficPolicyCommentResponse parameters.

Required: Yes

**[TrafficPolicy](#API_UpdateTrafficPolicyComment_ResponseSyntax)**

A complex type that contains settings for the specified traffic policy.

Type: [TrafficPolicy](https://docs.aws.amazon.com/Route53/latest/APIReference/API_TrafficPolicy.html) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**ConcurrentModification**

Another user submitted a request to create, update, or delete the object at the same
time that you did. Retry the request.

**message**

HTTP Status Code: 400

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

This example illustrates one usage of UpdateTrafficPolicyComment.

```

POST /2013-04-01/trafficpolicy/12345678-abcd-9876-fedc-1a2b3c4de5f6/42 HTTP/1.1
<?xml version="1.0" encoding="UTF-8"?>
<UpdateTrafficPolicyCommentRequest xmlns="https://route53.amazonaws.com/doc/2013-04-01/">
   <Comment>Updated comment</Comment>
</UpdateTrafficPolicyCommentRequest>
```

### Example Response

This example illustrates one usage of UpdateTrafficPolicyComment.

```nohighlight

HTTP/1.1 200 OK
<?xml version="1.0" encoding="UTF-8"?>
<UpdateTrafficPolicyCommentResponse xmlns="https://route53.amazonaws.com/doc/2013-04-01/">
   <Id>12345678-abcd-9876-fedc-1a2b3c4de5f6</Id>
   <VersionNumber>42</VersionNumber>
   <Name>MyTrafficPolicy</Name>
   <Type>A</Type>
   <Document>definition of the traffic policy</Document>
   <Comment>Updated comment</Comment>
</UpdateTrafficPolicyCommentResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/route53-2013-04-01/UpdateTrafficPolicyComment)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/route53-2013-04-01/UpdateTrafficPolicyComment)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/route53-2013-04-01/UpdateTrafficPolicyComment)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/route53-2013-04-01/UpdateTrafficPolicyComment)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/route53-2013-04-01/UpdateTrafficPolicyComment)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/route53-2013-04-01/UpdateTrafficPolicyComment)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/route53-2013-04-01/UpdateTrafficPolicyComment)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/route53-2013-04-01/UpdateTrafficPolicyComment)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/route53-2013-04-01/UpdateTrafficPolicyComment)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/route53-2013-04-01/UpdateTrafficPolicyComment)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

UpdateHostedZoneFeatures

UpdateTrafficPolicyInstance
