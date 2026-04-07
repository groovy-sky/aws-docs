# DisableHostedZoneDNSSEC

Disables DNSSEC signing in a specific hosted zone. This action does not deactivate any
key-signing keys (KSKs) that are active in the hosted zone.

## Request Syntax

```nohighlight

POST /2013-04-01/hostedzone/Id/disable-dnssec HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[Id](#API_DisableHostedZoneDNSSEC_RequestSyntax)**

A unique string used to identify a hosted zone.

Length Constraints: Maximum length of 32.

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
<?xml version="1.0" encoding="UTF-8"?>
<DisableHostedZoneDNSSECResponse>
   <ChangeInfo>
      <Comment>string</Comment>
      <Id>string</Id>
      <Status>string</Status>
      <SubmittedAt>timestamp</SubmittedAt>
   </ChangeInfo>
</DisableHostedZoneDNSSECResponse>
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in XML format by the service.

**[DisableHostedZoneDNSSECResponse](#API_DisableHostedZoneDNSSEC_ResponseSyntax)**

Root level tag for the DisableHostedZoneDNSSECResponse parameters.

Required: Yes

**[ChangeInfo](#API_DisableHostedZoneDNSSEC_ResponseSyntax)**

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

**KeySigningKeyInParentDSRecord**

The key-signing key (KSK) is specified in a parent DS record.

HTTP Status Code: 400

**NoSuchHostedZone**

No hosted zone exists with the ID that you specified.

**message**

HTTP Status Code: 404

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/route53-2013-04-01/DisableHostedZoneDNSSEC)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/route53-2013-04-01/DisableHostedZoneDNSSEC)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/route53-2013-04-01/DisableHostedZoneDNSSEC)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/route53-2013-04-01/DisableHostedZoneDNSSEC)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/route53-2013-04-01/DisableHostedZoneDNSSEC)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/route53-2013-04-01/DisableHostedZoneDNSSEC)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/route53-2013-04-01/DisableHostedZoneDNSSEC)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/route53-2013-04-01/DisableHostedZoneDNSSEC)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/route53-2013-04-01/DisableHostedZoneDNSSEC)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/route53-2013-04-01/DisableHostedZoneDNSSEC)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DeleteVPCAssociationAuthorization

DisassociateVPCFromHostedZone
