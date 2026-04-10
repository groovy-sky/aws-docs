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

[CancelUpdateStack](api-cancelupdatestack.md) \| [ContinueUpdateRollback](api-continueupdaterollback.md) \| [CreateStack](api-createstack.md) \| [DeleteStack](api-deletestack.md) \| [DescribeStacks](api-describestacks.md) \| [ListStacks](api-liststacks.md) \| [RollbackStack](api-rollbackstack.md) \| [UpdateStack](api-updatestack.md)

Stack events: [DescribeStackEvents](api-describestackevents.md)

Stack resources: [DescribeStackResource](api-describestackresource.md) \| [DescribeStackResources](api-describestackresources.md) \| [ListStackResources](api-liststackresources.md)

Stack drift: [DescribeStackDriftDetectionStatus](api-describestackdriftdetectionstatus.md) \| [DescribeStackResourceDrifts](api-describestackresourcedrifts.md) \| [DetectStackDrift](api-detectstackdrift.md) \| [DetectStackResourceDrift](api-detectstackresourcedrift.md)

Stack operations: [ListExports](api-listexports.md) \| [ListImports](api-listimports.md) \|
[SignalResource](api-signalresource.md) \| [UpdateTerminationProtection](api-updateterminationprotection.md)

Stack policies: [GetStackPolicy](api-getstackpolicy.md) \| [SetStackPolicy](api-setstackpolicy.md)

Templates: [GetTemplate](api-gettemplate.md) \| [GetTemplateSummary](api-gettemplatesummary.md) \|
[ValidateTemplate](api-validatetemplate.md)

**Change set actions**

If you need to make changes to the running resources in a stack, you
update the stack. Before making changes to your resources, you can generate a change
set, which is summary of your proposed changes. Change sets allow you to see how your
changes might impact your running resources, especially for critical resources, before
implementing them.

[CreateChangeSet](api-createchangeset.md) \| [DeleteChangeSet](api-deletechangeset.md) \| [DescribeChangeSet](api-describechangeset.md) \| [DescribeChangeSetHooks](api-describechangesethooks.md) \| [ExecuteChangeSet](api-executechangeset.md) \| [ListChangeSets](api-listchangesets.md)

**StackSets actions**

CloudFormation StackSets lets you create a collection, or _stack_
_set_, of stacks that can automatically and safely provision a common set of
AWS resources across multiple AWS accounts and multiple AWS Regions from a single
CloudFormation template. When you create a StackSet, CloudFormation provisions a stack in each
of the specified accounts and AWS Regions by using the supplied CloudFormation template
and parameters. Stack sets let you manage a common set of AWS resources in a selection
of accounts and AWS Regions in a single operation.

[CreateStackSet](api-createstackset.md) \| [DeleteStackSet](api-deletestackset.md) \| [DescribeStackSet](api-describestackset.md) \| [ListStackSets](api-liststacksets.md) \| [UpdateStackSet](api-updatestackset.md)

Stack instances: [CreateStackInstances](api-createstackinstances.md) \| [DeleteStackInstances](api-deletestackinstances.md) \| [DescribeStackInstance](api-describestackinstance.md) \| [ListStackInstances](api-liststackinstances.md)

StackSet operations: [DescribeStackSetOperation](api-describestacksetoperation.md) \| [ListStackSetOperations](api-liststacksetoperations.md) \| [ListStackSetOperationResults](api-liststacksetoperationresults.md) \|
[StopStackSetOperation](api-stopstacksetoperation.md)

**IaC Generator actions**

Use CloudFormation IaC generator to scan your existing AWS resources and
generate CloudFormation templates from them.

For more information, see [Generating templates from\
existing resources](../../../../services/cloudformation/latest/userguide/generate-iac.md) in the _AWS CloudFormation User Guide_.

Generated templates: [CreateGeneratedTemplate](api-creategeneratedtemplate.md) \| [DeleteGeneratedTemplate](api-deletegeneratedtemplate.md) \| [DescribeGeneratedTemplate](api-describegeneratedtemplate.md) \|
[GetGeneratedTemplate](api-getgeneratedtemplate.md) \| [ListGeneratedTemplates](api-listgeneratedtemplates.md) \|
[UpdateGeneratedTemplate](api-updategeneratedtemplate.md)

Resource scanning: [DescribeResourceScan](api-describeresourcescan.md) \| [ListResourceScans](api-listresourcescans.md) \| [ListResourceScanRelatedResources](api-listresourcescanrelatedresources.md) \|
[ListResourceScanResources](api-listresourcescanresources.md) \| [StartResourceScan](api-startresourcescan.md)

**Stack refactoring actions**

Use CloudFormation stack refactoring to move resources between stacks or
reorganize your stack architecture.

For more information, see [Refactoring\
stacks](../../../../services/cloudformation/latest/userguide/stack-refactoring.md) in the _AWS CloudFormation User Guide_.

[CreateStackRefactor](api-createstackrefactor.md) \| [DescribeStackRefactor](api-describestackrefactor.md) \|
[ExecuteStackRefactor](api-executestackrefactor.md) \| [ListStackRefactors](api-liststackrefactors.md) \| [ListStackRefactorActions](api-liststackrefactoractions.md)

**CloudFormation registry actions**

The CloudFormation registry serves as a centralized hub for managing
extensions that can be used with CloudFormation templates, including CloudFormation
Hooks, resource types, and modules.

For more information, see [Managing extensions with the\
CloudFormation registry](../../../../services/cloudformation/latest/userguide/registry.md) in the _AWS CloudFormation User Guide_.

Discover and activate: [ActivateType](api-activatetype.md) \| [DeactivateType](api-deactivatetype.md) \| [DescribeType](api-describetype.md) \| [ListTypes](api-listtypes.md)

Configure: [BatchDescribeTypeConfigurations](api-batchdescribetypeconfigurations.md) \| [SetTypeConfiguration](api-settypeconfiguration.md)

Manage versions: [ListTypeVersions](api-listtypeversions.md) \| [SetTypeDefaultVersion](api-settypedefaultversion.md)

Monitor results: [GetHookResult](api-gethookresult.md) \| [ListHookResults](api-listhookresults.md)

Register privately: [DescribeTypeRegistration](api-describetyperegistration.md) \| [DeregisterType](api-deregistertype.md) \| [ListTypeRegistrations](api-listtyperegistrations.md) \| [RegisterType](api-registertype.md)

Publish: [DescribePublisher](api-describepublisher.md) \| [PublishType](api-publishtype.md) \| [RegisterPublisher](api-registerpublisher.md) \| [TestType](api-testtype.md)

This document was last published on April 9, 2026.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Actions

All content copied from https://docs.aws.amazon.com/.
