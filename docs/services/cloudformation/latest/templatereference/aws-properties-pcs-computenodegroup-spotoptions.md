This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::PCS::ComputeNodeGroup SpotOptions

Additional configuration when you specify `SPOT` as the
`purchaseOption` for the `CreateComputeNodeGroup` API
action.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AllocationStrategy" : String
}

```

### YAML

```yaml

  AllocationStrategy: String

```

## Properties

`AllocationStrategy`

The Amazon EC2 allocation strategy AWS PCS uses to provision EC2 instances. AWS PCS supports **lowest price**, **capacity optimized**, and **price capacity**
**optimized**. For more information, see [Use allocation strategies to determine how EC2 Fleet or Spot Fleet fulfills Spot and\
On-Demand capacity](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/spot-fleet-allocation-strategy.html) in the _Amazon Elastic Compute Cloud User_
_Guide_. If you don't provide this option, it defaults to **price capacity optimized**.

_Required_: No

_Type_: String

_Allowed values_: `lowest-price | capacity-optimized | price-capacity-optimized`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

SlurmCustomSetting

AWS::PCS::Queue
