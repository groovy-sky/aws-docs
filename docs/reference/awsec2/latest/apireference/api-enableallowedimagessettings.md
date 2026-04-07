# EnableAllowedImagesSettings

Enables Allowed AMIs for your account in the specified AWS Region. Two values are
accepted:

- `enabled`: The image criteria in your Allowed AMIs settings are applied. As
a result, only AMIs matching these criteria are discoverable and can be used by your
account to launch instances.

- `audit-mode`: The image criteria in your Allowed AMIs settings are not
applied. No restrictions are placed on AMI discoverability or usage. Users in your account
can launch instances using any public AMI or AMI shared with your account.

The purpose of `audit-mode` is to indicate which AMIs will be affected when
Allowed AMIs is `enabled`. In `audit-mode`, each AMI displays either
`"ImageAllowed": true` or `"ImageAllowed": false` to indicate
whether the AMI will be discoverable and available to users in the account when Allowed
AMIs is enabled.

###### Note

The Allowed AMIs feature does not restrict the AMIs owned by your account. Regardless of
the criteria you set, the AMIs created by your account will always be discoverable and
usable by users in your account.

For more information, see [Control the discovery and use of AMIs in\
Amazon EC2 with Allowed AMIs](../../../../services/ec2/latest/userguide/ec2-allowed-amis.md) in
_Amazon EC2 User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**AllowedImagesSettingsState**

Specify `enabled` to apply the image criteria specified by the Allowed AMIs
settings. Specify `audit-mode` so that you can check which AMIs will be allowed or
not allowed by the image criteria.

Type: String

Valid Values: `enabled | audit-mode`

Required: Yes

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is
`DryRunOperation`. Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

## Response Elements

The following elements are returned by the service.

**allowedImagesSettingsState**

Returns `enabled` or `audit-mode` if the request succeeds;
otherwise, it returns an error.

Type: String

Valid Values: `enabled | audit-mode`

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/EnableAllowedImagesSettings)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/EnableAllowedImagesSettings)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/enableallowedimagessettings.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/enableallowedimagessettings.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/enableallowedimagessettings.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/enableallowedimagessettings.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/enableallowedimagessettings.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/enableallowedimagessettings.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/EnableAllowedImagesSettings)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/enableallowedimagessettings.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

EnableAddressTransfer

EnableAwsNetworkPerformanceMetricSubscription
