# Welcome

Amazon Augmented AI (Amazon A2I) adds the benefit of human judgment to any machine learning
application. When an AI application can't evaluate data with a high degree of confidence,
human reviewers can take over. This human review is called a human review workflow. To create
and start a human review workflow, you need three resources: a _worker task_
_template_, a _flow definition_, and a _human_
_loop_.

For information about these resources and prerequisites for using Amazon A2I, see [Get Started with\
Amazon Augmented AI](../../../../services/sagemaker/latest/dg/a2i-getting-started.md) in the Amazon SageMaker Developer Guide.

This API reference includes information about API actions and data types that you can use
to interact with Amazon A2I programmatically. Use this guide to:

- Start a human loop with the `StartHumanLoop` operation when using
Amazon A2I with a _custom task type_. To learn more about the
difference between custom and built-in task types, see [Use Task Types](../../../../services/sagemaker/latest/dg/a2i-task-types-general.md). To learn
how to start a human loop using this API, see [Create and Start a Human Loop for a Custom Task Type](../../../../services/sagemaker/latest/dg/a2i-start-human-loop-a2i-instructions-starthumanloop.md) in the
Amazon SageMaker Developer Guide.

- Manage your human loops. You can list all human loops that you have created, describe
individual human loops, and stop and delete human loops. To learn more, see [Monitor and Manage Your Human Loop](../../../../services/sagemaker/latest/dg/a2i-monitor-humanloop-results.md) in the Amazon SageMaker Developer Guide.

Amazon A2I integrates APIs from various AWS services to create and start human review
workflows for those services. To learn how Amazon A2I uses these APIs, see [Use APIs in\
Amazon A2I](../../../../services/sagemaker/latest/dg/a2i-api-references.md) in the Amazon SageMaker Developer Guide.

This document was last published on April 5, 2026.

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

Actions
