This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMRContainers::Endpoint EMREKSConfiguration

The `EMREKSConfiguration` property type specifies Property description not available. for an [AWS::EMRContainers::Endpoint](aws-resource-emrcontainers-endpoint.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Classification" : String,
  "Configurations" : [ EMREKSConfiguration, ... ],
  "Properties" : {Key: Value, ...}
}

```

### YAML

```yaml

  Classification: String
  Configurations:
    - EMREKSConfiguration
  Properties:
    Key: Value

```

## Properties

`Classification`

Property description not available.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Configurations`

Property description not available.

_Required_: No

_Type_: Array of [EMREKSConfiguration](aws-properties-emrcontainers-endpoint-emreksconfiguration.md)

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Properties`

Property description not available.

_Required_: No

_Type_: Object of String

_Pattern_: `.*`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ContainerLogRotationConfiguration

MonitoringConfiguration

All content copied from https://docs.aws.amazon.com/.
