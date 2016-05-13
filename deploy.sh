
#!/bin/bash



FOUND_STRING='cloud-go-ref';


ECHO "Testing kubectl to see if we have an existing deployment..."

output=$(kubectl describe deployments)


if [ "${output/$FOUND_STRING}" = "$output" ] ; then
  echo "DID NOT FIND THE DEPLOYMENT"
  #
  # add the deployment...
  #
  $(kubectl expose deployment cloud-go-ref --type="LoadBalancer")
else
  echo "FOUND THE DEPLOYMENT"
  $(kubectl patch deployment cloud-go-ref -p '{"spec":{"template":{"spec":{"containers":[{"name":"cloud-go-ref","image":"gcr.io/$GCLOUD_PROJECT/cloud-go-ref:'"$CIRCLE_SHA1"'"}]}}}}')
fi


#deployments.extensions "cloud-go-ref" not found
