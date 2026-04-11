This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecurityHub::FindingAggregator

The `AWS::SecurityHub::FindingAggregator` resource enables cross-Region aggregation.
When cross-Region aggregation is enabled, you can aggregate findings, finding updates, insights, control compliance statuses, and
security scores from one or more linked Regions to a single aggregation Region. You can then view and manage all of this data from the
aggregation Region. For more details about cross-Region aggregation, see [Cross-Region aggregation](../../../securityhub/latest/userguide/finding-aggregation.md) in the
_AWS Security Hub CSPM User Guide_

This resource must be created in the Region that you want to designate as your aggregation Region.

Cross-Region aggregation is also a prerequisite for using [central configuration](../../../securityhub/latest/userguide/central-configuration-intro.md) in Security Hub CSPM.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SecurityHub::FindingAggregator",
  "Properties" : {
      "RegionLinkingMode" : String,
      "Regions" : [ String, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::SecurityHub::FindingAggregator
Properties:
  RegionLinkingMode: String
  Regions:
    - String

```

## Properties

`RegionLinkingMode`

Indicates whether to aggregate findings from all of the available Regions in the current partition. Also determines whether to automatically aggregate findings from new Regions as Security Hub supports them and you opt into them.

The selected option also determines how to use the Regions provided in the Regions list.

In CloudFormation, the options for this property are as follows:

- `ALL_REGIONS` \- Indicates to aggregate findings from all of the Regions where Security Hub is enabled. When you choose this option, Security Hub also automatically aggregates findings from new Regions as Security Hub supports them and you opt into them.

- `ALL_REGIONS_EXCEPT_SPECIFIED` \- Indicates to aggregate findings from all of the Regions where Security Hub is enabled, except for the Regions listed in the `Regions` parameter. When you choose this option, Security Hub also automatically aggregates findings from new Regions as Security Hub supports them and you opt into them.

- `SPECIFIED_REGIONS` \- Indicates to aggregate findings only from the Regions listed in the `Regions` parameter. Security Hub does not automatically aggregate findings from new Regions.

_Required_: Yes

_Type_: String

_Allowed values_: `ALL_REGIONS | ALL_REGIONS_EXCEPT_SPECIFIED | SPECIFIED_REGIONS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Regions`

If `RegionLinkingMode` is `ALL_REGIONS_EXCEPT_SPECIFIED`, then this is a space-separated list of Regions that do not aggregate findings to the aggregation Region.

If `RegionLinkingMode` is `SPECIFIED_REGIONS`, then this is a space-separated list of Regions that do aggregate findings to the aggregation Region.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the finding aggregator. For example, `arn:aws:securityhub:us-east-1:123456789012:finding-aggregator/a1b2c3d4-5678-90ab-cdef-EXAMPLE11111`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`FindingAggregationRegion`

The home Region. Findings generated in linked Regions are replicated and sent to the home Region.

`FindingAggregatorArn`

The ARN of the finding aggregator. You use the finding aggregator ARN to retrieve details for, update, and delete the finding aggregator.

## Examples

### Configuring Security Hub CSPM cross-Region aggregation

The following example configures cross-Region aggregation. The region in which
the resource is created is the aggregation Region. In this example, `us-east-2` and `us-west-1`
contribute data to the aggregation Region.

#### JSON

```json

{
	"Description": "Example template to configure Security Hub cross-Region aggregation",
	"Resources": {
		"SecurityHubFindingAggregator": {
			"Type": "AWS::SecurityHub::FindingAggregator",
			"Properties": {
				"RegionLinkingMode": "SPECIFIED_REGIONS",
				"Regions": ["us-west-1", "us-east-2"]
			}
		}
	}
}
```

#### YAML

```yaml

Description: Example template to configure Security Hub cross-Region aggregation
Resources:
  SecurityHubFindingAggregator:
    Type: 'AWS::SecurityHub::FindingAggregator'
    Properties:
      RegionLinkingMode: 'SPECIFIED_REGIONS'
      Regions:
        - "us-west-1"
        - "us-east-2"
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::SecurityHub::DelegatedAdmin

AWS::SecurityHub::Hub

All content copied from https://docs.aws.amazon.com/.
