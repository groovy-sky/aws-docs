# VpcIngressConnection

The AWS App Runner resource that specifies an App Runner endpoint for incoming traffic. It establishes a connection between a VPC interface endpoint and a App Runner
service, to make your App Runner service accessible from only within an Amazon VPC.

## Contents

**AccountId**

The Account Id you use to create the VPC Ingress Connection resource.

Type: String

Length Constraints: Fixed length of 12.

Pattern: `[0-9]{12}`

Required: No

**CreatedAt**

The time when the VPC Ingress Connection was created. It's in the Unix time stamp format.

- Type: Timestamp

- Required: Yes

Type: Timestamp

Required: No

**DeletedAt**

The time when the App Runner service was deleted. It's in the Unix time stamp format.

- Type: Timestamp

- Required: No

Type: Timestamp

Required: No

**DomainName**

The domain name associated with the VPC Ingress Connection resource.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `[A-Za-z0-9*.-]{1,255}`

Required: No

**IngressVpcConfiguration**

Specifications for the customer’s VPC and related PrivateLink VPC endpoint that are used to associate with the VPC Ingress Connection resource.

Type: [IngressVpcConfiguration](api-ingressvpcconfiguration.md) object

Required: No

**ServiceArn**

The Amazon Resource Name (ARN) of the service associated with the VPC Ingress Connection.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1011.

Pattern: `arn:aws(-[\w]+)*:[a-z0-9-\\.]{0,63}:[a-z0-9-\\.]{0,63}:[0-9]{12}:(\w|\/|-){1,1011}`

Required: No

**Status**

The current status of the VPC Ingress Connection.
The VPC Ingress Connection displays one of the following statuses: `AVAILABLE`, `PENDING_CREATION`, `PENDING_UPDATE`, `PENDING_DELETION`, `FAILED_CREATION`, `FAILED_UPDATE`, `FAILED_DELETION`, and `DELETED`..

Type: String

Valid Values: `AVAILABLE | PENDING_CREATION | PENDING_UPDATE | PENDING_DELETION | FAILED_CREATION | FAILED_UPDATE | FAILED_DELETION | DELETED`

Required: No

**VpcIngressConnectionArn**

The Amazon Resource Name (ARN) of the VPC Ingress Connection.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1011.

Pattern: `arn:aws(-[\w]+)*:[a-z0-9-\\.]{0,63}:[a-z0-9-\\.]{0,63}:[0-9]{12}:(\w|\/|-){1,1011}`

Required: No

**VpcIngressConnectionName**

The customer-provided VPC Ingress Connection name.

Type: String

Length Constraints: Minimum length of 4. Maximum length of 40.

Pattern: `[A-Za-z0-9][A-Za-z0-9\-_]{3,39}`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apprunner-2020-05-15/vpcingressconnection.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apprunner-2020-05-15/vpcingressconnection.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apprunner-2020-05-15/vpcingressconnection.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VpcDNSTarget

VpcIngressConnectionSummary

All content copied from https://docs.aws.amazon.com/.
