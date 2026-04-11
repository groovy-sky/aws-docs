This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Events::Connection ResourceParameters

The parameters for EventBridge to use when invoking the resource endpoint.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ResourceAssociationArn" : String,
  "ResourceConfigurationArn" : String
}

```

### YAML

```yaml

  ResourceAssociationArn: String
  ResourceConfigurationArn: String

```

## Properties

`ResourceAssociationArn`

For connections to private APIs, the Amazon Resource Name (ARN) of the resource association EventBridge created between the connection and the private API's resource configuration.

###### Note

The value of this property is set by EventBridge. Any value you specify in your template is ignored.

_Required_: No

_Type_: String

_Pattern_: `^arn:[a-z0-9\-]+:vpc-lattice:[a-zA-Z0-9\-]+:\d{12}:servicenetworkresourceassociation/snra-[0-9a-z]{17}$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceConfigurationArn`

The Amazon Resource Name (ARN) of the Amazon VPC Lattice resource configuration for the resource endpoint.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:[a-z0-9f\-]+:vpc-lattice:[a-zA-Z0-9\-]+:\d{12}:resourceconfiguration/rcfg-[0-9a-z]{17}$`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Parameter

AWS::Events::Endpoint

All content copied from https://docs.aws.amazon.com/.
