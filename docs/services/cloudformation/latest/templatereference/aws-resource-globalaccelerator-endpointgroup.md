This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GlobalAccelerator::EndpointGroup

The `AWS::GlobalAccelerator::EndpointGroup` resource is a Global Accelerator resource type that contains information about
how you create an endpoint group for the specified listener. An endpoint group is a collection of endpoints in one AWS Region.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::GlobalAccelerator::EndpointGroup",
  "Properties" : {
      "EndpointConfigurations" : [ EndpointConfiguration, ... ],
      "EndpointGroupRegion" : String,
      "HealthCheckIntervalSeconds" : Integer,
      "HealthCheckPath" : String,
      "HealthCheckPort" : Integer,
      "HealthCheckProtocol" : String,
      "ListenerArn" : String,
      "PortOverrides" : [ PortOverride, ... ],
      "ThresholdCount" : Integer,
      "TrafficDialPercentage" : Number
    }
}

```

### YAML

```yaml

Type: AWS::GlobalAccelerator::EndpointGroup
Properties:
  EndpointConfigurations:
    - EndpointConfiguration
  EndpointGroupRegion: String
  HealthCheckIntervalSeconds: Integer
  HealthCheckPath: String
  HealthCheckPort: Integer
  HealthCheckProtocol: String
  ListenerArn: String
  PortOverrides:
    - PortOverride
  ThresholdCount: Integer
  TrafficDialPercentage: Number

```

## Properties

`EndpointConfigurations`

The list of endpoint objects.

_Required_: No

_Type_: Array of [EndpointConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-globalaccelerator-endpointgroup-endpointconfiguration.html)

_Minimum_: `0`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EndpointGroupRegion`

The AWS Regions where the endpoint group is located.

_Required_: Yes

_Type_: String

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`HealthCheckIntervalSeconds`

The time—10 seconds or 30 seconds—between health checks for each endpoint. The default value is 30.

_Required_: No

_Type_: Integer

_Minimum_: `10`

_Maximum_: `30`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HealthCheckPath`

If the protocol is HTTP/S, then this value provides the ping path that Global Accelerator uses for the destination on the
endpoints for health checks. The default is slash (/).

_Required_: No

_Type_: String

_Pattern_: `^/[-a-zA-Z0-9@:%_\\+.~#?&/=]*$`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HealthCheckPort`

The port that Global Accelerator uses to perform health checks on endpoints that are part of this endpoint group.

The default port is the port for the listener that this endpoint group is associated with. If the listener port is a
list, Global Accelerator uses the first specified port in the list of ports.

_Required_: No

_Type_: Integer

_Minimum_: `-1`

_Maximum_: `65535`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HealthCheckProtocol`

The protocol that Global Accelerator uses to perform health checks on endpoints that are part of this endpoint group. The default
value is TCP.

_Required_: No

_Type_: String

_Allowed values_: `TCP | HTTP | HTTPS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ListenerArn`

The Amazon Resource Name (ARN) of the listener.

_Required_: Yes

_Type_: String

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PortOverrides`

Allows you to override the destination ports used to route traffic to an endpoint.
Using a port override lets you map a list of external destination ports (that your
users send traffic to) to a list of internal destination ports that you want an application
endpoint to receive traffic on.

_Required_: No

_Type_: Array of [PortOverride](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-globalaccelerator-endpointgroup-portoverride.html)

_Minimum_: `0`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ThresholdCount`

The number of consecutive health checks required to set the state of a healthy endpoint to unhealthy, or to set an
unhealthy endpoint to healthy. The default value is 3.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TrafficDialPercentage`

The percentage of traffic to send to an AWS Regions. Additional traffic is distributed to other endpoint groups for
this listener.

Use this action to increase (dial up) or decrease (dial down) traffic to a specific Region. The percentage is
applied to the traffic that would otherwise have been routed to the Region based on optimal routing.

The default value is 100.

_Required_: No

_Type_: Number

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ARN of the endpoint group, such as
`arn:aws:globalaccelerator::012345678901:accelerator/1234abcd-abcd-1234-abcd-1234abcdefgh/listener/0123vxyz/endpoint-group/098765zyxwvu`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`EndpointGroupArn`

The ARN of the endpoint group, such as
`arn:aws:globalaccelerator::012345678901:accelerator/1234abcd-abcd-1234-abcd-1234abcdefgh/listener/0123vxyz/endpoint-group/098765zyxwvu`.

## Examples

### Add an endpoint group

These are examples to specify an endpoint group.

#### JSON

```json

"Resources": {
        "EndpointGroup": {
        "Type": "AWS::GlobalAccelerator::EndpointGroup",
        "Properties": {
            "ListenerArn": {
                "Ref": "Listener"
            },
            "EndpointGroupRegion": {
                "Ref": "EndpointRegion"
            },
            "TrafficDialPercentage": 100.0,
            "EndpointConfigurations": [
                {
                    "EndpointId": {
                        "Ref": "EndpointId"
                    }
                }
            ]
        }
    }
}
```

#### YAML

```yaml

EndpointGroup:
  Type: AWS::GlobalAccelerator::EndpointGroup
  Properties:
    ListenerArn:
      Ref: Listener
    EndpointGroupRegion:
      Ref: EndpointRegion
    TrafficDialPercentage: 100
    EndpointConfigurations:
    - EndpointId:
        Ref: EndpointId
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

EndpointConfiguration
