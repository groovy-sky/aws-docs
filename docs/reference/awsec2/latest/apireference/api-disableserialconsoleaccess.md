# DisableSerialConsoleAccess

Disables access to the EC2 serial console of all instances for your account. By default,
access to the EC2 serial console is disabled for your account. For more information, see
[Manage account access to the EC2 serial console](../../../../services/ec2/latest/userguide/configure-access-to-serial-console.md#serial-console-account-access) in the _Amazon EC2_
_User Guide_.

## Request Parameters

For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request, and provides an error response. If you have the required permissions, the error response is `DryRunOperation`. Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**serialConsoleAccessEnabled**

If `true`, access to the EC2 serial console of all instances is enabled for
your account. If `false`, access to the EC2 serial console of all instances
is disabled for your account.

Type: Boolean

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/disableserialconsoleaccess.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/disableserialconsoleaccess.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/disableserialconsoleaccess.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/disableserialconsoleaccess.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/disableserialconsoleaccess.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/disableserialconsoleaccess.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/disableserialconsoleaccess.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/disableserialconsoleaccess.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/disableserialconsoleaccess.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/disableserialconsoleaccess.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DisableRouteServerPropagation

DisableSnapshotBlockPublicAccess
