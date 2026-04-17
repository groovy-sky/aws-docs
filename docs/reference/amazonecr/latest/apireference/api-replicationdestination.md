---
title: "ReplicationDestination"
---

# ReplicationDestination

An array of objects representing the destination for a replication rule.

## Contents

**region**

The Region to replicate to.

Type: String

Length Constraints: Minimum length of 2. Maximum length of 25.

Pattern: `[0-9a-z-]{2,25}`

Required: Yes

**registryId**

The AWS account ID of the Amazon ECR private registry to replicate to. When configuring
cross-Region replication within your own registry, specify your own account ID.

Type: String

Pattern: `[0-9]{12}`

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ecr-2015-09-21/ReplicationDestination)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ecr-2015-09-21/ReplicationDestination)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ecr-2015-09-21/ReplicationDestination)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ReplicationConfiguration

ReplicationRule

All content copied from https://docs.aws.amazon.com/.
