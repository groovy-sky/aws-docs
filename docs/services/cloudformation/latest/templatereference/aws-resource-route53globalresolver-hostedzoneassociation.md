This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Route53GlobalResolver::HostedZoneAssociation

Associates a Route 53 private hosted zone with a Route 53 Global Resolver resource. This allows the
resolver to resolve DNS queries for the private hosted zone from anywhere globally.

###### Important

Route 53 Global Resolver is a global service that supports resolvers in multiple AWS Regions but you must specify the
US East (Ohio) Region to create, update, or otherwise work with Route 53 Global Resolver resources. That is, for example,
specify
`--region us-east-2`
on AWS CLI commands.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Route53GlobalResolver::HostedZoneAssociation",
  "Properties" : {
      "HostedZoneId" : String,
      "Name" : String,
      "ResourceArn" : String
    }
}

```

### YAML

```yaml

Type: AWS::Route53GlobalResolver::HostedZoneAssociation
Properties:
  HostedZoneId: String
  Name: String
  ResourceArn: String

```

## Properties

`HostedZoneId`

The ID of the hosted zone.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `32`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the hosted zone association.

_Required_: Yes

_Type_: String

_Pattern_: `(?!^[0-9]+$)([a-zA-Z0-9-_' ']+)`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceArn`

The Amazon Resource Name (ARN) of the resource associated with the hosted zone.

_Required_: Yes

_Type_: String

_Pattern_: `arn:[-.a-z0-9]{1,63}:[-.a-z0-9]{1,63}:[-.a-z0-9]{0,63}:[-.a-z0-9]{0,63}:[^/].{0,1023}`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

### Fn::GetAtt

`CreatedAt`

The date and time when the hosted zone association was created.

`HostedZoneAssociationId`

ID of the private hosted zone association.

`HostedZoneName`

The name of the hosted zone.

`Status`

The current status of the hosted zone association.

`UpdatedAt`

The date and time when the hosted zone association was last updated.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Next
