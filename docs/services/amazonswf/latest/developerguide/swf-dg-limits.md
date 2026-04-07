# Amazon SWF Quotas

Amazon SWF places quotas on the sizes of certain workflow parameters, such as on the number of
domains per account and on the size of the workflow execution history. These quotas are designed
to prevent erroneous workflows from consuming all of the resources of the system, but are not
hard limits. If you find that your application is frequently exceeding these quotas, you can
[request a service quota increase](#swf-dg-limits-how-to-increase).

###### Contents

- [General Account Quotas for Amazon SWF](https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dg-limits.html#swf-dg-limits-general)

- [Quotas on Workflow Executions](https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dg-limits.html#swf-dg-limits-workflow-executions)

- [Quotas on Task Executions](https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dg-limits.html#swf-dg-limits-tasks)

- [Amazon SWF throttling quotas](https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dg-limits.html#swf-throttling-limits)

  - [Throttling quotas for all Regions](https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dg-limits.html#throttle-limits-all-regions)

  - [Decision quotas for all Regions](https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dg-limits.html#decision-quota-limits-all-regions)

  - [Workflow-level quotas](https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dg-limits.html#wf-wflow-level-quotas)
- [Requesting a quota increase](https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dg-limits.html#swf-dg-limits-how-to-increase)

## General Account Quotas for Amazon SWF

- **Maximum registered domains** – 100

This quota includes both registered and deprecated domains.

- **Maximum workflow and activity types** – 10,000
each per domain

This quota includes both registered and deprecated types.

- **API call quota** – Beyond infrequent spikes,
applications may be throttled if they make a large number of API calls in a very short
period of time.

- **Maximum request size** – 1 MB per request

This is the _total_ data size per Amazon SWF API request, including the
request header and all other associated request data.

- **Truncated responses for _Count_**
**APIs** – Indicates that an internal quota was reached and that the
response is not the full count.

Some queries will internally reach the 1 MB quota mentioned above before returning a
full response. The following can return a truncated response instead of the full
count.

- [CountClosedWorkflowExecutions](https://docs.aws.amazon.com/amazonswf/latest/apireference/API_CountClosedWorkflowExecutions.html)

- [CountOpenWorkflowExecutions](https://docs.aws.amazon.com/amazonswf/latest/apireference/API_CountOpenWorkflowExecutions.html)

- [CountPendingActivityTasks](https://docs.aws.amazon.com/amazonswf/latest/apireference/API_CountPendingActivityTasks.html)

- [CountPendingDecisionTasks](https://docs.aws.amazon.com/amazonswf/latest/apireference/API_CountPendingDecisionTasks.html)

For each of these, if the `truncated` response is set to true, the count is
less than the full amount. This internal quota can not be increased.

- **Maximum number of tags** – 50 tags per
resource.

Attempting to add tags beyond 50 will result in a 400 error,
`TooManyTagsFault`.

## Quotas on Workflow Executions

- **Maximum open workflow executions** – 100,000 per
domain

This count includes child workflow executions.

- **Maximum workflow execution time** – 1 year. This is a hard quota that can't be changed.

- **Maximum workflow execution history size** –
25,000 events. This is a hard quota that can't be changed.

Best practice is to structure each workflow such that its history does not grow beyond
10,000 events. Because the decider has to fetch the workflow history, a smaller history
allows the decider to complete more quickly. If using the [Flow Framework](https://docs.aws.amazon.com/amazonswf/latest/developerguide/resources.html#aws-flow-framework-documentation), you can use
ContinueAsNew to continue a workflow with a fresh history.

- **Maximum open child workflow executions** – 1,000
per workflow execution

If your use case requires you to go beyond these quotas, you can use features Amazon SWF
provides to continue executions and structure your applications using [child workflow](https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-adv-child-workflows.html) executions. If you find that you
still need a quota increase, see [Requesting a quota increase](#swf-dg-limits-how-to-increase).

## Quotas on Task Executions

- **Maximum pollers per task list** – 1,000 per task
list

You can have a maximum of 1,000 pollers which simultaneously poll a particular task
list. If you go over 1,000, you receive a `LimitExceededException`.

###### Note

While the maximum is 1,000, you might encounter `LimitExceededException` errors well before this quota. This error does not mean your tasks are being delayed. Instead, it means that you have the maximum amount of idle pollers on a task list. Amazon SWF sets this limit to save resources on both the client and server side. Setting the limit prevents an excessive number of pollers from waiting unnecessarily. You can reduce the `LimitExceededException` errors by using multiple task lists to distribute polling.

- **Maximum tasks scheduled per second** – 2,000 per
task list

You can schedule a maximum of 2,000 tasks per second on a particular task list. If
you exceed 2,000, your `ScheduleActivityTask` decisions will fail with
`ACTIVITY_CREATION_RATE_EXCEEDED` error.

###### Note

While the maximum is 2,000, you might encounter
`ACTIVITY_CREATION_RATE_EXCEEDED` errors well before this quota. To reduce
these errors, use multiple task lists to distribute the load.

- **Maximum task execution time** – 1 year
(constrained by workflow execution time maximum)

You can configure [activity timeouts](swf-timeout-types.md) to cause
a timeout event to occur if a particular stage of your [activity task](https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-tasks.html) execution takes too long.

- **Maximum time SWF will keep a task in the queue**
– 1 year (constrained by workflow execution time quota)

You can configure default [activity timeouts](swf-timeout-types.md)
during activity registration that will cause a timeout event to occur if a particular
stage of your [activity task](https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-tasks.html) execution takes too long.
You can also override the default activity timeouts when you schedule an activity task in
your decider code.

- **Maximum open activity tasks** – 1,000 per
workflow execution.

This quota includes both activity tasks that have been scheduled and those being
processed by workers.

- **Maximum open timers** – 1,000 per workflow
execution

- **Maximum input/result data size** – 32,768
characters

This quota affects activity or workflow execution result data, input data when
scheduling activity tasks or workflow executions, and input sent with a [workflow execution signal](https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-adv-signals.html).

- **Maximum decisions in a decision task response** –
varies

Due to the 1 MB quota on the [maximum API request\
size](#swf-dg-limits-general), the number of decisions returned in a single call to `RespondDecisionTaskCompleted` will be limited according to the size of
the data used by each decision, including the size of any input data provided to scheduled
activity tasks or to workflow executions.

## Amazon SWF throttling quotas

In addition to the service quotas described previously, certain Amazon SWF API calls and
decision events are throttled to maintain service bandwidth, using a [token bucket](https://en.wikipedia.org/wiki/Token_bucket) scheme. If your rate
of requests consistently exceeds the rates that are listed here, you can [request a throttle quota increase](#swf-dg-limits-how-to-increase).

The throttling and decision quotas are same across all regions.

### Throttling quotas for all Regions

The following quotas are applicable at individual account-levels. You can also request an increase to the following quotas. For information about doing this, see
[Requesting a quota increase](#swf-dg-limits-how-to-increase).

API nameBucket sizeRefill rate per second`CountClosedWorkflowExecutions`20006`CountOpenWorkflowExecutions`20006`CountPendingActivityTasks`2006`CountPendingDecisionTasks`2006`DeleteActivityType`2006`DeleteWorkflowType`2006`DeprecateActivityType`2006`DeprecateDomain`1006`DeprecateWorkflowType`2006`DescribeActivityType`20006`DescribeDomain`2006`DescribeWorkflowExecution`20006`DescribeWorkflowType`20006`GetWorkflowExecutionHistory`200060`ListActivityTypes`2006`ListClosedWorkflowExecutions`2006`ListDomains`1006`ListOpenWorkflowExecutions`20048`ListTagsForResource`5030`ListWorkflowTypes`2006`PollForActivityTask`2000200`PollForDecisionTask`2000200`RecordActivityTaskHeartbeat`2000160`RegisterActivityType`20060`RegisterDomain`1006`RegisterWorkflowType`20060`RequestCancelWorkflowExecution`200030`RespondActivityTaskCanceled`2000200`RespondActivityTaskCompleted`2000200`RespondActivityTaskFailed`2000200`RespondDecisionTaskCompleted`2000200`SignalWorkflowExecution`200030`StartWorkflowExecution`2000200`TagResource`5030`TerminateWorkflowExecution`200060`UndeprecateActivityType`2006`UndeprecateDomain`1006`UndeprecateWorkflowType`2006`UntagResource`5030

### Decision quotas for all Regions

The following quotas are applicable at individual account-levels. You can also request an increase to the following quotas. For information about doing this, see
[Requesting a quota increase](#swf-dg-limits-how-to-increase).

API nameBucket sizeRefill rate per second`RequestCancelExternalWorkflowExecution`1200120`ScheduleActivityTask`1000200`SignalExternalWorkflowExecution`1200120`StartChildWorkflowExecution`50012`StartTimer`2000200

### Workflow-level quotas

The following quotas are applicable at workflow-levels and can't be increased.

API nameBucket sizeRefill rate per second`GetWorkflowExecutionHistory`400200`SignalWorkflowExecution`10001000`RecordActivityTaskHeartbeat`10001000`RequestCancelWorkflowExecution`200200

## Requesting a quota increase

For more information, see
[AWS service quotas](https://docs.aws.amazon.com/general/latest/gr/aws_service_limits.html) in the
_AWS General Reference_.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Handling errors

Additional resources
