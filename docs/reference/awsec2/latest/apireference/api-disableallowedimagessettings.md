# DisableAllowedImagesSettings

Disables Allowed AMIs for your account in the specified AWS Region. When set to
`disabled`, the image criteria in your Allowed AMIs settings do not apply, and no
restrictions are placed on AMI discoverability or usage. Users in your account can launch
instances using any public AMI or AMI shared with your account.

###### Note

The Allowed AMIs feature does not restrict the AMIs owned by your account. Regardless of
the criteria you set, the AMIs created by your account will always be discoverable and
usable by users in your account.

For more information, see [Control the discovery and use of AMIs in\
Amazon EC2 with Allowed AMIs](../../../../services/ec2/latest/userguide/ec2-allowed-amis.md) in
_Amazon EC2 User Guide_.

## Request Parameters

For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is
`DryRunOperation`. Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

## Response Elements

The following elements are returned by the service.

**allowedImagesSettingsState**

Returns `disabled` if the request succeeds; otherwise, it returns an
error.

Type: String

Valid Values: `disabled`

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/DisableAllowedImagesSettings)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/DisableAllowedImagesSettings)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/disableallowedimagessettings.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/disableallowedimagessettings.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/disableallowedimagessettings.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/disableallowedimagessettings.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/disableallowedimagessettings.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/disableallowedimagessettings.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/DisableAllowedImagesSettings)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/disableallowedimagessettings.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DisableAddressTransfer

DisableAwsNetworkPerformanceMetricSubscription
