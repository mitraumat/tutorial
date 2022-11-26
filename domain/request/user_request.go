package request

type UserUpdateRequest struct {
    Username  string "json:username, validate=required"
    Name string "json:name, validate=required"
}