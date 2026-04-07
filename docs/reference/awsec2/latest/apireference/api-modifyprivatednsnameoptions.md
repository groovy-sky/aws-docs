# ModifyPrivateDnsNameOptions

Modifies the options for instance hostnames for the specified instance.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually
making the request, and provides an error response. If you have the required
permissions, the error response is `DryRunOperation`. Otherwise, it is
`UnauthorizedOperation`.

Type: Boolean

Required: No

**EnableResourceNameDnsAAAARecord**

Indicates whether to respond to DNS queries for instance hostnames with DNS AAAA
records.

Type: Boolean

Required: No

**EnableResourceNameDnsARecord**

Indicates whether to respond to DNS queries for instance hostnames with DNS A
records.

Type: Boolean

Required: No

**InstanceId**

The ID of the instance.

Type: String

Required: Yes

**PrivateDnsHostnameType**

The type of hostname for EC2 instances. For IPv4 only subnets, an instance DNS name
must be based on the instance IPv4 address. For IPv6 only subnets, an instance DNS name
must be based on the instance ID. For dual-stack subnets, you can specify whether DNS
names use the instance IPv4 address or the instance ID.

Type: String

Valid Values: `ip-name | resource-name`

Required: No

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**return**

Returns `true` if the request succeeds; otherwise, it returns an
error.

Type: Boolean

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/ModifyPrivateDnsNameOptions)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/ModifyPrivateDnsNameOptions)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/modifyprivatednsnameoptions.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/modifyprivatednsnameoptions.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/modifyprivatednsnameoptions.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/modifyprivatednsnameoptions.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/modifyprivatednsnameoptions.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/modifyprivatednsnameoptions.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/ModifyPrivateDnsNameOptions)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/modifyprivatednsnameoptions.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ModifyNetworkInterfaceAttribute

ModifyPublicIpDnsNameOptions
