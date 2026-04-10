# Using service-linked roles for CloudTrail

AWS CloudTrail uses AWS Identity and Access Management (IAM) [service-linked roles](../../../iam/latest/userguide/id-roles-terms-and-concepts.md#iam-term-service-linked-role). A service-linked role is a unique type of IAM role that is
linked directly to CloudTrail. Service-linked roles are predefined by CloudTrail and
include all the permissions that the service requires to call other AWS services on your
behalf.

###### Topics

- [Using roles for creating and managing CloudTrail organization trails and CloudTrail Lake organization event data stores in CloudTrail](using-service-linked-roles-create-slr-for-org-trails.md)

- [Supported Regions for CloudTrail service-linked roles](#slr-regions-create-slr-for-org-trails)

- [Using roles for creating and managing CloudTrail event context in CloudTrail](using-service-linked-roles-create-slr-for-context-management.md)

- [Supported Regions for CloudTrail service-linked roles](#slr-regions-create-slr-for-context-management)

## Supported Regions for CloudTrail service-linked roles

CloudTrail supports using service-linked roles in all of the AWS Regions where CloudTrail and Organizations
are both available. For more information, see [AWS Regions and endpoints](../../../../general/latest/gr/rande.md) in the _AWS General Reference_.

## Supported Regions for CloudTrail service-linked roles

CloudTrail supports using service-linked roles in all of the Regions where CloudTrail and EventBridge are available. For more information, see [AWS Regions and endpoints](../../../../general/latest/gr/rande.md).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Troubleshooting

Organization trails and event data stores role

All content copied from https://docs.aws.amazon.com/.
