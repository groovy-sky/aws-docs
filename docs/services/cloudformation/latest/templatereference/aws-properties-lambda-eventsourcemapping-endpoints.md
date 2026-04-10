This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lambda::EventSourceMapping Endpoints

The list of bootstrap servers for your Kafka brokers in the following format: `"KafkaBootstrapServers": ["abc.xyz.com:xxxx","abc2.xyz.com:xxxx"]`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "KafkaBootstrapServers" : [ String, ... ]
}

```

### YAML

```yaml

  KafkaBootstrapServers:
    - String

```

## Properties

`KafkaBootstrapServers`

The list of bootstrap servers for your Kafka brokers in the following format: `"KafkaBootstrapServers": ["abc.xyz.com:xxxx","abc2.xyz.com:xxxx"]`.

_Required_: No

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `300 | 10`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DocumentDBEventSourceConfig

Filter

All content copied from https://docs.aws.amazon.com/.
