provider "artifactory" {
  username = "admin"
  password = "password"
}

/*
resource "artifactory_permission" "first3" {
    repo {
      repositories = ["one", "two"]
      include_patterns = ["one", "two"]
      exclude_patterns = ["tree", "four"]
      actions {
        groups = jsonencode({
          "dev": ["res","test"],
          "prod": ["res","test"],
        })
        users = jsonencode({
          "meyer": ["res","test"],
          "ole": ["res","test"],
        })
      } 
    }
    
}*/



resource "artifactory_repository" "first1" {
    key = "first1"
    rclass = "local"
    package_type = "npm"
    description = "hej4"
}