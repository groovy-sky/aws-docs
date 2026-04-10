This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MSK::Configuration

Creates a new MSK configuration. To see an example of how to use this operation, first save the following text to a file and name the file `config-file.txt`.

```

auto.create.topics.enable = true

zookeeper.connection.timeout.ms = 1000

log.roll.ms = 604800000
```

Now run the following Python 3.6 script in the folder where you saved `config-file.txt`. This script uses the properties specified in `config-file.txt` to create a configuration named `SalesClusterConfiguration`. This configuration can work with Apache Kafka versions 1.1.1 and 2.1.0.

```PYTHON

import boto3

client = boto3.client('kafka')

config_file = open('config-file.txt', 'r')

server_properties = config_file.read()

response = client.create_configuration(
    Name='SalesClusterConfiguration',
    Description='The configuration to use on all sales clusters.',
    KafkaVersions=['1.1.1', '2.1.0'],
    ServerProperties=server_properties
)

print(response)
```

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::MSK::Configuration",
  "Properties" : {
      "Description" : String,
      "KafkaVersionsList" : [ String, ... ],
      "LatestRevision" : LatestRevision,
      "Name" : String,
      "ServerProperties" : String
    }
}

```

### YAML

```yaml

Type: AWS::MSK::Configuration
Properties:
  Description: String
  KafkaVersionsList:
    - String
  LatestRevision:
    LatestRevision
  Name: String
  ServerProperties: String

```

## Properties

`Description`

The description of the configuration.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KafkaVersionsList`

The [versions of Apache Kafka](../../../msk/latest/developerguide/supported-kafka-versions.md) with which you can use this MSK configuration.

When you update the `KafkaVersionsList` property, CloudFormation recreates a new configuration with the updated property before deleting the old configuration. Such an update requires a [resource replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement). To successfully update `KafkaVersionsList`, you must also update the `Name` property in the same operation.

If your configuration is attached with any clusters created using the AWS Management Console or AWS CLI, you'll need to manually delete the old configuration from the console after the update completes.

For more information, see [Can’t update KafkaVersionsList in MSK configuration](../../../msk/latest/developerguide/troubleshooting.md#troubleshoot-kafkaversionslist-cfn-update-failure) in the _Amazon MSK Developer Guide_.

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LatestRevision`

Latest revision of the MSK configuration.

_Required_: No

_Type_: [LatestRevision](aws-properties-msk-configuration-latestrevision.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the configuration. Configuration names are strings that match the regex "^\[0-9A-Za-z\]\[0-9A-Za-z-\]{0,}$".

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ServerProperties`

Contents of the `server.properties` file. When using the console, the SDK, or the AWS CLI, the contents of `server.properties` can be in plaintext.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you provide the logical ID of this resource to the `Ref` intrinsic function, it returns the MSK configuration.

### Fn::GetAtt

`Fn::GetAtt` returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

`Arn`

The Amazon Resource Name (ARN) of the configuration.

`LatestRevision.CreationTime`

The time when the configuration revision was created.

`LatestRevision.Description`

The description of the configuration revision.

`LatestRevision.Revision`

The revision number.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::MSK::ClusterPolicy

LatestRevision

All content copied from https://docs.aws.amazon.com/.
