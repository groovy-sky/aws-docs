This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::Runtime Code

The source code configuration that specifies the location and details of the code to be executed.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "S3" : S3Location
}

```

### YAML

```yaml

  S3:
    S3Location

```

## Properties

`S3`

The Amazon Amazon S3 object that contains the source code for the agent runtime.

_Required_: No

_Type_: [S3Location](aws-properties-bedrockagentcore-runtime-s3location.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ClaimMatchValueType

CodeConfiguration

All content copied from https://docs.aws.amazon.com/.
