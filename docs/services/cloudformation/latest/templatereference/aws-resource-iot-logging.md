This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::Logging

Configure logging.

###### Note

If you already set the log function of AWS IoT Core, you can't deploy the AWS Cloud Development Kit (AWS CDK) to change the logging settings. You can change the logging settings by either:

- Importing the existing logging resource into your CloudFormation stack, such as with the [infrastructure as code generator (IaC generator)](../userguide/generate-iac.md).

- Calling `aws iot set-v2-logging-options --disable-all-logs` before creating a new CloudFormation stack. This command disables all AWS IoT logging. As a result, no AWS IoT logs will be delivered to Amazon CloudWatch until you re-enable logging.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IoT::Logging",
  "Properties" : {
      "AccountId" : String,
      "DefaultLogLevel" : String,
      "EventConfigurations" : [ EventConfiguration, ... ],
      "RoleArn" : String
    }
}

```

### YAML

```yaml

Type: AWS::IoT::Logging
Properties:
  AccountId: String
  DefaultLogLevel: String
  EventConfigurations:
    - EventConfiguration
  RoleArn: String

```

## Properties

`AccountId`

The account ID.

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9]{12}$`

_Minimum_: `12`

_Maximum_: `12`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DefaultLogLevel`

The default log level. Valid Values: `DEBUG | INFO | ERROR | WARN | DISABLED`

_Required_: Yes

_Type_: String

_Allowed values_: `ERROR | WARN | INFO | DEBUG | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EventConfigurations`

Configurations for event-based logging that specifies which event types to log and their logging settings. Overrides account-level logging for the specified event.

_Required_: No

_Type_: Array of [EventConfiguration](aws-properties-iot-logging-eventconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The role ARN used for the log.

_Required_: Yes

_Type_: String

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the log ID. For example:

`{"Ref": "Log-12345"}`

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TimeoutConfig

EventConfiguration

All content copied from https://docs.aws.amazon.com/.
