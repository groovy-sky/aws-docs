# InstanceConfiguration

Describes the runtime configuration of an AWS App Runner service instance (scaling unit).

## Contents

**Cpu**

The number of CPU units reserved for each instance of your App Runner service.

Default: `1 vCPU`

Type: String

Length Constraints: Minimum length of 3. Maximum length of 9.

Pattern: `256|512|1024|2048|4096|(0.25|0.5|1|2|4) vCPU`

Required: No

**InstanceRoleArn**

The Amazon Resource Name (ARN) of an IAM role that provides permissions to your App Runner service. These are permissions that your code needs when it calls
any AWS APIs.

Type: String

Length Constraints: Minimum length of 29. Maximum length of 1024.

Pattern: `arn:(aws|aws-us-gov|aws-cn|aws-iso|aws-iso-b):iam::[0-9]{12}:(role|role\/service-role)\/[\w+=,.@\-/]{1,1000}`

Required: No

**Memory**

The amount of memory, in MB or GB, reserved for each instance of your App Runner service.

Default: `2 GB`

Type: String

Length Constraints: Minimum length of 3. Maximum length of 6.

Pattern: `512|1024|2048|3072|4096|6144|8192|10240|12288|(0.5|1|2|3|4|6|8|10|12) GB`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apprunner-2020-05-15/instanceconfiguration.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apprunner-2020-05-15/instanceconfiguration.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apprunner-2020-05-15/instanceconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IngressVpcConfiguration

ListVpcIngressConnectionsFilter

All content copied from https://docs.aws.amazon.com/.
