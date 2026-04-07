# ReplaceImageCriteriaInAllowedImagesSettings

Sets or replaces the criteria for Allowed AMIs.

###### Note

The Allowed AMIs feature does not restrict the AMIs owned by your account. Regardless of
the criteria you set, the AMIs created by your account will always be discoverable and
usable by users in your account.

For more information, see [Control the discovery and use of AMIs in\
Amazon EC2 with Allowed AMIs](../../../../services/ec2/latest/userguide/ec2-allowed-amis.md) in
_Amazon EC2 User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is
`DryRunOperation`. Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**ImageCriterion.N**

The list of criteria that are evaluated to determine whether AMIs are discoverable and
usable in the account in the specified AWS Region.

Type: Array of [ImageCriterionRequest](api-imagecriterionrequest.md) objects

Required: No

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**return**

Returns `true` if the request succeeds; otherwise, it returns an error.

Type: Boolean

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/ReplaceImageCriteriaInAllowedImagesSettings)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/ReplaceImageCriteriaInAllowedImagesSettings)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/ReplaceImageCriteriaInAllowedImagesSettings)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/ec2-2016-11-15/ReplaceImageCriteriaInAllowedImagesSettings)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/ReplaceImageCriteriaInAllowedImagesSettings)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/ec2-2016-11-15/ReplaceImageCriteriaInAllowedImagesSettings)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/ec2-2016-11-15/ReplaceImageCriteriaInAllowedImagesSettings)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/ec2-2016-11-15/ReplaceImageCriteriaInAllowedImagesSettings)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/ReplaceImageCriteriaInAllowedImagesSettings)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/ReplaceImageCriteriaInAllowedImagesSettings)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ReplaceIamInstanceProfileAssociation

ReplaceNetworkAclAssociation
