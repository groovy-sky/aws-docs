# InstanceEventWindowDisassociationRequest

The targets to disassociate from the specified event window.

## Contents

**DedicatedHostId.N**

The IDs of the Dedicated Hosts to disassociate from the event window.

Type: Array of strings

Required: No

**InstanceId.N**

The IDs of the instances to disassociate from the event window.

Type: Array of strings

Required: No

**InstanceTag.N**

The instance tags to disassociate from the event window. Any instances associated with
the tags will be disassociated from the event window.

Type: Array of [Tag](api-tag.md) objects

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/instanceeventwindowdisassociationrequest.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/instanceeventwindowdisassociationrequest.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/instanceeventwindowdisassociationrequest.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

InstanceEventWindowAssociationTarget

InstanceEventWindowStateChange
