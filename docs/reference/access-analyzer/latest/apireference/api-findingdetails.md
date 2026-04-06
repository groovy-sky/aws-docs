# FindingDetails

Contains information about an external access or unused access finding. Only one
parameter can be used in a `FindingDetails` object.

## Contents

###### Important

This data type is a UNION, so only one of the following members can be specified when used or returned.

**externalAccessDetails**

The details for an external access analyzer finding.

Type: [ExternalAccessDetails](api-externalaccessdetails.md) object

Required: No

**internalAccessDetails**

The details for an internal access analyzer finding. This contains information about
access patterns identified within your AWS organization or account.

Type: [InternalAccessDetails](api-internalaccessdetails.md) object

Required: No

**unusedIamRoleDetails**

The details for an unused access analyzer finding with an unused IAM role finding
type.

Type: [UnusedIamRoleDetails](api-unusediamroledetails.md) object

Required: No

**unusedIamUserAccessKeyDetails**

The details for an unused access analyzer finding with an unused IAM user access key
finding type.

Type: [UnusedIamUserAccessKeyDetails](api-unusediamuseraccesskeydetails.md) object

Required: No

**unusedIamUserPasswordDetails**

The details for an unused access analyzer finding with an unused IAM user password
finding type.

Type: [UnusedIamUserPasswordDetails](api-unusediamuserpassworddetails.md) object

Required: No

**unusedPermissionDetails**

The details for an unused access analyzer finding with an unused permission finding
type.

Type: [UnusedPermissionDetails](api-unusedpermissiondetails.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/accessanalyzer-2019-11-01/FindingDetails)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/accessanalyzer-2019-11-01/FindingDetails)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/accessanalyzer-2019-11-01/FindingDetails)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

FindingAggregationAccountDetails

FindingSource
