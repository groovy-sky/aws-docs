This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lightsail::Distribution InputOrigin

`InputOrigin` is a property of the [AWS::Lightsail::Distribution](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-distribution.html) resource. It describes the origin resource of an
Amazon Lightsail content delivery network (CDN) distribution.

An origin can be a instance, bucket, or load balancer. A distribution pulls content from
an origin, caches it, and serves it to viewers through a worldwide network of edge servers.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Name" : String,
  "ProtocolPolicy" : String,
  "RegionName" : String
}

```

### YAML

```yaml

  Name: String
  ProtocolPolicy: String
  RegionName: String

```

## Properties

`Name`

The name of the origin resource.

_Required_: No

_Type_: String

_Pattern_: `\w[\w\-]*\w`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProtocolPolicy`

The protocol that your Amazon Lightsail distribution uses when establishing
a connection with your origin to pull content.

_Required_: No

_Type_: String

_Allowed values_: `http-only | https-only`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RegionName`

The AWS Region name of the origin resource.

_Required_: No

_Type_: String

_Allowed values_: `us-east-1 | us-east-2 | us-west-1 | us-west-2 | eu-west-1 | eu-west-2 | eu-west-3 | eu-central-1 | ca-central-1 | ap-south-1 | ap-southeast-1 | ap-southeast-2 | ap-northeast-1 | ap-northeast-2 | eu-north-1 | ap-southeast-3`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

HeaderObject

QueryStringObject
