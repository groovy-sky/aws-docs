# Amazon Connect Agent Workspace Contact API

The Amazon Connect SDK provides an `ContactClient` which serves as an interface that your app
in the Amazon Connect agent workspace can use to subscribe to contact events and make contact data requests.

The `ContactClient` accepts an optional constructor argument, `
        ConnectClientConfig` which itself is defined as:

```

            export type ConnectClientConfig = {
                context?: ModuleContext;
                provider?: AmazonConnectProvider;
             };

```

If you do not provide a value for this config, then the client will default to using the **AmazonConnectProvider** set in the global provider scope. You can also manually
configure this using **setGlobalProvider**.

You can instantiate the agent client as follows:

```

            import { ContactClient } from "@amazon-connect/contact";
            const contactClient = new ContactClient({ provider });

```

###### Note

You must first instantiate the [AmazonConnectApp](../../../../reference/agentworkspace/latest/devguide/getting-started-initialize-sdk.md) which initializes the default AmazonConnectProvider and returns `
            { provider } `. This is the recommended option.

Alternatively, providing a constructor argument:

```

            import { ContactClient } from "@amazon-connect/contact";

            const contactClient = new ContactClient({
                context: sampleContext,
                provider: sampleProvider
        });

```

The following sections describe API calls for working with the Contact API.

###### Contents

- [accept()](3p-apps-contact-requests-accept.md)

- [addParticipant()](3p-apps-contact-requests-addparticipant.md)

- [clear()](3p-apps-contact-requests-clear.md)

- [onCleared()](3p-apps-contact-requests-clearedsubscribing.md)

- [offCleared()](3p-apps-contact-requests-clearedunsubscribing.md)

- [onConnected()](3p-apps-contact-events-connected-sub.md)

- [offConnected()](3p-apps-contact-events-connected-unsub.md)

- [disconnectParticipant()](3p-apps-contact-requests-disconnectparticipant.md)

- [engagePreviewContact()](3p-apps-contact-requests-engagepreviewcontact.md)

- [getAttribute()](3p-apps-contact-requests-getattribute.md)

- [getAttributes()](3p-apps-contact-requests-getattributes.md)

- [getChannelType()](3p-apps-contact-requests-getchanneltype.md)

- [getContact()](3p-apps-contact-requests-getcontact.md)

- [getInitialContactId()](3p-apps-contact-requests-getinitialcontactid.md)

- [getParticipant()](3p-apps-contact-requests-getparticipant.md)

- [getParticipantState()](3p-apps-contact-requests-getparticipantstate.md)

- [getPreviewConfiguration()](3p-apps-contact-requests-getpreviewconfiguration.md)

- [getQueue()](3p-apps-contact-requests-getqueue.md)

- [getQueueTimestamp()](3p-apps-contact-requests-getqueuetimestamp.md)

- [getStateDuration()](3p-apps-contact-requests-getstateduration.md)

- [isPreviewMode()](3p-apps-contact-requests-ispreviewmode.md)

- [listContacts()](3p-apps-contact-requests-listcontacts.md)

- [listParticipants()](3p-apps-contact-requests-listparticipants.md)

- [onMissed()](3p-apps-contact-events-missed-sub.md)

- [offMissed()](3p-apps-contact-events-missed-unsub.md)

- [offIncoming()](3p-apps-contact-requests-off-incoming.md)

- [onIncoming()](3p-apps-contact-requests-on-incoming.md)

- [onParticipantAdded()](3p-apps-contact-events-participantadded-sub.md)

- [offParticipantAdded()](3p-apps-contact-events-participantadded-unsub.md)

- [onParticipantDisconnected()](3p-apps-contact-events-participantdisconnected-sub.md)

- [offParticipantDisconnected()](3p-apps-contact-events-participantdisconnected-unsub.md)

- [onParticipantStateChanged()](3p-apps-contact-events-participantstatechanged-sub.md)

- [onStartingAcw()](3p-apps-contact-events-startingacw-sub.md)

- [offStartingAcw()](3p-apps-contact-events-startingacw-unsub.md)

- [transfer()](3p-apps-contact-requests-transfer.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

launchApp()

accept()

All content copied from https://docs.aws.amazon.com/.
