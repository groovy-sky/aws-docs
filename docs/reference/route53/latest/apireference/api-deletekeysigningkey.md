# DeleteKeySigningKey

Deletes a key-signing key (KSK). Before you can delete a KSK, you must deactivate it.
The KSK must be deactivated before you can delete it regardless of whether the hosted
zone is enabled for DNSSEC signing.

You can use [DeactivateKeySigningKey](api-deactivatekeysigningkey.md) to deactivate the key before you delete it.

Use [GetDNSSEC](api-getdnssec.md) to verify that the KSK is in an `INACTIVE`
status.

## Request Syntax

```nohighlight

DELETE /2013-04-01/keysigningkey/HostedZoneId/Name HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[HostedZoneId](#API_DeleteKeySigningKey_RequestSyntax)**

A unique string used to identify a hosted zone.

Length Constraints: Maximum length of 32.

Required: Yes

**[Name](#API_DeleteKeySigningKey_RequestSyntax)**

A string used to identify a key-signing key (KSK).

Length Constraints: Minimum length of 3. Maximum length of 128.

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
<?xml version="1.0" encoding="UTF-8"?>
<DeleteKeySigningKeyResponse>
   <ChangeInfo>
      <Comment>string</Comment>
      <Id>string</Id>
      <Status>string</Status>
      <SubmittedAt>timestamp</SubmittedAt>
   </ChangeInfo>
</DeleteKeySigningKeyResponse>
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in XML format by the service.

**[DeleteKeySigningKeyResponse](#API_DeleteKeySigningKey_ResponseSyntax)**

Root level tag for the DeleteKeySigningKeyResponse parameters.

Required: Yes

**[ChangeInfo](#API_DeleteKeySigningKey_ResponseSyntax)**

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

**InvalidSigningStatus**

Your hosted zone status isn't valid for this operation. In the hosted zone, change the
status to enable `DNSSEC` or disable `DNSSEC`.

HTTP Status Code: 400

**NoSuchKeySigningKey**

The specified key-signing key (KSK) doesn't exist.

HTTP Status Code: 404

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/route53-2013-04-01/DeleteKeySigningKey)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/route53-2013-04-01/DeleteKeySigningKey)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/route53-2013-04-01/DeleteKeySigningKey)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/route53-2013-04-01/DeleteKeySigningKey)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/route53-2013-04-01/DeleteKeySigningKey)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/route53-2013-04-01/DeleteKeySigningKey)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/route53-2013-04-01/DeleteKeySigningKey)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/route53-2013-04-01/DeleteKeySigningKey)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/route53-2013-04-01/DeleteKeySigningKey)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/route53-2013-04-01/DeleteKeySigningKey)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DeleteHostedZone

DeleteQueryLoggingConfig
