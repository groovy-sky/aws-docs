# DeleteGroup

Deletes a group so that all users and sub groups that belong to the group can no
longer access documents only available to that group. For example, after deleting the
group "Summer Interns", all interns who belonged to that group no longer see intern-only
documents in their chat results.

If you want to delete, update, or replace users or sub groups of a group, you need to
use the `PutGroup` operation. For example, if a user in the group
"Engineering" leaves the engineering team and another user takes their place, you
provide an updated list of users or sub groups that belong to the "Engineering" group
when calling `PutGroup`.

## Request Syntax

```nohighlight

DELETE /applications/applicationId/indices/indexId/groups/groupName?dataSourceId=dataSourceId HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[applicationId](#API_DeleteGroup_RequestSyntax)**

The identifier of the application in which the group mapping belongs.

Length Constraints: Fixed length of 36.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9-]{35}`

Required: Yes

**[dataSourceId](#API_DeleteGroup_RequestSyntax)**

The identifier of the data source linked to the group

A group can be tied to multiple data sources. You can delete a group from accessing
documents in a certain data source. For example, the groups "Research", "Engineering",
and "Sales and Marketing" are all tied to the company's documents stored in the data
sources Confluence and Salesforce. You want to delete "Research" and "Engineering"
groups from Salesforce, so that these groups cannot access customer-related documents
stored in Salesforce. Only "Sales and Marketing" should access documents in the
Salesforce data source.

Length Constraints: Fixed length of 36.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9-]{35}`

**[groupName](#API_DeleteGroup_RequestSyntax)**

The name of the group you want to delete.

Length Constraints: Minimum length of 1. Maximum length of 1024.

Pattern: `\P{C}*`

Required: Yes

**[indexId](#API_DeleteGroup_RequestSyntax)**

The identifier of the index you want to delete the group from.

Length Constraints: Fixed length of 36.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9-]{35}`

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```

HTTP/1.1 200

```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**AccessDeniedException**

You don't have access to perform this action. Make sure you have the required
permission policies and user accounts and try again.

HTTP Status Code: 403

**ConflictException**

You are trying to perform an action that conflicts with the current status of your
resource. Fix any inconsistencies with your resources and try again.

**message**

The message describing a `ConflictException`.

**resourceId**

The identifier of the resource affected.

**resourceType**

The type of the resource affected.

HTTP Status Code: 409

**InternalServerException**

An issue occurred with the internal server used for your Amazon Q Business service. Wait
some minutes and try again, or contact [Support](http://aws.amazon.com/contact-us) for help.

HTTP Status Code: 500

**ResourceNotFoundException**

The application or plugin resource you want to use doesn’t exist. Make sure you have
provided the correct resource and try again.

**message**

The message describing a `ResourceNotFoundException`.

**resourceId**

The identifier of the resource affected.

**resourceType**

The type of the resource affected.

HTTP Status Code: 404

**ThrottlingException**

The request was denied due to throttling. Reduce the number of requests and try
again.

HTTP Status Code: 429

**ValidationException**

The input doesn't meet the constraints set by the Amazon Q Business service. Provide the
correct input and try again.

**fields**

The input field(s) that failed validation.

**message**

The message describing the `ValidationException`.

**reason**

The reason for the `ValidationException`.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/qbusiness-2023-11-27/deletegroup.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/qbusiness-2023-11-27/deletegroup.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qbusiness-2023-11-27/deletegroup.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/qbusiness-2023-11-27/deletegroup.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qbusiness-2023-11-27/deletegroup.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/qbusiness-2023-11-27/deletegroup.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/qbusiness-2023-11-27/deletegroup.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/qbusiness-2023-11-27/deletegroup.md)

- [AWS SDK for Python](../../../goto/boto3/qbusiness-2023-11-27/deletegroup.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qbusiness-2023-11-27/deletegroup.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteDataSource

DeleteIndex

All content copied from https://docs.aws.amazon.com/.
