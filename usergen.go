package walkhub

import (
	"errors"
	"net/http"

	"github.com/lib/pq"
	"github.com/nbio/hitch"
	"github.com/tamasd/ab"
)

// AUTOGENERATED DO NOT EDIT

func NewUser(Name string, Mail string) *User {
	e := &User{
		Name: Name,
		Mail: Mail,
	}

	// HOOK: newUser()

	return e
}

func EmptyUser() *User {
	return &User{}
}

var _ ab.Validator = &User{}

func (e *User) Validate() error {
	var err error

	// HOOK: validateUser()

	return err
}

func (e *User) GetID() string {
	return e.UUID
}

var UserNotFoundError = errors.New("user not found")

const userFields = "u.uuid, u.name, u.mail, u.admin, u.created, u.lastseen"

func selectUserFromQuery(db ab.DB, query string, args ...interface{}) ([]*User, error) {
	// HOOK: beforeUserSelect()

	entities := []*User{}

	rows, err := db.Query(query, args...)

	if err != nil {
		return entities, err
	}

	for rows.Next() {
		e := EmptyUser()

		if err = rows.Scan(&e.UUID, &e.Name, &e.Mail, &e.Admin, &e.Created, &e.LastSeen); err != nil {
			return []*User{}, err
		}

		entities = append(entities, e)
	}

	// HOOK: afterUserSelect()

	return entities, err
}

func selectSingleUserFromQuery(db ab.DB, query string, args ...interface{}) (*User, error) {
	entities, err := selectUserFromQuery(db, query, args...)
	if err != nil {
		return nil, err
	}

	if len(entities) > 0 {
		return entities[0], nil
	}

	return nil, nil
}

func (e *User) Insert(db ab.DB) error {
	// HOOK: beforeUserInsert()

	err := db.QueryRow("INSERT INTO \"user\"(name, mail, admin, created, lastseen) VALUES($1, $2, $3, $4, $5) RETURNING uuid", e.Name, e.Mail, e.Admin, e.Created, e.LastSeen).Scan(&e.UUID)

	// HOOK: afterUserInsert()

	return err
}

func (e *User) Update(db ab.DB) error {
	// HOOK: beforeUserUpdate()

	result, err := db.Exec("UPDATE \"user\" SET name = $1, mail = $2, admin = $3, created = $4, lastseen = $5 WHERE uuid = $6", e.Name, e.Mail, e.Admin, e.Created, e.LastSeen, e.UUID)
	if err != nil {
		return err
	}

	aff, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if aff != 1 {
		return UserNotFoundError
	}

	// HOOK: afterUserUpdate()

	return nil
}

func (e *User) Delete(db ab.DB) error {
	// HOOK: beforeUserDelete()

	res, err := db.Exec("DELETE FROM \"user\" WHERE uuid = $1", e.UUID)
	if err != nil {
		return err
	}

	aff, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if aff != 1 {
		return UserNotFoundError
	}

	// HOOK: afterUserDelete()

	return nil
}

func LoadUser(db ab.DB, UUID string) (*User, error) {
	// HOOK: beforeUserLoad()

	e, err := selectSingleUserFromQuery(db, "SELECT "+userFields+" FROM \"user\" u WHERE u.uuid = $1", UUID)

	// HOOK: afterUserLoad()

	return e, err
}

func LoadAllUser(db ab.DB, start, limit int) ([]*User, error) {
	// HOOK: beforeUserLoadAll()

	entities, err := selectUserFromQuery(db, "SELECT "+userFields+" FROM \"user\" u ORDER BY UUID DESC LIMIT $1 OFFSET $2", limit, start)

	// HOOK: afterUserLoadAll()

	return entities, err
}

type UserService struct {
}

func (s *UserService) Register(h *hitch.Hitch) error {
	var err error

	postMiddlewares := []func(http.Handler) http.Handler{}

	getMiddlewares := []func(http.Handler) http.Handler{}

	putMiddlewares := []func(http.Handler) http.Handler{}

	deleteMiddlewares := []func(http.Handler) http.Handler{}

	// HOOK: beforeUserServiceRegister()

	if err != nil {
		return err
	}

	h.Post("/api/user", s.userPostHandler(), postMiddlewares...)

	h.Get("/api/user/:id", s.userGetHandler(), getMiddlewares...)

	h.Put("/api/user/:id", s.userPutHandler(), putMiddlewares...)

	h.Delete("/api/user/:id", s.userDeleteHandler(), deleteMiddlewares...)

	afterUserServiceRegister(h)

	return err
}

