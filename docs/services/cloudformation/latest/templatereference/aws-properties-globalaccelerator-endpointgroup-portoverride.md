This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GlobalAccelerator::EndpointGroup PortOverride

Override specific listener ports used to route traffic to endpoints that are part of an endpoint group.
For example, you can create a port override in which the listener
receives user traffic on ports 80 and 443, but your accelerator routes that traffic to ports 1080
and 1443, respectively, on the endpoints.

For more information, see [Port overrides](https://docs.aws.amazon.com/global-accelerator/latest/dg/about-endpoint-groups-port-override.html) in the _AWS Global Accelerator Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EndpointPort" : Integer,
  "ListenerPort" : Integer
}

```

### YAML

```yaml

  EndpointPort: Integer
  ListenerPort: Integer

```

## Properties

`EndpointPort`

The endpoint port that you want a listener port to be mapped to. This is the port on the endpoint,
such as the Application Load Balancer or Amazon EC2 instance.

_Required_: Yes

_Type_: Integer

_Minimum_: `0`

_Maximum_: `65535`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ListenerPort`

The listener port that you want to map to a specific endpoint port. This is the port that user traffic
arrives to the Global Accelerator on.

_Required_: Yes

_Type_: Integer

_Minimum_: `0`

_Maximum_: `65535`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

EndpointConfiguration

AWS::GlobalAccelerator::Listener
