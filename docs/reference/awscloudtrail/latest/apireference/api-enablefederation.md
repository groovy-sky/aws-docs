# EnableFederation

###### Important

CloudTrail Lake will no longer be open to new customers starting May 31, 2026. If you would like to use CloudTrail Lake, sign up prior to that date. Existing customers can continue to use the service as normal. For more information, see [CloudTrail Lake availability change](../../../../services/awscloudtrail/latest/userguide/cloudtrail-lake-service-availability-change.md).

Enables Lake query federation on the specified event data store. Federating an event data store lets you view the metadata associated with the event data store in the AWS Glue
[Data Catalog](../../../../services/glue/latest/dg/components-overview.md#data-catalog-intro) and run
SQL queries against your event data using Amazon Athena. The table metadata stored in the AWS Glue Data Catalog
lets the Athena query engine know how to find, read, and process the data that you want to query.

When you enable Lake query federation, CloudTrail
creates a managed database named `aws:cloudtrail` (if the database doesn't already exist) and a managed federated table in
the AWS Glue Data Catalog. The event data store ID is used for the table name. CloudTrail registers the role ARN and event data store in
[AWS Lake Formation](../../../../services/awscloudtrail/latest/userguide/query-federation-lake-formation.md), the service responsible for allowing fine-grained access control
of the federated resources in the AWS Glue Data Catalog.

For more information about Lake query federation, see [Federate an event data store](../../../../services/awscloudtrail/latest/userguide/query-federation.md).

## Request Syntax

```nohighlight

{
   "EventDataStore": "string",
   "FederationRoleArn": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[EventDataStore](#API_EnableFederation_RequestSyntax)**

The ARN (or ID suffix of the ARN) of the event data store for which you want to enable Lake query federation.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 256.

Pattern: `^[a-zA-Z0-9._/\-:]+$`

Required: Yes

**[FederationRoleArn](#API_EnableFederation_RequestSyntax)**

The ARN of the federation role to use for the event data store. AWS services like AWS Lake Formation use this federation role to access data for the federated event
data store. The federation role must exist in your account and provide the [required minimum permissions](../../../../services/awscloudtrail/latest/userguide/query-federation.md#query-federation-permissions-role).

Type: String

Length Constraints: Minimum length of 3. Maximum length of 125.

Pattern: `^[a-zA-Z0-9._/\-:@=\+,\.]+$`

Required: Yes

## Response Syntax

```nohighlight

{
   "EventDataStoreArn": "string",
   "FederationRoleArn": "string",
   "FederationStatus": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[EventDataStoreArn](#API_EnableFederation_ResponseSyntax)**

The ARN of the event data store for which you enabled Lake query federation.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 256.

Pattern: `^[a-zA-Z0-9._/\-:]+$`

**[FederationRoleArn](#API_EnableFederation_ResponseSyntax)**

The ARN of the federation role.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 125.

Pattern: `^[a-zA-Z0-9._/\-:@=\+,\.]+$`

**[FederationStatus](#API_EnableFederation_ResponseSyntax)**

The federation status.

Type: String

Valid Values: `ENABLING | ENABLED | DISABLING | DISABLED`

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**AccessDeniedException**

You do not have sufficient access to perform this action.

HTTP Status Code: 400

**CloudTrailAccessNotEnabledException**

This exception is thrown when trusted access has not been enabled between AWS CloudTrail and AWS Organizations. For more information, see [How to enable or disable trusted access](../../../../services/organizations/latest/userguide/orgs-integrate-services.md#orgs_how-to-enable-disable-trusted-access) in the _AWS Organizations User Guide_ and [Prepare For Creating a Trail For Your Organization](../../../../services/awscloudtrail/latest/userguide/creating-an-organizational-trail-prepare.md) in the _AWS CloudTrail User Guide_.

HTTP Status Code: 400

**ConcurrentModificationException**

You are trying to update a resource when another request is in progress. Allow sufficient wait time for the previous request to complete, then retry your request.

HTTP Status Code: 400

**EventDataStoreARNInvalidException**

The specified event data store ARN is not valid or does not map to an event data store
in your account.

HTTP Status Code: 400

**EventDataStoreFederationEnabledException**

You cannot delete the event data store because Lake query federation is enabled. To delete the event data store, run the `DisableFederation` operation to
disable Lake query federation on the event data store.

HTTP Status Code: 400

**EventDataStoreNotFoundException**

The specified event data store was not found.

HTTP Status Code: 400

**InactiveEventDataStoreException**

The event data store is inactive.

HTTP Status Code: 400

**InsufficientDependencyServiceAccessPermissionException**

This exception is thrown when the IAM identity that is used to create
the organization resource lacks one or more required permissions for creating an
organization resource in a required service.

HTTP Status Code: 400

**InvalidParameterException**

The request includes a parameter that is not valid.

HTTP Status Code: 400

**NoManagementAccountSLRExistsException**

This exception is thrown when the management account does not have a service-linked
role.

HTTP Status Code: 400

**NotOrganizationMasterAccountException**

This exception is thrown when the AWS account making the request to
create or update an organization trail or event data store is not the management account
for an organization in AWS Organizations. For more information, see [Prepare For Creating a Trail For Your Organization](../../../../services/awscloudtrail/latest/userguide/creating-an-organizational-trail-prepare.md) or [Organization event data stores](../../../../services/awscloudtrail/latest/userguide/cloudtrail-lake-organizations.md).

HTTP Status Code: 400

**OperationNotPermittedException**

This exception is thrown when the requested operation is not permitted.

HTTP Status Code: 400

**OrganizationNotInAllFeaturesModeException**

This exception is thrown when AWS Organizations is not configured to support all
features. All features must be enabled in Organizations to support creating an
organization trail or event data store.

HTTP Status Code: 400

**OrganizationsNotInUseException**

This exception is thrown when the request is made from an AWS account
that is not a member of an organization. To make this request, sign in using the
credentials of an account that belongs to an organization.

HTTP Status Code: 400

**UnsupportedOperationException**

This exception is thrown when the requested operation is not supported.

HTTP Status Code: 400

## Examples

### Example

The following example shows how to enable CloudTrail Lake federation on an event data store.

```

{
   "EventDataStore": "arn:aws:cloudtrail:us-east-2:123456789012:eventdatastore/EXAMPLE-f852-4e8f-8bd1-bcf6cEXAMPLE",
   "FederationRoleArn": "arn:aws:iam::123456789012:role/FederationRole"
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/cloudtrail-2013-11-01/enablefederation.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/cloudtrail-2013-11-01/enablefederation.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/cloudtrail-2013-11-01/enablefederation.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/cloudtrail-2013-11-01/enablefederation.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/cloudtrail-2013-11-01/enablefederation.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/cloudtrail-2013-11-01/enablefederation.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/cloudtrail-2013-11-01/enablefederation.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/cloudtrail-2013-11-01/enablefederation.md)

- [AWS SDK for Python](../../../../services/goto/boto3/cloudtrail-2013-11-01/enablefederation.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/cloudtrail-2013-11-01/enablefederation.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DisableFederation

GenerateQuery

All content copied from https://docs.aws.amazon.com/.
