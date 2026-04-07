# DisableDomainAutoRenew

This operation disables automatic renewal of domain registration for the specified
domain.

## Request Syntax

```nohighlight

{
   "DomainName": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[DomainName](#API_domains_DisableDomainAutoRenew_RequestSyntax)**

The name of the domain that you want to disable automatic renewal for.

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

**UnsupportedTLD**

Amazon Route 53 does not support this top-level domain (TLD).

**message**

Amazon Route 53 does not support this top-level domain (TLD).

HTTP Status Code: 400

## Examples

### DisableDomainAutoRenew Example

This example illustrates one usage of DisableDomainAutoRenew.

#### Sample Request

```

POST / HTTP/1.1
host:route53domains.us-east-1.amazonaws.com
x-amz-date:20140711T205230Z
authorization:AWS4-HMAC-SHA256
              Credential=AKIAIOSFODNN7EXAMPLE/20140711/us-east-1/route53domains/aws4_request,
              SignedHeaders=content-length;content-type;host;user-agent;x-amz-date;x-amz-target,
              Signature=[calculated-signature]
x-amz-target:Route53Domains_v20140515.DisableDomainAutoRenew
user-agent:aws-sdk-java/1.8.3 Linux/2.6.18-164.el5PAE Java_HotSpot (TM )_Server_VM/24.60-b09/1.7.0_60
content-type:application/x-amz-json-1.1
content-length:[number of characters in the JSON string]
{
   "DomainName":"example.com"
}
```

#### Sample Response

```

HTTP/1.1 200
Content-Length:[number of characters in the JSON string]{}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/route53domains-2014-05-15/DisableDomainAutoRenew)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/route53domains-2014-05-15/DisableDomainAutoRenew)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/route53domains-2014-05-15/DisableDomainAutoRenew)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/route53domains-2014-05-15/DisableDomainAutoRenew)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/route53domains-2014-05-15/DisableDomainAutoRenew)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/route53domains-2014-05-15/DisableDomainAutoRenew)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/route53domains-2014-05-15/DisableDomainAutoRenew)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/route53domains-2014-05-15/DisableDomainAutoRenew)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/route53domains-2014-05-15/DisableDomainAutoRenew)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/route53domains-2014-05-15/DisableDomainAutoRenew)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DeleteTagsForDomain

DisableDomainTransferLock
