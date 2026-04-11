# Gateway

A gateway is an AWS Backup Gateway appliance that runs on the customer's network
to provide seamless connectivity to backup storage in the AWS Cloud.

## Contents

**GatewayArn**

The Amazon Resource Name (ARN) of the gateway. Use the `ListGateways` operation
to return a list of gateways for your account and AWS Region.

Type: String

Length Constraints: Minimum length of 50. Maximum length of 180.

Pattern: `arn:(aws|aws-cn|aws-us-gov):backup-gateway(:[a-zA-Z-0-9]+){3}\/[a-zA-Z-0-9]+`

Required: No

**GatewayDisplayName**

The display name of the gateway.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 100.

Pattern: `[a-zA-Z0-9-]*`

Required: No

**GatewayType**

The type of the gateway.

Type: String

Valid Values: `BACKUP_VM`

Required: No

**HypervisorId**

The hypervisor ID of the gateway.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 100.

Required: No

**LastSeenTime**

The last time AWS Backup gateway communicated with the gateway, in Unix format and
UTC time.

Type: Timestamp

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/backup-gateway-2021-01-01/gateway.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/backup-gateway-2021-01-01/gateway.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/backup-gateway-2021-01-01/gateway.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BandwidthRateLimitInterval

GatewayDetails

All content copied from https://docs.aws.amazon.com/.
