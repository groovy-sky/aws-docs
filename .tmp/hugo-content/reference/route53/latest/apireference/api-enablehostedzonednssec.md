# EnableHostedZoneDNSSEC

Enables DNSSEC signing in a specific hosted zone.

## Request Syntax

```nohighlight

POST /2013-04-01/hostedzone/Id/enable-dnssec HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[Id](#API_EnableHostedZoneDNSSEC_RequestSyntax)**

A unique string used to identify a hosted zone.

Length Constraints: Maximum length of 32.

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
<?xml version="1.0" encoding="UTF-8"?>
<EnableHostedZoneDNSSECResponse>
   <ChangeInfo>
      <Comment>string</Comment>
      <Id>string</Id>
      <Status>string</Status>
      <SubmittedAt>timestamp</SubmittedAt>
   </ChangeInfo>
</EnableHostedZoneDNSSECResponse>
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in XML format by the service.

**[EnableHostedZoneDNSSECResponse](#API_EnableHostedZoneDNSSEC_ResponseSyntax)**

Root level tag for the EnableHostedZoneDNSSECResponse parameters.

Required: Yes

**[ChangeInfo](#API_EnableHostedZoneDNSSEC_ResponseSyntax)**

A complex type that describes change information about changes made to your hosted
zone.

Type: [ChangeInfo](api-changeinfo.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**ConcurrentModification**

Another user submitted a request to create, update, or delete the object at the same
time that you did. Retry the request.

**message**

HTTP Status Code: 400

**DNSSECNotFound**

The hosted zone doesn't have any DNSSEC resources.

HTTP Status Code: 400

**HostedZonePartiallyDelegated**

The hosted zone nameservers don't match the parent nameservers. The hosted zone and
parent must have the same nameservers.

HTTP Status Code: 400

**InvalidArgument**

Parameter name is not valid.

**message**

HTTP Status Code: 400

**InvalidInput**

The input is not valid.

**message**

HTTP Status Code: 400

**InvalidKeySigningKeyStatus**

The key-signing key (KSK) status isn't valid or another KSK has the status
`INTERNAL_FAILURE`.

HTTP Status Code: 400

**InvalidKMSArn**

The KeyManagementServiceArn that you specified isn't valid to use with DNSSEC
signing.

HTTP Status Code: 400

**KeySigningKeyWithActiveStatusNotFound**

A key-signing key (KSK) with `ACTIVE` status wasn't found.

HTTP Status Code: 400

**NoSuchHostedZone**

No hosted zone exists with the ID that you specified.

**message**

HTTP Status Code: 404

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/route53-2013-04-01/enablehostedzonednssec.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/route53-2013-04-01/enablehostedzonednssec.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/route53-2013-04-01/enablehostedzonednssec.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/route53-2013-04-01/enablehostedzonednssec.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/route53-2013-04-01/enablehostedzonednssec.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/route53-2013-04-01/enablehostedzonednssec.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/route53-2013-04-01/enablehostedzonednssec.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/route53-2013-04-01/enablehostedzonednssec.md)

- [AWS SDK for Python](../../../../services/goto/boto3/route53-2013-04-01/enablehostedzonednssec.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/route53-2013-04-01/enablehostedzonednssec.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DisassociateVPCFromHostedZone

GetAccountLimit
