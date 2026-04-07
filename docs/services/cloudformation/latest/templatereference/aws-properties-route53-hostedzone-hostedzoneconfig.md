This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Route53::HostedZone HostedZoneConfig

A complex type that contains an optional comment about your hosted zone. If you don't want to specify a comment, omit both the
`HostedZoneConfig` and `Comment` elements.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Comment" : String
}

```

### YAML

```yaml

  Comment: String

```

## Properties

`Comment`

Any comments that you want to include about the hosted zone.

_Required_: No

_Type_: String

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [Return values](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-hostedzone.html#aws-resource-route53-hostedzone-return-values)
in the topic
[AWS::Route53::HostedZone](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-hostedzone.html)

- [HostedZoneConfig](../../../../reference/route53/latest/apireference/api-hostedzoneconfig.md)
in the _Amazon Route 53 API Reference_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Route53::HostedZone

HostedZoneFeatures
