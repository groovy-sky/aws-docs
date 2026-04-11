This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KafkaConnect::Connector KafkaClusterClientAuthentication

The client authentication information used in order to authenticate with the Apache
Kafka cluster.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AuthenticationType" : String
}

```

### YAML

```yaml

  AuthenticationType: String

```

## Properties

`AuthenticationType`

The type of client authentication used to connect to the Apache Kafka cluster. Value
NONE means that no client authentication is used.

_Required_: Yes

_Type_: String

_Allowed values_: `NONE | IAM`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

KafkaCluster

KafkaClusterEncryptionInTransit

All content copied from https://docs.aws.amazon.com/.
