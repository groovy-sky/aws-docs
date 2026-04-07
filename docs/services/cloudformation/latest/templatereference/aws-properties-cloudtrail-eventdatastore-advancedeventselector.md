This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudTrail::EventDataStore AdvancedEventSelector

Advanced event selectors let you create fine-grained selectors for AWS CloudTrail management, data, and network activity events. They help you control costs by logging only those
events that are important to you. For more information about configuring advanced event selectors, see
the [Logging data events](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.html), [Logging network activity events](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-network-events-with-cloudtrail.html), and [Logging management events](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-management-events-with-cloudtrail.html) topics in the _AWS CloudTrail User Guide_.

You cannot apply both event selectors and advanced event selectors to a trail.

**Supported CloudTrail event record fields for management events**

- `eventCategory` (required)

- `eventSource`

- `readOnly`

The following additional fields are available for event data stores:

- `eventName`

- `eventType`

- `sessionCredentialFromConsole`

- `userIdentity.arn`

**Supported CloudTrail event record fields for data events**

- `eventCategory` (required)

- `eventName`

- `eventSource`

- `eventType`

- `resources.ARN`

- `resources.type` (required)

- `readOnly`

- `sessionCredentialFromConsole`

- `userIdentity.arn`

**Supported CloudTrail event record fields for network activity events**

- `eventCategory` (required)

- `eventSource` (required)

- `eventName`

- `errorCode` \- The only valid value for `errorCode` is `VpceAccessDenied`.

- `vpcEndpointId`

###### Note

For event data stores for CloudTrail Insights events, AWS Config configuration items, Audit Manager evidence, or events outside of AWS, the only supported field is
`eventCategory`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FieldSelectors" : [ AdvancedFieldSelector, ... ],
  "Name" : String
}

```

### YAML

```yaml

  FieldSelectors:
    - AdvancedFieldSelector
  Name: String

```

## Properties

`FieldSelectors`

Contains all selector statements in an advanced event selector.

_Required_: Yes

_Type_: Array of [AdvancedFieldSelector](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudtrail-eventdatastore-advancedfieldselector.html)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

An optional, descriptive name for an advanced event selector, such as "Log data events
for only two S3 buckets".

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::CloudTrail::EventDataStore

AdvancedFieldSelector
