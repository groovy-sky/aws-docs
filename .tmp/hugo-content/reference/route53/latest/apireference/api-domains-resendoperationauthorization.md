# ResendOperationAuthorization

Resend the form of authorization email for this operation.

## Request Syntax

```nohighlight

{
   "OperationId": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[OperationId](#API_domains_ResendOperationAuthorization_RequestSyntax)**

Operation ID.

Type: String

Length Constraints: Maximum length of 255.

Required: Yes

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidInput**

The requested item is not acceptable. For example, for APIs that accept a domain name,
the request might specify a domain name that doesn't belong to the account that
submitted the request. For `AcceptDomainTransferFromAnotherAwsAccount`, the
password might be invalid.

**message**

The requested item is not acceptable. For example, for an OperationId it might refer
to the ID of an operation that is already completed. For a domain name, it might not be
a valid domain name or belong to the requester account.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/route53domains-2014-05-15/resendoperationauthorization.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/route53domains-2014-05-15/resendoperationauthorization.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/route53domains-2014-05-15/resendoperationauthorization.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/route53domains-2014-05-15/resendoperationauthorization.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/route53domains-2014-05-15/resendoperationauthorization.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/route53domains-2014-05-15/resendoperationauthorization.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/route53domains-2014-05-15/resendoperationauthorization.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/route53domains-2014-05-15/resendoperationauthorization.md)

- [AWS SDK for Python](../../../../services/goto/boto3/route53domains-2014-05-15/resendoperationauthorization.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/route53domains-2014-05-15/resendoperationauthorization.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ResendContactReachabilityEmail

RetrieveDomainAuthCode
