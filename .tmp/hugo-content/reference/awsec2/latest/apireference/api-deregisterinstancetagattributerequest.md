# DeregisterInstanceTagAttributeRequest

Information about the tag keys to deregister for the current Region. You can either
specify individual tag keys or deregister all tag keys in the current Region. You must
specify either `IncludeAllTagsOfInstance` or `InstanceTagKeys` in the
request

## Contents

**IncludeAllTagsOfInstance**

Indicates whether to deregister all tag keys in the current Region. Specify
`false` to deregister all tag keys.

Type: Boolean

Required: No

**InstanceTagKey.N**

Information about the tag keys to deregister.

Type: Array of strings

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/deregisterinstancetagattributerequest.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/deregisterinstancetagattributerequest.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/deregisterinstancetagattributerequest.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DeprecationTimeConditionRequest

DescribeFastLaunchImagesSuccessItem
