This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisFirehose::DeliveryStream SecretsManagerConfiguration

The structure that defines how Firehose accesses the secret.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Enabled" : Boolean,
  "RoleARN" : String,
  "SecretARN" : String
}

```

### YAML

```yaml

  Enabled: Boolean
  RoleARN: String
  SecretARN: String

```

## Properties

`Enabled`

Specifies whether you want to use the secrets manager feature. When set as
`True` the secrets manager configuration overwrites the existing secrets in
the destination configuration. When it's set to `False` Firehose falls back to
the credentials in the destination configuration.

_Required_: Yes

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RoleARN`

Specifies the role that Firehose assumes when calling the Secrets Manager API operation. When you provide the role, it overrides any destination specific role defined in the destination configuration. If you do not provide the then we use the destination specific role. This parameter is required for Splunk.

_Required_: No

_Type_: String

_Pattern_: `arn:.*:iam::\d{12}:role/[a-zA-Z_0-9+=,.@\-_/]+`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SecretARN`

The ARN of the secret that stores your credentials. It must be in the same region as the
Firehose stream and the role. The secret ARN can reside in a different account than the Firehose stream and role as Firehose supports cross-account secret access. This parameter is required when **Enabled** is set to `True`.

_Required_: No

_Type_: String

_Pattern_: `arn:.*:secretsmanager:[a-zA-Z0-9\-]+:\d{12}:secret:[a-zA-Z0-9\-/_+=.@]+`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SchemaEvolutionConfiguration

Serializer

All content copied from https://docs.aws.amazon.com/.
