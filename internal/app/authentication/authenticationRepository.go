package catalogue

import (
	"fmt"

	"github.com/OPYS/internal/pkg/types"
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
)

type AuthRepository struct {
	db *sqlx.DB
}

func InitRepository(ConVar *sqlx.DB) *AuthRepository {
	return &AuthRepository{
		db: ConVar,
	}
}

func (cotr *AuthRepository) insertRegistretedUser(us types.UserInf) (err error) {
	tx := cotr.db.MustBegin()

	inq := `insert into users values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)`

	tx.MustExec(inq, us.Id, us.UserFirstname+" "+us.UserLastname, us.NormalizedUserName, us.Email, us.EmailConfirmed,
		us.NormalizedEmail, us.PasswordHash, us.SecurityStamp, us.ConcurrencyStamp,
		us.LockoutEnd, us.LockoutEnabled, us.AccessFailedCount)

	claimID := uuid.Must(uuid.NewV1())

	inq = `insert into userclaims(id,userid,claimtype,claimvalue) values($1,$2,$3,$4)`
	tx.MustExec(inq, claimID, us.Id, "firstname", us.UserFirstname)
	tx.MustExec(inq, claimID, us.Id, "lastname", us.UserLastname)
	tx.MustExec(inq, claimID, us.Id, "birthdate", us.Birthdate)

	err = tx.Commit()
	return err
}

func (cotr *AuthRepository) exists(email string) (res bool) {
	inq := `select exists(select 1 from users where email=$1)`
	cotr.db.QueryRowx(inq, email).Scan(&res)
	return res
}

func (cotr *AuthRepository) GetPasswordbyEmail(email string) (hashPassword string) {
	inq := `select passwordhash from users where email=$1`
	cotr.db.QueryRowx(inq, email).Scan(&hashPassword)
	return hashPassword
}

func (cotr *AuthRepository) GetClaimsData(email string) (answ map[string]string) {
	inq := `select id from users where email=$1`
	var id string
	cotr.db.QueryRowx(inq, email).Scan(&id)
	inq = `select claimtype, claimvalue from userclaims where userid=$1`
	rows, _ := cotr.db.Queryx(inq, id)
	answ = make(map[string]string)
	fmt.Println("okq")
	for rows.Next() {
		var claim types.ClaimsType
		rows.Scan(&claim.Type, &claim.Value)
		answ[claim.Type] = claim.Value
		// claims = append(claims, claim)
	}
	fmt.Println("okq")
	return answ
}
