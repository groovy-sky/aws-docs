# Basic workflow concepts in Amazon SWF

###### Note

The concepts in this chapter provide an overview of the Amazon Simple Workflow Service and describe its major features. If you are looking for examples, see [Working with Amazon SWF APIs](https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dg-using-swf-api.html).

Using the Amazon Simple Workflow Service (Amazon SWF), you can implement distributed, asynchronous applications as
_workflows_. Workflows coordinate and manage the execution of activities that can be run
asynchronously across multiple computing devices and that can feature both sequential and parallel processing.

When designing a workflow, you analyze your application to identify its component _tasks_.
In Amazon SWF, these tasks are represented by _activities_. The order in which activities are performed
is determined by the workflow's coordination logic.

###### Example workflow for an e-commerce application

The following figure shows an e-commerce order-processing workflow involving both people and automated processes:

![Illustrative e-commerce example workflow](https://docs.aws.amazon.com/images/amazonswf/latest/developerguide/images/swf-overview-workflow.png)

The e-commerce application workflow starts when a customer places an order, and includes four _tasks_:

1. Verify the order.

2. If the order is valid, charge the customer.

3. If the payment is made, ship the order.

4. If the order is shipped, save the order details.

The tasks in this workflow are _sequential_: an order must be verified before a credit card
can be charged; a credit card must be charged successfully before an order can be shipped; and an order must be
shipped before it can be recorded. Even so, because Amazon SWF supports distributed processes, these tasks can be carried
out in different locations. If the tasks are programmatic in nature, they can also be written in different
programming languages or using different tools.

In addition to sequential processing of tasks, Amazon SWF also supports workflows with parallel processing of tasks.
Parallel tasks are performed at the same time, and may be carried out independently by different applications or
human workers. Your workflow makes decisions about how to proceed once one or more of the parallel tasks have been
completed.

###### Additional concepts

- [Creating a workflow](https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dg-create-workflow.html)

- [Running workflows](https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-run-workflows.html)

- [Workflow history](https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-about-workflow-history.html)

- [Object identifiers](https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-obj-ident.html)

- [Domains](https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-domains.html)

- [Actors](swf-dev-actors.md)

- [Tasks](https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-tasks.html)

- [Task lists](https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-task-lists.html)

- [Workflow execution closure](https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-workflow-exec-closure.html)

- [Workflow execution life cycle](https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-workflow-exec-lifecycle.html)

- [Polling for tasks](https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-comm-proto.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Working in the console

Creating a workflow
