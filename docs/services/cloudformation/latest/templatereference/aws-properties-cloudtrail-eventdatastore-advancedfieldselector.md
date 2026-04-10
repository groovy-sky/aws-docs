This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudTrail::EventDataStore AdvancedFieldSelector

A single selector statement in an advanced event selector.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EndsWith" : [ String, ... ],
  "Equals" : [ String, ... ],
  "Field" : String,
  "NotEndsWith" : [ String, ... ],
  "NotEquals" : [ String, ... ],
  "NotStartsWith" : [ String, ... ],
  "StartsWith" : [ String, ... ]
}

```

### YAML

```yaml

  EndsWith:
    - String
  Equals:
    - String
  Field: String
  NotEndsWith:
    - String
  NotEquals:
    - String
  NotStartsWith:
    - String
  StartsWith:
    - String

```

## Properties

`EndsWith`

An operator that includes events that match the last few characters of the event record
field specified as the value of `Field`.

_Required_: No

_Type_: Array of String

_Maximum_: `2048`

_Minimum_: `1 | 1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Equals`

An operator that includes events that match the exact value of the event record field
specified as the value of `Field`. This is the only valid operator that you can
use with the `readOnly`, `eventCategory`, and
`resources.type` fields.

_Required_: No

_Type_: Array of String

_Maximum_: `2048`

_Minimum_: `1 | 1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Field`

A field in a CloudTrail event record on which to filter events to be logged. For
event data stores for CloudTrail Insights events, AWS Config configuration items, Audit Manager evidence, or events outside of AWS, the field is used only for
selecting events as filtering is not supported.

For CloudTrail management events, supported fields include
`eventCategory` (required), `eventSource`, and
`readOnly`. The following additional fields are available for event data
stores: `eventName`, `eventType`,
`sessionCredentialFromConsole`, and `userIdentity.arn`.

For CloudTrail data events, supported fields include `eventCategory`
(required), `eventName`, `eventSource`, `eventType`, `resources.type` (required),
`readOnly`, `resources.ARN`, `sessionCredentialFromConsole`, and `userIdentity.arn`.

For CloudTrail network activity events, supported fields include `eventCategory` (required), `eventSource` (required), `eventName`,
`errorCode`, and `vpcEndpointId`.

For event data stores for CloudTrail Insights events, AWS Config configuration items, Audit Manager evidence, or events outside of AWS, the only supported field is
`eventCategory`.

###### Note

Selectors don't support the use of wildcards like `*` . To match multiple values with a single condition,
you may use `StartsWith`, `EndsWith`, `NotStartsWith`, or `NotEndsWith` to explicitly match the beginning or end of the event field.

- **`readOnly`** \- This is an optional field that is only used for management events and data events. This field can be set to
`Equals` with a value of `true` or `false`. If you do
not add this field, CloudTrail logs both `read` and
`write` events. A value of `true` logs only
`read` events. A value of `false` logs only
`write` events.

- **`eventSource`** \- This field is only used for management events, data events, and network activity events.

For management events for trails, this is an optional field that can be set to `NotEquals` `kms.amazonaws.com` to exclude KMS management events, or `NotEquals` `rdsdata.amazonaws.com` to exclude RDS management events.

For data events for trails, this is an optional field that you can use to include or
exclude any event source and can use any operator.

For management and data events for event data stores, this is an optional field that you can use to include or
exclude any event source and can use any operator.

For network activity events, this is a required field that only uses the
`Equals` operator. Set this field to the event source for which you want to
log network activity events. If you want to log network activity events for multiple
event sources, you must create a separate field selector for each event
source. For a list of services supporting network activity events, see [Logging network activity events](../../../awscloudtrail/latest/userguide/logging-network-events-with-cloudtrail.md) in the _AWS CloudTrail User Guide_.

- **`eventName`** \- This is an optional field that is only used for data events, management events (for event data stores only), and network activity events. You can use any operator with
`eventName`. You can use it to ﬁlter in or ﬁlter out specific events. You can have
multiple values for this ﬁeld, separated by commas.

- **`eventCategory`** \- This field is required and
must be set to `Equals`.

- For CloudTrail management events, the value
must be `Management`.

- For CloudTrail data events, the value
must be `Data`.

- For CloudTrail network activity events, the value
must be `NetworkActivity`.

The following are used only for event data stores:

- For CloudTrail Insights events, the value
must be `Insight`.

- For AWS Config
configuration items, the value must be `ConfigurationItem`.

- For Audit Manager evidence, the value must be `Evidence`.

- For events outside of AWS, the value must be `ActivityAuditLog`.

- **`eventType`** \- For event data stores, this is an optional
field available for event data stores to filter management and
data events on the event type. For trails, this is an optional field to filter data events on the event type. For information about available event types, see
[CloudTrail record contents](../../../awscloudtrail/latest/userguide/cloudtrail-event-reference-record-contents.md#ct-event-type) in the _AWS CloudTrail user_
_guide_.

- **`errorCode`** \- This ﬁeld is only used to filter CloudTrail network activity events
and is optional. This is the error code to filter on. Currently, the only valid `errorCode` is `VpceAccessDenied`.
`errorCode` can only use the `Equals` operator.

- **`sessionCredentialFromConsole`** \- For event data stores, this
is an optional field used to filter management and data events based on whether the events originated from an AWS Management Console session. For trails, this is an optional field used to filter data events. `sessionCredentialFromConsole` can only use the
`Equals` and `NotEquals` operators.

- **`resources.type`** \- This ﬁeld is
required for CloudTrail data events. `resources.type` can only
use the `Equals` operator.

For a list of available resource types for data events, see [Data events](../../../awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.md#logging-data-events) in the _AWS CloudTrail User Guide_.

You can have only one `resources.type` ﬁeld per selector. To log events on more than one resource type, add another selector.

- **`resources.ARN`** \- The `resources.ARN` is an optional field for
data events. You can use any
operator with `resources.ARN`, but if you use `Equals` or
`NotEquals`, the value must exactly match the ARN of a valid resource
of the type you've speciﬁed in the template as the value of resources.type. To log all data events for all objects in a specific S3 bucket,
use the `StartsWith` operator, and include only the bucket ARN as the matching value.

For more information about the ARN formats of data event
resources, see [Actions, resources, and condition\
keys for AWS services](../../../service-authorization/latest/reference/reference-policies-actions-resources-contextkeys.md) in the _Service Authorization Reference_.

###### Note

You can't use the `resources.ARN` field to filter resource types that do not have ARNs.

- **`userIdentity.arn`** \- For event data stores, this is an
optional field used to filter management and data events for actions taken by specific IAM identities. For trails,
this is an optional field used to filter data events. You can use any operator with
`userIdentity.arn`. For more information on the userIdentity element,
see [CloudTrail userIdentity element](../../../awscloudtrail/latest/userguide/cloudtrail-event-reference-user-identity.md) in the _AWS CloudTrail User Guide_.

- **`vpcEndpointId`** \- This ﬁeld is only used to filter CloudTrail network activity events
and is optional. This field identifies the VPC endpoint that the request passed through. You can use any operator with `vpcEndpointId`.

_Required_: Yes

_Type_: String

_Pattern_: `([\w|\d|\.|_]+)`

_Minimum_: `1`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NotEndsWith`

An operator that excludes events that match the last few characters of the event record
field specified as the value of `Field`.

_Required_: No

_Type_: Array of String

_Maximum_: `2048`

_Minimum_: `1 | 1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NotEquals`

An operator that excludes events that match the exact value of the event record field
specified as the value of `Field`.

_Required_: No

_Type_: Array of String

_Maximum_: `2048`

_Minimum_: `1 | 1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NotStartsWith`

An operator that excludes events that match the first few characters of the event
record field specified as the value of `Field`.

_Required_: No

_Type_: Array of String

_Maximum_: `2048`

_Minimum_: `1 | 1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StartsWith`

An operator that includes events that match the first few characters of the event record
field specified as the value of `Field`.

_Required_: No

_Type_: Array of String

_Maximum_: `2048`

_Minimum_: `1 | 1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AdvancedEventSelector

ContextKeySelector

All content copied from https://docs.aws.amazon.com/.
