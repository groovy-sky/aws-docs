This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KafkaConnect::Connector KafkaClusterEncryptionInTransit

Details of encryption in transit to the Apache Kafka cluster.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EncryptionType" : String
}

```

### YAML

```yaml

  EncryptionType: String

```

## Properties

`EncryptionType`

The type of encryption in transit to the Apache Kafka cluster.

_Required_: Yes

_Type_: String

_Allowed values_: `PLAINTEXT | TLS`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

KafkaClusterClientAuthentication

LogDelivery

All content copied from https://docs.aws.amazon.com/.
