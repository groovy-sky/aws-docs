This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::NetworkManager::Device AWSLocation

Specifies a location in AWS.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SubnetArn" : String,
  "Zone" : String
}

```

### YAML

```yaml

  SubnetArn: String
  Zone: String

```

## Properties

`SubnetArn`

The Amazon Resource Name (ARN) of the subnet that the device is located in.

_Required_: No

_Type_: String

_Pattern_: `^arn:[^:]{1,63}:ec2:[^:]{0,63}:[^:]{0,63}:subnet\/subnet-[0-9a-f]{8,17}$|^$`

_Minimum_: `0`

_Maximum_: `500`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Zone`

The Zone that the device is located in. Specify the ID of an Availability Zone, Local
Zone, Wavelength Zone, or an Outpost.

_Required_: No

_Type_: String

_Pattern_: `[\s\S]*`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::NetworkManager::Device

Location
