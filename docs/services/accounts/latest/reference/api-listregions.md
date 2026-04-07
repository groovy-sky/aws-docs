# ListRegions

Lists all the Regions for a given account and their respective opt-in statuses.
Optionally, this list can be filtered by the `region-opt-status-contains`
parameter.

## Request Syntax

```nohighlight

POST /listRegions HTTP/1.1
Content-type: application/json

{
   "AccountId": "string",
   "MaxResults": number,
   "NextToken": "string",
   "RegionOptStatusContains": [ "string" ]
}
```

## URI Request Parameters

The request does not use any URI parameters.

## Request Body

The request accepts the following data in JSON format.

**[AccountId](#API_ListRegions_RequestSyntax)**

Specifies the 12-digit account ID number of the AWS account that you want to access
or modify with this operation. If you don't specify this parameter, it defaults to the
Amazon Web Services account of the identity used to call the operation. To use this
parameter, the caller must be an identity in the [organization's\
management account](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_getting-started_concepts.html#account) or a delegated administrator account. The specified
account ID must be a member account in the same organization. The organization must have
[all features\
enabled](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_org_support-all-features.html), and the organization must have [trusted access](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_integrate_services.html) enabled for
the Account Management service, and optionally a [delegated\
admin](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_getting-started_concepts.html#delegated-admin) account assigned.

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

**[MaxResults](#API_ListRegions_RequestSyntax)**

The total number of items to return in the command’s output. If the total number of
items available is more than the value specified, a `NextToken` is provided
in the command’s output. To resume pagination, provide the `NextToken` value
in the `starting-token` argument of a subsequent command. Do not use the
`NextToken` response element directly outside of the AWS CLI. For usage
examples, see [Pagination](https://docs.aws.amazon.com/cli/latest/userguide/pagination.html) in the _AWS Command Line Interface User_
_Guide_.

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 50.

Required: No

**[NextToken](#API_ListRegions_RequestSyntax)**

A token used to specify where to start paginating. This is the `NextToken`
from a previously truncated response. For usage examples, see [Pagination](https://docs.aws.amazon.com/cli/latest/userguide/pagination.html) in the
_AWS Command Line Interface User Guide_.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1000.

Required: No

**[RegionOptStatusContains](#API_ListRegions_RequestSyntax)**

A list of Region statuses (Enabling, Enabled, Disabling, Disabled, Enabled\_by\_default)
to use to filter the list of Regions for a given account. For example, passing in a
value of ENABLING will only return a list of Regions with a Region status of
ENABLING.

Type: Array of strings

Valid Values: `ENABLED | ENABLING | DISABLING | DISABLED | ENABLED_BY_DEFAULT`

Required: No

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "NextToken": "string",
   "Regions": [
      {
         "RegionName": "string",
         "RegionOptStatus": "string"
      }
   ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[NextToken](#API_ListRegions_ResponseSyntax)**

If there is more data to be returned, this will be populated. It should be passed into
the `next-token` request parameter of `list-regions`.

Type: String

**[Regions](#API_ListRegions_ResponseSyntax)**

This is a list of Regions for a given account, or if the filtered parameter was used,
a list of Regions that match the filter criteria set in the `filter`
parameter.

Type: Array of [Region](https://docs.aws.amazon.com/accounts/latest/reference/API_Region.html) objects

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

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/account-2021-02-01/ListRegions)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/account-2021-02-01/ListRegions)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/account-2021-02-01/ListRegions)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/account-2021-02-01/ListRegions)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/account-2021-02-01/ListRegions)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/account-2021-02-01/ListRegions)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/account-2021-02-01/ListRegions)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/account-2021-02-01/ListRegions)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/account-2021-02-01/ListRegions)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/account-2021-02-01/ListRegions)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

GetRegionOptStatus

PutAccountName
