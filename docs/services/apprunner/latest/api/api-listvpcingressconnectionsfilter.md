# ListVpcIngressConnectionsFilter

Returns a list of VPC Ingress Connections based on the filter provided. It can return either `ServiceArn` or `VpcEndpointId`, or both.

## Contents

**ServiceArn**

The Amazon Resource Name (ARN) of a service to filter by.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1011.

Pattern: `arn:aws(-[\w]+)*:[a-z0-9-\\.]{0,63}:[a-z0-9-\\.]{0,63}:[0-9]{12}:(\w|\/|-){1,1011}`

Required: No

**VpcEndpointId**

The ID of a VPC Endpoint to filter by.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 51200.

Pattern: `.*`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apprunner-2020-05-15/listvpcingressconnectionsfilter.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apprunner-2020-05-15/listvpcingressconnectionsfilter.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apprunner-2020-05-15/listvpcingressconnectionsfilter.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

InstanceConfiguration

NetworkConfiguration

All content copied from https://docs.aws.amazon.com/.
