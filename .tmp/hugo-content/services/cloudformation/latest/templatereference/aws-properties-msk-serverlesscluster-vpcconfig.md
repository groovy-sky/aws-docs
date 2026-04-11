This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MSK::ServerlessCluster VpcConfig

The `VpcConfig` property type specifies Property description not available. for an [AWS::MSK::ServerlessCluster](aws-resource-msk-serverlesscluster.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SecurityGroups" : [ String, ... ],
  "SubnetIds" : [ String, ... ]
}

```

### YAML

```yaml

  SecurityGroups:
    - String
  SubnetIds:
    - String

```

## Properties

`SecurityGroups`

Property description not available.

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SubnetIds`

Property description not available.

_Required_: Yes

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Sasl

AWS::MSK::Topic

All content copied from https://docs.aws.amazon.com/.
