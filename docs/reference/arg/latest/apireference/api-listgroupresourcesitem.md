---
title: "ListGroupResourcesItem"
---

# ListGroupResourcesItem

A structure returned by the [ListGroupResources](api-listgroupresources.md) operation that
contains identity and group membership status information for one of the resources in
the group.

## Contents

###### Note

In the following list, the required parameters are described first.

**Identifier**

A structure that contains the ARN of a resource and its resource type.

Type: [ResourceIdentifier](api-resourceidentifier.md) object

Required: No

**Status**

A structure that contains the status of this resource's membership in the
group.

###### Note

This field is present in the response only if the group is of type
`AWS::EC2::HostManagement`.

Type: [ResourceStatus](api-resourcestatus.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/resource-groups-2017-11-27/ListGroupResourcesItem)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/resource-groups-2017-11-27/ListGroupResourcesItem)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/resource-groups-2017-11-27/ListGroupResourcesItem)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ListGroupingStatusesFilter

ListTagSyncTasksFilter

All content copied from https://docs.aws.amazon.com/.
