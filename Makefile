GPC_PROJECT_ID=fbi-confidential
SERVICE_NAME=fiber
CONTAINER_NAME=eu.gcr.io/$(GPC_PROJECT_ID)/$(SERVICE_NAME)
run: build
    docker run -p 8080:8080 $(CONTAINER_NAME)
build:
	\tdocker build -t $(CONTAINER_NAME)	.
push: build
	\tdocker push $(CONTAINER_NAME)
deploy: build push
	gcloud run deploy $(SERVICE_NAME)\
        --project $(GPC_PROJECT_ID)\
        --allow-unauthenticated\
        -q\
        --region europe-west1\
        --platform managed\
        --memory 128Mi\
        --image $(CONTAINER_NAME)
test:
	go test ./...