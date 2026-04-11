This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Batch::SchedulingPolicy FairsharePolicy

The fair-share scheduling policy details.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ComputeReservation" : Number,
  "ShareDecaySeconds" : Number,
  "ShareDistribution" : [ ShareAttributes, ... ]
}

```

### YAML

```yaml

  ComputeReservation: Number
  ShareDecaySeconds: Number
  ShareDistribution:
    - ShareAttributes

```

## Properties

`ComputeReservation`

A value used to reserve some of the available maximum vCPU for share identifiers that
aren't already used.

The reserved ratio is
`(computeReservation/100)^ActiveFairShares`
where `
                            ActiveFairShares
                        ` is the number of active share
identifiers.

For example, a `computeReservation` value of 50 indicates that AWS Batch reserves
50% of the maximum available vCPU if there's only one share identifier. It reserves 25% if
there are two share identifiers. It reserves 12.5% if there are three share
identifiers. A `computeReservation` value of 25 indicates that AWS Batch should reserve
25% of the maximum available vCPU if there's only one share identifier, 6.25% if there are
two fair share identifiers, and 1.56% if there are three share identifiers.

The minimum value is 0 and the maximum value is 99.

_Required_: No

_Type_: Number

_Minimum_: `0`

_Maximum_: `99`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ShareDecaySeconds`

The amount of time (in seconds) to use to calculate a fair-share percentage for each
share identifier in use. A value of zero (0) indicates the default minimum time window (600 seconds).
The maximum supported value is 604800 (1 week).

The decay allows for more recently run jobs to have more weight than jobs that ran earlier.
Consider adjusting this number if you have jobs that (on average) run longer than ten minutes,
or a large difference in job count or job run times between share identifiers, and the allocation
of resources doesn't meet your needs.

_Required_: No

_Type_: Number

_Minimum_: `0`

_Maximum_: `604800`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ShareDistribution`

An array of `SharedIdentifier` objects that contain the weights for the
share identifiers for the fair-share policy. Share identifiers that aren't included have a
default weight of `1.0`.

_Required_: No

_Type_: Array of [ShareAttributes](aws-properties-batch-schedulingpolicy-shareattributes.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [FairsharePolicy](../../../../reference/batch/latest/apireference/api-fairsharepolicy.md) in the _AWS Batch API Reference_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Batch::SchedulingPolicy

QuotaSharePolicy

All content copied from https://docs.aws.amazon.com/.
