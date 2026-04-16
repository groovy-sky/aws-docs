---
title: "Amazon S3 log file permissions"
---

# Amazon S3 log file permissions

In addition to the required bucket policies, Amazon S3 uses access control lists (ACLs)
to manage access to the log files created by a flow log. By default, the bucket
owner has `FULL_CONTROL` permissions on each log file. The log delivery
owner, if different from the bucket owner, has no permissions. The log delivery
account has `READ` and `WRITE` permissions. For more
information, see [Access control list\
(ACL) overview](../../../s3/latest/userguide/acl-overview.md) in the _Amazon S3 User Guide_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Required key policy for use with SSE-KMS

Create a flow log that publishes to Amazon S3

All content copied from https://docs.aws.amazon.com/.
