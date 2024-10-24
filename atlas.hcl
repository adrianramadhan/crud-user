# atlas.hcl
env "local" {
  url = "postgres://postgres:password@localhost:5432/usercrudgolangdb?sslmode=disable"
  migration {
    dir = "file://migrations"
  }
}
