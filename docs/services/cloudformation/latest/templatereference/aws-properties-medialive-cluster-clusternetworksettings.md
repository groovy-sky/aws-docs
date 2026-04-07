This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Cluster ClusterNetworkSettings

The `ClusterNetworkSettings` property type specifies Property description not available. for an [AWS::MediaLive::Cluster](aws-resource-medialive-cluster.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DefaultRoute" : String,
  "InterfaceMappings" : [ InterfaceMapping, ... ]
}

```

### YAML

```yaml

  DefaultRoute: String
  InterfaceMappings:
    - InterfaceMapping

```

## Properties

`DefaultRoute`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InterfaceMappings`

Property description not available.

_Required_: No

_Type_: Array of [InterfaceMapping](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-medialive-cluster-interfacemapping.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::MediaLive::Cluster

InterfaceMapping
