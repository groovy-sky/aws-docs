# EnableImageBlockPublicAccess

Enables _block public access for AMIs_ at the account level in the
specified AWS Region. This prevents the public sharing of your AMIs. However, if you already
have public AMIs, they will remain publicly available.

The API can take up to 10 minutes to configure this setting. During this time, if you run
[GetImageBlockPublicAccessState](api-getimageblockpublicaccessstate.md), the response will be `unblocked`. When
the API has completed the configuration, the response will be
`block-new-sharing`.

For more information, see [Block\
public access to your AMIs](../../../../services/ec2/latest/userguide/block-public-access-to-amis.md) in the _Amazon EC2 User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is
`DryRunOperation`. Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**ImageBlockPublicAccessState**

Specify `block-new-sharing` to enable block public access for AMIs at the
account level in the specified Region. This will block any attempt to publicly share your AMIs
in the specified Region.

Type: String

Valid Values: `block-new-sharing`

Required: Yes

## Response Elements

The following elements are returned by the service.

**imageBlockPublicAccessState**

Returns `block-new-sharing` if the request succeeds; otherwise, it returns an
error.

Type: String

Valid Values: `block-new-sharing`

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example enables _block public access for AMIs_ at the account
level in the specified Region to prevent users in your account from publicly sharing your
AMIs in the specified Region. If you already have public AMIs, they will remain publicly
available.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=EnableImageBlockPublicAccess
&Region=us-east-1
&ImageBlockPublicAccessState=block-new-sharing
&AUTHPARAMS
```

#### Sample Response

```

<EnableImageBlockPublicAccessResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
  <requestId>11aabb229-4eac-35bd-99ed-be587EXAMPLE</requestId>
  <return>block-new-sharing</return>
</EnableImageBlockPublicAccessResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/enableimageblockpublicaccess.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/enableimageblockpublicaccess.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/enableimageblockpublicaccess.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/enableimageblockpublicaccess.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/enableimageblockpublicaccess.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/enableimageblockpublicaccess.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/enableimageblockpublicaccess.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/enableimageblockpublicaccess.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/enableimageblockpublicaccess.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/enableimageblockpublicaccess.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

EnableImage

EnableImageDeprecation
