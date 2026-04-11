# Managing subscriptions for an Amazon Q Business application using APIs

Amazon Q Business provides APIs to manage subscriptions in your Amazon Q Business application. You can use these APIs to implement your own
subscription management solution if you create a Amazon Q Business application environment
programmatically.

###### Note

As of Dec 17, 2024, Amazon Q Business will recognize all email addresses as
case-insensitive and recognize subaddresses as equivalent to the original email
address. For example, JohnDoe@example.com, johndoe@example.com, and
johndoe+work@example.com will be considered the same email address. For assistance
with applications or to report a concern, contact Support, sign into the
[AWS Support Center](https://console.aws.amazon.com/support/home) .

API actionAPI descriptionRelevant User Guide topic[CreateSubscription](../api-reference/api-createsubscription.md)Subscribes an IAM Identity Center user or a group to a pricing tier for an Amazon Q Business application

###### Note

If you're using IAM federation to manage user and group
access to Amazon Q Business, subscriptions are created
automatically when a user logs in to their Amazon Q Business application.

- [Creating an IAM Identity Center-integrated\
Amazon Q Business application](create-app.md)

- [Creating an IAM-federated Amazon Q Business application](create-application-iam.md)

[CancelSubscription](../api-reference/api-cancelsubscription.md)Unsubscribes a user or a group from their pricing tier in an Amazon Q Business application

- [Canceling user subscriptions for\
IAM Identity Center-integrated applications](manage-user-subscriptions.md#delete-user-subscriptions)

- [Canceling user subscriptions for\
IAM federated applications](manage-user-subscriptions-iam.md#delete-user-subscriptions-iam)

[ListSubscriptions](../api-reference/api-listsubscriptions.md)Lists all subscriptions created in an Amazon Q Business
application

- [Listing user subscriptions for\
IAM Identity Center-integrated applications](manage-user-subscriptions.md#list-user-subscriptions)

- [Listing user subscriptions for\
IAM federated applications](manage-user-subscriptions-iam.md#list-user-subscriptions-iam)

[UpdateSubscription](../api-reference/api-updatesubscription.md)Updates the pricing tier for an Amazon Q Business
subscription

- [Updating user subscriptions for\
IAM Identity Center-integrated applications](manage-user-subscriptions.md#update-user-subscriptions)

- [Updating user subscriptions for\
IAM federated applications](manage-user-subscriptions-iam.md#update-user-subscriptions-tier-iam)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

User and group management

Plugins

All content copied from https://docs.aws.amazon.com/.
