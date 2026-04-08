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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/route53-2013-04-01/deletekeysigningkey.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/route53-2013-04-01/deletekeysigningkey.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/route53-2013-04-01/deletekeysigningkey.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/route53-2013-04-01/deletekeysigningkey.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/route53-2013-04-01/deletekeysigningkey.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/route53-2013-04-01/deletekeysigningkey.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/route53-2013-04-01/deletekeysigningkey.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/route53-2013-04-01/deletekeysigningkey.md)

- [AWS SDK for Python](../../../../services/goto/boto3/route53-2013-04-01/deletekeysigningkey.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/route53-2013-04-01/deletekeysigningkey.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DeleteHostedZone

DeleteQueryLoggingConfig
