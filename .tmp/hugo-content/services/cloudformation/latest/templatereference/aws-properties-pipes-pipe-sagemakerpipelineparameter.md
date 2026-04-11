This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pipes::Pipe SageMakerPipelineParameter

Name/Value pair of a parameter to start execution of a SageMaker AI Model Building
Pipeline.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Name" : String,
  "Value" : String
}

```

### YAML

```yaml

  Name: String
  Value: String

```

## Properties

`Name`

Name of parameter to start execution of a SageMaker AI Model Building
Pipeline.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9](-*[a-zA-Z0-9])*|(\$(\.[\w/_-]+(\[(\d+|\*)\])*)*)$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

Value of parameter to start execution of a SageMaker AI Model Building
Pipeline.

_Required_: Yes

_Type_: String

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

S3LogDestination

SelfManagedKafkaAccessConfigurationCredentials

All content copied from https://docs.aws.amazon.com/.
