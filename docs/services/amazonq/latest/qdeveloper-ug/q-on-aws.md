# Using Amazon Q Developer on AWS apps and websites

Use Amazon Q Developer in the AWS Management Console, AWS Console Mobile Application, AWS Marketing website, AWS Documentation website,
and supported chat applications to ask questions about AWS. You can ask Amazon Q
about AWS architecture, best practices, support, and documentation. Amazon Q can also help
with code that you're writing with the AWS SDKs and AWS Command Line Interface (AWS CLI).

In the AWS Management Console, you can ask Amazon Q about your AWS resources and costs, contact Support
directly, and diagnose common console errors.

To quickly provide access to features of Amazon Q Developer on AWS, attach the `
            AmazonQDeveloperAccess`
AWS managed policy to the IAM identity using Amazon Q. For permissions needed for specific
features, see the topic for the feature you want to use.

###### Topics

- [Authenticating to your Amazon Q Developer Pro subscription](#qdevpro-authentication)

- [Chatting with Amazon Q Developer about AWS](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/chat-with-q.html)

- [Using Amazon Q Developer plugins](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/plugins.html)

- [Automating AWS services with Amazon Q Developer Console-to-Code](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/console-to-code.html)

- [Diagnosing common errors in the console with Amazon Q Developer](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/diagnose-console-errors.html)

- [Using Amazon Q Developer to chat with Support](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/support-chat.html)

## Authenticating to your Amazon Q Developer Pro subscription

To access Amazon Q at the Free tier, sign in to the AWS Management Console. Any Free tier features
are available as long as you have the required permissions.

To access Amazon Q at the Pro tier, sign to the console with IAM Identity Center. When you sign in
with IAM Identity Center, including authenticating through an external identity provider that is
connected to IAM Identity Center, you will automatically have access to the Pro tier if your IAM Identity Center
identity is subscribed to Amazon Q Developer Pro.

For more information on the Amazon Q Developer Pro tier, see [Tiers of service for Q Developer – Free and Pro](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/q-tiers.html).

###### Note

If you see an error message that starts with, `Your account has not been
                    configured to use an Amazon Q subscription`, see [Troubleshooting Amazon Q Developer Pro subscriptions](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/q-admin-setup-subscribe-troubleshooting.html) for troubleshooting
tips.

If you sign in to the AWS console with IAM or federation with IAM, then you will
be prompted to authenticate with IAM Identity Center when you reach a Free tier limit or attempt to
use a feature only available at the Pro tier.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Upgrade to Kiro

Chatting about AWS
