# Hypervisor

Represents the hypervisor's permissions to which the gateway will connect.

A hypervisor is hardware, software, or firmware that creates and manages virtual machines,
and allocates resources to them.

## Contents

**Host**

The server host of the hypervisor. This can be either an IP address or a fully-qualified
domain name (FQDN).

Type: String

Length Constraints: Minimum length of 3. Maximum length of 128.

Pattern: `.+`

Required: No

**HypervisorArn**

The Amazon Resource Name (ARN) of the hypervisor.

Type: String

Length Constraints: Minimum length of 50. Maximum length of 500.

Pattern: `arn:(aws|aws-cn|aws-us-gov):backup-gateway(:[a-zA-Z-0-9]+){3}\/[a-zA-Z-0-9]+`

Required: No

**KmsKeyArn**

The Amazon Resource Name (ARN) of the AWS Key Management Service used to encrypt the
hypervisor.

Type: String

Length Constraints: Minimum length of 50. Maximum length of 500.

Pattern: `(^arn:(aws|aws-cn|aws-us-gov):kms:([a-zA-Z0-9-]+):([0-9]+):(key|alias)/(\S+)$)|(^alias/(\S+)$)`

Required: No

**Name**

The name of the hypervisor.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 100.

Pattern: `[a-zA-Z0-9-]*`

Required: No

**State**

The state of the hypervisor.

Type: String

Valid Values: `PENDING | ONLINE | OFFLINE | ERROR`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/backup-gateway-2021-01-01/hypervisor.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/backup-gateway-2021-01-01/hypervisor.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/backup-gateway-2021-01-01/hypervisor.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GatewayDetails

HypervisorDetails

All content copied from https://docs.aws.amazon.com/.
