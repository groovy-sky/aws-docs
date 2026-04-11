# Amazon Connect Agent Workspace Email API

The Amazon Connect SDK provides an `EmailClient` which serves as an interface that your app
can use to subscribe to email contact events and make email contact requests.

The `EmailClient` accepts an optional constructor argument, `
        ConnectClientConfig` which itself is defined as:

```typescript

export type ConnectClientConfig = {
    context?: ModuleContext;
    provider?: AmazonConnectProvider;
};

```

If you do not provide a value for this config, then the client will default to using the **AmazonConnectProvider** set in the global provider scope. You can also manually
configure this using **setGlobalProvider**.

You can instantiate the agent client as follows:

```typescript

import { EmailClient } from "@amazon-connect/email";

const emailClient = new EmailClient({ provider });

```

###### Note

You must first instantiate the [AmazonConnectApp](../../../../reference/agentworkspace/latest/devguide/getting-started-initialize-sdk.md) which initializes the default AmazonConnectProvider and returns `
            { provider } `. This is the recommended option.

Alternatively, providing a constructor argument:

```typescript

import { EmailClient } from "@amazon-connect/email";

const emailClient = new EmailClient({
    context: sampleContext,
    provider: sampleProvider
});

```

###### Note

Third-party applications must be configured with Cross Contact scope in order to utilize the EmailClient APIs, and \* permission is required.

The following sections describe API calls for working with the Email API.

###### Contents

- [onAcceptedEmail()](3p-apps-email-requests-acceptedemail-subscribing.md)

- [offAcceptedEmail()](3p-apps-email-requests-acceptedemail-unsubscribing.md)

- [createDraftEmail()](3p-apps-email-requests-createdraftemail.md)

- [onDraftEmailCreated()](3p-apps-email-requests-draftemailcreated-subscribing.md)

- [offDraftEmailCreated()](3p-apps-email-requests-draftemailcreated-unsubscribing.md)

- [getEmailData()](3p-apps-email-requests-getemaildata.md)

- [getEmailThread()](3p-apps-email-requests-getemailthread.md)

- [sendEmail()](3p-apps-email-requests-sendemail.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

transfer()

onAcceptedEmail()

All content copied from https://docs.aws.amazon.com/.
