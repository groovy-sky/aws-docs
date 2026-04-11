# InstanceEventWindowAssociationRequest

One or more targets associated with the specified event window. Only one
_type_ of target (instance ID, instance tag, or Dedicated Host ID)
can be associated with an event window.

## Contents

**DedicatedHostId.N**

The IDs of the Dedicated Hosts to associate with the event window.

Type: Array of strings

Required: No

**InstanceId.N**

The IDs of the instances to associate with the event window. If the instance is on a
Dedicated Host, you can't specify the Instance ID parameter; you must use the Dedicated
Host ID parameter.

Type: Array of strings

Required: No

**InstanceTag.N**

The instance tags to associate with the event window. Any instances associated with the
tags will be associated with the event window.

Note that while you can't create tag keys beginning with `aws:`, you can
specify existing AWS managed tag keys (with the `aws:` prefix) when specifying
them as targets to associate with the event window.

Type: Array of [Tag](api-tag.md) objects

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/instanceeventwindowassociationrequest.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/instanceeventwindowassociationrequest.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/instanceeventwindowassociationrequest.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

InstanceEventWindow

InstanceEventWindowAssociationTarget