func userDBErrorConverter(err *pq.Error) ab.VerboseError {
	ve := ab.NewVerboseError(err.Message, err.Detail)

	// HOOK: convertUserDBError()

	return ve
}

func (s *UserService) userPostHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		entity := &User{}
		ab.MustDecode(r, entity)

		abort := false

		// HOOK: userPostValidation()

		if abort {
			return
		}

		if err := entity.Validate(); err != nil {
			ab.Fail(r, http.StatusBadRequest, err)
		}

		db := ab.GetDB(r)

		err := entity.Insert(db)
		ab.MaybeFail(r, http.StatusInternalServerError, ab.ConvertDBError(err, userDBErrorConverter))

		// HOOK: afterUserPostInsertHandler()

		if abort {
			return
		}

		ab.Render(r).SetCode(http.StatusCreated).JSON(entity)
	})
}

func (s *UserService) userGetHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := hitch.Params(r).ByName("id")
		db := ab.GetDB(r)
		abort := false
		loadFunc := LoadUser

		// HOOK: beforeUserGetHandler()

		if abort {
			return
		}

		entity, err := loadFunc(db, id)
		ab.MaybeFail(r, http.StatusInternalServerError, ab.ConvertDBError(err, userDBErrorConverter))
		if entity == nil {
			ab.Fail(r, http.StatusNotFound, nil)
		}

		// HOOK: afterUserGetHandler()

		if abort {
			return
		}

		ab.Render(r).JSON(entity)
	})
}

func (s *UserService) userPutHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := hitch.Params(r).ByName("id")

		entity := &User{}
		ab.MustDecode(r, entity)

		if err := entity.Validate(); entity.UUID != id || err != nil {
			ab.Fail(r, http.StatusBadRequest, err)
		}

		db := ab.GetDB(r)
		abort := false

		// HOOK: beforeUserPutUpdateHandler()

		if abort {
			return
		}

		err := entity.Update(db)
		ab.MaybeFail(r, http.StatusInternalServerError, ab.ConvertDBError(err, userDBErrorConverter))

		// HOOK: afterUserPutUpdateHandler()

		if abort {
			return
		}

		ab.Render(r).JSON(entity)
	})
}

func (s *UserService) userDeleteHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := hitch.Params(r).ByName("id")
		db := ab.GetDB(r)
		abort := false
		loadFunc := LoadUser

		// HOOK: beforeUserDeleteHandler()

		if abort {
			return
		}

		entity, err := loadFunc(db, id)
		ab.MaybeFail(r, http.StatusInternalServerError, ab.ConvertDBError(err, userDBErrorConverter))
		if entity == nil {
			ab.Fail(r, http.StatusNotFound, nil)
		}

		// HOOK: insideUserDeleteHandler()

		if abort {
			return
		}

		err = entity.Delete(db)
		ab.MaybeFail(r, http.StatusInternalServerError, ab.ConvertDBError(err, userDBErrorConverter))

		// HOOK: afterUserDeleteHandler()

		if abort {
			return
		}
	})
}

func (s *UserService) SchemaInstalled(db ab.DB) bool {
	found := ab.TableExists(db, "user")

	// HOOK: afterUserSchemaInstalled()

	return found
}

func (s *UserService) SchemaSQL() string {
	sql := "CREATE TABLE \"user\" (\n" +
		"\t\"uuid\" uuid DEFAULT uuid_generate_v4() NOT NULL,\n" +
		"\t\"name\" character varying NOT NULL,\n" +
		"\t\"mail\" character varying NOT NULL,\n" +
		"\t\"admin\" bool NOT NULL,\n" +
		"\t\"created\" timestamp with time zone NOT NULL,\n" +
		"\t\"lastseen\" timestamp with time zone NOT NULL,\n" +
		"\tCONSTRAINT user_pkey PRIMARY KEY (uuid)\n);\n"

	sql = afterUserSchemaSQL(sql)

	return sql
}
