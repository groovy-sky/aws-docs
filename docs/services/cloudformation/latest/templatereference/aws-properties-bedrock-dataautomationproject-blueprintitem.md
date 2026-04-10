This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::DataAutomationProject BlueprintItem

An abbreviated summary of a blueprint.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BlueprintArn" : String,
  "BlueprintStage" : String,
  "BlueprintVersion" : String
}

```

### YAML

```yaml

  BlueprintArn: String
  BlueprintStage: String
  BlueprintVersion: String

```

## Properties

`BlueprintArn`

The blueprint's ARN.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws(|-cn|-us-gov|-iso|-iso-[a-z]):bedrock:[a-zA-Z0-9-]*:(aws|[0-9]{12}):blueprint/(bedrock-data-automation-public-[a-zA-Z0-9-_]{1,30}|[a-zA-Z0-9-]{12,36})$`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BlueprintStage`

The blueprint's stage.

_Required_: No

_Type_: String

_Allowed values_: `DEVELOPMENT | LIVE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BlueprintVersion`

The blueprint's version.

_Required_: No

_Type_: String

_Pattern_: `^[0-9]*$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AudioStandardOutputConfiguration

ChannelLabelingConfiguration

All content copied from https://docs.aws.amazon.com/.
