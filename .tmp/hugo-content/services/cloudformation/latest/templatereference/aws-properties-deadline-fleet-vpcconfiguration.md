This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Deadline::Fleet VpcConfiguration

The configuration options for a service managed fleet's VPC.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ResourceConfigurationArns" : [ String, ... ]
}

```

### YAML

```yaml

  ResourceConfigurationArns:
    - String

```

## Properties

`ResourceConfigurationArns`

The ARNs of the VPC Lattice resource configurations attached to the fleet.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VCpuCountRange

AWS::Deadline::LicenseEndpoint

All content copied from https://docs.aws.amazon.com/.
