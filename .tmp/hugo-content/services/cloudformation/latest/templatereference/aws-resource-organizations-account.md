This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Organizations::Account

Creates an AWS account that is automatically a member of the
organization whose credentials made the request.

CloudFormation uses the [`CreateAccount`](../../../../reference/organizations/latest/apireference/api-createaccount.md) operation to create accounts. This is an
asynchronous request that AWS performs in the background. Because
`CreateAccount` operates asynchronously, it can return a successful
completion message even though account initialization might still be in progress. You
might need to wait a few minutes before you can successfully access the account. To
check the status of the request, do one of the following:

- Use the `Id` value of the `CreateAccountStatus` response
element from the `CreateAccount` operation to provide as a parameter
to the [`DescribeCreateAccountStatus`](../../../../reference/organizations/latest/apireference/api-describecreateaccountstatus.md) operation.

- Check the CloudTrail log for the `CreateAccountResult` event. For
information on using CloudTrail with AWS Organizations, see [Logging and monitoring in AWS Organizations](../../../organizations/latest/userguide/orgs-security-incident-response.md#orgs_cloudtrail-integration) in the _AWS Organizations User Guide_.

The user who calls the API to create an account must have the
`organizations:CreateAccount` permission. If you enabled all features in
the organization, AWS Organizations creates the required service-linked role named
`AWSServiceRoleForOrganizations`. For more information, see [AWS Organizations and service-linked roles](../../../organizations/latest/userguide/orgs-integrate-services.md#orgs_integrate_services-using_slrs) in the _AWS Organizations User Guide_.

If the request includes tags, then the requester must have the
`organizations:TagResource` permission.

AWS Organizations preconfigures the new member account with a role (named
`OrganizationAccountAccessRole` by default) that grants users in the
management account administrator permissions in the new member account. Principals in
the management account can assume the role. AWS Organizations clones the company
name and address information for the new account from the organization's management
account.

For more information about creating accounts, see [Creating a\
member account in your organization](../../../organizations/latest/userguide/orgs-manage-accounts-create.md) in the _AWS Organizations User Guide_.

This operation can be called only from the organization's management account.

**Deleting Account resources**

The default `DeletionPolicy` for resource
`AWS::Organizations::Account` is `Retain`. For more
information about how CloudFormation deletes resources, see [DeletionPolicy Attribute](../userguide/aws-attribute-deletionpolicy.md).

###### Important

- If you include multiple accounts in a single template, you must use the
`DependsOn` attribute on each account resource type so that
the accounts are created sequentially. If you create multiple accounts at
the same time, Organizations returns an error and the stack operation
fails.

- You can't modify the following list of `Account` resource
parameters using CloudFormation updates.

- AccountName

- Email

- RoleName

If you attempt to update the listed parameters, CloudFormation
will attempt the update, but you will receive an error message as those
updates are not supported from an Organizations management account or a
[registered delegated administrator](../userguide/stacksets-orgs-delegated-admin.md) account. Both the update and
the update roll-back will fail, so you must skip the account resource
update. To update parameters `AccountName` and
`Email`, you must sign in to the AWS Management Console as
the AWS account root user. For more information, see [Update\
the AWS account name, email address, or password for the\
root user](../../../accounts/latest/reference/manage-acct-update-root-user.md) in the _AWS Account Management Reference_
_Guide_.

- When you create an account in an organization using the AWS Organizations console, API, or AWS CLI commands, we don't
automatically collect the information required for the account to operate as
a standalone account. That includes collecting the payment method and
signing the end user license agreement (EULA). If you must remove an account
from your organization later, you can do so only after you provide the
missing information. For more information, see [Considerations before removing an account from an organization](../../../organizations/latest/userguide/orgs-manage-account-before-remove.md)
in the _AWS Organizations User Guide_.

- When you create an account in an organization using CloudFormation,
you can't specify a value for the `CreateAccount` operation
parameter `IamUserAccessToBilling`. The default value for
parameter `IamUserAccessToBilling` is `ALLOW`, and
IAM users and roles with the required permissions can
access billing information for the new account.

- If you get an exception that indicates `DescribeCreateAccountStatus
                              returns IN_PROGRESS state before time out`. You must check the
account creation status using the [`DescribeCreateAccountStatus`](../../../../reference/organizations/latest/apireference/api-describecreateaccountstatus.md) operation. If the
account state returns as `SUCCEEDED`, you can import the account
into CloudFormation management using [`resource import`](../userguide/resource-import.md).

- If you get an exception that indicates you have exceeded your account
quota for the organization, you can request an increase by using the [Service Quotas console](../../../organizations/latest/userguide/orgs-reference-limits.md).

- If you get an exception that indicates the operation failed because your
organization is still initializing, wait one hour and then try again. If the
error persists, contact [AWS Support](https://console.aws.amazon.com/support/home).

- We don't recommend that you use the `CreateAccount` operation
to create multiple temporary accounts. You can close accounts using the
[`CloseAccount`](../../../../reference/organizations/latest/apireference/api-closeaccount.md) operation or from the AWS Organizations console in the organization's management account. For
information on the requirements and process for closing an account, see
[Closing a member account in your organization](../../../organizations/latest/userguide/orgs-manage-accounts-close.md) in the _AWS Organizations User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Organizations::Account",
  "Properties" : {
      "AccountName" : String,
      "Email" : String,
      "ParentIds" : [ String, ... ],
      "RoleName" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Organizations::Account
Properties:
  AccountName: String
  Email: String
  ParentIds:
    - String
  RoleName: String
  Tags:
    - Tag

```

## Properties

`AccountName`

The account name given to the account when it was created.

_Required_: Yes

_Type_: String

_Pattern_: `[\u0020-\u007E]+`

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: Updates are not supported.

`Email`

The email address associated with the AWS account.

The [regex pattern](http://wikipedia.org/wiki/regex) for this parameter is a string of characters that represents a
standard internet email address.

_Required_: Yes

_Type_: String

_Pattern_: `[^\s@]+@[^\s@]+\.[^\s@]+`

_Minimum_: `6`

_Maximum_: `64`

_Update requires_: Updates are not supported.

`ParentIds`

The unique identifier (ID) of the root or organizational unit (OU) that you want to
create the new account in. If you don't specify this parameter, the
`ParentId` defaults to the root ID.

This parameter only accepts a string array with one string value.

The [regex pattern](http://wikipedia.org/wiki/regex) for a parent ID
string requires one of the following:

- **Root** \- A string that begins with "r-" followed
by from 4 to 32 lowercase letters or digits.

- **Organizational unit (OU)** \- A string that begins
with "ou-" followed by from 4 to 32 lowercase letters or digits (the ID of the
root that the OU is in). This string is followed by a second "-" dash and from 8
to 32 additional lowercase letters or digits.

_Required_: No

_Type_: Array of String

_Pattern_: `^(r-[0-9a-z]{4,32})|(ou-[0-9a-z]{4,32}-[a-z0-9]{8,32})$`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleName`

The name of an IAM role that AWS Organizations automatically preconfigures in the new member
account. This role trusts the management account, allowing users in the management
account to assume the role, as permitted by the management account administrator. The
role has administrator permissions in the new member account.

If you don't specify this parameter, the role name defaults to
`OrganizationAccountAccessRole`.

For more information about how to use this role to access the member account, see the
following links:

- [Creating the OrganizationAccountAccessRole in an invited member\
account](../../../organizations/latest/userguide/orgs-manage-accounts-access.md#orgs_manage_accounts_create-cross-account-role) in the _AWS Organizations User Guide_

- Steps 2 and 3 in [IAM Tutorial:\
Delegate access across AWS accounts using IAM roles](../../../iam/latest/userguide/tutorial-cross-account-with-roles.md) in the
_IAM User Guide_

The [regex pattern](http://wikipedia.org/wiki/regex) that
is used to validate this parameter. The pattern can include uppercase
letters, lowercase letters, digits with no spaces, and any of the following characters: =,.@-

_Required_: No

_Type_: String

_Pattern_: `[\w+=,.@-]{1,64}`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: Updates are not supported.

`Tags`

A list of tags that you want to attach to the newly created account. For each tag in
the list, you must specify both a tag key and a value. You can set the value to an empty
string, but you can't set it to `null`. For more information about tagging,
see [Tagging AWS Organizations\
resources](../../../organizations/latest/userguide/orgs-tagging.md) in the AWS Organizations User Guide.

###### Note

If any one of the tags is not valid or if you exceed the maximum allowed number of
tags for an account, then the entire request fails and the account is not
created.

_Required_: No

_Type_: Array of [Tag](aws-properties-organizations-account-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the `AccountId`. For example:
`123456789012`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AccountId`

Returns the unique identifier (ID) of the account. For example:
`123456789012`.

`Arn`

Returns the Amazon Resource Name (ARN) of the account. For example:
`arn:aws:organizations::111111111111:account/o-exampleorgid/555555555555`.

`JoinedMethod`

Returns the method by which the account joined the organization. For example:
`INVITED | CREATED`.

`JoinedTimestamp`

Returns the date the account became a part of the organization. For example:
`2016-11-24T11:11:48-08:00`.

`State`

Each state represents a specific phase in the account lifecycle. Use this information
to manage account access, automate workflows, or trigger actions based on account state
changes.

For more information about account states and their implications, see [Monitor the state of your AWS accounts](../../../organizations/latest/userguide/orgs-manage-accounts-account-state.md) in the
_AWS Organizations User Guide_.

`Status`

Returns the status of the account in the organization. For example: `ACTIVE |
                SUSPENDED | PENDING_CLOSURE`.

## See also

- [Creating a member account in your organization](../../../organizations/latest/userguide/orgs-manage-accounts-create.md) in the
_AWS Organizations User Guide_.

- [CreateAccount](../../../../reference/organizations/latest/apireference/api-createaccount.md) in the _AWS Organizations API_
_Reference Guide_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS Organizations

Tag

All content copied from https://docs.aws.amazon.com/.
