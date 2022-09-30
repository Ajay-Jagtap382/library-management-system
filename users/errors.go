package users

import "errors"

var (
	errCreateSuperadmin      = errors.New("Admin does not have the access to create superadmin")
	errNotAccess             = errors.New("You don't have the access")
	errEmptyID               = errors.New("User ID must be present")
	errEmptyFirstName        = errors.New("User first name must be present")
	errEmptyLastName         = errors.New("User last name must be present")
	errNoUsers               = errors.New("No user present")
	errTakenUser             = errors.New("User has not returned the books which are issued to him please first return it then we can delete the profile.")
	errNoUserId              = errors.New("User is not present")
	errEmptyPassword         = errors.New("Password cannot be empty")
	errWrongPassword         = errors.New("Wrong Password")
	errValideGender          = errors.New("Enter a valid gender")
	errEmptyEmail            = errors.New("Email must be present")
	errEmptyMobNo            = errors.New("Mob no must be present")
	errEmptyRole             = errors.New("Role must be present")
	errRoleType              = errors.New("Enter a valid Role type from user and admin ")
	errNotValidMail          = errors.New("Email is not valid")
	errInvalidMobNo          = errors.New("Mob Number is not valid")
	errInvalidLastName       = errors.New("Last Name is not valid")
	errInvalidFirstName      = errors.New("First Name is not valid")
	errMinimumLengthPassword = errors.New("Password length should be grater than 6 characters")
	errWrongUser             = errors.New("Token id and the entered id does not matched")
)
