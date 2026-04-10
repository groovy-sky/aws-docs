This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecurityAgent::AgentSpace CodeReviewSettings

The code review settings for an agent space, controlling which types of scanning are enabled.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ControlsScanning" : Boolean,
  "GeneralPurposeScanning" : Boolean
}

```

### YAML

```yaml

  ControlsScanning: Boolean
  GeneralPurposeScanning: Boolean

```

## Properties

`ControlsScanning`

Indicates whether controls scanning is enabled for code reviews.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GeneralPurposeScanning`

Indicates whether general-purpose scanning is enabled for code reviews.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWSResources

GitHubCapabilitiesResource

All content copied from https://docs.aws.amazon.com/.
