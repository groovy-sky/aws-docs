This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::VirtualNode ServiceDiscovery

An object that represents the service discovery information for a virtual node.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AWSCloudMap" : AwsCloudMapServiceDiscovery,
  "DNS" : DnsServiceDiscovery
}

```

### YAML

```yaml

  AWSCloudMap:
    AwsCloudMapServiceDiscovery
  DNS:
    DnsServiceDiscovery

```

## Properties

`AWSCloudMap`

Specifies any AWS Cloud Map information for the virtual node.

_Required_: No

_Type_: [AwsCloudMapServiceDiscovery](aws-properties-appmesh-virtualnode-awscloudmapservicediscovery.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DNS`

Specifies the DNS information for the virtual node.

_Required_: No

_Type_: [DnsServiceDiscovery](aws-properties-appmesh-virtualnode-dnsservicediscovery.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PortMapping

SubjectAlternativeNameMatchers

All content copied from https://docs.aws.amazon.com/.
