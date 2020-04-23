https://www.terraform.io/docs/extend/writing-custom-providers.html

go build -o terraform-provider-artifactory


docker run --name artifactory2 -p 8081:8081 -p 8082:8082 docker.bintray.io/jfrog/artifactory-pro