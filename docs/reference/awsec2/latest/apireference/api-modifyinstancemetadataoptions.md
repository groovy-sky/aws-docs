# ModifyInstanceMetadataOptions

Modify the instance metadata parameters on a running or stopped instance. When you
modify the parameters on a stopped instance, they are applied when the instance is
started. When you modify the parameters on a running instance, the API responds with a
state of “pending”. After the parameter modifications are successfully applied to the
instance, the state of the modifications changes from “pending” to “applied” in
subsequent describe-instances API calls. For more information, see [Instance metadata and user data](../../../../services/ec2/latest/userguide/ec2-instance-metadata.md) in the
_Amazon EC2 User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually
making the request, and provides an error response. If you have the required
permissions, the error response is `DryRunOperation`. Otherwise, it is
`UnauthorizedOperation`.

Type: Boolean

Required: No

**HttpEndpoint**

Enables or disables the HTTP metadata endpoint on your instances. If this parameter is
not specified, the existing state is maintained.

If you specify a value of `disabled`, you cannot access your instance
metadata.

Type: String

Valid Values: `disabled | enabled`

Required: No

**HttpProtocolIpv6**

Enables or disables the IPv6 endpoint for the instance metadata service.
Applies only if you enabled the HTTP metadata endpoint.

Type: String

Valid Values: `disabled | enabled`

Required: No

**HttpPutResponseHopLimit**

The desired HTTP PUT response hop limit for instance metadata requests. The larger the
number, the further instance metadata requests can travel. If no parameter is specified,
the existing state is maintained.

Possible values: Integers from 1 to 64

Type: Integer

Required: No

**HttpTokens**

Indicates whether IMDSv2 is required.

- `optional` \- IMDSv2 is optional. You can choose whether to send a
session token in your instance metadata retrieval requests. If you retrieve
IAM role credentials without a session token, you receive the IMDSv1 role
credentials. If you retrieve IAM role credentials using a valid session token,
you receive the IMDSv2 role credentials.

- `required` \- IMDSv2 is required. You must send a session token
in your instance metadata retrieval requests. With this option, retrieving the
IAM role credentials always returns IMDSv2 credentials; IMDSv1 credentials are
not available.

Default:

- If the value of `ImdsSupport` for the Amazon Machine Image (AMI)
for your instance is `v2.0` and the account level default is set to
`no-preference`, the default is `required`.

- If the value of `ImdsSupport` for the Amazon Machine Image (AMI)
for your instance is `v2.0`, but the account level default is set to
`V1 or V2`, the default is `optional`.

The default value can also be affected by other combinations of parameters. For more
information, see [Order of precedence for instance metadata options](../../../../services/ec2/latest/userguide/configuring-instance-metadata-options.md#instance-metadata-options-order-of-precedence) in the
_Amazon EC2 User Guide_.

Type: String

Valid Values: `optional | required`

Required: No

**InstanceId**

The ID of the instance.

Type: String

Required: Yes

**InstanceMetadataTags**

Set to `enabled` to allow access to instance tags from the instance
metadata. Set to `disabled` to turn off access to instance tags from the
instance metadata. For more information, see [View tags for your EC2\
instances using instance metadata](../../../../services/ec2/latest/userguide/work-with-tags-in-imds.md).

Type: String

Valid Values: `disabled | enabled`

Required: No

## Response Elements

The following elements are returned by the service.

**instanceId**

The ID of the instance.

Type: String

**instanceMetadataOptions**

The metadata options for the instance.

Type: [InstanceMetadataOptionsResponse](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_InstanceMetadataOptionsResponse.html) object

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example 1: Turn on token requirement

The following example disables access to the instance metadata unless a
session token is used in the instance metadata request header. To turn on token
requirement, specify `required` for `HttpTokens`.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=ModifyInstanceMetadataOptions
&InstanceId=i-1234567890abcdef0
&HttpTokens=required
&AUTHPARAMS
```

#### Sample Response

```

<ModifyInstanceMetadataOptions xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>59dbff89-35bd-4eac-99ed-be587EXAMPLE</requestId>
    <instanceId>i-1234567890abcdef0</instanceId>
    <MetadataOptions>
          <state>pending</state>
          <HttpTokens>required</HttpTokens>
          <HttpPutResponseHopLimit>1</HttpPutResponseHopLimit>
          <HttpEndpoint>enabled</HttpEndpoint>
    </MetadataOptions>
</ModifyInstanceMetadataOptions>
```

### Example 2: Turn off access to instance metadata

The following example disables access to the instance metadata by changing the
HTTP endpoint state to disabled. To turn off access to instance metadata,
specify `disabled` for `HttpEndpoint`.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=ModifyInstanceMetadataOptions
&InstanceId=i-1234567890abcdef0
&HttpEndpoint=disabled
&AUTHPARAMS
```

#### Sample Response

```

<ModifyInstanceMetadataOptions xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>59dbff89-35bd-4eac-99ed-be587EXAMPLE</requestId>
    <instanceId>i-1234567890abcdef0</instanceId>
    <MetadataOptions>
          <state>pending</state>
          <HttpTokens>required</HttpTokens>
          <HttpPutResponseHopLimit>1</HttpPutResponseHopLimit>
          <HttpEndpoint>disabled</HttpEndpoint>
    </MetadataOptions>
</ModifyInstanceMetadataOptions>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/ModifyInstanceMetadataOptions)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/ModifyInstanceMetadataOptions)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/ModifyInstanceMetadataOptions)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/ec2-2016-11-15/ModifyInstanceMetadataOptions)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/ModifyInstanceMetadataOptions)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/ec2-2016-11-15/ModifyInstanceMetadataOptions)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/ec2-2016-11-15/ModifyInstanceMetadataOptions)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/ec2-2016-11-15/ModifyInstanceMetadataOptions)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/ModifyInstanceMetadataOptions)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/ModifyInstanceMetadataOptions)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ModifyInstanceMetadataDefaults

ModifyInstanceNetworkPerformanceOptions
