# FailedResource

A resource that failed to be added to or removed from a group.

## Contents

###### Note

In the following list, the required parameters are described first.

**ErrorCode**

The error code associated with the failure.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Required: No

**ErrorMessage**

The error message text associated with the failure.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: No

**ResourceArn**

The Amazon resource name (ARN) of the resource that failed to be added or removed.

Type: String

Pattern: `arn:aws(-[a-z]+)*:[a-z0-9\-]*:([a-z]{2}(-[a-z]+)+-\d{1})?:([0-9]{12})?:.+`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/resource-groups-2017-11-27/failedresource.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/resource-groups-2017-11-27/failedresource.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/resource-groups-2017-11-27/failedresource.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AccountSettings

Group

All content copied from https://docs.aws.amazon.com/.
