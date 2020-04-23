call go build -o terraform-provider-artifactory.exe
call terraform init
call terraform validate
call terraform apply -auto-approve 