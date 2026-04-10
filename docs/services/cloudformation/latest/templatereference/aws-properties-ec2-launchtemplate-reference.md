This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::LaunchTemplate Reference

Specifies an instance family to use as the baseline reference for CPU performance.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "InstanceFamily" : String
}

```

### YAML

```yaml

  InstanceFamily: String

```

## Properties

`InstanceFamily`

The instance family to use as a baseline reference.

###### Note

Ensure that you specify the correct value for the instance family. The instance
family is everything before the period ( `.`) in the instance type name. For
example, in the instance type `c6i.large`, the instance family is
`c6i`, not `c6`. For more information, see [Amazon EC2\
instance type naming conventions](../../../ec2/latest/instancetypes/instance-type-names.md) in _Amazon EC2 Instance_
_Types_.

The following instance families are _not supported_ for performance
protection:

- `c1`

- `g3` \| `g3s`

- `hpc7g`

- `m1` \| `m2`

- `mac1` \| `mac2` \| `mac2-m1ultra` \|
`mac2-m2` \| `mac2-m2pro`

- `p3dn` \| `p4d` \| `p5`

- `t1`

- `u-12tb1` \| `u-18tb1` \| `u-24tb1` \|
`u-3tb1` \| `u-6tb1` \| `u-9tb1` \|
`u7i-12tb` \| `u7in-16tb` \| `u7in-24tb` \|
`u7in-32tb`

If you enable performance protection by specifying a supported instance family, the
returned instance types will exclude the above unsupported instance families.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PrivateIpAdd

SpotOptions

All content copied from https://docs.aws.amazon.com/.
