# Managing stacks across accounts and Regions with StackSets

CloudFormation StackSets extends the capability of stacks by allowing you to create, update, or delete
stacks across multiple accounts and AWS Regions with a single operation. Using an
administrator account, you define and manage a CloudFormation template, and use the template as
the basis for provisioning stacks into selected target accounts across specified
AWS Regions.

![A StackSet is a collection of resources in a template, that's deployed across multiple accounts and regions.](https://docs.aws.amazon.com/images/AWSCloudFormation/latest/UserGuide/images/stack_set_conceptual_sv.png)

This section helps you get started using StackSets, and answers common questions about
how to work with and troubleshoot StackSet creation, updates, and deletion.

###### Topics

- [StackSets concepts](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-concepts.html)

- [Prerequisites for using CloudFormation StackSets](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-prereqs.html)

- [Get started with StackSets using a sample template](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-getting-started.html)

- [Create CloudFormation StackSets with self-managed permissions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-getting-started-create-self-managed.html)

- [Create CloudFormation StackSets with service-managed permissions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-orgs-associate-stackset-with-org.html)

- [Enable or disable automatic deployments for StackSets in AWS Organizations](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-orgs-manage-auto-deployment.html)

- [Update CloudFormation StackSets](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-update.html)

- [Add stacks to CloudFormation StackSets](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stackinstances-create.html)

- [Override parameter values on stacks within your CloudFormation StackSet](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stackinstances-override.html)

- [Delete stacks from CloudFormation StackSets](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stackinstances-delete.html)

- [Delete CloudFormation StackSets](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-delete.html)

- [Prevent failed StackSets deployments using target account gates](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-account-gating.html)

- [Choose the Concurrency Mode for CloudFormation StackSets](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/concurrency-mode.html)

- [Performing drift detection on CloudFormation StackSets](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-drift.html)

- [Import stacks into CloudFormation StackSets](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-import.html)

- [Best practices for using CloudFormation StackSets](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-bestpractices.html)

- [CloudFormation StackSets sample templates](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-sampletemplates.html)

- [Troubleshooting CloudFormation StackSets](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-troubleshooting.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Upload local artifacts to an S3
bucket

StackSets concepts
