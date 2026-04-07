# GetImageBlockPublicAccessState

Gets the current state of _block public access for AMIs_ at the account
level in the specified AWS Region.

For more information, see [Block\
public access to your AMIs](../../../../services/ec2/latest/userguide/block-public-access-to-amis.md) in the _Amazon EC2 User Guide_.

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

**imageBlockPublicAccessState**

The current state of block public access for AMIs at the account level in the specified
AWS Region.

Possible values:

- `block-new-sharing` \- Any attempt to publicly share your AMIs in the
specified Region is blocked.

- `unblocked` \- Your AMIs in the specified Region can be publicly
shared.

Type: String

**managedBy**

The entity that manages the state for block public access for AMIs. Possible values
include:

- `account` \- The state is managed by the account.

- `declarative-policy` \- The state is managed by a declarative policy and
can't be modified by the account.

Type: String

Valid Values: `account | declarative-policy`

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example gets the state of _block public access for AMIs_ at
the account level in the specified Region to see whether the public sharing of your AMIs
is blocked in your account. The value for the response is either
`block-new-sharing` or `unblocked`.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=GetImageBlockPublicAccessState
&Region=us-east-1
&AUTHPARAMS
```

#### Sample Response

```

<GetImageBlockPublicAccessStateResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
  <requestId>11aabb229-4eac-35bd-99ed-be587EXAMPLE</requestId>
  <return>block-new-sharing</return>
</GetImageBlockPublicAccessStateResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/GetImageBlockPublicAccessState)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/GetImageBlockPublicAccessState)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/getimageblockpublicaccessstate.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/getimageblockpublicaccessstate.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/getimageblockpublicaccessstate.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/getimageblockpublicaccessstate.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/getimageblockpublicaccessstate.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/getimageblockpublicaccessstate.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/GetImageBlockPublicAccessState)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/getimageblockpublicaccessstate.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

GetImageAncestry

GetInstanceMetadataDefaults
