# Dockerfile

# Get latest golang docker image.
FROM golang:latest

# Create a directory inside the container to store our web-app and then make it working directory.
RUN mkdir -p /go/src/api
WORKDIR /go/src/api

# Copy the web-app directory into the container.
COPY . /go/src/api

# Download and install third party dependencies into the container.
RUN go get github.com/codegangsta/gin
RUN go get -d -v ./...
RUN go install -v ./...

#RUN go-wrapper download
##RUN go-wrapper install

# Set the PORT environment variable
ENV PORT 8080

# Expose port 8080 to the host so that outer-world can access your application
EXPOSE 8080

# Tell Docker what command to run when the container starts
#CMD ["go-wrapper", "run"]

#ENTRYPOINT ["./api"]
CMD ["api", "run"]


# to build, run and push
#docker build -t <hub-user>/<repo-name>[:<tag>]
#docker build -t colemind/assessment-api:dev-api .

#docker run -it --rm --name assessment-api api
#docker run --rm -d  -p 8080:8080/tcp api:latest

#docker push <hub-user>/<repo-name>:<tag>
#docker push colemind/assessment-api:tagname