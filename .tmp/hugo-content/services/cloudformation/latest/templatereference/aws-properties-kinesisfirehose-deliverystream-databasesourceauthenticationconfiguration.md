This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisFirehose::DeliveryStream DatabaseSourceAuthenticationConfiguration

The structure to configure the authentication methods for Firehose to connect to source database endpoint.

Amazon Data Firehose is in preview release and is subject to change.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SecretsManagerConfiguration" : SecretsManagerConfiguration
}

```

### YAML

```yaml

  SecretsManagerConfiguration:
    SecretsManagerConfiguration

```

## Properties

`SecretsManagerConfiguration`

Property description not available.

_Required_: Yes

_Type_: [SecretsManagerConfiguration](aws-properties-kinesisfirehose-deliverystream-secretsmanagerconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Databases

DatabaseSourceConfiguration

All content copied from https://docs.aws.amazon.com/.
