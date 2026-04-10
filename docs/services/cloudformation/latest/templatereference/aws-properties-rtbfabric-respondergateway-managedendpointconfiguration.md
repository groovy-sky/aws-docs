This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RTBFabric::ResponderGateway ManagedEndpointConfiguration

Describes the configuration of a managed endpoint.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AutoScalingGroupsConfiguration" : AutoScalingGroupsConfiguration,
  "EksEndpointsConfiguration" : EksEndpointsConfiguration
}

```

### YAML

```yaml

  AutoScalingGroupsConfiguration:
    AutoScalingGroupsConfiguration
  EksEndpointsConfiguration:
    EksEndpointsConfiguration

```

## Properties

`AutoScalingGroupsConfiguration`

Describes the configuration of an auto scaling group.

_Required_: No

_Type_: [AutoScalingGroupsConfiguration](aws-properties-rtbfabric-respondergateway-autoscalinggroupsconfiguration.md)

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`EksEndpointsConfiguration`

Describes the configuration of an Amazon Elastic Kubernetes Service endpoint.

_Required_: No

_Type_: [EksEndpointsConfiguration](aws-properties-rtbfabric-respondergateway-eksendpointsconfiguration.md)

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EksEndpointsConfiguration

Tag

All content copied from https://docs.aws.amazon.com/.
