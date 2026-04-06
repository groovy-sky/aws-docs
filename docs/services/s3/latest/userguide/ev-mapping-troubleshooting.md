# Amazon EventBridge mapping and troubleshooting

The following table describes how Amazon S3 event types are mapped to Amazon EventBridge event
types.

S3 event type Amazon EventBridge detail type

[ObjectCreated:Put](../api/api-putobject.md)

[ObjectCreated:Post](../api/restobjectpost.md)

[ObjectCreated:Copy](../api/api-copyobject.md)

[ObjectCreated:CompleteMultipartUpload](../api/api-completemultipartupload.md)

Object Created

ObjectRemoved:Delete

ObjectRemoved:DeleteMarkerCreated

LifecycleExpiration:Delete

LifecycleExpiration:DeleteMarkerCreated

Object Deleted

[ObjectRestore:Post](../api/api-restoreobject.md)

Object Restore Initiated

ObjectRestore:Completed

Object Restore Completed

ObjectRestore:Delete

Object Restore Expired

LifecycleTransition

Object Storage Class Changed

IntelligentTiering

Object Access Tier Changed

[ObjectTagging:Put](../api/api-putobjecttagging.md)

Object Tags Added

[ObjectTagging:Delete](../api/api-deleteobjecttagging.md)

Object Tags Deleted

[ObjectAcl:Put](../api/api-putobjectacl.md)

Object ACL Updated

## Amazon EventBridge troubleshooting

For information about how to troubleshoot EventBridge, see [Troubleshooting Amazon EventBridge](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-troubleshooting.html) in
the _Amazon EventBridge User Guide_.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

EventBridge event message structure

Monitoring your storage activity and usage with S3 Storage Lens
