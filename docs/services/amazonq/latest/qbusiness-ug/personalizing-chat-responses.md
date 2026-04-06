# Personalizing Amazon Q Business chat responses

Amazon Q Business uses location and job-related information from your IAM Identity Center
instance to generate personalized responses that are relevant to your end user. For
example, if your end user asks "What are the company holidays for this year?", Amazon Q Business might list region-specific holidays based on their country.

To enable response personalization, add **Address** information and
**Job related information** for users in the IAM Identity Center instance that
connects to your Amazon Q Business application. For more information, see [Add\
users](https://docs.aws.amazon.com/singlesignon/latest/userguide/addusers.html) in the IAM Identity Center User Guide.

User personalization data that's in your IAM Identity Center instance is connected to your Amazon Q Business application environment, and responses are personalized by default. You can
deactivate response personalization at any time by using the [Admin controls and guardrails](guardrails-global-controls.md#guardrails-global-controls-customizing) feature in Amazon Q Business.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Upload files and chat

Quick prompts
