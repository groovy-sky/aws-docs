# PutAccountConfiguration

Adds or modifies account-level configurations in ACM.

The supported configuration option is `DaysBeforeExpiry`. This option
specifies the number of days prior to certificate expiration when ACM starts
generating `EventBridge` events. ACM sends one event per day per
certificate until the certificate expires. By default, accounts receive events starting
45 days before certificate expiration.

## Request Syntax

```nohighlight

{
   "ExpiryEvents": {
      "DaysBeforeExpiry": number
   },
   "IdempotencyToken": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

###### Note

In the following list, the required parameters are described first.

**[IdempotencyToken](#API_PutAccountConfiguration_RequestSyntax)**

Customer-chosen string used to distinguish between calls to
`PutAccountConfiguration`. Idempotency tokens time out after one hour. If
you call `PutAccountConfiguration` multiple times with the same unexpired
idempotency token, ACM treats it as the same request and returns the original result.
If you change the idempotency token for each call, ACM treats each call as a new
request.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 32.

Pattern: `\w+`

Required: Yes

**[ExpiryEvents](#API_PutAccountConfiguration_RequestSyntax)**

Specifies expiration events associated with an account.

Type: [ExpiryEventsConfiguration](api-expiryeventsconfiguration.md) object

Required: No

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**AccessDeniedException**

You do not have access required to perform this action.

HTTP Status Code: 400

**ConflictException**

You are trying to update a resource or configuration that is already being created or
updated. Wait for the previous operation to finish and try again.

HTTP Status Code: 400

**ThrottlingException**

The request was denied because it exceeded a quota.

**throttlingReasons**

One or more reasons why the request was throttled.

HTTP Status Code: 400

**ValidationException**

The supplied input failed to satisfy constraints of an AWS service.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/acm-2015-12-08/putaccountconfiguration.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/acm-2015-12-08/putaccountconfiguration.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/acm-2015-12-08/putaccountconfiguration.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/acm-2015-12-08/putaccountconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/acm-2015-12-08/putaccountconfiguration.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/acm-2015-12-08/putaccountconfiguration.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/acm-2015-12-08/putaccountconfiguration.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/acm-2015-12-08/putaccountconfiguration.md)

- [AWS SDK for Python](../../../../services/goto/boto3/acm-2015-12-08/putaccountconfiguration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/acm-2015-12-08/putaccountconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ListTagsForCertificate

RemoveTagsFromCertificate

All content copied from https://docs.aws.amazon.com/.
