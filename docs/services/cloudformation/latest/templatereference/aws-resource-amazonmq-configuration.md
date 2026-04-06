This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AmazonMQ::Configuration

Creates a new configuration for the specified configuration name. Amazon MQ uses
the default configuration (the engine type and version).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::AmazonMQ::Configuration",
  "Properties" : {
      "AuthenticationStrategy" : String,
      "Data" : String,
      "Description" : String,
      "EngineType" : String,
      "EngineVersion" : String,
      "Name" : String,
      "Tags" : [ TagsEntry, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::AmazonMQ::Configuration
Properties:
  AuthenticationStrategy: String
  Data: String
  Description: String
  EngineType: String
  EngineVersion: String
  Name: String
  Tags:
    - TagsEntry

```

## Properties

`AuthenticationStrategy`

Optional. The authentication strategy associated with the configuration. The
default is `SIMPLE`.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Data`

Amazon MQ for Active MQ: The base64-encoded XML configuration.
Amazon MQ for RabbitMQ: the base64-encoded Cuttlefish configuration.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the configuration.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EngineType`

Required. The type of broker engine. Currently, Amazon MQ supports `ACTIVEMQ` and `RABBITMQ`.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EngineVersion`

The broker engine version. Defaults to the latest available version for the specified broker engine type. For more information, see the
[ActiveMQ version management](../../../amazon-mq/latest/developer-guide/activemq-version-management.md)
and the [RabbitMQ version management](../../../amazon-mq/latest/developer-guide/rabbitmq-version-management.md)
sections in the Amazon MQ Developer Guide.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

Required. The name of the configuration. This value can contain only alphanumeric
characters, dashes, periods, underscores, and tildes (- . \_ ~). This value must be
1-150 characters long.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Create tags when creating the configuration.

_Required_: No

_Type_: Array of [TagsEntry](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-amazonmq-configuration-tagsentry.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon MQ configuration ID. For example:

`c-1234a5b6-78cd-901e-2fgh-3i45j6k178l9`

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the Amazon MQ configuration.

`arn:aws:mq:us-east-2:123456789012:configuration:MyConfigurationDevelopment:c-1234a5b6-78cd-901e-2fgh-3i45j6k178l9`

`Id`

The ID of the Amazon MQ configuration.

`c-1234a5b6-78cd-901e-2fgh-3i45j6k178l9`

`Revision`

The revision number of the configuration.

`1`

## Examples

### Amazon MQ Configuration

#### JSON

```json

{
  "Description": "Create an Amazon MQ for ActiveMQ configuration",
    "Configuration1": {
      "Type": "AWS::AmazonMQ::Configuration",
      "Properties": {
        "Data": {
          "Fn::Base64": "<?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"yes\"?>\n<broker xmlns=\"http://activemq.apache.org/schema/core\" start=\"false\">\n  <destinationPolicy>\n    <policyMap>\n      <policyEntries>\n        <policyEntry topic=\">\">\n          <pendingMessageLimitStrategy>\n            <constantPendingMessageLimitStrategy limit=\"3000\"/>\n          </pendingMessageLimitStrategy>\n        </policyEntry>\n      </policyEntries>\n    </policyMap>\n  </destinationPolicy>\n  <plugins>\n  </plugins>\n</broker>\n"
        },
        "EngineType": "ACTIVEMQ",
        "EngineVersion": "5.15.0",
        "Name": "my-configuration-1"
      }
   }
}
```

#### YAML

```yaml

---
Description: "Create an Amazon MQ for ActiveMQ configuration"
Resources:
  Configuration:
    Type: "AWS::AmazonMQ::Configuration"
    Properties:
      Data:
        ? "Fn::Base64"
        : |
            <?xml version="1.0" encoding="UTF-8" standalone="yes"?>
            <broker xmlns="http://activemq.apache.org/schema/core" start="false">
              <destinationPolicy>
                <policyMap>
                  <policyEntries>
                    <policyEntry topic=">">
                      <pendingMessageLimitStrategy>
                        <constantPendingMessageLimitStrategy limit="3000"/>
                      </pendingMessageLimitStrategy>
                    </policyEntry>
                  </policyEntries>
                </policyMap>
              </destinationPolicy>
              <plugins>
              </plugins>
            </broker>
      EngineType: ACTIVEMQ
      EngineVersion: "5.15.0"
      Name: my-configuration-1
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

User

TagsEntry
