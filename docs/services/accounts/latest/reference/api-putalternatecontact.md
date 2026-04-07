# PutAlternateContact

Modifies the specified alternate contact attached to an AWS account.

For complete details about how to use the alternate contact operations, see [Update the alternate contacts for your AWS account](manage-acct-update-contact-alternate.md).

###### Note

Before you can update the alternate contact information for an
AWS account that is managed by AWS Organizations, you must first enable integration between AWS Account Management
and Organizations. For more information, see [Enable trusted access for \
AWS Account Management](using-orgs-trusted-access.md).

## Request Syntax

```nohighlight

POST /putAlternateContact HTTP/1.1
Content-type: application/json

{
   "AccountId": "string",
   "AlternateContactType": "string",
   "EmailAddress": "string",
   "Name": "string",
   "PhoneNumber": "string",
   "Title": "string"
}
```

## URI Request Parameters

The request does not use any URI parameters.

## Request Body

The request accepts the following data in JSON format.

**[AccountId](#API_PutAlternateContact_RequestSyntax)**

Specifies the 12 digit account ID number of the AWS account that
you want to access or modify with this operation.

If you do not specify this parameter, it defaults to the AWS account of the
identity used to call the operation.

To use this parameter, the caller must be an identity in the [organization's management account](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_getting-started_concepts.html#account) or a delegated administrator account, and
the specified account ID must be a member account in the same organization. The
organization must have [all features \
enabled](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_org_support-all-features.html), and the organization must have [trusted access](https://docs.aws.amazon.com/organizations/latest/userguide/services-that-can-integrate-account.html) enabled for the
Account Management service, and optionally a [delegated\
administrator](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_getting-started_concepts.html#delegated-admin) account
assigned.

###### Note

The management account can't specify its own `AccountId`; it must call
the operation in standalone context by not including the `AccountId`
parameter.

To call this operation on an account that is not a member of an organization, then
don't specify this parameter, and call the operation using an identity belonging to
the account whose contacts you wish to retrieve or modify.

Type: String

Pattern: `\d{12}`

Required: No

**[AlternateContactType](#API_PutAlternateContact_RequestSyntax)**

Specifies which alternate contact you want to create or update.

Type: String

Valid Values: `BILLING | OPERATIONS | SECURITY`

Required: Yes

**[EmailAddress](#API_PutAlternateContact_RequestSyntax)**

Specifies an email address for the alternate contact.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 254.

Pattern: `[\s]*[\w+=.#|!&-]+@[\w.-]+\.[\w]+[\s]*`

Required: Yes

**[Name](#API_PutAlternateContact_RequestSyntax)**

Specifies a name for the alternate contact.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

Required: Yes

**[PhoneNumber](#API_PutAlternateContact_RequestSyntax)**

Specifies a phone number for the alternate contact.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 25.

Pattern: `[\s0-9()+-]+`

Required: Yes

**[Title](#API_PutAlternateContact_RequestSyntax)**

Specifies a title for the alternate contact.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 50.

Required: Yes

## Response Syntax

```

HTTP/1.1 200

```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](https://docs.aws.amazon.com/accounts/latest/reference/CommonErrors.html).

**AccessDeniedException**

The operation failed because the calling identity doesn't have the minimum required
permissions.

**errorType**

The value populated to the `x-amzn-ErrorType` response header by API Gateway.

HTTP Status Code: 403

**InternalServerException**

The operation failed because of an error internal to AWS. Try your operation again
later.

**errorType**

The value populated to the `x-amzn-ErrorType` response header by API Gateway.

HTTP Status Code: 500

**TooManyRequestsException**

The operation failed because it was called too frequently and exceeded a throttle
limit.

**errorType**

The value populated to the `x-amzn-ErrorType` response header by API Gateway.

HTTP Status Code: 429

**ValidationException**

The operation failed because one of the input parameters was invalid.

**fieldList**

The field where the invalid entry was detected.

**message**

The message that informs you about what was invalid about the request.

**reason**

The reason that validation failed.

HTTP Status Code: 400

## Examples

### Example 1

The following example sets the billing alternate contact for the account whose
credentials are used to call the operation.

#### Sample Request

```

POST / HTTP/1.1
X-Amz-Target: AWSAccountV20210201.PutAlternateContact

{
    "AlternateContactType": "Billing",
    "Name": "Carlos Salazar",
    "Title": "CFO",
    "EmailAddress": "carlos@example.com",
    "PhoneNumber": "206-555-0199"
}
```

#### Sample Response

```

HTTP/1.1 200 OK
Content-Type: application/json
```

### Example 2

The following example sets or overwrites the billing alternate contact for the
specified member account in an organization. You must use credentials from
either the organization's management account or from the Account Management service's
delegated admin account.

#### Sample Request

```

POST / HTTP/1.1
X-Amz-Target: AWSAccountV20210201.PutAlternateContact

{
    "AccountId": "123456789012",
    "AlternateContactType": "Billing",
    "Name": "Carlos Salazar",
    "Title": "CFO",
    "EmailAddress": "carlos@example.com",
    "PhoneNumber": "206-555-0199"
}
```

#### Sample Response

```

HTTP/1.1 200 OK
Content-Type: application/json
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/account-2021-02-01/PutAlternateContact)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/account-2021-02-01/PutAlternateContact)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/account-2021-02-01/PutAlternateContact)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/account-2021-02-01/PutAlternateContact)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/account-2021-02-01/PutAlternateContact)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/account-2021-02-01/PutAlternateContact)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/account-2021-02-01/PutAlternateContact)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/account-2021-02-01/PutAlternateContact)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/account-2021-02-01/PutAlternateContact)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/account-2021-02-01/PutAlternateContact)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

PutAccountName

PutContactInformation
