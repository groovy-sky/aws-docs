# QueryError

A two-part error structure that can occur in `ListGroupResources` or
`SearchResources`.

## Contents

###### Note

In the following list, the required parameters are described first.

**ErrorCode**

Specifies the error code that was raised.

Type: String

Valid Values: `CLOUDFORMATION_STACK_INACTIVE | CLOUDFORMATION_STACK_NOT_EXISTING | CLOUDFORMATION_STACK_UNASSUMABLE_ROLE | RESOURCE_TYPE_NOT_SUPPORTED`

Required: No

**Message**

A message that explains the `ErrorCode`.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/resource-groups-2017-11-27/queryerror.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/resource-groups-2017-11-27/queryerror.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/resource-groups-2017-11-27/queryerror.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PendingResource

ResourceFilter

All content copied from https://docs.aws.amazon.com/.
