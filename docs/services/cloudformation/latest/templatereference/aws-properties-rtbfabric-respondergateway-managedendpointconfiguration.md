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

_Type_: [AutoScalingGroupsConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-rtbfabric-respondergateway-autoscalinggroupsconfiguration.html)

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`EksEndpointsConfiguration`

Describes the configuration of an Amazon Elastic Kubernetes Service endpoint.

_Required_: No

_Type_: [EksEndpointsConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-rtbfabric-respondergateway-eksendpointsconfiguration.html)

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

EksEndpointsConfiguration

Tag
