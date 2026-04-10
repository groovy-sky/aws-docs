This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MSK::Cluster LoggingInfo

You can configure your MSK cluster to send broker logs to different destination types. This is a container for the configuration details related to broker logs.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BrokerLogs" : BrokerLogs
}

```

### YAML

```yaml

  BrokerLogs:
    BrokerLogs

```

## Properties

`BrokerLogs`

You can configure your MSK cluster to send broker logs to different destination types. This configuration specifies the details of these destinations.

_Required_: Yes

_Type_: [BrokerLogs](aws-properties-msk-cluster-brokerlogs.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

JmxExporter

NodeExporter

All content copied from https://docs.aws.amazon.com/.
