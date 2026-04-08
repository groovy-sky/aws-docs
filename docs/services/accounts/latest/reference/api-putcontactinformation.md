# PutContactInformation

Updates the primary contact information of an AWS account.

For complete details about how to use the primary contact operations, see [Update the primary contact for your AWS account](manage-acct-update-contact-primary.md).

## Request Syntax

```nohighlight

POST /putContactInformation HTTP/1.1
Content-type: application/json

{
   "AccountId": "string",
   "ContactInformation": {
      "AddressLine1": "string",
      "AddressLine2": "string",
      "AddressLine3": "string",
      "City": "string",
      "CompanyName": "string",
      "CountryCode": "string",
      "DistrictOrCounty": "string",
      "FullName": "string",
      "PhoneNumber": "string",
      "PostalCode": "string",
      "StateOrRegion": "string",
      "WebsiteUrl": "string"
   }
}
```

## URI Request Parameters

The request does not use any URI parameters.

## Request Body

The request accepts the following data in JSON format.

**[AccountId](#API_PutContactInformation_RequestSyntax)**

Specifies the 12-digit account ID number of the AWS account that you want to access
or modify with this operation. If you don't specify this parameter, it defaults to the
Amazon Web Services account of the identity used to call the operation. To use this
parameter, the caller must be an identity in the [organization's\
management account](../../../organizations/latest/userguide/orgs-getting-started-concepts.md#account) or a delegated administrator account. The specified
account ID must be a member account in the same organization. The organization must have
[all features\
enabled](../../../organizations/latest/userguide/orgs-manage-org-support-all-features.md), and the organization must have [trusted access](../../../organizations/latest/userguide/services-that-can-integrate-account.md)
enabled for the Account Management service, and optionally a [delegated\
administrator](../../../organizations/latest/userguide/orgs-getting-started-concepts.md#delegated-admin) account assigned.

###### Note

The management account can't specify its own `AccountId`. It must call
the operation in standalone context by not including the `AccountId`
parameter.

To call this operation on an account that is not a member of an organization, don't
specify this parameter. Instead, call the operation using an identity belonging to the
account whose contacts you wish to retrieve or modify.

Type: String

Pattern: `\d{12}`

Required: No

**[ContactInformation](#API_PutContactInformation_RequestSyntax)**

Contains the details of the primary contact information associated with an
AWS account.

Type: [ContactInformation](api-contactinformation.md) object

Required: Yes

## Response Syntax

```

HTTP/1.1 200

```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

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

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/account-2021-02-01/putcontactinformation.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/account-2021-02-01/putcontactinformation.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/account-2021-02-01/putcontactinformation.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/account-2021-02-01/putcontactinformation.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/account-2021-02-01/putcontactinformation.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/account-2021-02-01/putcontactinformation.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/account-2021-02-01/putcontactinformation.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/account-2021-02-01/putcontactinformation.md)

- [AWS SDK for Python](../../../goto/boto3/account-2021-02-01/putcontactinformation.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/account-2021-02-01/putcontactinformation.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PutAlternateContact

StartPrimaryEmailUpdate

All content copied from https://docs.aws.amazon.com/.
