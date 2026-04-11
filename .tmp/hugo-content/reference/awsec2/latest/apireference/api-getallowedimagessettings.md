# GetAllowedImagesSettings

Gets the current state of the Allowed AMIs setting and the list of Allowed AMIs criteria
at the account level in the specified Region.

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

**imageCriterionSet**

The list of criteria for images that are discoverable and usable in the account in the
specified AWS Region.

Type: Array of [ImageCriterion](api-imagecriterion.md) objects

**managedBy**

The entity that manages the Allowed AMIs settings. Possible values include:

- `account` \- The Allowed AMIs settings is managed by the account.

- `declarative-policy` \- The Allowed AMIs settings is managed by a
declarative policy and can't be modified by the account.

Type: String

Valid Values: `account | declarative-policy`

**requestId**

The ID of the request.

Type: String

**state**

The current state of the Allowed AMIs setting at the account level in the specified AWS
Region.

Possible values:

- `disabled`: All AMIs are allowed.

- `audit-mode`: All AMIs are allowed, but the `ImageAllowed` field
is set to `true` if the AMI would be allowed with the current list of criteria
if allowed AMIs was enabled.

- `enabled`: Only AMIs matching the image criteria are discoverable and
available for use.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/getallowedimagessettings.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/getallowedimagessettings.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/getallowedimagessettings.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/getallowedimagessettings.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/getallowedimagessettings.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/getallowedimagessettings.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/getallowedimagessettings.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/getallowedimagessettings.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/getallowedimagessettings.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/getallowedimagessettings.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

GetActiveVpnTunnelStatus

GetAssociatedEnclaveCertificateIamRoles
