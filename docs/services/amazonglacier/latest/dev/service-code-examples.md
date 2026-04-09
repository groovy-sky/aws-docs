**This page is only for existing customers of the Amazon Glacier service using Vaults and the original REST API from 2012.**

If you're looking for archival storage solutions, we recommend using the Amazon Glacier storage classes in Amazon S3, S3 Glacier Instant Retrieval, S3 Glacier Flexible Retrieval, and S3 Glacier Deep Archive. To learn more about these storage options, see [Amazon Glacier storage classes](https://aws.amazon.com/s3/storage-classes/glacier).

Amazon Glacier (original standalone vault-based service) is no longer accepting new customers. Amazon Glacier is a standalone service with its own APIs that stores data in vaults and is distinct from Amazon S3 and the Amazon S3 Glacier storage classes. Your existing data will remain secure and accessible in Amazon Glacier indefinitely. No migration is required. For low-cost, long-term archival storage, AWS recommends the [Amazon S3 Glacier storage classes](https://aws.amazon.com/s3/storage-classes/glacier), which deliver a superior customer experience with S3 bucket-based APIs, full AWS Region availability, lower costs, and AWS service integration. If you want enhanced capabilities, consider migrating to Amazon S3 Glacier storage classes by using our [AWS Solutions Guidance for transferring data from Amazon Glacier vaults to Amazon S3 Glacier storage classes](https://aws.amazon.com/solutions/guidance/data-transfer-from-amazon-s3-glacier-vaults-to-amazon-s3).

# Code examples for Amazon Glacier using AWS SDKs

The following code examples show how to use Amazon Glacier with an AWS software development kit (SDK).

_Actions_ are code excerpts from larger programs and must be run in context. While actions show you how to call individual service functions, you can see actions in context in their related scenarios.

_Scenarios_ are code examples that show you how to accomplish specific tasks by calling multiple functions within a service or combined with other AWS services.

For a complete list of AWS SDK developer guides and code examples, see
[Using Amazon Glacier with an AWS SDK](../../../../reference/amazonglacier/latest/dev/sdk-general-information-section.md).
This topic also includes information about getting started and details about previous SDK versions.

###### Code examples

- [Basics](service-code-examples-basics.md)

  - [Hello Amazon Glacier](example-glacier-hello-section.md)

  - [Actions](service-code-examples-actions.md)

    - [AddTagsToVault](example-glacier-addtagstovault-section.md)

    - [CreateVault](example-glacier-createvault-section.md)

    - [DeleteArchive](example-glacier-deletearchive-section.md)

    - [DeleteVault](example-glacier-deletevault-section.md)

    - [DeleteVaultNotifications](example-glacier-deletevaultnotifications-section.md)

    - [DescribeJob](example-glacier-describejob-section.md)

    - [DescribeVault](example-glacier-describevault-section.md)

    - [GetJobOutput](example-glacier-getjoboutput-section.md)

    - [GetVaultNotifications](example-glacier-getvaultnotifications-section.md)

    - [InitiateJob](example-glacier-initiatejob-section.md)

    - [ListJobs](example-glacier-listjobs-section.md)

    - [ListTagsForVault](example-glacier-listtagsforvault-section.md)

    - [ListVaults](example-glacier-listvaults-section.md)

    - [SetVaultNotifications](example-glacier-setvaultnotifications-section.md)

    - [UploadArchive](example-glacier-uploadarchive-section.md)

    - [UploadMultipartPart](example-glacier-uploadmultipartpart-section.md)
- [Scenarios](service-code-examples-scenarios.md)

  - [Archive a file, get notifications, and initiate a job](example-glacier-usage-uploadnotifyinitiate-section.md)

  - [Get archive content and delete the archive](example-glacier-usage-retrievedelete-section.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Using the AWS SDK for .NET

Basics

All content copied from https://docs.aws.amazon.com/.
