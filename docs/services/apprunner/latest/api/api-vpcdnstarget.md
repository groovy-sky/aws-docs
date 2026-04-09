# VpcDNSTarget

DNS Target record for a custom domain of this Amazon VPC.

## Contents

**DomainName**

The domain name of your target DNS that is associated with the Amazon VPC.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `[A-Za-z0-9*.-]{1,255}`

Required: No

**VpcId**

The ID of the Amazon VPC that is associated with the custom domain name of the target DNS.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 51200.

Pattern: `.*`

Required: No

**VpcIngressConnectionArn**

The Amazon Resource Name (ARN) of the VPC Ingress Connection that is associated with your service.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1011.

Pattern: `arn:aws(-[\w]+)*:[a-z0-9-\\.]{0,63}:[a-z0-9-\\.]{0,63}:[0-9]{12}:(\w|\/|-){1,1011}`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apprunner-2020-05-15/vpcdnstarget.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apprunner-2020-05-15/vpcdnstarget.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apprunner-2020-05-15/vpcdnstarget.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VpcConnector

VpcIngressConnection

All content copied from https://docs.aws.amazon.com/.
