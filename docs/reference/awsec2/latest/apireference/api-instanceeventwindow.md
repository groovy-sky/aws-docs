# InstanceEventWindow

The event window.

## Contents

**associationTarget**

One or more targets associated with the event window.

Type: [InstanceEventWindowAssociationTarget](api-instanceeventwindowassociationtarget.md) object

Required: No

**cronExpression**

The cron expression defined for the event window.

Type: String

Required: No

**instanceEventWindowId**

The ID of the event window.

Type: String

Required: No

**name**

The name of the event window.

Type: String

Required: No

**state**

The current state of the event window.

Type: String

Valid Values: `creating | deleting | active | deleted`

Required: No

**TagSet.N**

The instance tags associated with the event window.

Type: Array of [Tag](api-tag.md) objects

Required: No

**TimeRangeSet.N**

One or more time ranges defined for the event window.

Type: Array of [InstanceEventWindowTimeRange](api-instanceeventwindowtimerange.md) objects

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/instanceeventwindow.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/instanceeventwindow.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/instanceeventwindow.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

InstanceCreditSpecificationRequest

InstanceEventWindowAssociationRequest
