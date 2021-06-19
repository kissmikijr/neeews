#Deploy cloudfunction
gcloud functions deploy $1 --runtime go113 --trigger-http --allow-unauthenticated --region europe-west6
#Create new api config
# gcloud api-gateway api-configs create $1 --api=neeews-api --openapi-spec=openapi-functions.yaml --pr oject=graphic-charter-314415 --backend-auth-service-account=graphic-charter-314415@appspot.gserviceaccount.com 