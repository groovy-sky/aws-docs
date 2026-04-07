# RegisterInstanceTagAttributeRequest

Information about the tag keys to register for the current Region. You can either
specify individual tag keys or register all tag keys in the current Region. You must
specify either `IncludeAllTagsOfInstance` or `InstanceTagKeys` in the
request

## Contents

**IncludeAllTagsOfInstance**

Indicates whether to register all tag keys in the current Region. Specify
`true` to register all tag keys.

Type: Boolean

Required: No

**InstanceTagKey.N**

The tag keys to register.

Type: Array of strings

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/registerinstancetagattributerequest.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/registerinstancetagattributerequest.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/registerinstancetagattributerequest.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

RegisteredInstance

RemoveIpamOperatingRegion
