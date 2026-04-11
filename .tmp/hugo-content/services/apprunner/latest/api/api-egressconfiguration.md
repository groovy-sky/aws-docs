# EgressConfiguration

Describes configuration settings related to outbound network traffic of an AWS App Runner service.

## Contents

**EgressType**

The type of egress configuration.

Set to `DEFAULT` for access to resources hosted on public networks.

Set to `VPC` to associate your service to a custom VPC specified by `VpcConnectorArn`.

Type: String

Valid Values: `DEFAULT | VPC`

Required: No

**VpcConnectorArn**

The Amazon Resource Name (ARN) of the App Runner VPC connector that you want to associate with your App Runner service. Only valid when `EgressType =
        VPC`.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1011.

Pattern: `arn:aws(-[\w]+)*:[a-z0-9-\\.]{0,63}:[a-z0-9-\\.]{0,63}:[0-9]{12}:(\w|\/|-){1,1011}`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apprunner-2020-05-15/egressconfiguration.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apprunner-2020-05-15/egressconfiguration.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apprunner-2020-05-15/egressconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CustomDomain

EncryptionConfiguration

All content copied from https://docs.aws.amazon.com/.
