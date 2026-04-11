This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DevOpsAgent::Association SlackTransmissionTarget

Defines the Slack channels where different types of agent notifications will be sent.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IncidentResponseTarget" : SlackChannel
}

```

### YAML

```yaml

  IncidentResponseTarget:
    SlackChannel

```

## Properties

`IncidentResponseTarget`

Destination for AWS DevOps Agent Incident Response.

_Required_: Yes

_Type_: [SlackChannel](aws-properties-devopsagent-association-slackchannel.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SlackConfiguration

SourceAwsConfiguration

All content copied from https://docs.aws.amazon.com/.
