# Create a new private key

export PROJECT_ID=gcpnext-s9-demo
export SVC_ACCT=datastore-guestbook-access

gcloud iam service-accounts keys create ./private-key.json --iam-account=${SVC_ACCT}@${PROJECT_ID}.iam.gserviceaccount.com
