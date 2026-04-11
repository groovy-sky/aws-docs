This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Route53::HostedZone HostedZoneTag

A complex type that contains information about a tag that you want to add or edit for
the specified health check or hosted zone.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Key" : String,
  "Value" : String
}

```

### YAML

```yaml

  Key: String
  Value: String

```

## Properties

`Key`

The value of `Key` depends on the operation that you want to
perform:

- **Add a tag to a health check or hosted zone**:
`Key` is the name that you want to give the new tag.

- **Edit a tag**: `Key` is the name of
the tag that you want to change the `Value` for.

- **Delete a key**: `Key` is the name
of the tag you want to remove.

- **Give a name to a health check**: Edit the
default `Name` tag. In the Amazon Route 53 console, the list of your
health checks includes a **Name** column that lets
you see the name that you've given to each health check.

_Required_: Yes

_Type_: String

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The value of `Value` depends on the operation that you want to
perform:

- **Add a tag to a health check or hosted zone**:
`Value` is the value that you want to give the new tag.

- **Edit a tag**: `Value` is the new
value that you want to assign the tag.

_Required_: Yes

_Type_: String

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [Return values](../userguide/aws-resource-route53-hostedzone.md#aws-resource-route53-hostedzone-return-values)
in the topic
[AWS::Route53::HostedZone](../userguide/aws-resource-route53-hostedzone.md)

- [Tag](../../../../reference/route53/latest/apireference/api-tag.md)
in the _Amazon Route 53 API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HostedZoneFeatures

QueryLoggingConfig

All content copied from https://docs.aws.amazon.com/.
