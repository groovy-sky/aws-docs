# Prerequisites for using CloudFormation StackSets

StackSets extend the functionality of stacks, so you can create, update, or delete
stacks across multiple accounts and Regions with a single operation.

Because StackSets perform stack operations across multiple accounts, before you can
create your first StackSet you need the necessary permissions defined in your
AWS accounts.

You can manage StackSets using _self-managed_ or
_service-managed_ permissions.

- For _self-managed_ StackSets, you must create and manage
IAM roles in each target account and AWS Region. For more information, see [Grant self-managed\
permissions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-prereqs-self-managed.html).

- For _service-managed_ StackSets, you don't need to manually
create and manage IAM roles in each account; AWS handles the role creation and
permissions for you. For more information, see [Activate trusted\
access](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-orgs-activate-trusted-access.html).

###### Topics

- [Prepare to perform StackSet operations in AWS Regions that are disabled by default](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-opt-in-regions.html)

- [Grant self-managed permissions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-prereqs-self-managed.html)

- [Activate trusted access for StackSets with AWS Organizations](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-orgs-activate-trusted-access.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

StackSets concepts

Enable AWS Regions that are
disabled by default
