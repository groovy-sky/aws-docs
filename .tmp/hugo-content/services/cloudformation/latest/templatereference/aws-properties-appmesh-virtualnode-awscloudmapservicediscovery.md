This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::VirtualNode AwsCloudMapServiceDiscovery

An object that represents the AWS Cloud Map service discovery information for
your virtual node.

###### Note

AWS Cloud Map is not available in the eu-south-1 Region.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Attributes" : [ AwsCloudMapInstanceAttribute, ... ],
  "IpPreference" : String,
  "NamespaceName" : String,
  "ServiceName" : String
}

```

### YAML

```yaml

  Attributes:
    - AwsCloudMapInstanceAttribute
  IpPreference: String
  NamespaceName: String
  ServiceName: String

```

## Properties

`Attributes`

A string map that contains attributes with values that you can use to filter instances
by any custom attribute that you specified when you registered the instance. Only instances
that match all of the specified key/value pairs will be returned.

_Required_: No

_Type_: Array of [AwsCloudMapInstanceAttribute](aws-properties-appmesh-virtualnode-awscloudmapinstanceattribute.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IpPreference`

The preferred IP version that this virtual node uses. Setting the IP preference on the
virtual node only overrides the IP preference set for the mesh on this specific
node.

_Required_: No

_Type_: String

_Allowed values_: `IPv6_PREFERRED | IPv4_PREFERRED | IPv4_ONLY | IPv6_ONLY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NamespaceName`

The HTTP name of the AWS Cloud Map namespace to use.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServiceName`

The name of the AWS Cloud Map service to use.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AwsCloudMapInstanceAttribute

Backend

All content copied from https://docs.aws.amazon.com/.
