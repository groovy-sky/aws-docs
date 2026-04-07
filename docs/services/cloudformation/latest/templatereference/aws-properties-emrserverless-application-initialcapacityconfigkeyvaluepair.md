This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMRServerless::Application InitialCapacityConfigKeyValuePair

The `InitialCapacityConfigKeyValuePair` property type specifies Property description not available. for an [AWS::EMRServerless::Application](aws-resource-emrserverless-application.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Key" : String,
  "Value" : InitialCapacityConfig
}

```

### YAML

```yaml

  Key: String
  Value:
    InitialCapacityConfig

```

## Properties

`Key`

Property description not available.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z]+[-_]*[a-zA-Z]+$`

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`Value`

Property description not available.

_Required_: Yes

_Type_: [InitialCapacityConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-emrserverless-application-initialcapacityconfig.html)

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

InitialCapacityConfig

InteractiveConfiguration
