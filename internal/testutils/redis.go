package testutils

func ExpectGetAllUsers(res map[string]string) Opt {
	return func(app *App) {
		app.redisMock.ExpectHGetAll(RedisHashTableData).SetVal(res)
	}
}

func ExpectGetUser(key, res string) Opt {
	return func(app *App) {
		exp := app.redisMock.ExpectHGet(RedisHashTableData, key)

		if res != "" {
			exp.SetVal(res)
			return
		}

		exp.RedisNil()
	}
}

func ExpectCreateUser(key, val string) Opt {
	return func(app *App) {
		app.redisMock.ExpectHSet(RedisHashTableData, key, val).SetVal(1)
	}
}

func ExpectGetUserLastLogin(key, res string) Opt {
	return func(app *App) {
		exp := app.redisMock.ExpectHGet(RedisHashTableLogin, key)

		if res != "" {
			exp.SetVal(res)
			return
		}

		exp.RedisNil()
	}
}

func ExpectSetUserLastLogin(key string) Opt {
	return func(app *App) {
		app.redisMock.Regexp().ExpectHSet(RedisHashTableLogin, key, `^[0-9]+$`).SetVal(1)
	}
}

func ExpectDeleteUser(key string) Opt {
	return func(app *App) {
		app.redisMock.
			ExpectHDel(RedisHashTableData, key).
			SetVal(1)
	}
}
