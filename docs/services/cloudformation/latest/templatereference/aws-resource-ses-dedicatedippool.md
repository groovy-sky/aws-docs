This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::DedicatedIpPool

Create a new pool of dedicated IP addresses. A pool can include one or more dedicated
IP addresses that are associated with your AWS account. You can
associate a pool with a configuration set. When you send an email that uses that
configuration set, the message is sent from one of the addresses in the associated
pool.

###### Important

You can't delete dedicated IP pools that have a `STANDARD` scaling mode
with one or more dedicated IP addresses. This constraint doesn't apply to dedicated
IP pools that have a `MANAGED` scaling mode.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SES::DedicatedIpPool",
  "Properties" : {
      "PoolName" : String,
      "ScalingMode" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::SES::DedicatedIpPool
Properties:
  PoolName: String
  ScalingMode: String
  Tags:
    - Tag

```

## Properties

`PoolName`

The name of the dedicated IP pool that the IP address is associated with.

_Required_: No

_Type_: String

_Pattern_: `^[a-z0-9_-]{0,64}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ScalingMode`

The type of scaling mode.

The following options are available:

- `STANDARD` \- The customer controls which IPs are part of the
dedicated IP pool.

- `MANAGED` \- The reputation and number of IPs are automatically
managed by Amazon SES.

The `STANDARD` option is selected by default if no value is
specified.

###### Note

Updating _ScalingMode_ doesn't require a replacement if you're
updating its value from `STANDARD` to `MANAGED`. However,
updating _ScalingMode_ from `MANAGED` to
`STANDARD` is not supported.

_Required_: No

_Type_: String

_Pattern_: `^(STANDARD|MANAGED)$`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`Tags`

An object that defines the tags (keys and values) that you want to associate with the
pool.

_Required_: No

_Type_: Array of [Tag](aws-properties-ses-dedicatedippool-tag.md)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource name.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Tag

All content copied from https://docs.aws.amazon.com/.
