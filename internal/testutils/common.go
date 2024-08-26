package testutils

import (
	"fmt"
	"time"
)

const (
	RedisHashTableData  = "test_redis_table_data"
	RedisHashTableLogin = "test_redis_table_login"

	Login    = "login@example.com"
	Password = "password"

	SimpleOkResponse   = `{"result":{"status":"ok"},"error":null}`
	BadUserResponse    = `{"result":null,"error":"wrong login or password"}`
	UserExistsResponse = `{"result":null,"error":"user with same login exists"}`
	NotFoundResponse   = `{"result":null,"error":"entity not found"}`

	WorkerCommitInterval = 200 * time.Millisecond
)

func MakeUserBody() string {
	const template = `{"login": "%s", "password": "%s"}`

	return fmt.Sprintf(template, Login, Password)
}

func MakeUserBodyWithLogin(login string) string {
	const template = `{"login": "%s", "password": "%s"}`

	return fmt.Sprintf(template, login, Password)
}

func MakeUserUpdateBody(newPassword string) string {
	const template = `{"login": "%s", "old_password": "%s", "new_password": "%s"}`

	return fmt.Sprintf(template, Login, Password, newPassword)
}
