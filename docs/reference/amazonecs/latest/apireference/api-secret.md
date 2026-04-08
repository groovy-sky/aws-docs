# Secret

An object representing the secret to expose to your container. Secrets can be exposed
to a container in the following ways:

- To inject sensitive data into your containers as environment variables, use
the `secrets` container definition parameter.

- To reference sensitive information in the log configuration of a container,
use the `secretOptions` container definition parameter.

For more information, see [Specifying\
sensitive data](../../../../services/amazonecs/latest/developerguide/specifying-sensitive-data.md) in the _Amazon Elastic Container Service Developer_
_Guide_.

## Contents

**name**

The name of the secret.

Type: String

Required: Yes

**valueFrom**

The secret to expose to the container. The supported values are either the full ARN of
the AWS Secrets Manager secret or the full ARN of the parameter in the SSM Parameter
Store.

For information about the require AWS Identity and Access Management permissions,
see [Required IAM permissions for Amazon ECS secrets](../../../../services/amazonecs/latest/developerguide/specifying-sensitive-data-secrets-secrets-iam.md) (for Secrets Manager) or
[Required IAM permissions for Amazon ECS secrets](../../../../services/amazonecs/latest/developerguide/specifying-sensitive-data-parameters.md) (for Systems Manager
Parameter store) in the _Amazon Elastic Container Service Developer_
_Guide_.

###### Note

If the SSM Parameter Store parameter exists in the same Region as the task you're
launching, then you can use either the full ARN or name of the parameter. If the
parameter exists in a different Region, then the full ARN must be specified.

Type: String

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ecs-2014-11-13/secret.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecs-2014-11-13/secret.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecs-2014-11-13/secret.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

Scale

Service
