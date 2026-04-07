This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppRunner::AutoScalingConfiguration

The `AWS::AppRunner::AutoScalingConfiguration` resource is an AWS App Runner resource type that specifies an App Runner
automatic scaling configuration.

App Runner requires this resource to set
non-default auto scaling settings for instances used to process the web requests. You can share an auto scaling configuration across multiple services.

Create multiple revisions of a configuration by calling this action multiple times using the same `AutoScalingConfigurationName`. The call
returns incremental `AutoScalingConfigurationRevision` values. When you create a service and configure an auto scaling configuration resource,
the service uses the latest active revision of the auto scaling configuration by default. You can optionally configure the service to use a specific
revision.

Configure a higher `MinSize` to increase the spread of your App Runner service over more Availability Zones in the AWS Region. The tradeoff is
a higher minimal cost.

Configure a lower `MaxSize` to control your cost. The tradeoff is lower responsiveness during peak demand.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::AppRunner::AutoScalingConfiguration",
  "Properties" : {
      "AutoScalingConfigurationName" : String,
      "MaxConcurrency" : Integer,
      "MaxSize" : Integer,
      "MinSize" : Integer,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::AppRunner::AutoScalingConfiguration
Properties:
  AutoScalingConfigurationName: String
  MaxConcurrency: Integer
  MaxSize: Integer
  MinSize: Integer
  Tags:
    - Tag

```

## Properties

`AutoScalingConfigurationName`

The customer-provided auto scaling configuration name. It can be used in multiple revisions of a configuration.

_Required_: No

_Type_: String

_Pattern_: `[A-Za-z0-9][A-Za-z0-9\-_]{3,31}`

_Minimum_: `4`

_Maximum_: `32`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MaxConcurrency`

The maximum number of concurrent requests that an instance processes. If the number of concurrent requests exceeds this limit, App Runner scales the service
up.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MaxSize`

The maximum number of instances that a service scales up to. At most `MaxSize` instances actively serve traffic for your service.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MinSize`

The minimum number of instances that App Runner provisions for a service. The service always has at least `MinSize` provisioned instances. Some
of them actively serve traffic. The rest of them (provisioned and inactive instances) are a cost-effective compute capacity reserve and are ready to be
quickly activated. You pay for memory usage of all the provisioned instances. You pay for CPU usage of only the active subset.

App Runner temporarily doubles the number of provisioned instances during deployments, to maintain the same capacity for both old and new code.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

A list of metadata items that you can associate with your auto scaling configuration resource. A tag is a key-value pair.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-apprunner-autoscalingconfiguration-tag.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When the logical ID of this resource is provided to the `Ref` intrinsic function, `Ref` returns the resource name.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AutoScalingConfigurationArn`

The Amazon Resource Name (ARN) of this auto scaling configuration.

`AutoScalingConfigurationRevision`

The revision of this auto scaling configuration. It's unique among all the active configurations that share the same `AutoScalingConfigurationName`.

`Latest`

It's set to true for the configuration with the highest `Revision` among all configurations that share the same `AutoScalingConfigurationName`.
It's set to false otherwise. App Runner temporarily doubles the number of provisioned instances during deployments, to maintain the same capacity for both old and new code.

## Examples

### Auto Scaling configuration

This example illustrates setting up an automatic scaling configuration for your App Runner service.

#### JSON

```json

{
  "Type": "AWS::AppRunner::AutoScalingConfiguration",
  "Properties": {
    "AutoScalingConfigurationName": "AUTO_SCALING_CONFIGURATION_NAME",
    "MaxConcurrency": 100,
    "MaxSize": 25,
    "MinSize": 1
  }
}
```

#### YAML

```yaml

Type: AWS::AppRunner::AutoScalingConfiguration
Properties:
  AutoScalingConfigurationName: "AUTO_SCALING_CONFIGURATION_NAME"
  MaxConcurrency: 100
  MaxSize: 25
  MinSize: 1

```

## See also

- [Managing App Runner automatic scaling](https://docs.aws.amazon.com/apprunner/latest/dg/manage-autoscaling.html) in the
_AWS App Runner Developer Guide_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS App Runner

Tag
