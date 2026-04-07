# Creating a HIT

In the previous section, you set up your AWS account and
installed the programming tools you need to use Amazon Mechanical Turk.
The following topics describe how to write, test, publish,
and manage a Human Intelligence Task (HIT).

###### Topics

- [Workflow](#workflow)

## Workflow

The following procedure gives you an overview of creating, testing, publishing, and
managing a HIT.

Workflow for Requesters

1

Define your HIT.

Construct your question in one of the
[Question and Answer Data Structure](https://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_QuestionAnswerDataArticle.html).
The Question paramater accepts a XML string or HTML string.

2

Create HIT.

Build a new HIT with the
[CreateHIT](https://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_CreateHITOperation.html)
operation or [CreateHITwithHITType](https://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_CreateHITWithHITTypeOperation.html)
operation.

Provide [Title,\
Description, Keywords, and Question Details](https://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/Concepts_HITsArticle.html) as outlined in the
documentation

For code samples demonstrating how to use the _Amazon Mechanical Turk_
_Requester API_, go to [AWSLabs on\
Github.](https://github.com/awslabs/mturk-code-samples)

For language-specific support about using _Amazon Mechanical Turk Requester API_ go to [AWS SDKs](https://aws.amazon.com/tools)

3

Test your HIT.

Publish your HIT on the _Amazon Mechanical Turk Developer_
_Sandbox_. The _Amazon Mechanical Turk_
_Developer Sandbox_ is a simulated environment that
allows you to view your HIT as it would appear to Workers.

For more information about the _Amazon Mechanical_
_Turk Developer Sandbox_ and how to use it, go to the
[Amazon Mechanical\
Turk Developer Sandbox](https://requestersandbox.mturk.com/).

4

Publish your HIT on the Amazon Mechanical Turk production system.

This step makes your HIT available to Workers.

5

Workers accept your HIT and complete the assignment.

You can view the status of your HITs with [AWS Shell](https://github.com/awslabs/aws-shell)
or the [AWS CLI](https://aws.amazon.com/cli).

6

Process the assignment results.

When a Worker completes an assignment, you can view the
results, output the results to a file, and accept or reject the
work. Accepting the work means that you agree to pay the
Worker.

Use [List Assignments for a HIT](https://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_ListAssignmentsForHITOperation.html) to get results
once
Workers have completed the Assignment. Then you can process the
results with the [ApproveAssignment](https://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_ApproveAssignmentOperation.html) operation and [RejectAssignment](https://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_RejectAssignmentOperation.html) operation.

7

Manage your HIT.

You can extend the completion time for your HIT, expire the HIT
early, add additional assignments, modify the HIT properties, or
block Workers whose work does not meet your standards.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Setting Up Accounts and Tools

Implementing Amazon Mechanical Turk
