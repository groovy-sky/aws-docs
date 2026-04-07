This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::TaskDefinition ProxyConfiguration

The configuration details for the App Mesh proxy.

For tasks that use the EC2 launch type, the container instances require at least
version 1.26.0 of the container agent and at least version 1.26.0-1 of the
`ecs-init` package to use a proxy configuration. If your container
instances are launched from the Amazon ECS optimized AMI version `20190301`
or later, then they contain the required versions of the container agent and
`ecs-init`. For more information, see [Amazon ECS-optimized\
Linux AMI](../../../amazonecs/latest/developerguide/ecs-optimized-ami.md)

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ContainerName" : String,
  "ProxyConfigurationProperties" : [ KeyValuePair, ... ],
  "Type" : String
}

```

### YAML

```yaml

  ContainerName: String
  ProxyConfigurationProperties:
    - KeyValuePair
  Type: String

```

## Properties

`ContainerName`

The name of the container that will serve as the App Mesh proxy.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ProxyConfigurationProperties`

The set of network configuration parameters to provide the Container Network Interface
(CNI) plugin, specified as key-value pairs.

- `IgnoredUID` \- (Required) The user ID (UID) of the proxy container
as defined by the `user` parameter in a container definition. This is
used to ensure the proxy ignores its own traffic. If `IgnoredGID` is
specified, this field can be empty.

- `IgnoredGID` \- (Required) The group ID (GID) of the proxy container
as defined by the `user` parameter in a container definition. This is
used to ensure the proxy ignores its own traffic. If `IgnoredUID` is
specified, this field can be empty.

- `AppPorts` \- (Required) The list of ports that the application
uses. Network traffic to these ports is forwarded to the
`ProxyIngressPort` and `ProxyEgressPort`.

- `ProxyIngressPort` \- (Required) Specifies the port that incoming
traffic to the `AppPorts` is directed to.

- `ProxyEgressPort` \- (Required) Specifies the port that outgoing
traffic from the `AppPorts` is directed to.

- `EgressIgnoredPorts` \- (Required) The egress traffic going to the
specified ports is ignored and not redirected to the
`ProxyEgressPort`. It can be an empty list.

- `EgressIgnoredIPs` \- (Required) The egress traffic going to the
specified IP addresses is ignored and not redirected to the
`ProxyEgressPort`. It can be an empty list.

_Required_: No

_Type_: Array of [KeyValuePair](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ecs-taskdefinition-keyvaluepair.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Type`

The proxy type. The only supported value is `APPMESH`.

_Required_: No

_Type_: String

_Allowed values_: `APPMESH`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

PortMapping

RepositoryCredentials
