# Welcome

AWS CloudFormation allows you to create and manage AWS infrastructure deployments predictably and
repeatedly. You can use CloudFormation to leverage AWS products, such as Amazon Elastic Compute Cloud, Amazon Elastic Block Store,
Amazon Simple Notification Service, Elastic Load Balancing, and Amazon EC2 Auto Scaling to build highly reliable, highly scalable, cost-effective
applications without creating or configuring the underlying AWS infrastructure.

With CloudFormation, you declare all your resources and dependencies in a template file. The
template defines a collection of resources as a single unit called a stack. CloudFormation creates
and deletes all member resources of the stack together and manages all dependencies between
the resources for you.

For more information about CloudFormation, see the [AWS CloudFormation product page](http://aws.amazon.com/cloudformation).

CloudFormation makes use of other AWS products. If you need additional technical information
about a specific AWS product, you can find the product's technical documentation at [docs.aws.amazon.com](../../../../general/index.md).

**Stack actions**

When you use CloudFormation, you manage related resources as a single unit
called a stack. You create, update, and delete a collection of resources by creating,
updating, and deleting stacks. All the resources in a stack are defined by the stack's
template.

[CancelUpdateStack](https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_CancelUpdateStack.html) \| [ContinueUpdateRollback](api-continueupdaterollback.md) \| [CreateStack](api-createstack.md) \| [DeleteStack](api-deletestack.md) \| [DescribeStacks](api-describestacks.md) \| [ListStacks](api-liststacks.md) \| [RollbackStack](https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_RollbackStack.html) \| [UpdateStack](api-updatestack.md)

Stack events: [DescribeStackEvents](https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_DescribeStackEvents.html)

Stack resources: [DescribeStackResource](api-describestackresource.md) \| [DescribeStackResources](api-describestackresources.md) \| [ListStackResources](https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_ListStackResources.html)

Stack drift: [DescribeStackDriftDetectionStatus](https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_DescribeStackDriftDetectionStatus.html) \| [DescribeStackResourceDrifts](https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_DescribeStackResourceDrifts.html) \| [DetectStackDrift](https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_DetectStackDrift.html) \| [DetectStackResourceDrift](https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_DetectStackResourceDrift.html)

Stack operations: [ListExports](https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_ListExports.html) \| [ListImports](https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_ListImports.html) \|
[SignalResource](https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_SignalResource.html) \| [UpdateTerminationProtection](https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_UpdateTerminationProtection.html)

Stack policies: [GetStackPolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_GetStackPolicy.html) \| [SetStackPolicy](api-setstackpolicy.md)

Templates: [GetTemplate](api-gettemplate.md) \| [GetTemplateSummary](https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_GetTemplateSummary.html) \|
[ValidateTemplate](https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_ValidateTemplate.html)

**Change set actions**

If you need to make changes to the running resources in a stack, you
update the stack. Before making changes to your resources, you can generate a change
set, which is summary of your proposed changes. Change sets allow you to see how your
changes might impact your running resources, especially for critical resources, before
implementing them.

[CreateChangeSet](api-createchangeset.md) \| [DeleteChangeSet](https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_DeleteChangeSet.html) \| [DescribeChangeSet](api-describechangeset.md) \| [DescribeChangeSetHooks](https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_DescribeChangeSetHooks.html) \| [ExecuteChangeSet](api-executechangeset.md) \| [ListChangeSets](https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_ListChangeSets.html)

**StackSets actions**

CloudFormation StackSets lets you create a collection, or _stack_
_set_, of stacks that can automatically and safely provision a common set of
AWS resources across multiple AWS accounts and multiple AWS Regions from a single
CloudFormation template. When you create a StackSet, CloudFormation provisions a stack in each
of the specified accounts and AWS Regions by using the supplied CloudFormation template
and parameters. Stack sets let you manage a common set of AWS resources in a selection
of accounts and AWS Regions in a single operation.

[CreateStackSet](https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_CreateStackSet.html) \| [DeleteStackSet](https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_DeleteStackSet.html) \| [DescribeStackSet](https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_DescribeStackSet.html) \| [ListStackSets](https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_ListStackSets.html) \| [UpdateStackSet](https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_UpdateStackSet.html)

Stack instances: [CreateStackInstances](https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_CreateStackInstances.html) \| [DeleteStackInstances](https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_DeleteStackInstances.html) \| [DescribeStackInstance](https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_DescribeStackInstance.html) \| [ListStackInstances](https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_ListStackInstances.html)

StackSet operations: [DescribeStackSetOperation](https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_DescribeStackSetOperation.html) \| [ListStackSetOperations](https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_ListStackSetOperations.html) \| [ListStackSetOperationResults](https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_ListStackSetOperationResults.html) \|
[StopStackSetOperation](https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_StopStackSetOperation.html)

**IaC Generator actions**

Use CloudFormation IaC generator to scan your existing AWS resources and
generate CloudFormation templates from them.

For more information, see [Generating templates from\
existing resources](../../../../services/cloudformation/latest/userguide/generate-iac.md) in the _AWS CloudFormation User Guide_.

Generated templates: [CreateGeneratedTemplate](https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_CreateGeneratedTemplate.html) \| [DeleteGeneratedTemplate](https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_DeleteGeneratedTemplate.html) \| [DescribeGeneratedTemplate](https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_DescribeGeneratedTemplate.html) \|
[GetGeneratedTemplate](https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_GetGeneratedTemplate.html) \| [ListGeneratedTemplates](https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_ListGeneratedTemplates.html) \|
[UpdateGeneratedTemplate](https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_UpdateGeneratedTemplate.html)

Resource scanning: [DescribeResourceScan](https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_DescribeResourceScan.html) \| [ListResourceScans](https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_ListResourceScans.html) \| [ListResourceScanRelatedResources](https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_ListResourceScanRelatedResources.html) \|
[ListResourceScanResources](https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_ListResourceScanResources.html) \| [StartResourceScan](https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_StartResourceScan.html)

**Stack refactoring actions**

Use CloudFormation stack refactoring to move resources between stacks or
reorganize your stack architecture.

For more information, see [Refactoring\
stacks](../../../../services/cloudformation/latest/userguide/stack-refactoring.md) in the _AWS CloudFormation User Guide_.

[CreateStackRefactor](https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_CreateStackRefactor.html) \| [DescribeStackRefactor](https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_DescribeStackRefactor.html) \|
[ExecuteStackRefactor](https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_ExecuteStackRefactor.html) \| [ListStackRefactors](https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_ListStackRefactors.html) \| [ListStackRefactorActions](https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_ListStackRefactorActions.html)

**CloudFormation registry actions**

The CloudFormation registry serves as a centralized hub for managing
extensions that can be used with CloudFormation templates, including CloudFormation
Hooks, resource types, and modules.

For more information, see [Managing extensions with the\
CloudFormation registry](../../../../services/cloudformation/latest/userguide/registry.md) in the _AWS CloudFormation User Guide_.

Discover and activate: [ActivateType](https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_ActivateType.html) \| [DeactivateType](https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_DeactivateType.html) \| [DescribeType](api-describetype.md) \| [ListTypes](https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_ListTypes.html)

Configure: [BatchDescribeTypeConfigurations](https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_BatchDescribeTypeConfigurations.html) \| [SetTypeConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_SetTypeConfiguration.html)

Manage versions: [ListTypeVersions](https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_ListTypeVersions.html) \| [SetTypeDefaultVersion](https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_SetTypeDefaultVersion.html)

Monitor results: [GetHookResult](https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_GetHookResult.html) \| [ListHookResults](https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_ListHookResults.html)

Register privately: [DescribeTypeRegistration](https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_DescribeTypeRegistration.html) \| [DeregisterType](api-deregistertype.md) \| [ListTypeRegistrations](https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_ListTypeRegistrations.html) \| [RegisterType](https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_RegisterType.html)

Publish: [DescribePublisher](https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_DescribePublisher.html) \| [PublishType](https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_PublishType.html) \| [RegisterPublisher](https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_RegisterPublisher.html) \| [TestType](https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_TestType.html)

This document was last published on April 5, 2026.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Actions
