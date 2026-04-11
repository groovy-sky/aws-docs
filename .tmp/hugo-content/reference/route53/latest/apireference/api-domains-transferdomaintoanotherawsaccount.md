# TransferDomainToAnotherAwsAccount

Transfers a domain from the current AWS account to another AWS account. Note the following:

- The AWS account that you're transferring the domain to must
accept the transfer. If the other account doesn't accept the transfer within 3
days, we cancel the transfer. See [AcceptDomainTransferFromAnotherAwsAccount](api-domains-acceptdomaintransferfromanotherawsaccount.md).

- You can cancel the transfer before the other account accepts it. See [CancelDomainTransferToAnotherAwsAccount](api-domains-canceldomaintransfertoanotherawsaccount.md).

- The other account can reject the transfer. See [RejectDomainTransferFromAnotherAwsAccount](api-domains-rejectdomaintransferfromanotherawsaccount.md).

###### Important

When you transfer a domain from one AWS account to another, Route
53 doesn't transfer the hosted zone that is associated with the domain. DNS
resolution isn't affected if the domain and the hosted zone are owned by separate
accounts, so transferring the hosted zone is optional. For information about
transferring the hosted zone to another AWS account, see [Migrating a\
Hosted Zone to a Different AWS Account](../../../../services/route53/latest/developerguide/hosted-zones-migrating.md) in the
_Amazon Route 53 Developer Guide_.

Use either [ListOperations](api-domains-listoperations.md) or [GetOperationDetail](api-domains-getoperationdetail.md) to determine whether the operation succeeded. [GetOperationDetail](api-domains-getoperationdetail.md) provides additional information, for example,
`Domain Transfer from Aws Account 111122223333 has been cancelled`.

## Request Syntax

```nohighlight

{
   "AccountId": "string",
   "DomainName": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[AccountId](#API_domains_TransferDomainToAnotherAwsAccount_RequestSyntax)**

The account ID of the AWS account that you want to transfer the domain
to, for example, `111122223333`.

Type: String

Length Constraints: Fixed length of 12.

Pattern: `^(\d{12})$`

Required: Yes

**[DomainName](#API_domains_TransferDomainToAnotherAwsAccount_RequestSyntax)**

The name of the domain that you want to transfer from the current AWS account to another account.

Type: String

Length Constraints: Maximum length of 255.

Required: Yes

## Response Syntax

```nohighlight

{
   "OperationId": "string",
   "Password": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[OperationId](#API_domains_TransferDomainToAnotherAwsAccount_ResponseSyntax)**

Identifier for tracking the progress of the request. To query the operation status,
use [GetOperationDetail](api-domains-getoperationdetail.md).

Type: String

Length Constraints: Maximum length of 255.

**[Password](#API_domains_TransferDomainToAnotherAwsAccount_ResponseSyntax)**

To finish transferring a domain to another AWS account, the account
that the domain is being transferred to must submit an [AcceptDomainTransferFromAnotherAwsAccount](api-domains-acceptdomaintransferfromanotherawsaccount.md) request. The request must include
the value of the `Password` element that was returned in the
`TransferDomainToAnotherAwsAccount` response.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**DuplicateRequest**

The request is already in progress for the domain.

**message**

The request is already in progress for the domain.

**requestId**

ID of the request operation.

HTTP Status Code: 400

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

**OperationLimitExceeded**

The number of operations or jobs running exceeded the allowed threshold for the
account.

**message**

The number of operations or jobs running exceeded the allowed threshold for the
account.

HTTP Status Code: 400

**UnsupportedTLD**

Amazon Route 53 does not support this top-level domain (TLD).

**message**

Amazon Route 53 does not support this top-level domain (TLD).

HTTP Status Code: 400

## Examples

### TransferDomainToAnotherAwsAccount Example

This example illustrates one usage of TransferDomainToAnotherAwsAccount.

#### Sample Request

```

POST / HTTP/1.1
host:route53domains.us-east-1.amazonaws.com
x-amz-date:20200311T205230Z
authorization:AWS4-HMAC-SHA256
              Credential=AKIAIOSFODNN7EXAMPLE/20140711/us-east-1/route53domains/aws4_request,
              SignedHeaders=content-length;content-type;host;user-agent;x-amz-date;x-amz-target,
              Signature=[calculated-signature]
x-amz-target:Route53Domains_v20140515.TransferDomainToAnotherAwsAccount
user-agent:aws-sdk-java/1.8.3 Linux/2.6.18-164.el5PAE Java_HotSpot (TM )_Server_VM/24.60-b09/1.7.0_60
content-type:application/x-amz-json-1.1
content-length:[number of characters in the JSON string]
{
    "DomainName":"example.com",
    "AccountId":"111122223333"
}
```

#### Sample Response

```

HTTP/1.1 200
Content-Length:[number of characters in the JSON string]
{
    "OperationId":"308c56712-faa4-40fe-94c8-b4230example",
    "Password":"CeT2Zxs~Example"
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/route53domains-2014-05-15/transferdomaintoanotherawsaccount.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/route53domains-2014-05-15/transferdomaintoanotherawsaccount.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/route53domains-2014-05-15/transferdomaintoanotherawsaccount.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/route53domains-2014-05-15/transferdomaintoanotherawsaccount.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/route53domains-2014-05-15/transferdomaintoanotherawsaccount.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/route53domains-2014-05-15/transferdomaintoanotherawsaccount.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/route53domains-2014-05-15/transferdomaintoanotherawsaccount.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/route53domains-2014-05-15/transferdomaintoanotherawsaccount.md)

- [AWS SDK for Python](../../../../services/goto/boto3/route53domains-2014-05-15/transferdomaintoanotherawsaccount.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/route53domains-2014-05-15/transferdomaintoanotherawsaccount.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

TransferDomain

UpdateDomainContact
