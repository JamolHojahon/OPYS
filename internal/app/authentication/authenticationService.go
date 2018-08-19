package catalogue

import (
	"crypto/md5"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/OPYS/internal/pkg/types"

	uuid "github.com/satori/go.uuid"
)

type AuthService struct {
	repo *AuthRepository
}

func InitService(aRep *AuthRepository) *AuthService {
	return &AuthService{
		repo: aRep,
	}
}

var emailVal = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,255}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,255}[a-zA-Z0-9])?)*$")
var nameValidation = regexp.MustCompile("^[a-zA-Z]{3,255}$") // username validation

func (srv *AuthService) TestuserForSignIn(us *types.UserLogin) (error, map[string]string) {
	var forerr map[string]string
	if res := srv.repo.exists(us.Email); !res {
		return errors.New("Email doesnt exists"), forerr
	}
	if !srv.CheckPassword(us.Password, us.Email) {
		return errors.New("Wrong Password!"), forerr
	}
	booboo := srv.repo.GetClaimsData(us.Email) //!
	return nil, booboo
}

func (srv *AuthService) CreateUser(us *types.InfoForReg) (error, types.UserInf) {
	var forErr types.UserInf
	if !emailVal.MatchString(us.Email) {
		return errors.New("Wrong email"), forErr
	}
	if us.Password != us.ConfirmPassword {
		return errors.New("Password does not match!"), forErr
	}
	if ans := srv.repo.exists(us.Email); ans {
		return errors.New("Email exists!"), forErr
	}

	myUser, err := srv.CreateNewUserInUsers(*us)

	return err, myUser
}

func (srv *AuthService) HashIt(someStr string) string {
	h := md5.New()
	h.Write([]byte(someStr))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func (srv *AuthService) CreateNewUserInUsers(us types.InfoForReg) (forerr types.UserInf, err error) {
	var firstname, lastname, birthdate string
	for _, val := range us.Claims {
		switch val.Type {
		case "firstname":
			if !nameValidation.MatchString(val.Value) {
				return forerr, errors.New("Wrong Feeld Feilds")
			}
			firstname = val.Value
		case "lastname":
			if !nameValidation.MatchString(val.Value) {
				return forerr, errors.New("ERROR With entered text")
			}
			lastname = val.Value
		case "birthdate":
			birthdate = val.Value
		default:
			return forerr, errors.New("FuckuTimNazar!AxTiPidrila")
		}
	}
	if firstname != "" || lastname != "" || birthdate != "" {
		var newUs = types.UserInf{
			Id:                 uuid.Must(uuid.NewV1()), // generating uuid
			UserFirstname:      firstname,
			UserLastname:       lastname,
			NormalizedUserName: strings.ToUpper(firstname + " " + lastname),
			Email:              us.Email,
			EmailConfirmed:     true,
			NormalizedEmail:    strings.ToUpper(us.Email),
			PasswordHash:       srv.HashIt(us.Password),
			SecurityStamp:      srv.HashIt(firstname), // must generate something
			ConcurrencyStamp:   srv.HashIt(firstname), // olso
			LockoutEnd:         time.Time.Format(time.Now(), "01-02-2006"),
			LockoutEnabled:     false,
			AccessFailedCount:  0,
			Birthdate:          birthdate,
		}
		err := srv.repo.insertRegistretedUser(newUs)
		if err == nil {
			return newUs, nil
		}
	}
	return forerr, errors.New("Fields filled Wrong!<flname,birthdate>")

}

func (srv *AuthService) CheckPassword(password, email string) bool {
	hashPass := srv.repo.GetPasswordbyEmail(email)
	return hashPass == srv.HashIt(password)
}
