# ModifyInstanceMetadataDefaults

Modifies the default instance metadata service (IMDS) settings at the account level in
the specified AWS  Region.

###### Note

To remove a parameter's account-level default setting, specify
`no-preference`. If an account-level setting is cleared with
`no-preference`, then the instance launch considers the other
instance metadata settings. For more information, see [Order of precedence for instance metadata options](../../../../services/ec2/latest/userguide/configuring-instance-metadata-options.md#instance-metadata-options-order-of-precedence) in the
_Amazon EC2 User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the operation, without actually making the
request, and provides an error response. If you have the required permissions, the error response is
`DryRunOperation`. Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**HttpEndpoint**

Enables or disables the IMDS endpoint on an instance. When disabled, the instance
metadata can't be accessed.

Type: String

Valid Values: `disabled | enabled | no-preference`

Required: No

**HttpPutResponseHopLimit**

The maximum number of hops that the metadata token can travel. To indicate no
preference, specify `-1`.

Possible values: Integers from `1` to `64`, and `-1`
to indicate no preference

Type: Integer

Required: No

**HttpTokens**

Indicates whether IMDSv2 is required.

- `optional` – IMDSv2 is optional, which means that you can
use either IMDSv2 or IMDSv1.

- `required` – IMDSv2 is required, which means that IMDSv1 is
disabled, and you must use IMDSv2.

Type: String

Valid Values: `optional | required | no-preference`

Required: No

**HttpTokensEnforced**

Specifies whether to enforce the requirement of IMDSv2 on an instance at the time of
launch. When enforcement is enabled, the instance can't launch unless IMDSv2
( `HttpTokens`) is set to `required`. For more information, see
[Enforce IMDSv2 at the account level](../../../../services/ec2/latest/userguide/configuring-imds-new-instances.md#enforce-imdsv2-at-the-account-level) in the
_Amazon EC2 User Guide_.

Type: String

Valid Values: `disabled | enabled | no-preference`

Required: No

**InstanceMetadataTags**

Enables or disables access to an instance's tags from the instance metadata. For more
information, see [View tags for your EC2\
instances using instance metadata](../../../../services/ec2/latest/userguide/work-with-tags-in-imds.md) in the
_Amazon EC2 User Guide_.

Type: String

Valid Values: `disabled | enabled | no-preference`

Required: No

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**return**

If the request succeeds, the response returns `true`. If the request fails,
no response is returned, and instead an error message is returned.

Type: Boolean

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/modifyinstancemetadatadefaults.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/modifyinstancemetadatadefaults.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/modifyinstancemetadatadefaults.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/modifyinstancemetadatadefaults.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/modifyinstancemetadatadefaults.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/modifyinstancemetadatadefaults.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/modifyinstancemetadatadefaults.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/modifyinstancemetadatadefaults.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/modifyinstancemetadatadefaults.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/modifyinstancemetadatadefaults.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ModifyInstanceMaintenanceOptions

ModifyInstanceMetadataOptions
