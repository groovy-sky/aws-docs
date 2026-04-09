# Managing Amazon Q Business admin controls and guardrails

To manage Amazon Q Business admin controls and guardrails, you can take the
following actions:

###### Note

You can't create or delete guardrail global controls. You can only update
existing global controls in your application environment.

###### Actions

- [Deleting topic controls](#guardrails-update)

- [Getting topic control properties](#topic-control-properties)

## Deleting topic controls

To delete configured chat controls, you can use AWS Management Console or the [DeleteChatControlsConfiguration](../api-reference/api-deletechatcontrolsconfiguration.md)
API operation. The following tabs provide a procedure for the console and code
examples for the AWS CLI.

Console

**To delete topic controls**

1. Sign in to the AWS Management Console and open the Amazon Q Business console.

2. In **Applications**, select the name of
    your application environment from the list of applications.

3. From the left navigation menu, choose
    **Enhancements**, and then choose
    **Guardrails**.

4. In **Guardrails**, from **Topic**
**specific controls**, choose the topic control
    you want to delete, and then choose
    **Delete**.

5. In the dialog box, type `delete` to
    confirm your action.

The console displays a successful deletion message when
    the plugin deletion process is finished.

AWS CLI

**To delete a topic specific**
**control**

```nohighlight

aws qbusiness delete-chat-controls-configuration \
--application-id application-id
```

## Getting topic control properties

To get the details of Amazon Q Business topic controls, you can use
either the AWS Management Console or the [GetChatControlsConfiguration](../api-reference/api-getchatcontrolsconfiguration.md) API
operation. The following tabs provide a procedure for the console and code
examples for the AWS CLI.

Console

**To get configured details for admin controls**
**and guardrails**

1. Sign in to the AWS Management Console and open the Amazon Q Business console.

2. From the Amazon Q Business console, in
    **Applications**, select the name of
    your application environment from the list of applications.

3. From the left navigation menu, choose
    **Enhancements**, and then choose
    **Admin controls and**
**guardrails**.

You will find the details of your configured
    **Global controls** and **Topic**
**specific controls** on the page.

AWS CLI

**To get admin controls and guardrails**
**details**

```nohighlight

aws qbusiness get-chat-control-configuration \
--application-id application-id
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Using topic-level controls

Response customization

All content copied from https://docs.aws.amazon.com/.
