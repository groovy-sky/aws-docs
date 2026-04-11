# PendingResource

A structure that identifies a resource that is currently pending addition to the group
as a member. Adding a resource to a resource group happens asynchronously as a
background task and this one isn't completed yet.

## Contents

###### Note

In the following list, the required parameters are described first.

**ResourceArn**

The Amazon resource name (ARN) of the resource that's in a pending state.

Type: String

Pattern: `arn:aws(-[a-z]+)*:[a-z0-9\-]*:([a-z]{2}(-[a-z]+)+-\d{1})?:([0-9]{12})?:.+`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/resource-groups-2017-11-27/pendingresource.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/resource-groups-2017-11-27/pendingresource.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/resource-groups-2017-11-27/pendingresource.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ListTagSyncTasksFilter

QueryError

All content copied from https://docs.aws.amazon.com/.
