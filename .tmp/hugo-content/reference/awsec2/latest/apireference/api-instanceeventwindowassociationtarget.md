# InstanceEventWindowAssociationTarget

One or more targets associated with the event window.

## Contents

**DedicatedHostIdSet.N**

The IDs of the Dedicated Hosts associated with the event window.

Type: Array of strings

Required: No

**InstanceIdSet.N**

The IDs of the instances associated with the event window.

Type: Array of strings

Required: No

**TagSet.N**

The instance tags associated with the event window. Any instances associated with the
tags will be associated with the event window.

Note that while you can't create tag keys beginning with `aws:`, you can
specify existing AWS managed tag keys (with the `aws:` prefix) when specifying
them as targets to associate with the event window.

Type: Array of [Tag](api-tag.md) objects

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/instanceeventwindowassociationtarget.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/instanceeventwindowassociationtarget.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/instanceeventwindowassociationtarget.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

InstanceEventWindowAssociationRequest

InstanceEventWindowDisassociationRequest
