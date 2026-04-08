# UpdateHostedZoneFeatures

Updates the features configuration for a hosted zone. This operation allows you to enable or disable specific features for your hosted zone, such as accelerated recovery.

Accelerated recovery enables you to update DNS records in your public hosted zone even when the us-east-1 region is unavailable.

## Request Syntax

```nohighlight

POST /2013-04-01/hostedzone/Id/features HTTP/1.1
<?xml version="1.0" encoding="UTF-8"?>
<UpdateHostedZoneFeaturesRequest xmlns="https://route53.amazonaws.com/doc/2013-04-01/">
   <EnableAcceleratedRecovery>boolean</EnableAcceleratedRecovery>
</UpdateHostedZoneFeaturesRequest>
```

## URI Request Parameters

The request uses the following URI parameters.

**[Id](#API_UpdateHostedZoneFeatures_RequestSyntax)**

The ID of the hosted zone for which you want to update features. This is the unique identifier for your hosted zone.

Length Constraints: Maximum length of 32.

Required: Yes

## Request Body

The request accepts the following data in XML format.

**[UpdateHostedZoneFeaturesRequest](#API_UpdateHostedZoneFeatures_RequestSyntax)**

Root level tag for the UpdateHostedZoneFeaturesRequest parameters.

Required: Yes

**[EnableAcceleratedRecovery](#API_UpdateHostedZoneFeatures_RequestSyntax)**

Specifies whether to enable accelerated recovery for the hosted zone. Set to `true` to enable accelerated recovery, or `false` to disable it.

Type: Boolean

Required: No

## Response Syntax

```

HTTP/1.1 200

```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidInput**

The input is not valid.

**message**

HTTP Status Code: 400

**LimitsExceeded**

This operation can't be completed because the current account has reached the
limit on the resource you are trying to create. To request a higher limit, [create a case](http://aws.amazon.com/route53-request) with the AWS Support
Center.

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

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/route53-2013-04-01/updatehostedzonefeatures.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/route53-2013-04-01/updatehostedzonefeatures.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/route53-2013-04-01/updatehostedzonefeatures.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/route53-2013-04-01/updatehostedzonefeatures.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/route53-2013-04-01/updatehostedzonefeatures.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/route53-2013-04-01/updatehostedzonefeatures.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/route53-2013-04-01/updatehostedzonefeatures.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/route53-2013-04-01/updatehostedzonefeatures.md)

- [AWS SDK for Python](../../../../services/goto/boto3/route53-2013-04-01/updatehostedzonefeatures.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/route53-2013-04-01/updatehostedzonefeatures.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

UpdateHostedZoneComment

UpdateTrafficPolicyComment
