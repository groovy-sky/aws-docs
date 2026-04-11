# UpdateHostedZoneComment

Updates the comment for a specified hosted zone.

## Request Syntax

```nohighlight

POST /2013-04-01/hostedzone/Id HTTP/1.1
<?xml version="1.0" encoding="UTF-8"?>
<UpdateHostedZoneCommentRequest xmlns="https://route53.amazonaws.com/doc/2013-04-01/">
   <Comment>string</Comment>
</UpdateHostedZoneCommentRequest>
```

## URI Request Parameters

The request uses the following URI parameters.

**[Id](#API_UpdateHostedZoneComment_RequestSyntax)**

The ID for the hosted zone that you want to update the comment for.

Length Constraints: Maximum length of 32.

Required: Yes

## Request Body

The request accepts the following data in XML format.

**[UpdateHostedZoneCommentRequest](#API_UpdateHostedZoneComment_RequestSyntax)**

Root level tag for the UpdateHostedZoneCommentRequest parameters.

Required: Yes

**[Comment](#API_UpdateHostedZoneComment_RequestSyntax)**

The new comment for the hosted zone. If you don't specify a value for
`Comment`, Amazon Route 53 deletes the existing value of the
`Comment` element, if any.

Type: String

Length Constraints: Maximum length of 256.

Required: No

## Response Syntax

```nohighlight

HTTP/1.1 200
<?xml version="1.0" encoding="UTF-8"?>
<UpdateHostedZoneCommentResponse>
   <HostedZone>
      <CallerReference>string</CallerReference>
      <Config>
         <Comment>string</Comment>
         <PrivateZone>boolean</PrivateZone>
      </Config>
      <Features>
         <AcceleratedRecoveryStatus>string</AcceleratedRecoveryStatus>
         <FailureReasons>
            <AcceleratedRecovery>string</AcceleratedRecovery>
         </FailureReasons>
      </Features>
      <Id>string</Id>
      <LinkedService>
         <Description>string</Description>
         <ServicePrincipal>string</ServicePrincipal>
      </LinkedService>
      <Name>string</Name>
      <ResourceRecordSetCount>long</ResourceRecordSetCount>
   </HostedZone>
</UpdateHostedZoneCommentResponse>
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in XML format by the service.

**[UpdateHostedZoneCommentResponse](#API_UpdateHostedZoneComment_ResponseSyntax)**

Root level tag for the UpdateHostedZoneCommentResponse parameters.

Required: Yes

**[HostedZone](#API_UpdateHostedZoneComment_ResponseSyntax)**

A complex type that contains the response to the `UpdateHostedZoneComment`
request.

Type: [HostedZone](api-hostedzone.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidInput**

The input is not valid.

**message**

HTTP Status Code: 400

**NoSuchHostedZone**

No hosted zone exists with the ID that you specified.

**message**

HTTP Status Code: 404

**PriorRequestNotComplete**

If Amazon Route 53 can't process a request before the next request arrives, it will
reject subsequent requests for the same hosted zone and return an `HTTP 400
				error` ( `Bad request`). If Route 53 returns this error repeatedly
for the same request, we recommend that you wait, in intervals of increasing duration,
before you try the request again.

HTTP Status Code: 400

## Examples

### Example Request

This example illustrates one usage of UpdateHostedZoneComment.

```

POST /2013-04-01/hostedzone/hosted zone ID HTTP/1.1
<?xml version="1.0" encoding="UTF-8"?>
<UpdateHostedZoneCommentRequest xmlns="https://route53.amazonaws.com/doc/2013-04-01/">
   <Comment>for internal testing</Comment>
</UpdateHostedZoneCommentRequest>
```

### Example Response

This example illustrates one usage of UpdateHostedZoneComment.

```

HTTP/1.1 200 OK
<?xml version="1.0" encoding="UTF-8"?>
<UpdateHostedZoneCommentResponse xmlns="https://route53.amazonaws.com/doc/2013-04-01/">
   <HostedZone>
      <Id>/hostedzone/Z1D633PJN98FT9</Id>
      <Name>example.com</Name>
      <CallerReference>2014-10-15T01:36:41.958Z</CallerReference>
      <Config>
         <Comment>for internal testing</Comment>
         <PrivateZone>false</PrivateZone>
      </Config>
      <ResourceRecordSetCount>42</ResourceRecordSetCount>
   </HostedZone>
</UpdateHostedZoneCommentResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/route53-2013-04-01/updatehostedzonecomment.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/route53-2013-04-01/updatehostedzonecomment.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/route53-2013-04-01/updatehostedzonecomment.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/route53-2013-04-01/updatehostedzonecomment.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/route53-2013-04-01/updatehostedzonecomment.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/route53-2013-04-01/updatehostedzonecomment.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/route53-2013-04-01/updatehostedzonecomment.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/route53-2013-04-01/updatehostedzonecomment.md)

- [AWS SDK for Python](../../../../services/goto/boto3/route53-2013-04-01/updatehostedzonecomment.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/route53-2013-04-01/updatehostedzonecomment.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

UpdateHealthCheck

UpdateHostedZoneFeatures
