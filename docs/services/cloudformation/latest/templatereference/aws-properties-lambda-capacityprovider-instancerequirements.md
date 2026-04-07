This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lambda::CapacityProvider InstanceRequirements

Specifications that define the characteristics and constraints for compute instances used by the capacity provider.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AllowedInstanceTypes" : [ String, ... ],
  "Architectures" : [ String, ... ],
  "ExcludedInstanceTypes" : [ String, ... ]
}

```

### YAML

```yaml

  AllowedInstanceTypes:
    - String
  Architectures:
    - String
  ExcludedInstanceTypes:
    - String

```

## Properties

`AllowedInstanceTypes`

A list of EC2 instance types that the capacity provider is allowed to use. If not specified, all compatible instance types are allowed.

_Required_: No

_Type_: Array of String

_Minimum_: `1 | 0`

_Maximum_: `30 | 400`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Architectures`

A list of supported CPU architectures for compute instances. Valid values include `x86_64` and `arm64`.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ExcludedInstanceTypes`

A list of EC2 instance types that the capacity provider should not use, even if they meet other requirements.

_Required_: No

_Type_: Array of String

_Minimum_: `1 | 0`

_Maximum_: `30 | 400`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Examples

### Instance requirement configuration

Configure instance types and architecture.

#### YAML

```yaml

InstanceRequirements:
    AllowedInstanceTypes:
        - c5.4xlarge
    ExcludedInstanceTypes:
        - r6g.xlarge
    Architecture:
        - x86_64

```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CapacityProviderVpcConfig

Tag
