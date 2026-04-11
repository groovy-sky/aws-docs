This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Logs::Transformer ParseToOCSF

This processor converts logs into [Open Cybersecurity Schema\
Framework (OCSF)](https://ocsf.io/) events.

For more information about this processor including examples, see [parseToOCSF](../../../amazoncloudwatch/latest/logs/cloudwatch-logs-transformation.md#CloudWatch-Logs-Transformation-parseToOCSF) in the _CloudWatch Logs User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EventSource" : String,
  "MappingVersion" : String,
  "OcsfVersion" : String,
  "Source" : String
}

```

### YAML

```yaml

  EventSource: String
  MappingVersion: String
  OcsfVersion: String
  Source: String

```

## Properties

`EventSource`

Specify the service or process that produces the log events that will be converted with
this processor.

_Required_: Yes

_Type_: String

_Allowed values_: `CloudTrail | Route53Resolver | VPCFlow | EKSAudit | AWSWAF`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MappingVersion`

The version of the OCSF mapping to use for parsing log data.

_Required_: No

_Type_: String

_Pattern_: `^v\d+\.\d+(\.\d+)?$`

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OcsfVersion`

Specify which version of the OCSF schema to use for the transformed log events.

_Required_: Yes

_Type_: String

_Allowed values_: `V1.1 | V1.5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Source`

The path to the field in the log event that you want to parse. If you omit this value, the
whole log message is parsed.

_Required_: No

_Type_: String

_Pattern_: `^.*[a-zA-Z0-9]+.*$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ParseRoute53

ParseVPC

All content copied from https://docs.aws.amazon.com/.
